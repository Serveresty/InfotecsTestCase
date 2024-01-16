package cerr

import "errors"

var SenderWalletNotFound = errors.New("Sender's wallet not found")
var ReceiverWalletNotFound = errors.New("Receiver's wallet not found")
var LessMoney = errors.New("The transfer amount is more than available on the balance")
var WalletNotFound = errors.New("Wallet not found")
