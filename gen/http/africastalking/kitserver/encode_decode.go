// Code generated by goa v3.1.3, DO NOT EDIT.
//
// africastalking go-kit HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/wondenge/at-go/design

package server

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/wondenge/at-go/gen/http/africastalking/server"
	goahttp "goa.design/goa/v3/http"
)

// EncodeSendBulkSMSResponse returns a go-kit EncodeResponseFunc suitable for
// encoding africastalking SendBulkSMS responses.
func EncodeSendBulkSMSResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeSendBulkSMSResponse(encoder)
}

// DecodeSendBulkSMSRequest returns a go-kit DecodeRequestFunc suitable for
// decoding africastalking SendBulkSMS requests.
func DecodeSendBulkSMSRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeSendBulkSMSRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeSendPremiumSMSResponse returns a go-kit EncodeResponseFunc suitable
// for encoding africastalking SendPremiumSMS responses.
func EncodeSendPremiumSMSResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeSendPremiumSMSResponse(encoder)
}

// DecodeSendPremiumSMSRequest returns a go-kit DecodeRequestFunc suitable for
// decoding africastalking SendPremiumSMS requests.
func DecodeSendPremiumSMSRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeSendPremiumSMSRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeFetchSMSResponse returns a go-kit EncodeResponseFunc suitable for
// encoding africastalking FetchSMS responses.
func EncodeFetchSMSResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeFetchSMSResponse(encoder)
}

// DecodeFetchSMSRequest returns a go-kit DecodeRequestFunc suitable for
// decoding africastalking FetchSMS requests.
func DecodeFetchSMSRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeFetchSMSRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeNewCheckoutTokenResponse returns a go-kit EncodeResponseFunc suitable
// for encoding africastalking NewCheckoutToken responses.
func EncodeNewCheckoutTokenResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeNewCheckoutTokenResponse(encoder)
}

// DecodeNewCheckoutTokenRequest returns a go-kit DecodeRequestFunc suitable
// for decoding africastalking NewCheckoutToken requests.
func DecodeNewCheckoutTokenRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeNewCheckoutTokenRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeNewPremiumSubscriptionResponse returns a go-kit EncodeResponseFunc
// suitable for encoding africastalking NewPremiumSubscription responses.
func EncodeNewPremiumSubscriptionResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeNewPremiumSubscriptionResponse(encoder)
}

// DecodeNewPremiumSubscriptionRequest returns a go-kit DecodeRequestFunc
// suitable for decoding africastalking NewPremiumSubscription requests.
func DecodeNewPremiumSubscriptionRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeNewPremiumSubscriptionRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeFetchPremiumSubscriptionResponse returns a go-kit EncodeResponseFunc
// suitable for encoding africastalking FetchPremiumSubscription responses.
func EncodeFetchPremiumSubscriptionResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeFetchPremiumSubscriptionResponse(encoder)
}

// DecodeFetchPremiumSubscriptionRequest returns a go-kit DecodeRequestFunc
// suitable for decoding africastalking FetchPremiumSubscription requests.
func DecodeFetchPremiumSubscriptionRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeFetchPremiumSubscriptionRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodePurgePremiumSubscriptionResponse returns a go-kit EncodeResponseFunc
// suitable for encoding africastalking PurgePremiumSubscription responses.
func EncodePurgePremiumSubscriptionResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodePurgePremiumSubscriptionResponse(encoder)
}

// DecodePurgePremiumSubscriptionRequest returns a go-kit DecodeRequestFunc
// suitable for decoding africastalking PurgePremiumSubscription requests.
func DecodePurgePremiumSubscriptionRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodePurgePremiumSubscriptionRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeMakeCallResponse returns a go-kit EncodeResponseFunc suitable for
// encoding africastalking MakeCall responses.
func EncodeMakeCallResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeMakeCallResponse(encoder)
}

// DecodeMakeCallRequest returns a go-kit DecodeRequestFunc suitable for
// decoding africastalking MakeCall requests.
func DecodeMakeCallRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeMakeCallRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeTransferCallResponse returns a go-kit EncodeResponseFunc suitable for
// encoding africastalking TransferCall responses.
func EncodeTransferCallResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeTransferCallResponse(encoder)
}

// DecodeTransferCallRequest returns a go-kit DecodeRequestFunc suitable for
// decoding africastalking TransferCall requests.
func DecodeTransferCallRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeTransferCallRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeSayResponse returns a go-kit EncodeResponseFunc suitable for encoding
// africastalking Say responses.
func EncodeSayResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeSayResponse(encoder)
}

