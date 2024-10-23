package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func verifySignature(from string, signatureHex string, message string) (verified bool, err error) {
	signature, err := hexutil.Decode(signatureHex)
	if err != nil {
		return
	}

	messageHash := accounts.TextHash([]byte(message))
	if signature[crypto.RecoveryIDOffset] == 27 || signature[crypto.RecoveryIDOffset] == 28 {
		signature[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	sigPublicKeyECDSA, err := crypto.SigToPub(messageHash, signature)
	if err != nil {
		err = fmt.Errorf("sigToPub: %w", err)
		return
	}

	if common.HexToAddress(from) != crypto.PubkeyToAddress(*sigPublicKeyECDSA) {
		log.Println("pub key mismatch")
		return false, nil
	}

	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)

	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	verified = crypto.VerifySignature(sigPublicKeyBytes, messageHash, signatureNoRecoverID)
	return
}
