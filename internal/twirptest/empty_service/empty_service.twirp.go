// Code generated by protoc-gen-twirp v5.9.0, DO NOT EDIT.
// source: empty_service.proto

/*
Package empty_service is a generated twirp stub package.
This code was generated with github.com/twitchtv/twirp/protoc-gen-twirp v5.9.0.

It is generated from these files:
	empty_service.proto
*/
package empty_service

import bytes "bytes"
import context "context"
import fmt "fmt"
import ioutil "io/ioutil"
import http "net/http"
import strconv "strconv"

import jsonpb "github.com/golang/protobuf/jsonpb"
import proto "github.com/golang/protobuf/proto"
import twirp "github.com/twitchtv/twirp"
import ctxsetters "github.com/twitchtv/twirp/ctxsetters"

// Imports only used by utility functions:
import io "io"
import json "encoding/json"
import url "net/url"

// ===============
// Empty Interface
// ===============

type Empty interface {
}

// =====================
// Empty Protobuf Client
// =====================

type emptyProtobufClient struct {
	client HTTPClient
	urls   [0]string
	opts   twirp.ClientOptions
}

// NewEmptyProtobufClient creates a Protobuf client that implements the Empty interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewEmptyProtobufClient(addr string, client HTTPClient, opt ...twirp.ClientOption) Empty {
	var httpClient HTTPClient = client
	if c, ok := client.(*http.Client); ok {
		httpClient = withoutRedirects(c)
	}

	opts := twirp.ClientOptions{}
	for _, o := range opt {
		o(&opts)
	}

	urls := [0]string{}

	return &emptyProtobufClient{
		client: httpClient,
		urls:   urls,
		opts:   opts,
	}
}

// =================
// Empty JSON Client
// =================

type emptyJSONClient struct {
	client HTTPClient
	urls   [0]string
	opts   twirp.ClientOptions
}

// NewEmptyJSONClient creates a JSON client that implements the Empty interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewEmptyJSONClient(addr string, client HTTPClient, opt ...twirp.ClientOption) Empty {
	var httpClient HTTPClient = client
	if c, ok := client.(*http.Client); ok {
		httpClient = withoutRedirects(c)
	}

	opts := twirp.ClientOptions{}
	for _, o := range opt {
		o(&opts)
	}

	urls := [0]string{}

	return &emptyJSONClient{
		client: httpClient,
		urls:   urls,
		opts:   opts,
	}
}

// ====================
// Empty Server Handler
// ====================

type emptyServer struct {
	Empty
	hooks *twirp.ServerHooks
}

