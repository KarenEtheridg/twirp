// Code generated by protoc-gen-twirp v5.12.1, DO NOT EDIT.
// source: service_method_same_name.proto

/*
Package service_method_same_name is a generated twirp stub package.
This code was generated with github.com/twitchtv/twirp/protoc-gen-twirp v5.12.1.

It is generated from these files:
	service_method_same_name.proto
*/
package service_method_same_name

import bytes "bytes"
import strings "strings"
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

// ==============
// Echo Interface
// ==============

type Echo interface {
	Echo(context.Context, *Msg) (*Msg, error)
}

// ====================
// Echo Protobuf Client
// ====================

type echoProtobufClient struct {
	client HTTPClient
	urls   [1]string
	opts   twirp.ClientOptions
}

// NewEchoProtobufClient creates a Protobuf client that implements the Echo interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewEchoProtobufClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) Echo {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// <baseURL>[/twirp]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	if !clientOpts.SkipPathPrefix {
		serviceURL += "/twirp"
	}
	serviceURL += "/Echo"
	urls := [1]string{
		serviceURL + "/Echo",
	}

	return &echoProtobufClient{
		client: client,
		urls:   urls,
		opts:   clientOpts,
	}
}

func (c *echoProtobufClient) Echo(ctx context.Context, in *Msg) (*Msg, error) {
	ctx = ctxsetters.WithPackageName(ctx, "")
	ctx = ctxsetters.WithServiceName(ctx, "Echo")
	ctx = ctxsetters.WithMethodName(ctx, "Echo")
	out := new(Msg)
	ctx, err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// ================
// Echo JSON Client
// ================

type echoJSONClient struct {
	client HTTPClient
	urls   [1]string
	opts   twirp.ClientOptions
}

// NewEchoJSONClient creates a JSON client that implements the Echo interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewEchoJSONClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) Echo {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// <baseURL>[/twirp]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	if !clientOpts.SkipPathPrefix {
		serviceURL += "/twirp"
	}
	serviceURL += "/Echo"
	urls := [1]string{
		serviceURL + "/Echo",
	}

	return &echoJSONClient{
		client: client,
		urls:   urls,
		opts:   clientOpts,
	}
}

