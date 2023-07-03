// Package veil provides a simple key-value storage mechanism with encryption for sensitive data.
package veil

import (
	"encoding/json"
	"errors"
	"github.com/amalmadhu06/veil/cipher"
	"io"
	"os"
	"sync"
)

// Vile represents a key-value store with encryption capabilities.
type Vile struct {
	encodingKey string            // The key used for encryption/decryption.
	filepath    string            // The file path to store the data.
	mutex       sync.Mutex        // Mutex to synchronize access to the key-value store.
	keyValues   map[string]string // The actual key-value data.
}

// NewVile creates a new instance of Vile with the provided encoding key and file path.
func NewVile(encodingKey, filepath string) *Vile {
	return &Vile{
		encodingKey: encodingKey,
		filepath:    filepath,
	}
}

// Set sets the value for the given key in the key-value store.
// It encrypts the data and saves it to the file.
func (v *Vile) Set(key, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	// Load existing data from file, if any.
	err := v.load()
	if err != nil {
		return err
	}

	// Update the value for the given key.
	v.keyValues[key] = value

	// Save the updated key-value store to the file.
	err = v.save()
	return err
}

// Get retrieves the value for the given key from the key-value store.
// It decrypts the data from the file and returns the value.
func (v *Vile) Get(key string) (string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	// Load data from file.
	err := v.load()
	if err != nil {
		return "", err
	}

	// Retrieve the value for the given key.
	value, ok := v.keyValues[key]
	if !ok {
		return "", errors.New("secret: no value for that key")
	}
	return value, nil
}

// save encrypts and saves the key-value store to the file.
func (v *Vile) save() error {
	// Open the file for writing.
	f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	// Create an encrypted writer.
	w, err := cipher.EncryptWriter(v.encodingKey, f)
	if err != nil {
		return err
	}

	// Write the encrypted key-value store to the writer.
	return v.writeKeyValues(w)
}

// load reads and decrypts the key-value store from the file.
func (v *Vile) load() error {
	// Open the file for reading.
	f, err := os.Open(v.filepath)
	if err != nil {
		// If the file doesn't exist, create an empty key-value store.
		v.keyValues = make(map[string]string)
		return nil
	}
	defer f.Close()

	// Create a decrypted reader.
	r, err := cipher.DecryptReader(v.encodingKey, f)
	if err != nil {
		return err
	}

	// Read the decrypted key-value store from the reader.
	return v.readKeyValues(r)
}

// writeKeyValues writes the key-value store to the given writer.
func (v *Vile) writeKeyValues(w io.Writer) error {
	// Create a JSON encoder for writing the key-value store.
	enc := json.NewEncoder(w)

	// Encode the key-value store to JSON and write it to the writer.
	return enc.Encode(v.keyValues)
}

// readKeyValues reads the key-value store from the given reader.
func (v *Vile) readKeyValues(r io.Reader) error {
	// Create a JSON decoder for reading the key-value store.
	dec := json.NewDecoder(r)

	// Decode the JSON data and store it in the key-value store.
	return dec.Decode(&v.keyValues)
}
