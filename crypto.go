package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/nanjishidu/gomini/gocrypto"
)

type Instance func() Interface
type Interface interface {
	Encrypt(key, s string) (string, error)
	Decrypt(key, s string) (string, error)
}

var adapters = make(map[string]Instance)

func Register(name string, adapter Instance) {
	adapters[name] = adapter
}
func NewInstance(adapterName string) (adapter Instance, err error) {
	adapter, ok := adapters[adapterName]
	if !ok {
		err = fmt.Errorf("instance: unknown adapter name %q (forgot to import?)", adapterName)
		return
	}
	return
}

func init() {
	Register("aes-cfb", NewAesCFB)
	Register("aes-cbc", NewAesCBC)
	Register("des-cfb", NewDesCFB)
	Register("des-cbc", NewDesCBC)
	Register("3des-cfb", NewTripleDesCFB)
	Register("3des-cbc", NewTripleDesCBC)

}

func NewAesCFB() Interface {
	return new(AesCFB)
}

func NewAesCBC() Interface {
	return new(AesCBC)
}

func NewDesCFB() Interface {
	return new(DesCFB)
}

func NewDesCBC() Interface {
	return new(DesCBC)
}
func NewTripleDesCFB() Interface {
	return new(TripleDesCFB)
}

func NewTripleDesCBC() Interface {
	return new(TripleDesCBC)
}

type AesCFB struct{}

// aes CFB encrypt
func (m *AesCFB) Encrypt(key, s string) (string, error) {
	//set aeskey
	err := gocrypto.SetAesKey(gocrypto.Md5(key))
	if err != nil {
		return "", err
	}
	ace, err := gocrypto.AesCFBEncrypt([]byte(kpassSign + s))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ace), nil
}

// aes CFB decrypt
func (m *AesCFB) Decrypt(key, s string) (string, error) {
	err := gocrypto.SetAesKey(gocrypto.Md5(key))
	if err != nil {
		return "", err
	}
	bsd, err := base64.StdEncoding.DecodeString(strings.Replace(s, kpassSign, "", -1))
	if err != nil {
		return "", err
	}
	acd, err := gocrypto.AesCFBDecrypt(bsd)
	if !strings.Contains(string(acd), kpassSign) {
		return "", errors.New("aesKey is error")
	}
	if err != nil {
		return "", err
	}
	return strings.Replace(string(acd), kpassSign, "", -1), err
}

type AesCBC struct{}

// aes CFB encrypt
func (m *AesCBC) Encrypt(key, s string) (string, error) {
	//set aeskey
	err := gocrypto.SetAesKey(gocrypto.Md5(key))
	if err != nil {
		return "", err
	}
	ace, err := gocrypto.AesCBCEncrypt([]byte(kpassSign + s))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ace), nil
}

// aes CFB decrypt
func (m *AesCBC) Decrypt(key, s string) (string, error) {
	err := gocrypto.SetAesKey(gocrypto.Md5(key))
	if err != nil {
		return "", err
	}
	bsd, err := base64.StdEncoding.DecodeString(strings.Replace(s, kpassSign, "", -1))
	if err != nil {
		return "", err
	}
	acd, err := gocrypto.AesCBCDecrypt(bsd)
	if !strings.Contains(string(acd), kpassSign) {
		return "", errors.New("aesKey is error")
	}
	if err != nil {
		return "", err
	}
	return strings.Replace(string(acd), kpassSign, "", -1), err
}

type DesCBC struct{}

// des CBC encrypt
func (m *DesCBC) Encrypt(key, s string) (string, error) {
	//set aeskey
	err := gocrypto.SetDesKey(gocrypto.Md5(key)[0:8])
	if err != nil {
		return "", err
	}
	ace, err := gocrypto.DesCBCEncrypt([]byte(kpassSign + s))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ace), nil
}

