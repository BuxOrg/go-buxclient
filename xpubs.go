package buxclient

import (
	"context"

	"github.com/BuxOrg/bux"
)

// NewXpub registers a new xpub - admin key needed
func (b *BuxClient) NewXpub(ctx context.Context, rawXPub string, metadata *bux.Metadata) error {
	return b.transport.NewXpub(ctx, rawXPub, metadata)
}

// GetXPub gets the current xpub
func (b *BuxClient) GetXPub(ctx context.Context) (*bux.Xpub, error) {
	return b.transport.GetXPub(ctx)
}

// UpdateXPubMetadata update the metadata of the logged in xpub
func (b *BuxClient) UpdateXPubMetadata(ctx context.Context, metadata *bux.Metadata) (*bux.Xpub, error) {
	return b.transport.UpdateXPubMetadata(ctx, metadata)
}