func NewEmptyServer(svc Empty, hooks *twirp.ServerHooks) TwirpServer {
	return &emptyServer{
		Empty: svc,
		hooks: hooks,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *emptyServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// EmptyPathPrefix is used for all URL paths on a twirp Empty server.
// Requests are always: POST EmptyPathPrefix/method
// It can be used in an HTTP mux to route twirp requests along with non-twirp requests on other routes.
const EmptyPathPrefix = "/twirp/twirp.internal.twirptest.emptyservice.Empty/"

func (s *emptyServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "twirp.internal.twirptest.emptyservice")
	ctx = ctxsetters.WithServiceName(ctx, "Empty")
	ctx = ctxsetters.WithResponseWriter(ctx, resp)

	var err error
	ctx, err = callRequestReceived(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	if req.Method != "POST" {
		msg := fmt.Sprintf("unsupported method %q (only POST is allowed)", req.Method)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}

	switch req.URL.Path {
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}
}

func (s *emptyServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor0, 0
}

func (s *emptyServer) ProtocGenTwirpVersion() string {
	return "v5.9.0"
}

func (s *emptyServer) PathPrefix() string {
	return EmptyPathPrefix
}

// =====
// Utils
// =====

// HTTPClient is the interface used by generated clients to send HTTP requests.
// It is fulfilled by *(net/http).Client, which is sufficient for most users.
// Users can provide their own implementation for special retry policies.
//
// HTTPClient implementations should not follow redirects. Redirects are
// automatically disabled if *(net/http).Client is passed to client
// constructors. See the withoutRedirects function in this file for more
// details.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// TwirpServer is the interface generated server structs will support: they're
// HTTP handlers with additional methods for accessing metadata about the
// service. Those accessors are a low-level API for building reflection tools.
// Most people can think of TwirpServers as just http.Handlers.
type TwirpServer interface {
	http.Handler
	// ServiceDescriptor returns gzipped bytes describing the .proto file that
	// this service was generated from. Once unzipped, the bytes can be
	// unmarshalled as a
	// github.com/golang/protobuf/protoc-gen-go/descriptor.FileDescriptorProto.
	//
	// The returned integer is the index of this particular service within that
	// FileDescriptorProto's 'Service' slice of ServiceDescriptorProtos. This is a
	// low-level field, expected to be used for reflection.
	ServiceDescriptor() ([]byte, int)
	// ProtocGenTwirpVersion is the semantic version string of the version of
	// twirp used to generate this file.
	ProtocGenTwirpVersion() string
	// PathPrefix returns the HTTP URL path prefix for all methods handled by this
	// service. This can be used with an HTTP mux to route twirp requests
	// alongside non-twirp requests on one HTTP listener.
	PathPrefix() string
}

// WriteError writes an HTTP response with a valid Twirp error format (code, msg, meta).
// Useful outside of the Twirp server (e.g. http middleware), but does not trigger hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func WriteError(resp http.ResponseWriter, err error) {
	writeError(context.Background(), resp, err, nil)
}

// writeError writes Twirp errors in the response and triggers hooks.
func writeError(ctx context.Context, resp http.ResponseWriter, err error, hooks *twirp.ServerHooks) {
	// Non-twirp errors are wrapped as Internal (default)
	twerr, ok := err.(twirp.Error)
	if !ok {
		twerr = twirp.InternalErrorWith(err)
	}

	statusCode := twirp.ServerHTTPStatusFromErrorCode(twerr.Code())
	ctx = ctxsetters.WithStatusCode(ctx, statusCode)
	ctx = callError(ctx, hooks, twerr)

	respBody := marshalErrorToJSON(twerr)

	resp.Header().Set("Content-Type", "application/json") // Error responses are always JSON
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBody)))
	resp.WriteHeader(statusCode) // set HTTP status code and send response

	_, writeErr := resp.Write(respBody)
	if writeErr != nil {
		// We have three options here. We could log the error, call the Error
		// hook, or just silently ignore the error.
		//
		// Logging is unacceptable because we don't have a user-controlled
		// logger; writing out to stderr without permission is too rude.
		//
		// Calling the Error hook would confuse users: it would mean the Error
		// hook got called twice for one request, which is likely to lead to
		// duplicated log messages and metrics, no matter how well we document
		// the behavior.
		//
		// Silently ignoring the error is our least-bad option. It's highly
		// likely that the connection is broken and the original 'err' says
		// so anyway.
		_ = writeErr
	}

	callResponseSent(ctx, hooks)
}

// urlBase helps ensure that addr specifies a scheme. If it is unparsable
// as a URL, it returns addr unchanged.
func urlBase(addr string) string {
	// If the addr specifies a scheme, use it. If not, default to
	// http. If url.Parse fails on it, return it unchanged.
	url, err := url.Parse(addr)
	if err != nil {
		return addr
	}
	if url.Scheme == "" {
		url.Scheme = "http"
	}
	return url.String()
}

// getCustomHTTPReqHeaders retrieves a copy of any headers that are set in
// a context through the twirp.WithHTTPRequestHeaders function.
// If there are no headers set, or if they have the wrong type, nil is returned.
func getCustomHTTPReqHeaders(ctx context.Context) http.Header {
	header, ok := twirp.HTTPRequestHeaders(ctx)
	if !ok || header == nil {
		return nil
	}
	copied := make(http.Header)
	for k, vv := range header {
		if vv == nil {
			copied[k] = nil
			continue
		}
		copied[k] = make([]string, len(vv))
		copy(copied[k], vv)
	}
	return copied
}

// newRequest makes an http.Request from a client, adding common headers.
func newRequest(ctx context.Context, url string, reqBody io.Reader, contentType string) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if customHeader := getCustomHTTPReqHeaders(ctx); customHeader != nil {
		req.Header = customHeader
	}
	req.Header.Set("Accept", contentType)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Twirp-Version", "v5.9.0")
	return req, nil
}

// JSON serialization for errors
type twerrJSON struct {
	Code string            `json:"code"`
	Msg  string            `json:"msg"`
	Meta map[string]string `json:"meta,omitempty"`
}