// DecodeSayRequest returns a go-kit DecodeRequestFunc suitable for decoding
// africastalking Say requests.
func DecodeSayRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeSayRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodePlayResponse returns a go-kit EncodeResponseFunc suitable for encoding
// africastalking Play responses.
func EncodePlayResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodePlayResponse(encoder)
}

// DecodePlayRequest returns a go-kit DecodeRequestFunc suitable for decoding
// africastalking Play requests.
func DecodePlayRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodePlayRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeGetDigitsResponse returns a go-kit EncodeResponseFunc suitable for
// encoding africastalking GetDigits responses.
func EncodeGetDigitsResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeGetDigitsResponse(encoder)
}

// DecodeGetDigitsRequest returns a go-kit DecodeRequestFunc suitable for
// decoding africastalking GetDigits requests.
func DecodeGetDigitsRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeGetDigitsRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeDialResponse returns a go-kit EncodeResponseFunc suitable for encoding
// africastalking Dial responses.
func EncodeDialResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeDialResponse(encoder)
}

// DecodeDialRequest returns a go-kit DecodeRequestFunc suitable for decoding
// africastalking Dial requests.
func DecodeDialRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeDialRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeRecordResponse returns a go-kit EncodeResponseFunc suitable for
// encoding africastalking Record responses.
func EncodeRecordResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeRecordResponse(encoder)
}

// DecodeRecordRequest returns a go-kit DecodeRequestFunc suitable for decoding
// africastalking Record requests.
func DecodeRecordRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeRecordRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeEnqueueResponse returns a go-kit EncodeResponseFunc suitable for
// encoding africastalking Enqueue responses.
func EncodeEnqueueResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeEnqueueResponse(encoder)
}

// DecodeEnqueueRequest returns a go-kit DecodeRequestFunc suitable for
// decoding africastalking Enqueue requests.
func DecodeEnqueueRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeEnqueueRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeDequeueResponse returns a go-kit EncodeResponseFunc suitable for
// encoding africastalking Dequeue responses.
func EncodeDequeueResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeDequeueResponse(encoder)
}

// DecodeDequeueRequest returns a go-kit DecodeRequestFunc suitable for
// decoding africastalking Dequeue requests.
func DecodeDequeueRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeDequeueRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeRedirectResponse returns a go-kit EncodeResponseFunc suitable for
// encoding africastalking Redirect responses.
func EncodeRedirectResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeRedirectResponse(encoder)
}

// DecodeRedirectRequest returns a go-kit DecodeRequestFunc suitable for
// decoding africastalking Redirect requests.
func DecodeRedirectRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeRedirectRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeRejectResponse returns a go-kit EncodeResponseFunc suitable for
// encoding africastalking Reject responses.
func EncodeRejectResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeRejectResponse(encoder)
}

// DecodeRejectRequest returns a go-kit DecodeRequestFunc suitable for decoding
// africastalking Reject requests.
func DecodeRejectRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeRejectRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeQueueResponse returns a go-kit EncodeResponseFunc suitable for
// encoding africastalking Queue responses.
func EncodeQueueResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeQueueResponse(encoder)
}

// DecodeQueueRequest returns a go-kit DecodeRequestFunc suitable for decoding
// africastalking Queue requests.
func DecodeQueueRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeQueueRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeUploadMediaResponse returns a go-kit EncodeResponseFunc suitable for
// encoding africastalking UploadMedia responses.
func EncodeUploadMediaResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeUploadMediaResponse(encoder)
}

// DecodeUploadMediaRequest returns a go-kit DecodeRequestFunc suitable for
// decoding africastalking UploadMedia requests.
func DecodeUploadMediaRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeUploadMediaRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeMobileCheckoutResponse returns a go-kit EncodeResponseFunc suitable
// for encoding africastalking MobileCheckout responses.
func EncodeMobileCheckoutResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeMobileCheckoutResponse(encoder)
}

// DecodeMobileCheckoutRequest returns a go-kit DecodeRequestFunc suitable for
// decoding africastalking MobileCheckout requests.
func DecodeMobileCheckoutRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeMobileCheckoutRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeMobileB2CResponse returns a go-kit EncodeResponseFunc suitable for
// encoding africastalking MobileB2C responses.
func EncodeMobileB2CResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeMobileB2CResponse(encoder)
}

// DecodeMobileB2CRequest returns a go-kit DecodeRequestFunc suitable for
// decoding africastalking MobileB2C requests.
func DecodeMobileB2CRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeMobileB2CRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeMobileB2BResponse returns a go-kit EncodeResponseFunc suitable for
// encoding africastalking MobileB2B responses.
func EncodeMobileB2BResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeMobileB2BResponse(encoder)
}

