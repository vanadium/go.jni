// This file was auto-generated by the veyron vdl tool.
// Source: fortune.vdl

package fortune

import (
	"veyron.io/veyron/veyron2/services/security/access"

	// The non-user imports are prefixed with "__" to prevent collisions.
	__io "io"
	__veyron2 "veyron.io/veyron/veyron2"
	__context "veyron.io/veyron/veyron2/context"
	__ipc "veyron.io/veyron/veyron2/ipc"
	__vdlutil "veyron.io/veyron/veyron2/vdl/vdlutil"
	__wiretype "veyron.io/veyron/veyron2/wiretype"
)

// TODO(toddw): Remove this line once the new signature support is done.
// It corrects a bug where __wiretype is unused in VDL pacakges where only
// bootstrap types are used on interfaces.
const _ = __wiretype.TypeIDInvalid

// FortuneClientMethods is the client interface
// containing Fortune methods.
//
// Fortune allows clients to Get and Add fortune strings.
type FortuneClientMethods interface {
	// Add stores a fortune in the set used by Get.
	Add(ctx __context.T, Fortune string, opts ...__ipc.CallOpt) error
	// Get returns a random fortune.
	Get(__context.T, ...__ipc.CallOpt) (Fortune string, err error)
	// StreamingGet returns a stream that can be used to obtain fortunes.
	StreamingGet(__context.T, ...__ipc.CallOpt) (FortuneStreamingGetCall, error)
}

// FortuneClientStub adds universal methods to FortuneClientMethods.
type FortuneClientStub interface {
	FortuneClientMethods
	__ipc.UniversalServiceMethods
}

// FortuneClient returns a client stub for Fortune.
func FortuneClient(name string, opts ...__ipc.BindOpt) FortuneClientStub {
	var client __ipc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(__ipc.Client); ok {
			client = clientOpt
		}
	}
	return implFortuneClientStub{name, client}
}

type implFortuneClientStub struct {
	name   string
	client __ipc.Client
}

func (c implFortuneClientStub) c(ctx __context.T) __ipc.Client {
	if c.client != nil {
		return c.client
	}
	return __veyron2.RuntimeFromContext(ctx).Client()
}

func (c implFortuneClientStub) Add(ctx __context.T, i0 string, opts ...__ipc.CallOpt) (err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Add", []interface{}{i0}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (c implFortuneClientStub) Get(ctx __context.T, opts ...__ipc.CallOpt) (o0 string, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Get", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implFortuneClientStub) StreamingGet(ctx __context.T, opts ...__ipc.CallOpt) (ocall FortuneStreamingGetCall, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "StreamingGet", nil, opts...); err != nil {
		return
	}
	ocall = &implFortuneStreamingGetCall{Call: call}
	return
}

func (c implFortuneClientStub) Signature(ctx __context.T, opts ...__ipc.CallOpt) (o0 __ipc.ServiceSignature, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

// FortuneStreamingGetClientStream is the client stream for Fortune.StreamingGet.
type FortuneStreamingGetClientStream interface {
	// RecvStream returns the receiver side of the Fortune.StreamingGet client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() string
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
	// SendStream returns the send side of the Fortune.StreamingGet client stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending, or if Send is called after Close or Cancel.  Blocks if
		// there is no buffer space; will unblock when buffer space is available or
		// after Cancel.
		Send(item bool) error
		// Close indicates to the server that no more items will be sent; server
		// Recv calls will receive io.EOF after all sent items.  This is an optional
		// call - e.g. a client might call Close if it needs to continue receiving
		// items from the server after it's done sending.  Returns errors
		// encountered while closing, or if Close is called after Cancel.  Like
		// Send, blocks if there is no buffer space available.
		Close() error
	}
}

// FortuneStreamingGetCall represents the call returned from Fortune.StreamingGet.
type FortuneStreamingGetCall interface {
	FortuneStreamingGetClientStream
	// Finish performs the equivalent of SendStream().Close, then blocks until
	// the server is done, and returns the positional return values for the call.
	//
	// Finish returns immediately if Cancel has been called; depending on the
	// timing the output could either be an error signaling cancelation, or the
	// valid positional return values from the server.
	//
	// Calling Finish is mandatory for releasing stream resources, unless Cancel
	// has been called or any of the other methods return an error.  Finish should
	// be called at most once.
	Finish() (total int32, err error)
	// Cancel cancels the RPC, notifying the server to stop processing.  It is
	// safe to call Cancel concurrently with any of the other stream methods.
	// Calling Cancel after Finish has returned is a no-op.
	Cancel()
}