// marshalErrorToJSON returns JSON from a twirp.Error, that can be used as HTTP error response body.
// If serialization fails, it will use a descriptive Internal error instead.
func marshalErrorToJSON(twerr twirp.Error) []byte {
	// make sure that msg is not too large
	msg := twerr.Msg()
	if len(msg) > 1e6 {
		msg = msg[:1e6]
	}

	tj := twerrJSON{
		Code: string(twerr.Code()),
		Msg:  msg,
		Meta: twerr.MetaMap(),
	}

	buf, err := json.Marshal(&tj)
	if err != nil {
		buf = []byte("{\"type\": \"" + twirp.Internal + "\", \"msg\": \"There was an error but it could not be serialized into JSON\"}") // fallback
	}

	return buf
}

// errorFromResponse builds a twirp.Error from a non-200 HTTP response.
// If the response has a valid serialized Twirp error, then it's returned.
// If not, the response status code is used to generate a similar twirp
// error. See twirpErrorFromIntermediary for more info on intermediary errors.
func errorFromResponse(resp *http.Response) twirp.Error {
	statusCode := resp.StatusCode
	statusText := http.StatusText(statusCode)

	if isHTTPRedirect(statusCode) {
		// Unexpected redirect: it must be an error from an intermediary.
		// Twirp clients don't follow redirects automatically, Twirp only handles
		// POST requests, redirects should only happen on GET and HEAD requests.
		location := resp.Header.Get("Location")
		msg := fmt.Sprintf("unexpected HTTP status code %d %q received, Location=%q", statusCode, statusText, location)
		return twirpErrorFromIntermediary(statusCode, msg, location)
	}

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return wrapInternal(err, "failed to read server error response body")
	}

	var tj twerrJSON
	dec := json.NewDecoder(bytes.NewReader(respBodyBytes))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&tj); err != nil || tj.Code == "" {
		// Invalid JSON response; it must be an error from an intermediary.
		msg := fmt.Sprintf("Error from intermediary with HTTP status code %d %q", statusCode, statusText)
		return twirpErrorFromIntermediary(statusCode, msg, string(respBodyBytes))
	}

	errorCode := twirp.ErrorCode(tj.Code)
	if !twirp.IsValidErrorCode(errorCode) {
		msg := "invalid type returned from server error response: " + tj.Code
		return twirp.InternalError(msg)
	}

	twerr := twirp.NewError(errorCode, tj.Msg)
	for k, v := range tj.Meta {
		twerr = twerr.WithMeta(k, v)
	}
	return twerr
}

// twirpErrorFromIntermediary maps HTTP errors from non-twirp sources to twirp errors.
// The mapping is similar to gRPC: https://github.com/grpc/grpc/blob/master/doc/http-grpc-status-mapping.md.
// Returned twirp Errors have some additional metadata for inspection.
func twirpErrorFromIntermediary(status int, msg string, bodyOrLocation string) twirp.Error {
	var code twirp.ErrorCode
	if isHTTPRedirect(status) { // 3xx
		code = twirp.Internal
	} else {
		switch status {
		case 400: // Bad Request
			code = twirp.Internal
		case 401: // Unauthorized
			code = twirp.Unauthenticated
		case 403: // Forbidden
			code = twirp.PermissionDenied
		case 404: // Not Found
			code = twirp.BadRoute
		case 429, 502, 503, 504: // Too Many Requests, Bad Gateway, Service Unavailable, Gateway Timeout
			code = twirp.Unavailable
		default: // All other codes
			code = twirp.Unknown
		}
	}

	twerr := twirp.NewError(code, msg)
	twerr = twerr.WithMeta("http_error_from_intermediary", "true") // to easily know if this error was from intermediary
	twerr = twerr.WithMeta("status_code", strconv.Itoa(status))
	if isHTTPRedirect(status) {
		twerr = twerr.WithMeta("location", bodyOrLocation)
	} else {
		twerr = twerr.WithMeta("body", bodyOrLocation)
	}
	return twerr
}

func isHTTPRedirect(status int) bool {
	return status >= 300 && status <= 399
}

// wrapInternal wraps an error with a prefix as an Internal error.
// The original error cause is accessible by github.com/pkg/errors.Cause.
func wrapInternal(err error, prefix string) twirp.Error {
	return twirp.InternalErrorWith(&wrappedError{prefix: prefix, cause: err})
}

type wrappedError struct {
	prefix string
	cause  error
}

func (e *wrappedError) Cause() error  { return e.cause }
func (e *wrappedError) Error() string { return e.prefix + ": " + e.cause.Error() }

