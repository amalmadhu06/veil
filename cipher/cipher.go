package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
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
