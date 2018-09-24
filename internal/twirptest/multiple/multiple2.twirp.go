// Code generated by protoc-gen-twirp v5.5.0, DO NOT EDIT.
// source: multiple2.proto

package multiple

import bytes "bytes"
import strings "strings"
import context "context"
import fmt "fmt"
import ioutil "io/ioutil"
import http "net/http"

import jsonpb "github.com/golang/protobuf/jsonpb"
import proto "github.com/golang/protobuf/proto"
import twirp "github.com/twitchtv/twirp"
import ctxsetters "github.com/twitchtv/twirp/ctxsetters"

// ==============
// Svc2 Interface
// ==============

type Svc2 interface {
	Send(context.Context, *Msg2) (*Msg2, error)

	SamePackageProtoImport(context.Context, *Msg1) (*Msg1, error)
}

// ====================
// Svc2 Protobuf Client
// ====================

type svc2ProtobufClient struct {
	client HTTPClient
	urls   [2]string
}

// NewSvc2ProtobufClient creates a Protobuf client that implements the Svc2 interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewSvc2ProtobufClient(addr string, client HTTPClient) Svc2 {
	prefix := urlBase(addr) + Svc2PathPrefix
	urls := [2]string{
		prefix + "Send",
		prefix + "SamePackageProtoImport",
	}
	if httpClient, ok := client.(*http.Client); ok {
		return &svc2ProtobufClient{
			client: withoutRedirects(httpClient),
			urls:   urls,
		}
	}
	return &svc2ProtobufClient{
		client: client,
		urls:   urls,
	}
}

func (c *svc2ProtobufClient) Send(ctx context.Context, in *Msg2) (*Msg2, error) {
	ctx = ctxsetters.WithPackageName(ctx, "twirp.internal.twirptest.multiple")
	ctx = ctxsetters.WithServiceName(ctx, "Svc2")
	ctx = ctxsetters.WithMethodName(ctx, "Send")
	out := new(Msg2)
	err := doProtobufRequest(ctx, c.client, c.urls[0], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *svc2ProtobufClient) SamePackageProtoImport(ctx context.Context, in *Msg1) (*Msg1, error) {
	ctx = ctxsetters.WithPackageName(ctx, "twirp.internal.twirptest.multiple")
	ctx = ctxsetters.WithServiceName(ctx, "Svc2")
	ctx = ctxsetters.WithMethodName(ctx, "SamePackageProtoImport")
	out := new(Msg1)
	err := doProtobufRequest(ctx, c.client, c.urls[1], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ================
// Svc2 JSON Client
// ================

type svc2JSONClient struct {
	client HTTPClient
	urls   [2]string
}

// NewSvc2JSONClient creates a JSON client that implements the Svc2 interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewSvc2JSONClient(addr string, client HTTPClient) Svc2 {
	prefix := urlBase(addr) + Svc2PathPrefix
	urls := [2]string{
		prefix + "Send",
		prefix + "SamePackageProtoImport",
	}
	if httpClient, ok := client.(*http.Client); ok {
		return &svc2JSONClient{
			client: withoutRedirects(httpClient),
			urls:   urls,
		}
	}
	return &svc2JSONClient{
		client: client,
		urls:   urls,
	}
}

func (c *svc2JSONClient) Send(ctx context.Context, in *Msg2) (*Msg2, error) {
	ctx = ctxsetters.WithPackageName(ctx, "twirp.internal.twirptest.multiple")
	ctx = ctxsetters.WithServiceName(ctx, "Svc2")
	ctx = ctxsetters.WithMethodName(ctx, "Send")
	out := new(Msg2)
	err := doJSONRequest(ctx, c.client, c.urls[0], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *svc2JSONClient) SamePackageProtoImport(ctx context.Context, in *Msg1) (*Msg1, error) {
	ctx = ctxsetters.WithPackageName(ctx, "twirp.internal.twirptest.multiple")
	ctx = ctxsetters.WithServiceName(ctx, "Svc2")
	ctx = ctxsetters.WithMethodName(ctx, "SamePackageProtoImport")
	out := new(Msg1)
	err := doJSONRequest(ctx, c.client, c.urls[1], in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ===================
// Svc2 Server Handler
// ===================

type svc2Server struct {
	Svc2
	hooks *twirp.ServerHooks
}

func NewSvc2Server(svc Svc2, hooks *twirp.ServerHooks) TwirpServer {
	return &svc2Server{
		Svc2:  svc,
		hooks: hooks,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *svc2Server) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// Svc2PathPrefix is used for all URL paths on a twirp Svc2 server.
// Requests are always: POST Svc2PathPrefix/method
// It can be used in an HTTP mux to route twirp requests along with non-twirp requests on other routes.
const Svc2PathPrefix = "/twirp/twirp.internal.twirptest.multiple.Svc2/"

func (s *svc2Server) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "twirp.internal.twirptest.multiple")
	ctx = ctxsetters.WithServiceName(ctx, "Svc2")
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
	case "/twirp/twirp.internal.twirptest.multiple.Svc2/Send":
		s.serveSend(ctx, resp, req)
		return
	case "/twirp/twirp.internal.twirptest.multiple.Svc2/SamePackageProtoImport":
		s.serveSamePackageProtoImport(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}
}

func (s *svc2Server) serveSend(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveSendJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveSendProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *svc2Server) serveSendJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Send")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(Msg2)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		err = wrapErr(err, "failed to parse request json")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	// Call service method
	var respContent *Msg2
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if r := recover(); r != nil {
				s.writeError(ctx, resp, twirp.InternalError("Internal service panic"))
				panic(r)
			}
		}()
		respContent, err = s.Svc2.Send(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Msg2 and nil error while calling Send. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		err = wrapErr(err, "failed to marshal json response")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)

	respBytes := buf.Bytes()
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *svc2Server) serveSendProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Send")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		err = wrapErr(err, "failed to read request body")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}
	reqContent := new(Msg2)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		err = wrapErr(err, "failed to parse request proto")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	// Call service method
	var respContent *Msg2
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if r := recover(); r != nil {
				s.writeError(ctx, resp, twirp.InternalError("Internal service panic"))
				panic(r)
			}
		}()
		respContent, err = s.Svc2.Send(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Msg2 and nil error while calling Send. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		err = wrapErr(err, "failed to marshal proto response")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *svc2Server) serveSamePackageProtoImport(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveSamePackageProtoImportJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveSamePackageProtoImportProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *svc2Server) serveSamePackageProtoImportJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "SamePackageProtoImport")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(Msg1)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		err = wrapErr(err, "failed to parse request json")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	// Call service method
	var respContent *Msg1
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if r := recover(); r != nil {
				s.writeError(ctx, resp, twirp.InternalError("Internal service panic"))
				panic(r)
			}
		}()
		respContent, err = s.Svc2.SamePackageProtoImport(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Msg1 and nil error while calling SamePackageProtoImport. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		err = wrapErr(err, "failed to marshal json response")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)

	respBytes := buf.Bytes()
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *svc2Server) serveSamePackageProtoImportProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "SamePackageProtoImport")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		err = wrapErr(err, "failed to read request body")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}
	reqContent := new(Msg1)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		err = wrapErr(err, "failed to parse request proto")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	// Call service method
	var respContent *Msg1
	func() {
		defer func() {
			// In case of a panic, serve a 500 error and then panic.
			if r := recover(); r != nil {
				s.writeError(ctx, resp, twirp.InternalError("Internal service panic"))
				panic(r)
			}
		}()
		respContent, err = s.Svc2.SamePackageProtoImport(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Msg1 and nil error while calling SamePackageProtoImport. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		err = wrapErr(err, "failed to marshal proto response")
		s.writeError(ctx, resp, twirp.InternalErrorWith(err))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *svc2Server) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor1, 0
}