func (c *echoJSONClient) Echo(ctx context.Context, in *Msg) (*Msg, error) {
	ctx = ctxsetters.WithPackageName(ctx, "")
	ctx = ctxsetters.WithServiceName(ctx, "Echo")
	ctx = ctxsetters.WithMethodName(ctx, "Echo")
	out := new(Msg)
	ctx, err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// ===================
// Echo Server Handler
// ===================

type echoServer struct {
	Echo
	hooks *twirp.ServerHooks
}

func NewEchoServer(svc Echo, hooks *twirp.ServerHooks) TwirpServer {
	return &echoServer{
		Echo:  svc,
		hooks: hooks,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *echoServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// EchoPathPrefix is a convenience constant that could used to identify URL paths.
// Should be used with caution, it only matches routes generated by Twirp Go clients,
// that add a "/twirp" prefix by default, and use CamelCase service and method names.
// More info: https://twitchtv.github.io/twirp/docs/routing.html
const EchoPathPrefix = "/twirp/Echo/"

func (s *echoServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "")
	ctx = ctxsetters.WithServiceName(ctx, "Echo")
	ctx = ctxsetters.WithResponseWriter(ctx, resp)

	var err error
	ctx, err = callRequestReceived(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	if req.Method != "POST" {
		msg := fmt.Sprintf("unsupported method %q (only POST is allowed)", req.Method)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}

	// Verify path format: [<prefix>]/<package>.<Service>/<Method>
	parts := strings.Split(req.URL.Path, "/")
	if len(parts) < 2 || parts[len(parts)-2] != "Echo" {
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}

	// Route by <Method>
	switch parts[len(parts)-1] {
	case "Echo":
		s.serveEcho(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
}

func (s *echoServer) serveEcho(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveEchoJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveEchoProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *echoServer) serveEchoJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Echo")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(Msg)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *Msg
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.Echo.Echo(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Msg and nil error while calling Echo. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	respBytes := buf.Bytes()
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *echoServer) serveEchoProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Echo")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to read request body"))
		return
	}
	reqContent := new(Msg)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *Msg
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.Echo.Echo(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Msg and nil error while calling Echo. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *echoServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor0, 0
}

func (s *echoServer) ProtocGenTwirpVersion() string {
	return "v5.12.1"
}

func (s *echoServer) PathPrefix() string {
	return EchoPathPrefix
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

// sanitizeBaseURL parses the the baseURL, and adds the "http" scheme if needed.
// If the URL is unparsable, the baseURL is returned unchaged.
func sanitizeBaseURL(baseURL string) string {
	u, err := url.Parse(baseURL)
	if err != nil {
		return baseURL // invalid URL will fail later when making requests
	}
	if u.Scheme == "" {
		u.Scheme = "http"
	}
	return u.String()
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
	req.Header.Set("Twirp-Version", "v5.12.1")
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
		return twirp.InternalError(msg).WithMeta("body", string(respBodyBytes))
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

func (e *wrappedError) Error() string { return e.prefix + ": " + e.cause.Error() }
func (e *wrappedError) Unwrap() error { return e.cause } // for go1.13 + errors.Is/As
func (e *wrappedError) Cause() error  { return e.cause } // for github.com/pkg/errors

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

// internalWithCause is a Twirp Internal error wrapping an original error cause,
// but the original error message is not exposed on Msg(). The original error
// can be checked with go1.13+ errors.Is/As, and also by (github.com/pkg/errors).Unwrap
type internalWithCause struct {
	msg   string
	cause error
}

func (e *internalWithCause) Unwrap() error                               { return e.cause } // for go1.13 + errors.Is/As
func (e *internalWithCause) Cause() error                                { return e.cause } // for github.com/pkg/errors
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
func doProtobufRequest(ctx context.Context, client HTTPClient, hooks *twirp.ClientHooks, url string, in, out proto.Message) (_ context.Context, err error) {
	reqBodyBytes, err := proto.Marshal(in)
	if err != nil {
		return ctx, wrapInternal(err, "failed to marshal proto request")
	}
	reqBody := bytes.NewBuffer(reqBodyBytes)
	if err = ctx.Err(); err != nil {
		return ctx, wrapInternal(err, "aborted because context was done")
	}

	req, err := newRequest(ctx, url, reqBody, "application/protobuf")
	if err != nil {
		return ctx, wrapInternal(err, "could not build request")
	}
	ctx, err = callClientRequestPrepared(ctx, hooks, req)
	if err != nil {
		return ctx, err
	}

	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return ctx, wrapInternal(err, "failed to do request")
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = wrapInternal(cerr, "failed to close response body")
		}
	}()

	if err = ctx.Err(); err != nil {
		return ctx, wrapInternal(err, "aborted because context was done")
	}

	if resp.StatusCode != 200 {
		return ctx, errorFromResponse(resp)
	}

	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ctx, wrapInternal(err, "failed to read response body")
	}
	if err = ctx.Err(); err != nil {
		return ctx, wrapInternal(err, "aborted because context was done")
	}

	if err = proto.Unmarshal(respBodyBytes, out); err != nil {
		return ctx, wrapInternal(err, "failed to unmarshal proto response")
	}
	return ctx, nil
}

// doJSONRequest makes a JSON request to the remote Twirp service.
func doJSONRequest(ctx context.Context, client HTTPClient, hooks *twirp.ClientHooks, url string, in, out proto.Message) (_ context.Context, err error) {
	reqBody := bytes.NewBuffer(nil)
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(reqBody, in); err != nil {
		return ctx, wrapInternal(err, "failed to marshal json request")
	}
	if err = ctx.Err(); err != nil {
		return ctx, wrapInternal(err, "aborted because context was done")
	}

	req, err := newRequest(ctx, url, reqBody, "application/json")
	if err != nil {
		return ctx, wrapInternal(err, "could not build request")
	}
	ctx, err = callClientRequestPrepared(ctx, hooks, req)
	if err != nil {
		return ctx, err
	}

	req = req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		return ctx, wrapInternal(err, "failed to do request")
	}

	defer func() {
		cerr := resp.Body.Close()
		if err == nil && cerr != nil {
			err = wrapInternal(cerr, "failed to close response body")
		}
	}()

	if err = ctx.Err(); err != nil {
		return ctx, wrapInternal(err, "aborted because context was done")
	}

	if resp.StatusCode != 200 {
		return ctx, errorFromResponse(resp)
	}

	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(resp.Body, out); err != nil {
		return ctx, wrapInternal(err, "failed to unmarshal json response")
	}
	if err = ctx.Err(); err != nil {
		return ctx, wrapInternal(err, "aborted because context was done")
	}
	return ctx, nil
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
	// 96 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x8d, 0xcf, 0x4d, 0x2d, 0xc9, 0xc8, 0x4f, 0x89, 0x2f, 0x4e, 0xcc, 0x4d, 0x8d,
	0xcf, 0x4b, 0xcc, 0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x62, 0xe5, 0x62, 0xf6, 0x2d,
	0x4e, 0x37, 0x92, 0xe1, 0x62, 0x71, 0x4d, 0xce, 0xc8, 0x17, 0x12, 0x81, 0xd2, 0x2c, 0x7a, 0xbe,
	0xc5, 0xe9, 0x52, 0x60, 0x52, 0x89, 0xc1, 0x49, 0x2a, 0x4a, 0x02, 0x97, 0x31, 0x49, 0x6c, 0x60,
	0x73, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4d, 0xcf, 0xa1, 0x77, 0x69, 0x00, 0x00, 0x00,
}