type implFortuneStreamingGetCall struct {
	__ipc.Call
	valRecv string
	errRecv error
}

func (c *implFortuneStreamingGetCall) RecvStream() interface {
	Advance() bool
	Value() string
	Err() error
} {
	return implFortuneStreamingGetCallRecv{c}
}

type implFortuneStreamingGetCallRecv struct {
	c *implFortuneStreamingGetCall
}

func (c implFortuneStreamingGetCallRecv) Advance() bool {
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implFortuneStreamingGetCallRecv) Value() string {
	return c.c.valRecv
}
func (c implFortuneStreamingGetCallRecv) Err() error {
	if c.c.errRecv == __io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implFortuneStreamingGetCall) SendStream() interface {
	Send(item bool) error
	Close() error
} {
	return implFortuneStreamingGetCallSend{c}
}

type implFortuneStreamingGetCallSend struct {
	c *implFortuneStreamingGetCall
}

func (c implFortuneStreamingGetCallSend) Send(item bool) error {
	return c.c.Send(item)
}
func (c implFortuneStreamingGetCallSend) Close() error {
	return c.c.CloseSend()
}
func (c *implFortuneStreamingGetCall) Finish() (o0 int32, err error) {
	if ierr := c.Call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

// FortuneServerMethods is the interface a server writer
// implements for Fortune.
//
// Fortune allows clients to Get and Add fortune strings.
type FortuneServerMethods interface {
	// Add stores a fortune in the set used by Get.
	Add(ctx __ipc.ServerContext, Fortune string) error
	// Get returns a random fortune.
	Get(__ipc.ServerContext) (Fortune string, Err error)
	// StreamingGet returns a stream that can be used to obtain fortunes.
	StreamingGet(FortuneStreamingGetContext) (total int32, err error)
}

// FortuneServerStubMethods is the server interface containing
// Fortune methods, as expected by ipc.Server.
// The only difference between this interface and FortuneServerMethods
// is the streaming methods.
type FortuneServerStubMethods interface {
	// Add stores a fortune in the set used by Get.
	Add(ctx __ipc.ServerContext, Fortune string) error
	// Get returns a random fortune.
	Get(__ipc.ServerContext) (Fortune string, Err error)
	// StreamingGet returns a stream that can be used to obtain fortunes.
	StreamingGet(*FortuneStreamingGetContextStub) (total int32, err error)
}

// FortuneServerStub adds universal methods to FortuneServerStubMethods.
type FortuneServerStub interface {
	FortuneServerStubMethods
	// Describe the Fortune interfaces.
	Describe__() []__ipc.InterfaceDesc
	// Signature will be replaced with Describe__.
	Signature(ctx __ipc.ServerContext) (__ipc.ServiceSignature, error)
}

// FortuneServer returns a server stub for Fortune.
// It converts an implementation of FortuneServerMethods into
// an object that may be used by ipc.Server.
func FortuneServer(impl FortuneServerMethods) FortuneServerStub {
	stub := implFortuneServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := __ipc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := __ipc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implFortuneServerStub struct {
	impl FortuneServerMethods
	gs   *__ipc.GlobState
}

func (s implFortuneServerStub) Add(ctx __ipc.ServerContext, i0 string) error {
	return s.impl.Add(ctx, i0)
}

func (s implFortuneServerStub) Get(ctx __ipc.ServerContext) (string, error) {
	return s.impl.Get(ctx)
}

func (s implFortuneServerStub) StreamingGet(ctx *FortuneStreamingGetContextStub) (int32, error) {
	return s.impl.StreamingGet(ctx)
}

func (s implFortuneServerStub) VGlob() *__ipc.GlobState {
	return s.gs
}

func (s implFortuneServerStub) Describe__() []__ipc.InterfaceDesc {
	return []__ipc.InterfaceDesc{FortuneDesc}
}

// FortuneDesc describes the Fortune interface.
var FortuneDesc __ipc.InterfaceDesc = descFortune

// descFortune hides the desc to keep godoc clean.
var descFortune = __ipc.InterfaceDesc{
	Name:    "Fortune",
	PkgPath: "veyron.io/jni/test/fortune",
	Doc:     "// Fortune allows clients to Get and Add fortune strings.",
	Methods: []__ipc.MethodDesc{
		{
			Name: "Add",
			Doc:  "// Add stores a fortune in the set used by Get.",
			InArgs: []__ipc.ArgDesc{
				{"Fortune", ``}, // string
			},
			OutArgs: []__ipc.ArgDesc{
				{"", ``}, // error
			},
			Tags: []__vdlutil.Any{access.Tag("Write")},
		},
		{
			Name: "Get",
			Doc:  "// Get returns a random fortune.",
			OutArgs: []__ipc.ArgDesc{
				{"Fortune", ``}, // string
				{"Err", ``},     // error
			},
			Tags: []__vdlutil.Any{access.Tag("Read")},
		},
		{
			Name: "StreamingGet",
			Doc:  "// StreamingGet returns a stream that can be used to obtain fortunes.",
			OutArgs: []__ipc.ArgDesc{
				{"total", ``}, // int32
				{"err", ``},   // error
			},
			Tags: []__vdlutil.Any{access.Tag("Read")},
		},
	},
}

func (s implFortuneServerStub) Signature(ctx __ipc.ServerContext) (__ipc.ServiceSignature, error) {
	// TODO(toddw): Replace with new Describe__ implementation.
	result := __ipc.ServiceSignature{Methods: make(map[string]__ipc.MethodSignature)}
	result.Methods["Add"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{
			{Name: "Fortune", Type: 3},
		},
		OutArgs: []__ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["Get"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{},
		OutArgs: []__ipc.MethodArgument{
			{Name: "Fortune", Type: 3},
			{Name: "Err", Type: 65},
		},
	}
	result.Methods["StreamingGet"] = __ipc.MethodSignature{
		InArgs: []__ipc.MethodArgument{},
		OutArgs: []__ipc.MethodArgument{
			{Name: "total", Type: 36},
			{Name: "err", Type: 65},
		},
		InStream:  2,
		OutStream: 3,
	}

	result.TypeDefs = []__vdlutil.Any{
		__wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}}

	return result, nil
}

// FortuneStreamingGetServerStream is the server stream for Fortune.StreamingGet.
type FortuneStreamingGetServerStream interface {
	// RecvStream returns the receiver side of the Fortune.StreamingGet server stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() bool
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
	// SendStream returns the send side of the Fortune.StreamingGet server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item string) error
	}
}

// FortuneStreamingGetContext represents the context passed to Fortune.StreamingGet.
type FortuneStreamingGetContext interface {
	__ipc.ServerContext
	FortuneStreamingGetServerStream
}

// FortuneStreamingGetContextStub is a wrapper that converts ipc.ServerCall into
// a typesafe stub that implements FortuneStreamingGetContext.
type FortuneStreamingGetContextStub struct {
	__ipc.ServerCall
	valRecv bool
	errRecv error
}

// Init initializes FortuneStreamingGetContextStub from ipc.ServerCall.
func (s *FortuneStreamingGetContextStub) Init(call __ipc.ServerCall) {
	s.ServerCall = call
}

// RecvStream returns the receiver side of the Fortune.StreamingGet server stream.
func (s *FortuneStreamingGetContextStub) RecvStream() interface {
	Advance() bool
	Value() bool
	Err() error
} {
	return implFortuneStreamingGetContextRecv{s}
}

type implFortuneStreamingGetContextRecv struct {
	s *FortuneStreamingGetContextStub
}

func (s implFortuneStreamingGetContextRecv) Advance() bool {
	s.s.errRecv = s.s.Recv(&s.s.valRecv)
	return s.s.errRecv == nil
}
func (s implFortuneStreamingGetContextRecv) Value() bool {
	return s.s.valRecv
}
func (s implFortuneStreamingGetContextRecv) Err() error {
	if s.s.errRecv == __io.EOF {
		return nil
	}
	return s.s.errRecv
}

// SendStream returns the send side of the Fortune.StreamingGet server stream.
func (s *FortuneStreamingGetContextStub) SendStream() interface {
	Send(item string) error
} {
	return implFortuneStreamingGetContextSend{s}
}

type implFortuneStreamingGetContextSend struct {
	s *FortuneStreamingGetContextStub
}

func (s implFortuneStreamingGetContextSend) Send(item string) error {
	return s.s.Send(item)
}
