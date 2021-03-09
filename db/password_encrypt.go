package db

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

/*EncryptPassword it is the routine to encrypt passwords*/
func EncryptPassword(pass string) (string, error) {
	/* costo := 8 //El algoritmo de encriptación hará (2 elevado al costo) pasados por el texto
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err */
	key := "1234567 1234567 1234567 1234567k"

	encryptedPassword, err := encrypt([]byte(key), []byte(pass))

	if err != nil {
		return "", err
	}

	return string(encryptedPassword), nil
}

/* DecryptPassword it is the routine to decrypt passwords */
func DecryptPassword(encryptedPass string) (string, error) {
	key := "1234567 1234567 1234567 1234567k"

	decryptedPassword, err := decrypt([]byte(key), []byte(encryptedPass))

	if err != nil {
		return "", err
	}

	return string(decryptedPassword), nil
}

//Encrypt receives a map of bytes of the text to encrypt and a key and returns the cyphered thext
func encrypt(key, text []byte) ([]byte, error) {
	//NewCipher creates and returns a new cipher.Block.
	//The key received should be the AES key, either 16/24/32 bytes to select AES-128/AES-192/AES-256
	//(AES it is a U.S. Federal Information Processing Standards for encripting)
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}
	//EncodeToString returns the base64 encoding of text
	b := base64.StdEncoding.EncodeToString(text)
	//We build a map of bytes with the AES block size in bytes + the lenght of our text encoded
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	//NewCFBEncrypter returns a Stream which encrypts with cipher feedback mode, using the given Block.
	//The iv must be the same length as the Block's block size.(IV = Inicialization Vector)
	cfb := cipher.NewCFBEncrypter(block, iv)
	//XORKeyStream adds the cipher to our text
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

func decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}