// ensurePanicResponses makes sure that rpc methods causing a panic still result in a Twirp Internal
// error response (status 500), and error hooks are properly called with the panic wrapped as an error.
// The panic is re-raised so it can be handled normally with middleware.
func ensurePanicResponses(ctx context.Context, resp http.ResponseWriter, hooks *twirp.ServerHooks) {
	if r := recover(); r != nil {
		// Wrap the panic as an error so it can be passed to error hooks.
		// The original error is accessible from error hooks, but not visible in the response.
		err := errFromPanic(r)
		twerr := &internalWithCause{msg: "Internal service panic", cause: err}
		// Actually write the error
		writeError(ctx, resp, twerr, hooks)
		// If possible, flush the error to the wire.
		f, ok := resp.(http.Flusher)
		if ok {
			f.Flush()
		}

		panic(r)
	}
}

// errFromPanic returns the typed error if the recovered panic is an error, otherwise formats as error.
func errFromPanic(p interface{}) error {
	if err, ok := p.(error); ok {
		return err
	}
	return fmt.Errorf("panic: %v", p)
}

// internalWithCause is a Twirp Internal error wrapping an original error cause, accessible
// by github.com/pkg/errors.Cause, but the original error message is not exposed on Msg().
type internalWithCause struct {
	msg   string
	cause error
}

func (e *internalWithCause) Cause() error                                { return e.cause }
func (e *internalWithCause) Error() string                               { return e.msg + ": " + e.cause.Error() }
func (e *internalWithCause) Code() twirp.ErrorCode                       { return twirp.Internal }
func (e *internalWithCause) Msg() string                                 { return e.msg }
func (e *internalWithCause) Meta(key string) string                      { return "" }
func (e *internalWithCause) MetaMap() map[string]string                  { return nil }
func (e *internalWithCause) WithMeta(key string, val string) twirp.Error { return e }

// malformedRequestError is used when the twirp server cannot unmarshal a request
func malformedRequestError(msg string) twirp.Error {
	return twirp.NewError(twirp.Malformed, msg)
}

// badRouteError is used when the twirp server cannot route a request
func badRouteError(msg string, method, url string) twirp.Error {
	err := twirp.NewError(twirp.BadRoute, msg)
	err = err.WithMeta("twirp_invalid_route", method+" "+url)
	return err
}

// withoutRedirects makes sure that the POST request can not be redirected.
// The standard library will, by default, redirect requests (including POSTs) if it gets a 302 or
// 303 response, and also 301s in go1.8. It redirects by making a second request, changing the
// method to GET and removing the body. This produces very confusing error messages, so instead we
// set a redirect policy that always errors. This stops Go from executing the redirect.
//
// We have to be a little careful in case the user-provided http.Client has its own CheckRedirect
// policy - if so, we'll run through that policy first.
//
// Because this requires modifying the http.Client, we make a new copy of the client and return it.
func withoutRedirects(in *http.Client) *http.Client {
	copy := *in
	copy.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if in.CheckRedirect != nil {
			// Run the input's redirect if it exists, in case it has side effects, but ignore any error it
			// returns, since we want to use ErrUseLastResponse.
			err := in.CheckRedirect(req, via)
			_ = err // Silly, but this makes sure generated code passes errcheck -blank, which some people use.
		}
		return http.ErrUseLastResponse
	}
	return &copy
}

// doProtobufRequest makes a Protobuf request to the remote Twirp service.
func doProtobufRequest(ctx context.Context, client HTTPClient, hooks *twirp.ClientHooks, url string, in, out proto.Message) (err error) {
	reqBodyBytes, err := proto.Marshal(in)
	if err != nil {
		return wrapInternal(err, "failed to marshal proto request")
	}
	reqBody := bytes.NewBuffer(reqBodyBytes)
	if err = ctx.Err(); err != nil {
		return wrapInternal(err, "aborted because context was done")
	}

	req, err := newRequest(ctx, url, reqBody, "application/protobuf")
	if err != nil {
		return wrapInternal(err, "could not build request")
	}
	ctx, err = callClientRequestPrepared(ctx, hooks, req)
	if err != nil {
		return wrapInternal(err, "failed to call RequestPrepared hook")
	}

	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return wrapInternal(err, "failed to do request")
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = wrapInternal(cerr, "failed to close response body")
		}
	}()

	if err = ctx.Err(); err != nil {
		return wrapInternal(err, "aborted because context was done")
	}

	if resp.StatusCode != 200 {
		return errorFromResponse(resp)
	}

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return wrapInternal(err, "failed to read response body")
	}
	if err = ctx.Err(); err != nil {
		return wrapInternal(err, "aborted because context was done")
	}

	if err = proto.Unmarshal(respBodyBytes, out); err != nil {
		return wrapInternal(err, "failed to unmarshal proto response")
	}
	return nil
}

