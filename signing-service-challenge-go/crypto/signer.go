package crypto

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

// Signer defines a contract for different types of signing implementations.
type Signer interface {
	Sign(dataToBeSigned []byte) ([]byte, error)
}

type RSASigner struct {
	marshaller      RSAMarshaler
	privateKeyBytes []byte
}

func NewRSASigner(privateKey []byte) RSASigner {
	return RSASigner{
		marshaller:      NewRSAMarshaler(),
		privateKeyBytes: privateKey,
	}
}

type ECCSigner struct {
	marshaller      ECCMarshaler
	privateKeyBytes []byte
}

func NewECCSigner(privateKey []byte) ECCSigner {
	return ECCSigner{
		marshaller:      NewECCMarshaler(),
		privateKeyBytes: privateKey,
	}
}

func getHashSum(dataToBeSigned []byte) ([]byte, error) {
	msgHash := sha256.New()
	_, err := msgHash.Write(dataToBeSigned)
	if err != nil {
		return nil, fmt.Errorf("failed to get hash sum: %w", err)
	}
	return msgHash.Sum(nil), nil
}

func (signer RSASigner) Sign(dataToBeSigned []byte) ([]byte, error) {
	keyPair, err := signer.marshaller.Unmarshal(signer.privateKeyBytes)
	if err != nil {
		return nil, err
	}
	return keyPair.Private.Sign(rand.Reader, dataToBeSigned, nil)
}

func (signer ECCSigner) Sign(dataToBeSigned []byte) ([]byte, error) {
	keyPair, err := signer.marshaller.Unmarshal(signer.privateKeyBytes)
	if err != nil {
		return nil, err
	}
	return keyPair.Private.Sign(rand.Reader, dataToBeSigned, nil)
}

type SignatureAlgorithm string

type SignatureAlgorithmRegistry struct {
	RSA           string
	ECDSA         string
	AlgorithmList []string
}

func NewSignatureAlgorithmRegistry() SignatureAlgorithmRegistry {
	rsa := "RSA"
	ecdsa := "ECDSA"
	return SignatureAlgorithmRegistry{
		RSA:           rsa,
		ECDSA:         ecdsa,
		AlgorithmList: []string{rsa, ecdsa},
	}
}