// des CBC decrypt
func (m *DesCBC) Decrypt(key, s string) (string, error) {
	err := gocrypto.SetDesKey(gocrypto.Md5(key)[0:8])
	if err != nil {
		return "", err
	}
	bsd, err := base64.StdEncoding.DecodeString(strings.Replace(s, kpassSign, "", -1))
	if err != nil {
		return "", err
	}
	acd, err := gocrypto.DesCBCDecrypt(bsd)
	if !strings.Contains(string(acd), kpassSign) {
		return "", errors.New("desKey is error")
	}
	if err != nil {
		return "", err
	}
	return strings.Replace(string(acd), kpassSign, "", -1), err
}

type DesCFB struct{}

// des CFB encrypt
func (m *DesCFB) Encrypt(key, s string) (string, error) {
	//set aeskey
	err := gocrypto.SetDesKey(gocrypto.Md5(key)[0:8])
	if err != nil {
		return "", err
	}
	ace, err := gocrypto.DesCFBEncrypt([]byte(kpassSign + s))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ace), nil
}

// des CFB decrypt
func (m *DesCFB) Decrypt(key, s string) (string, error) {
	err := gocrypto.SetDesKey(gocrypto.Md5(key)[0:8])
	if err != nil {
		return "", err
	}
	bsd, err := base64.StdEncoding.DecodeString(strings.Replace(s, kpassSign, "", -1))
	if err != nil {
		return "", err
	}
	acd, err := gocrypto.DesCFBDecrypt(bsd)
	if !strings.Contains(string(acd), kpassSign) {
		return "", errors.New("desKey is error")
	}
	if err != nil {
		return "", err
	}
	return strings.Replace(string(acd), kpassSign, "", -1), err
}

type TripleDesCBC struct{}

// triple des CBC encrypt
func (m *TripleDesCBC) Encrypt(key, s string) (string, error) {
	//set aeskey
	err := gocrypto.SetDesKey(gocrypto.Md5(key)[0:24])
	if err != nil {
		return "", err
	}
	ace, err := gocrypto.TripleDesCBCEncrypt([]byte(kpassSign + s))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ace), nil
}

//  triple des CBC decrypt
func (m *TripleDesCBC) Decrypt(key, s string) (string, error) {
	err := gocrypto.SetDesKey(gocrypto.Md5(key)[0:24])
	if err != nil {
		return "", err
	}
	bsd, err := base64.StdEncoding.DecodeString(strings.Replace(s, kpassSign, "", -1))
	if err != nil {
		return "", err
	}
	acd, err := gocrypto.TripleDesCBCDecrypt(bsd)
	if !strings.Contains(string(acd), kpassSign) {
		return "", errors.New("desKey is error")
	}
	if err != nil {
		return "", err
	}
	return strings.Replace(string(acd), kpassSign, "", -1), err
}

type TripleDesCFB struct{}

// triple des CFB encrypt
func (m *TripleDesCFB) Encrypt(key, s string) (string, error) {
	//set aeskey
	err := gocrypto.SetDesKey(gocrypto.Md5(key)[0:24])
	if err != nil {
		return "", err
	}
	ace, err := gocrypto.TripleDesCFBEncrypt([]byte(kpassSign + s))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ace), nil
}

// triple des CFB decrypt
func (m *TripleDesCFB) Decrypt(key, s string) (string, error) {
	err := gocrypto.SetDesKey(gocrypto.Md5(key)[0:24])
	if err != nil {
		return "", err
	}
	bsd, err := base64.StdEncoding.DecodeString(strings.Replace(s, kpassSign, "", -1))
	if err != nil {
		return "", err
	}
	acd, err := gocrypto.TripleDesCFBDecrypt(bsd)
	if !strings.Contains(string(acd), kpassSign) {
		return "", errors.New("desKey is error")
	}
	if err != nil {
		return "", err
	}
	return strings.Replace(string(acd), kpassSign, "", -1), err
}
