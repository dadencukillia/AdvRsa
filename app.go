package main

import (
	"context"
	"encoding/base64"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) AGenerateKeyPair() []string {
	priv, publ, _ := GenerateKeyPair(2048)
	strpriv := PrivateKeyToBytes(priv)
	strpubl, _ := PublicKeyToBytes(publ)
	encpriv := base64.StdEncoding.EncodeToString(strpriv)
	encpubl := base64.StdEncoding.EncodeToString(strpubl)
	return []string{encpubl, encpriv}
}

func (a *App) AEncryptMessage(msg, publKey string) []string {
	strPublicKey, err := base64.StdEncoding.DecodeString(publKey)
	if err != nil {
		return []string{"", "Пошкоджений публічний ключ!"}
	}
	publicKey, err := BytesToPublicKey(strPublicKey)
	if err != nil {
		return []string{"", "Пошкоджений публічний ключ!"}
	}
	encrypted, err := EncryptWithPublicKey([]byte(msg), publicKey)
	if err != nil {
		return []string{"", "Помилка шифрування!"}
	}
	return []string{base64.StdEncoding.EncodeToString(encrypted), "0"}
}

func (a *App) ADecryptMessage(msg, privKey string) []string {
	strPrivateKey, err := base64.StdEncoding.DecodeString(privKey)
	if err != nil {
		return []string{"", "Пошкоджений приватний ключ!"}
	}
	privateKey, err := BytesToPrivateKey(strPrivateKey)
	if err != nil {
		return []string{"", "Пошкоджений приватний ключ!"}
	}
	strEncMsg, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		return []string{"", "Помилка читання зашифрованого повідомлення!"}
	}
	text, err := DecryptWithPrivateKey(strEncMsg, privateKey)
	if err != nil {
		return []string{"", "Помилка розшифрування повідомлення!"}
	}
	return []string{string(text), "0"}
}
