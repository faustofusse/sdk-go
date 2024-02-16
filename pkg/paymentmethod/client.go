package paymentmethod

import (
	"context"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/baseclient"
)

const url = "https://api.mercadopago.com/v1/payment_methods"

// Client contains the methods to interact with the Payment Methods API.
type Client interface {
	// List lists all payment methods.
	// It is a get request to the endpoint: https://api.mercadopago.com/v1/payment_methods
	// Reference: https://www.mercadopago.com/developers/en/reference/payment_methods/_payment_methods/get/
	List(ctx context.Context) ([]Response, error)
}

// client is the implementation of Client.
type client struct {
	cfg *config.Config
}

// NewClient returns a new Payment Methods API Client.
func NewClient(cfg *config.Config) Client {
	return &client{cfg}
}

func (c *client) List(ctx context.Context) ([]Response, error) {
	res, err := baseclient.Get[[]Response](ctx, c.cfg, url)
	if err != nil {
		return nil, err
	}

	return res, nil
}
