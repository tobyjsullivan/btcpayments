package btc

import (
	"testing"
	"io"
	"github.com/tobyjsullivan/btckeygenie/btckey"
	"errors"
)

func TestGenerateWallet_Success(t *testing.T) {
	expectedWif := "5J1F7GHadZG3sCCKHCwg8Jvys9xUbFsjLnGec4H125Ny1V9nR6V"
	expectedAddress := "1PMycacnJaSqwwJqjawXBErnLsZ7RkXUAs"

	// Override generate key method
	generateBtcKey = func(rand io.Reader) (btckey.PrivateKey, error) {
		priv := &btckey.PrivateKey{}
		priv.FromWIF(expectedWif)
		return *priv, nil
	}

	result, err := GenerateWallet()

	if err != nil {
		t.Errorf("Expected no error but got: %s", err.Error())
	}

	if wif := result.GetPrivateKeyWIF(); wif != expectedWif {
		t.Errorf("Wrong private key. Expected: %s; Actual: %s", expectedWif, wif)
	}

	if address := result.GetAddress(); address != expectedAddress {
		t.Errorf("Wrong address. Expected: %s; Actual: %s", expectedAddress, address)
	}
}

func TestGenerateWallet_Error(t *testing.T) {
	expectedError := errors.New("Dummy error")

	// Override generate key method to return an error
	generateBtcKey = func(rand io.Reader) (btckey.PrivateKey, error) {
		priv := &btckey.PrivateKey{}
		return *priv, expectedError
	}

	result, err := GenerateWallet()

	if err != expectedError {
		t.Errorf("Did not receive expected error. Expected: %v, Actual: %v", expectedError, err)
	}

	if result != nil {
		t.Errorf("Expected a nil result but received: %v", result)
	}
}
