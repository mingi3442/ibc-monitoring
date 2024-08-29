package utils

import (
	"encoding/base64"
	"fmt"

	tx "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/gogoproto/proto"
)

func DecodeTxData(txData []byte) (string, error) {

	var protoTx tx.Tx
	err := proto.Unmarshal(txData, &protoTx)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal Tx data: %w", err)
	}

	decodedTx := "Transaction:\n"

	decodedTx += "  Body:\n"
	decodedTx += "    Memo: " + protoTx.Body.Memo + "\n"
	decodedTx += "    Messages:\n"
	for i, msg := range protoTx.Body.Messages {
		decodedTx += fmt.Sprintf("      Message %d: %s\n", i+1, msg.TypeUrl)

	}

	decodedTx += "  AuthInfo:\n"
	decodedTx += "    Fee:\n"
	for _, amount := range protoTx.AuthInfo.Fee.Amount {
		decodedTx += "      Amount: " + amount.Amount.String() + " " + amount.Denom + "\n"

	}
	decodedTx += fmt.Sprintf("      Gas Limit: %d\n", protoTx.AuthInfo.Fee.GasLimit)
	decodedTx += "    SignerInfos:\n"
	for i, signerInfo := range protoTx.AuthInfo.SignerInfos {
		decodedTx += fmt.Sprintf("      Signer %d:\n", i+1)
		decodedTx += "        Public Key: " + base64.StdEncoding.EncodeToString(signerInfo.PublicKey.GetValue()) + "\n"
		decodedTx += fmt.Sprintf("        Sequence: %d\n", signerInfo.Sequence)
	}

	decodedTx += "  Signatures:\n"
	for i, signature := range protoTx.Signatures {
		decodedTx += "    Signature " + fmt.Sprintf("%d: ", i+1) + base64.StdEncoding.EncodeToString(signature) + "\n"
	}

	return decodedTx, nil
}
