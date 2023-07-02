package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

// encryptStream creates and returns a cipher.Stream for encrypting data using
// AES-256 in CFB mode. It takes a key as a string and an initialization vector
// (IV) as a byte slice. Returns the cipher.Stream for encryption or an error
// if the cipher block creation fails.
func encryptStream(key string, iv []byte) (cipher.Stream, error) {
	block, err := newCipherBlock(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewCFBEncrypter(block, iv), nil
}

// Encrypt encrypts the provided plaintext using AES-256 in CFB mode with the given
// key. It returns the encrypted ciphertext as a hex-encoded string or an error if
// encryption fails.
func Encrypt(key, plaintext string) (string, error) {
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream, err := encryptStream(key, iv)
	if err != nil {
		return "", err
	}
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	return fmt.Sprintf("%x", ciphertext), nil
}

// EncryptWriter creates and returns a cipher.StreamWriter that encrypts data
// written to it using AES-256 in CFB mode. It takes a key as a string and an
// io.Writer where the encrypted data will be written. Returns the
// cipher.StreamWriter for encryption or an error if the initialization vector
// cannot be written.
func EncryptWriter(key string, w io.Writer) (*cipher.StreamWriter, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream, err := encryptStream(key, iv)
	if err != nil {
		return nil, err
	}
	n, err := w.Write(iv)
	if n != len(iv) || err != nil {
		return nil, errors.New("encrypt: unable to write full iv to writer")
	}
	return &cipher.StreamWriter{S: stream, W: w}, nil
}

// decryptStream creates and returns a cipher.Stream for decrypting data using
// AES-256 in CFB mode. It takes a key as a string and an initialization vector (IV)
// as a byte slice. Returns the cipher.Stream for decryption or an error if the
// cipher block creation fails.
func decryptStream(key string, iv []byte) (cipher.Stream, error) {
	block, err := newCipherBlock(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewCFBDecrypter(block, iv), nil
}

// Decrypt decrypts the provided ciphertext using AES-256 in CFB mode with the
// given key. It takes the encrypted ciphertext as a hex-encoded string and returns
// the decrypted plaintext or an error if decryption fails.
func Decrypt(key, cipherHex string) (string, error) {
	ciphertext, err := hex.DecodeString(cipherHex)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("encrypt: cipher too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream, err := decryptStream(key, iv)
	if err != nil {
		return "", err
	}

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext), nil
}

// DecryptReader creates and returns a cipher.StreamReader that decrypts data read
// from an io.Reader using AES-256 in CFB mode. It takes a key as a string and an
// io.Reader from which the encrypted data will be read. Returns the
// cipher.StreamReader for decryption or an error if the initialization vector
// cannot be read.
func DecryptReader(key string, r io.Reader) (*cipher.StreamReader, error) {
	iv := make([]byte, aes.BlockSize)
	n, err := r.Read(iv)
	if n < len(iv) || err != nil {
		return nil, errors.New("encrypt: unable to read the full iv")
	}
	stream, err := decryptStream(key, iv)
	if err != nil {
		return nil, err
	}
	return &cipher.StreamReader{S: stream, R: r}, nil
}

// newCipherBlock creates and returns a new AES cipher.Block using the provided key.
// It takes the key as a string and returns the cipher.Block or an error if key
// hashing fails.
func newCipherBlock(key string) (cipher.Block, error) {
	hasher := md5.New()
	fmt.Fprint(hasher, key)
	cipherKey := hasher.Sum(nil)
	return aes.NewCipher(cipherKey)
}
