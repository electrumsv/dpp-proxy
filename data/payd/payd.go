package payd

import (
	"context"
	"fmt"
	"net/http"

	"github.com/libsv/go-bt/v2"
	"github.com/pkg/errors"

	"github.com/libsv/pptcl"
	"github.com/libsv/pptcl/config"
	"github.com/libsv/pptcl/data"
	"github.com/libsv/pptcl/data/payd/models"
)

// Known endpoints for the payd wallet implementing the payment protocol interface.
const (
	urlPayments      = "%s/api/v1/payments?invoiceID=%s"
	urlOwner         = "%s/api/v1/owner"
	urlDestinations  = "%s/api/v1/destinations/:%s"
	urlFees          = "%s/api/v1/fees"
	protocolInsecure = "http"
	protocolSecure   = "https"
)

type payd struct {
	client data.HttpClient
	cfg    *config.PayD
}

// NewPayD will setup a new store that can interface with a payd wallet implementing
// the Payment Protocol Interface.
func NewPayD(cfg *config.PayD, client data.HttpClient) *payd {
	return &payd{
		cfg:    cfg,
		client: client,
	}
}

// PaymentCreate will post a request to payd to validate and add the txos to the wallet.
//
// If invalid a non 204 status code is returned.
func (p *payd) PaymentCreate(ctx context.Context, args pptcl.PaymentCreateArgs, req pptcl.PaymentCreate) error {
	paymentReq := models.PayDPaymentRequest{
		SPVEnvelope:    req.SPVEnvelope,
		ProofCallbacks: req.ProofCallbacks,
	}
	return p.client.Do(ctx, http.MethodPost, fmt.Sprintf(urlPayments, p.baseURL(), args.PaymentID), http.StatusNoContent, paymentReq, nil)
}

// Owner will return information regarding the owner of a payd wallet.
//
// In this example, the payd wallet has no auth, in proper implementations auth would
// be enabled and a cookie / oauth / bearer token etc would be passed down.
func (p *payd) Owner(ctx context.Context) (*pptcl.MerchantData, error) {
	var owner *pptcl.MerchantData
	if err := p.client.Do(ctx, http.MethodGet, fmt.Sprintf(urlOwner, p.baseURL()), http.StatusOK, nil, &owner); err != nil {
		return nil, errors.WithStack(err)
	}
	return owner, nil
}

// Outputs will return outputs for payment requests, the sender will then fulfil these outputs
// and send a tx for broadcast.
func (p *payd) Outputs(ctx context.Context, args pptcl.PaymentRequestArgs) ([]pptcl.Output, error) {
	var dd []models.Destination
	if err := p.client.Do(ctx, http.MethodGet, fmt.Sprintf(urlDestinations, p.baseURL(), args.PaymentID), http.StatusOK, nil, &dd); err != nil {
		return nil, errors.WithStack(err)
	}
	oo := make([]pptcl.Output, 0, len(dd))
	for _, d := range dd {
		oo = append(oo, pptcl.Output{
			Amount: d.Satoshis,
			Script: d.Script,
		})
	}
	return oo, nil
}

// Fees will return current fees that a payd wallet is using.
func (p *payd) Fees(ctx context.Context) (*bt.FeeQuote, error) {
	var fees *bt.FeeQuote
	if err := p.client.Do(ctx, http.MethodGet, fmt.Sprintf(urlFees, p.baseURL()), http.StatusOK, nil, &fees); err != nil {
		return nil, errors.WithStack(err)
	}
	return fees, nil
}

// baseURL will return http or https depending on if we're using TLS.
func (p *payd) baseURL() string {
	if p.cfg.Secure {
		return fmt.Sprintf("%s://%s%s", protocolSecure, p.cfg.Host, p.cfg.Port)
	}
	return fmt.Sprintf("%s://%s%s", protocolInsecure, p.cfg.Host, p.cfg.Port)
}