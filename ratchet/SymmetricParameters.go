package ratchet

import (
	"timtyndale/libsignal-protocol-go/ecc"
	"timtyndale/libsignal-protocol-go/keys/identity"
)

// SymmetricParameters describes the session parameters for sessions where
// both users are online, which doesn't use prekeys for setup.
type SymmetricParameters struct {
	OurBaseKey         *ecc.ECKeyPair
	OurRatchetKey      *ecc.ECKeyPair
	OurIdentityKeyPair *identity.KeyPair

	TheirBaseKey     ecc.ECPublicKeyable
	TheirRatchetKey  ecc.ECPublicKeyable
	TheirIdentityKey *identity.Key
}
