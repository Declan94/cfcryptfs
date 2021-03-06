package main

// helloCrypter implements corecrypter.CoreCrypter interface
type helloCrypter struct {
}

// EncryptedLen return encrypted length according to plaintext length
func (hc *helloCrypter) EncryptedLen(plainLen int) int {
	return plainLen
}

func (hc *helloCrypter) DecryptedLen(cipherLen int) int {
	return cipherLen
}

// Encrypt encrypt src to dest
func (hc *helloCrypter) Encrypt(dest, src []byte) error {
	copy(dest, src)
	return nil
}

// Decrypt decrypt src to dest
func (hc *helloCrypter) Decrypt(dest, src []byte) error {
	copy(dest, src)
	return nil
}
