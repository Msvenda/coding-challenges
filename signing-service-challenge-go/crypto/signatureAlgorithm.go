package crypto

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