func (s *svc2Server) ProtocGenTwirpVersion() string {
	return "v5.5.0"
}

var twirpFileDescriptor1 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x2d, 0xcd, 0x29,
	0xc9, 0x2c, 0xc8, 0x49, 0x35, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x2c, 0x29, 0xcf,
	0x2c, 0x2a, 0xd0, 0xcb, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x03, 0x73, 0x4b, 0x52,
	0x8b, 0x4b, 0xf4, 0x60, 0x2a, 0xa5, 0xe0, 0x7a, 0x0c, 0x21, 0x7a, 0x94, 0xd8, 0xb8, 0x58, 0x7c,
	0x8b, 0xd3, 0x8d, 0x8c, 0xce, 0x30, 0x72, 0xb1, 0x04, 0x97, 0x25, 0x1b, 0x09, 0x45, 0x70, 0xb1,
	0x04, 0xa7, 0xe6, 0xa5, 0x08, 0xa9, 0xeb, 0x11, 0x34, 0x4d, 0x0f, 0xa4, 0x53, 0x8a, 0x58, 0x85,
	0x42, 0x59, 0x5c, 0x62, 0xc1, 0x89, 0xb9, 0xa9, 0x01, 0x89, 0xc9, 0xd9, 0x89, 0xe9, 0xa9, 0x01,
	0x20, 0xeb, 0x3d, 0x73, 0x0b, 0xf2, 0x8b, 0x4a, 0x88, 0xb5, 0xcb, 0x90, 0x58, 0xbb, 0x0c, 0x9d,
	0xb8, 0xa2, 0x38, 0x60, 0x02, 0x49, 0x6c, 0x60, 0x9f, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x2b, 0x66, 0x12, 0x3c, 0x30, 0x01, 0x00, 0x00,
}
