#  Veil: CLI for Secure Secrets Management

Veil is a command-line interface (CLI) tool built in Golang that allows you to securely store and manage secrets. It provides a convenient way to store sensitive information such as API keys, passwords, or any other confidential data you need to keep safe.

## Features

Veil currently supports the following commands:

* `veil set` - Sets a new secret by providing a key-value pair.
* `veil get` - Retrieves the value associated with a specific key.

The following commands will be soon:
* `veil list` - Lists all the stored secrets.
* `veil rm` - Removes a specific secret.
* `veil edit` - Modifies the value of an existing secret.

## Installation

To use Veil, follow these installation steps:

1. Ensure you have Golang installed on your system. If not, you can download it from the official Golang website (https://golang.org/) and follow the installation instructions.

2. Clone the veil repository to your local machine:
```shell
git clone https://github.com/amalmadhu06/veil.git
```

3. Change to the veil project directory:
```shell
cd veil
```
4. Install Veil CLI:
```shell
go install 
```
5. Verify the installation
```shell
veil
```
If it displays Veil logo and description, installation is a success. Otherwise, please build veil and add the executable to your system's path

## Usage

To use Veil, follow the examples below for each command:

### Set a New Secret

To set a new secret, use the `set` command followed by the key-value pair:
```shell
veil set twitter_api_key <your_secret_api_key>
```

This command will encrypt and store the value associated with the specified key securely.

### Get a Secret

To retrieve a secret, use the `get` command followed by the key:
```shell
veil get twitter_api_key
```
The command will decrypt and display the value associated with the specified key.


This command will update the value associated with the specified key.

## Security

Vile ensures the security of your secrets by utilizing strong encryption algorithms. It employs the following packages from the Golang standard library:

* `crypto/aes` - Implements the AES encryption algorithm.
* `crypto/cipher` - Provides symmetric encryption and decryption operations.
* `crypto/md5` - Computes the MD5 hash of data.
* `crypto/rand` - Generates random numbers for cryptographic operations.
* `encoding/hex` - Encodes and decodes hexadecimal strings.

These encryption mechanisms guarantee that your secrets remain safe and can only be accessed by authorized individuals.

## Contributions

Contributions to Vile are welcome!
If you find any bugs, issues,
or have suggestions for new features,
please open an issue on the Vile GitHub repository
(https://github.com/amalmadhu06/vile) and follow the contribution guidelines.