// doJSONRequest makes a JSON request to the remote Twirp service.
func doJSONRequest(ctx context.Context, client HTTPClient, hooks *twirp.ClientHooks, url string, in, out proto.Message) (err error) {
	reqBody := bytes.NewBuffer(nil)
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(reqBody, in); err != nil {
		return wrapInternal(err, "failed to marshal json request")
	}
	if err = ctx.Err(); err != nil {
		return wrapInternal(err, "aborted because context was done")
	}

	req, err := newRequest(ctx, url, reqBody, "application/json")
	if err != nil {
		return wrapInternal(err, "could not build request")
	}
	ctx, err = callClientRequestPrepared(ctx, hooks, req)
	if err != nil {
		return wrapInternal(err, "failed to call RequestPrepared hook")
	}

	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return wrapInternal(err, "failed to do request")
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = wrapInternal(cerr, "failed to close response body")
		}
	}()

	if err = ctx.Err(); err != nil {
		return wrapInternal(err, "aborted because context was done")
	}

	if resp.StatusCode != 200 {
		return errorFromResponse(resp)
	}

	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(resp.Body, out); err != nil {
		return wrapInternal(err, "failed to unmarshal json response")
	}
	if err = ctx.Err(); err != nil {
		return wrapInternal(err, "aborted because context was done")
	}
	return nil
}

// Call twirp.ServerHooks.RequestReceived if the hook is available
func callRequestReceived(ctx context.Context, h *twirp.ServerHooks) (context.Context, error) {
	if h == nil || h.RequestReceived == nil {
		return ctx, nil
	}
	return h.RequestReceived(ctx)
}

// Call twirp.ServerHooks.RequestRouted if the hook is available
func callRequestRouted(ctx context.Context, h *twirp.ServerHooks) (context.Context, error) {
	if h == nil || h.RequestRouted == nil {
		return ctx, nil
	}
	return h.RequestRouted(ctx)
}

// Call twirp.ServerHooks.ResponsePrepared if the hook is available
func callResponsePrepared(ctx context.Context, h *twirp.ServerHooks) context.Context {
	if h == nil || h.ResponsePrepared == nil {
		return ctx
	}
	return h.ResponsePrepared(ctx)
}

// Call twirp.ServerHooks.ResponseSent if the hook is available
func callResponseSent(ctx context.Context, h *twirp.ServerHooks) {
	if h == nil || h.ResponseSent == nil {
		return
	}
	h.ResponseSent(ctx)
}

// Call twirp.ServerHooks.Error if the hook is available
func callError(ctx context.Context, h *twirp.ServerHooks, err twirp.Error) context.Context {
	if h == nil || h.Error == nil {
		return ctx
	}
	return h.Error(ctx, err)
}

func callClientResponseReceived(ctx context.Context, h *twirp.ClientHooks) {
	if h == nil || h.ResponseReceived == nil {
		return
	}
	h.ResponseReceived(ctx)
}

func callClientRequestPrepared(ctx context.Context, h *twirp.ClientHooks, req *http.Request) (context.Context, error) {
	if h == nil || h.RequestPrepared == nil {
		return ctx, nil
	}
	return h.RequestPrepared(ctx, req)
}

func callClientError(ctx context.Context, h *twirp.ClientHooks, err twirp.Error) {
	if h == nil || h.Error == nil {
		return
	}
	h.Error(ctx, err)
}

var twirpFileDescriptor0 = []byte{
	// 91 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xcd, 0x2d, 0x28,
	0xa9, 0x8c, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x52, 0x2d, 0x29, 0xcf, 0x2c, 0x2a, 0xd0, 0xcb, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1,
	0x03, 0x73, 0x4b, 0x52, 0x8b, 0x4b, 0xf4, 0xc0, 0xaa, 0xa1, 0x8a, 0x8d, 0xd8, 0xb9, 0x58, 0x5d,
	0x41, 0x7c, 0x27, 0xfe, 0x28, 0x5e, 0x14, 0x63, 0x92, 0xd8, 0xc0, 0xe6, 0x18, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x2a, 0xc1, 0xde, 0x92, 0x5e, 0x00, 0x00, 0x00,
}
