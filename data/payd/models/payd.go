package models

import (
	"github.com/libsv/go-bc/spv"

	"github.com/libsv/pptcl"
)

// PayDPaymentRequest is used to send a payment to PayD for valdiation and storage.
type PayDPaymentRequest struct {
	SPVEnvelope    *spv.Envelope
	ProofCallbacks map[string]pptcl.ProofCallback `json:"proofCallbacks"`
}

// Destination is a payment output with locking script.
type Destination struct {
	Script   string `json:"script"`
	Satoshis uint64 `json:"satohsis"`
}