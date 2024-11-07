package interop

import (
	"encoding/hex"
	"errors"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"golang.org/x/crypto/sha3"
)

// Account represents a blockchain account derived from a public key.
type Account struct {
	params        *chaincfg.Params
	PubKey        []byte
	AddressP2PK   string
	AddressP2PKH  string
	AddressP2WPKH string
	AddressP2TR   string
	AddressEth    string
}

// NewAccount creates a new Account instance for a given public key and network parameters.
func NewAccount(pubKey []byte, params *chaincfg.Params) *Account {
	return &Account{
		params: params,
		PubKey: pubKey,
	}
}

// GenerateAddresses generates all supported addresses (P2PK, P2PKH, P2WPKH, P2TR, and Ethereum).
func (a *Account) GenerateAddresses() error {
	if err := a.generateP2PK(); err != nil {
		return errors.New("failed to generate P2PK address: " + err.Error())
	}
	if err := a.generateP2PKH(); err != nil {
		return errors.New("failed to generate P2PKH address: " + err.Error())
	}
	if err := a.generateP2WPKH(); err != nil {
		return errors.New("failed to generate P2WPKH address: " + err.Error())
	}
	if err := a.generateP2TR(); err != nil {
		return errors.New("failed to generate P2TR address: " + err.Error())
	}
	if err := a.generateEth(); err != nil {
		return errors.New("failed to generate Ethereum address: " + err.Error())
	}
	return nil
}

// generateP2PK generates the Pay-to-PubKey (P2PK) address.
func (a *Account) generateP2PK() error {
	address, err := btcutil.NewAddressPubKey(a.PubKey, a.params)
	if err != nil {
		return err
	}
	a.AddressP2PK = address.EncodeAddress()
	return nil
}

// generateP2PKH generates the Pay-to-PubKey-Hash (P2PKH) address.
func (a *Account) generateP2PKH() error {
	pubKeyHash := btcutil.Hash160(a.PubKey)
	address, err := btcutil.NewAddressPubKeyHash(pubKeyHash, a.params)
	if err != nil {
		return err
	}
	a.AddressP2PKH = address.EncodeAddress()
	return nil
}

// generateP2WPKH generates the Pay-to-Witness-PubKey-Hash (P2WPKH) address.
func (a *Account) generateP2WPKH() error {
	pubKeyHash := btcutil.Hash160(a.PubKey)
	address, err := btcutil.NewAddressWitnessPubKeyHash(pubKeyHash, a.params)
	if err != nil {
		return err
	}
	a.AddressP2WPKH = address.EncodeAddress()
	return nil
}

// generateP2TR generates the Pay-to-Taproot (P2TR) address.
func (a *Account) generateP2TR() error {
	internalKey, err := btcec.ParsePubKey(a.PubKey)
	if err != nil {
		return err
	}
	taprootKey := txscript.ComputeTaprootKeyNoScript(internalKey)
	address, err := btcutil.NewAddressTaproot(taprootKey.SerializeCompressed()[1:], a.params)
	if err != nil {
		return err
	}
	a.AddressP2TR = address.EncodeAddress()
	return nil
}

// generateEth generates the Ethereum address.
func (a *Account) generateEth() error {
	pubKey, err := btcec.ParsePubKey(a.PubKey)
	if err != nil {
		return err
	}

	// Calculate Keccak-256 hash of the uncompressed public key.
	hash := sha3.NewLegacyKeccak256()
	hash.Write(pubKey.SerializeUncompressed())
	addressBytes := hash.Sum(nil)

	// Ethereum address is the last 20 bytes of the hash.
	a.AddressEth = "0x" + hex.EncodeToString(addressBytes[12:])
	return nil
}