// DecodeMobileB2BRequest returns a go-kit DecodeRequestFunc suitable for
// decoding africastalking MobileB2B requests.
func DecodeMobileB2BRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeMobileB2BRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeBankCheckoutResponse returns a go-kit EncodeResponseFunc suitable for
// encoding africastalking Bank Checkout responses.
func EncodeBankCheckoutResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeBankCheckoutResponse(encoder)
}

// DecodeBankCheckoutRequest returns a go-kit DecodeRequestFunc suitable for
// decoding africastalking Bank Checkout requests.
func DecodeBankCheckoutRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeBankCheckoutRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeBankCheckoutValidateResponse returns a go-kit EncodeResponseFunc
// suitable for encoding africastalking BankCheckoutValidate responses.
func EncodeBankCheckoutValidateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeBankCheckoutValidateResponse(encoder)
}

// DecodeBankCheckoutValidateRequest returns a go-kit DecodeRequestFunc
// suitable for decoding africastalking BankCheckoutValidate requests.
func DecodeBankCheckoutValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeBankCheckoutValidateRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeBankTransferResponse returns a go-kit EncodeResponseFunc suitable for
// encoding africastalking BankTransfer responses.
func EncodeBankTransferResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeBankTransferResponse(encoder)
}

// DecodeBankTransferRequest returns a go-kit DecodeRequestFunc suitable for
// decoding africastalking BankTransfer requests.
func DecodeBankTransferRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeBankTransferRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeCardCheckoutResponse returns a go-kit EncodeResponseFunc suitable for
// encoding africastalking CardCheckout responses.
func EncodeCardCheckoutResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeCardCheckoutResponse(encoder)
}

// DecodeCardCheckoutRequest returns a go-kit DecodeRequestFunc suitable for
// decoding africastalking CardCheckout requests.
func DecodeCardCheckoutRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeCardCheckoutRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeCardCheckoutValidateResponse returns a go-kit EncodeResponseFunc
// suitable for encoding africastalking CardCheckoutValidate responses.
func EncodeCardCheckoutValidateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeCardCheckoutValidateResponse(encoder)
}

// DecodeCardCheckoutValidateRequest returns a go-kit DecodeRequestFunc
// suitable for decoding africastalking CardCheckoutValidate requests.
func DecodeCardCheckoutValidateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeCardCheckoutValidateRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeWalletTransferResponse returns a go-kit EncodeResponseFunc suitable
// for encoding africastalking WalletTransfer responses.
func EncodeWalletTransferResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeWalletTransferResponse(encoder)
}

// DecodeWalletTransferRequest returns a go-kit DecodeRequestFunc suitable for
// decoding africastalking WalletTransfer requests.
func DecodeWalletTransferRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeWalletTransferRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeTopupStashResponse returns a go-kit EncodeResponseFunc suitable for
// encoding africastalking TopupStash responses.
func EncodeTopupStashResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeTopupStashResponse(encoder)
}

// DecodeTopupStashRequest returns a go-kit DecodeRequestFunc suitable for
// decoding africastalking TopupStash requests.
func DecodeTopupStashRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeTopupStashRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeSendAirtimeResponse returns a go-kit EncodeResponseFunc suitable for
// encoding africastalking SendAirtime responses.
func EncodeSendAirtimeResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeSendAirtimeResponse(encoder)
}

// DecodeSendAirtimeRequest returns a go-kit DecodeRequestFunc suitable for
// decoding africastalking SendAirtime requests.
func DecodeSendAirtimeRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeSendAirtimeRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodePublishIoTResponse returns a go-kit EncodeResponseFunc suitable for
// encoding africastalking PublishIoT responses.
func EncodePublishIoTResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodePublishIoTResponse(encoder)
}

// DecodePublishIoTRequest returns a go-kit DecodeRequestFunc suitable for
// decoding africastalking PublishIoT requests.
func DecodePublishIoTRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodePublishIoTRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeInitiateAppDataResponse returns a go-kit EncodeResponseFunc suitable
// for encoding africastalking InitiateAppData responses.
func EncodeInitiateAppDataResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeInitiateAppDataResponse(encoder)
}

// DecodeInitiateAppDataRequest returns a go-kit DecodeRequestFunc suitable for
// decoding africastalking InitiateAppData requests.
func DecodeInitiateAppDataRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeInitiateAppDataRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}

// EncodeGenerateResponse returns a go-kit EncodeResponseFunc suitable for
// encoding africastalking Generate responses.
func EncodeGenerateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeGenerateResponse(encoder)
}

// DecodeGenerateRequest returns a go-kit DecodeRequestFunc suitable for
// decoding africastalking Generate requests.
func DecodeGenerateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) kithttp.DecodeRequestFunc {
	dec := server.DecodeGenerateRequest(mux, decoder)
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		r = r.WithContext(ctx)
		return dec(r)
	}
}
