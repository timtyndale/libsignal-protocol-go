package store

import (
	"timtyndale/libsignal-protocol-go/groups/state/record"
	"timtyndale/libsignal-protocol-go/protocol"
)

type SenderKey interface {
	StoreSenderKey(senderKeyName *protocol.SenderKeyName, keyRecord *record.SenderKey)
	LoadSenderKey(senderKeyName *protocol.SenderKeyName) *record.SenderKey
}
