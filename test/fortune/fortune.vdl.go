// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: fortune.vdl

package fortune

import (
	"fmt"
	"io"
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/rpc"
	"v.io/v23/security/access"
	"v.io/v23/vdl"
	"v.io/v23/verror"
)

type ComplexErrorParam struct {
	Str  string
	Num  int32
	List []uint32
}

func (ComplexErrorParam) __VDLReflect(struct {
	Name string `vdl:"v.io/x/jni/test/fortune.ComplexErrorParam"`
}) {
}

func (m *ComplexErrorParam) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if __VDLType_fortune_v_io_x_jni_test_fortune_ComplexErrorParam == nil || __VDLTypefortune0 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Str")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromString(string(m.Str), vdl.StringType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Num")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget5.FromInt(int64(m.Num), vdl.Int32Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	keyTarget6, fieldTarget7, err := fieldsTarget1.StartField("List")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		listTarget8, err := fieldTarget7.StartList(__VDLTypefortune1, len(m.List))
		if err != nil {
			return err
		}
		for i, elem10 := range m.List {
			elemTarget9, err := listTarget8.StartElem(i)
			if err != nil {
				return err
			}
			if err := elemTarget9.FromUint(uint64(elem10), vdl.Uint32Type); err != nil {
				return err
			}
			if err := listTarget8.FinishElem(elemTarget9); err != nil {
				return err
			}
		}
		if err := fieldTarget7.FinishList(listTarget8); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget6, fieldTarget7); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *ComplexErrorParam) MakeVDLTarget() vdl.Target {
	return &ComplexErrorParamTarget{Value: m}
}

type ComplexErrorParamTarget struct {
	Value *ComplexErrorParam
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *ComplexErrorParamTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {
	if !vdl.Compatible(tt, __VDLType_fortune_v_io_x_jni_test_fortune_ComplexErrorParam) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_fortune_v_io_x_jni_test_fortune_ComplexErrorParam)
	}
	return t, nil
}
func (t *ComplexErrorParamTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Str":
		val, err := &vdl.StringTarget{Value: &t.Value.Str}, error(nil)
		return nil, val, err
	case "Num":
		val, err := &vdl.Int32Target{Value: &t.Value.Num}, error(nil)
		return nil, val, err
	case "List":
		val, err := &fortune5b5d75696e743332Target{Value: &t.Value.List}, error(nil)
		return nil, val, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct %v", name, __VDLType_fortune_v_io_x_jni_test_fortune_ComplexErrorParam)
	}
}
func (t *ComplexErrorParamTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *ComplexErrorParamTarget) FinishFields(_ vdl.FieldsTarget) error {
	return nil
}

type fortune5b5d75696e743332Target struct {
	Value *[]uint32
	vdl.TargetBase
	vdl.ListTargetBase
}

func (t *fortune5b5d75696e743332Target) StartList(tt *vdl.Type, len int) (vdl.ListTarget, error) {
	if !vdl.Compatible(tt, __VDLTypefortune1) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLTypefortune1)
	}
	if cap(*t.Value) < len {
		*t.Value = make([]uint32, len)
	} else {
		*t.Value = (*t.Value)[:len]
	}
	return t, nil
}
func (t *fortune5b5d75696e743332Target) StartElem(index int) (elem vdl.Target, _ error) {
	return &vdl.Uint32Target{Value: &(*t.Value)[index]}, error(nil)
}
func (t *fortune5b5d75696e743332Target) FinishElem(elem vdl.Target) error {
	return nil
}
func (t *fortune5b5d75696e743332Target) FinishList(elem vdl.ListTarget) error {
	return nil
}

func init() {
	vdl.Register((*ComplexErrorParam)(nil))
}

var __VDLTypefortune0 *vdl.Type = vdl.TypeOf((*ComplexErrorParam)(nil))
var __VDLTypefortune1 *vdl.Type = vdl.TypeOf([]uint32(nil))
var __VDLType_fortune_v_io_x_jni_test_fortune_ComplexErrorParam *vdl.Type = vdl.TypeOf(ComplexErrorParam{})

func __VDLEnsureNativeBuilt_fortune() {
}

var (
	ErrNoFortunes = verror.Register("v.io/x/jni/test/fortune.NoFortunes", verror.NoRetry, "{1:}{2:} no fortunes added")
	ErrComplex    = verror.Register("v.io/x/jni/test/fortune.Complex", verror.NoRetry, "{1:}{2:} this is a complex error with params {3} {4} {5}")
)

func init() {
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrNoFortunes.ID), "{1:}{2:} no fortunes added")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrComplex.ID), "{1:}{2:} this is a complex error with params {3} {4} {5}")
}

// NewErrNoFortunes returns an error with the ErrNoFortunes ID.
func NewErrNoFortunes(ctx *context.T) error {
	return verror.New(ErrNoFortunes, ctx)
}

// NewErrComplex returns an error with the ErrComplex ID.
func NewErrComplex(ctx *context.T, first ComplexErrorParam, second string, third int32) error {
	return verror.New(ErrComplex, ctx, first, second, third)
}

// FortuneClientMethods is the client interface
// containing Fortune methods.
//
// Fortune allows clients to Get and Add fortune strings.
type FortuneClientMethods interface {
	// Add stores a fortune in the set used by Get.
	Add(_ *context.T, Fortune string, _ ...rpc.CallOpt) error
	// Get returns the last-added fortune.
	Get(*context.T, ...rpc.CallOpt) (Fortune string, _ error)
	// ParameterizedGet returns the last-added fortune as a map (which is a parameterized
	// type in Java).
	ParameterizedGet(*context.T, ...rpc.CallOpt) (map[string]string, error)
	// StreamingGet returns a stream that can be used to obtain fortunes, and returns the
	// total number of items sent.
	StreamingGet(*context.T, ...rpc.CallOpt) (FortuneStreamingGetClientCall, error)
	// MultipleGet returns the same fortune twice.
	MultipleGet(*context.T, ...rpc.CallOpt) (Fortune string, Another string, _ error)
	// MultipleStreamingGet returns a stream that can be used to obtain fortunes, and returns
	// the total number of items sent, twice.
	MultipleStreamingGet(*context.T, ...rpc.CallOpt) (FortuneMultipleStreamingGetClientCall, error)
	// GetComplexError returns (always!) ErrComplex.
	GetComplexError(*context.T, ...rpc.CallOpt) error
	// NoTags is a method without tags.
	NoTags(*context.T, ...rpc.CallOpt) error
	// TestServerCall is a method used for testing that the server receives a
	// correct ServerCall.
	TestServerCall(*context.T, ...rpc.CallOpt) error
	// GetServerThread returns a name of the server thread that executes this method.
	GetServerThread(*context.T, ...rpc.CallOpt) (string, error)
}

// FortuneClientStub adds universal methods to FortuneClientMethods.
type FortuneClientStub interface {
	FortuneClientMethods
	rpc.UniversalServiceMethods
}

// FortuneClient returns a client stub for Fortune.
func FortuneClient(name string) FortuneClientStub {
	return implFortuneClientStub{name}
}

type implFortuneClientStub struct {
	name string
}

func (c implFortuneClientStub) Add(ctx *context.T, i0 string, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Add", []interface{}{i0}, nil, opts...)
	return
}

func (c implFortuneClientStub) Get(ctx *context.T, opts ...rpc.CallOpt) (o0 string, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Get", nil, []interface{}{&o0}, opts...)
	return
}

func (c implFortuneClientStub) ParameterizedGet(ctx *context.T, opts ...rpc.CallOpt) (o0 map[string]string, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "ParameterizedGet", nil, []interface{}{&o0}, opts...)
	return
}

func (c implFortuneClientStub) StreamingGet(ctx *context.T, opts ...rpc.CallOpt) (ocall FortuneStreamingGetClientCall, err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "StreamingGet", nil, opts...); err != nil {
		return
	}
	ocall = &implFortuneStreamingGetClientCall{ClientCall: call}
	return
}

func (c implFortuneClientStub) MultipleGet(ctx *context.T, opts ...rpc.CallOpt) (o0 string, o1 string, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "MultipleGet", nil, []interface{}{&o0, &o1}, opts...)
	return
}

func (c implFortuneClientStub) MultipleStreamingGet(ctx *context.T, opts ...rpc.CallOpt) (ocall FortuneMultipleStreamingGetClientCall, err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "MultipleStreamingGet", nil, opts...); err != nil {
		return
	}
	ocall = &implFortuneMultipleStreamingGetClientCall{ClientCall: call}
	return
}

func (c implFortuneClientStub) GetComplexError(ctx *context.T, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "GetComplexError", nil, nil, opts...)
	return
}

func (c implFortuneClientStub) NoTags(ctx *context.T, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "NoTags", nil, nil, opts...)
	return
}

func (c implFortuneClientStub) TestServerCall(ctx *context.T, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "TestServerCall", nil, nil, opts...)
	return
}

func (c implFortuneClientStub) GetServerThread(ctx *context.T, opts ...rpc.CallOpt) (o0 string, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "GetServerThread", nil, []interface{}{&o0}, opts...)
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
		// Send places the item onto the output stream.  Returns errors
		// encountered while sending, or if Send is called after Close or
		// the stream has been canceled.  Blocks if there is no buffer
		// space; will unblock when buffer space is available or after
		// the stream has been canceled.
		Send(item bool) error
		// Close indicates to the server that no more items will be sent;
		// server Recv calls will receive io.EOF after all sent items.
		// This is an optional call - e.g. a client might call Close if it
		// needs to continue receiving items from the server after it's
		// done sending.  Returns errors encountered while closing, or if
		// Close is called after the stream has been canceled.  Like Send,
		// blocks if there is no buffer space available.
		Close() error
	}
}

// FortuneStreamingGetClientCall represents the call returned from Fortune.StreamingGet.
type FortuneStreamingGetClientCall interface {
	FortuneStreamingGetClientStream
	// Finish performs the equivalent of SendStream().Close, then blocks until
	// the server is done, and returns the positional return values for the call.
	//
	// Finish returns immediately if the call has been canceled; depending on the
	// timing the output could either be an error signaling cancelation, or the
	// valid positional return values from the server.
	//
	// Calling Finish is mandatory for releasing stream resources, unless the call
	// has been canceled or any of the other methods return an error.  Finish should
	// be called at most once.
	Finish() (total int32, _ error)
}

type implFortuneStreamingGetClientCall struct {
	rpc.ClientCall
	valRecv string
	errRecv error
}

func (c *implFortuneStreamingGetClientCall) RecvStream() interface {
	Advance() bool
	Value() string
	Err() error
} {
	return implFortuneStreamingGetClientCallRecv{c}
}

type implFortuneStreamingGetClientCallRecv struct {
	c *implFortuneStreamingGetClientCall
}

func (c implFortuneStreamingGetClientCallRecv) Advance() bool {
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implFortuneStreamingGetClientCallRecv) Value() string {
	return c.c.valRecv
}
func (c implFortuneStreamingGetClientCallRecv) Err() error {
	if c.c.errRecv == io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implFortuneStreamingGetClientCall) SendStream() interface {
	Send(item bool) error
	Close() error
} {
	return implFortuneStreamingGetClientCallSend{c}
}

type implFortuneStreamingGetClientCallSend struct {
	c *implFortuneStreamingGetClientCall
}

func (c implFortuneStreamingGetClientCallSend) Send(item bool) error {
	return c.c.Send(item)
}
func (c implFortuneStreamingGetClientCallSend) Close() error {
	return c.c.CloseSend()
}
func (c *implFortuneStreamingGetClientCall) Finish() (o0 int32, err error) {
	err = c.ClientCall.Finish(&o0)
	return
}

// FortuneMultipleStreamingGetClientStream is the client stream for Fortune.MultipleStreamingGet.
type FortuneMultipleStreamingGetClientStream interface {
	// RecvStream returns the receiver side of the Fortune.MultipleStreamingGet client stream.
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
	// SendStream returns the send side of the Fortune.MultipleStreamingGet client stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors
		// encountered while sending, or if Send is called after Close or
		// the stream has been canceled.  Blocks if there is no buffer
		// space; will unblock when buffer space is available or after
		// the stream has been canceled.
		Send(item bool) error
		// Close indicates to the server that no more items will be sent;
		// server Recv calls will receive io.EOF after all sent items.
		// This is an optional call - e.g. a client might call Close if it
		// needs to continue receiving items from the server after it's
		// done sending.  Returns errors encountered while closing, or if
		// Close is called after the stream has been canceled.  Like Send,
		// blocks if there is no buffer space available.
		Close() error
	}
}

// FortuneMultipleStreamingGetClientCall represents the call returned from Fortune.MultipleStreamingGet.
type FortuneMultipleStreamingGetClientCall interface {
	FortuneMultipleStreamingGetClientStream
	// Finish performs the equivalent of SendStream().Close, then blocks until
	// the server is done, and returns the positional return values for the call.
	//
	// Finish returns immediately if the call has been canceled; depending on the
	// timing the output could either be an error signaling cancelation, or the
	// valid positional return values from the server.
	//
	// Calling Finish is mandatory for releasing stream resources, unless the call
	// has been canceled or any of the other methods return an error.  Finish should
	// be called at most once.
	Finish() (total int32, another int32, _ error)
}

type implFortuneMultipleStreamingGetClientCall struct {
	rpc.ClientCall
	valRecv string
	errRecv error
}

func (c *implFortuneMultipleStreamingGetClientCall) RecvStream() interface {
	Advance() bool
	Value() string
	Err() error
} {
	return implFortuneMultipleStreamingGetClientCallRecv{c}
}

type implFortuneMultipleStreamingGetClientCallRecv struct {
	c *implFortuneMultipleStreamingGetClientCall
}

func (c implFortuneMultipleStreamingGetClientCallRecv) Advance() bool {
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implFortuneMultipleStreamingGetClientCallRecv) Value() string {
	return c.c.valRecv
}
func (c implFortuneMultipleStreamingGetClientCallRecv) Err() error {
	if c.c.errRecv == io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implFortuneMultipleStreamingGetClientCall) SendStream() interface {
	Send(item bool) error
	Close() error
} {
	return implFortuneMultipleStreamingGetClientCallSend{c}
}

type implFortuneMultipleStreamingGetClientCallSend struct {
	c *implFortuneMultipleStreamingGetClientCall
}

func (c implFortuneMultipleStreamingGetClientCallSend) Send(item bool) error {
	return c.c.Send(item)
}
func (c implFortuneMultipleStreamingGetClientCallSend) Close() error {
	return c.c.CloseSend()
}
func (c *implFortuneMultipleStreamingGetClientCall) Finish() (o0 int32, o1 int32, err error) {
	err = c.ClientCall.Finish(&o0, &o1)
	return
}

// FortuneServerMethods is the interface a server writer
// implements for Fortune.
//
// Fortune allows clients to Get and Add fortune strings.
type FortuneServerMethods interface {
	// Add stores a fortune in the set used by Get.
	Add(_ *context.T, _ rpc.ServerCall, Fortune string) error
	// Get returns the last-added fortune.
	Get(*context.T, rpc.ServerCall) (Fortune string, _ error)
	// ParameterizedGet returns the last-added fortune as a map (which is a parameterized
	// type in Java).
	ParameterizedGet(*context.T, rpc.ServerCall) (map[string]string, error)
	// StreamingGet returns a stream that can be used to obtain fortunes, and returns the
	// total number of items sent.
	StreamingGet(*context.T, FortuneStreamingGetServerCall) (total int32, _ error)
	// MultipleGet returns the same fortune twice.
	MultipleGet(*context.T, rpc.ServerCall) (Fortune string, Another string, _ error)
	// MultipleStreamingGet returns a stream that can be used to obtain fortunes, and returns
	// the total number of items sent, twice.
	MultipleStreamingGet(*context.T, FortuneMultipleStreamingGetServerCall) (total int32, another int32, _ error)
	// GetComplexError returns (always!) ErrComplex.
	GetComplexError(*context.T, rpc.ServerCall) error
	// NoTags is a method without tags.
	NoTags(*context.T, rpc.ServerCall) error
	// TestServerCall is a method used for testing that the server receives a
	// correct ServerCall.
	TestServerCall(*context.T, rpc.ServerCall) error
	// GetServerThread returns a name of the server thread that executes this method.
	GetServerThread(*context.T, rpc.ServerCall) (string, error)
}

// FortuneServerStubMethods is the server interface containing
// Fortune methods, as expected by rpc.Server.
// The only difference between this interface and FortuneServerMethods
// is the streaming methods.
type FortuneServerStubMethods interface {
	// Add stores a fortune in the set used by Get.
	Add(_ *context.T, _ rpc.ServerCall, Fortune string) error
	// Get returns the last-added fortune.
	Get(*context.T, rpc.ServerCall) (Fortune string, _ error)
	// ParameterizedGet returns the last-added fortune as a map (which is a parameterized
	// type in Java).
	ParameterizedGet(*context.T, rpc.ServerCall) (map[string]string, error)
	// StreamingGet returns a stream that can be used to obtain fortunes, and returns the
	// total number of items sent.
	StreamingGet(*context.T, *FortuneStreamingGetServerCallStub) (total int32, _ error)
	// MultipleGet returns the same fortune twice.
	MultipleGet(*context.T, rpc.ServerCall) (Fortune string, Another string, _ error)
	// MultipleStreamingGet returns a stream that can be used to obtain fortunes, and returns
	// the total number of items sent, twice.
	MultipleStreamingGet(*context.T, *FortuneMultipleStreamingGetServerCallStub) (total int32, another int32, _ error)
	// GetComplexError returns (always!) ErrComplex.
	GetComplexError(*context.T, rpc.ServerCall) error
	// NoTags is a method without tags.
	NoTags(*context.T, rpc.ServerCall) error
	// TestServerCall is a method used for testing that the server receives a
	// correct ServerCall.
	TestServerCall(*context.T, rpc.ServerCall) error
	// GetServerThread returns a name of the server thread that executes this method.
	GetServerThread(*context.T, rpc.ServerCall) (string, error)
}

// FortuneServerStub adds universal methods to FortuneServerStubMethods.
type FortuneServerStub interface {
	FortuneServerStubMethods
	// Describe the Fortune interfaces.
	Describe__() []rpc.InterfaceDesc
}

// FortuneServer returns a server stub for Fortune.
// It converts an implementation of FortuneServerMethods into
// an object that may be used by rpc.Server.
func FortuneServer(impl FortuneServerMethods) FortuneServerStub {
	stub := implFortuneServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implFortuneServerStub struct {
	impl FortuneServerMethods
	gs   *rpc.GlobState
}

func (s implFortuneServerStub) Add(ctx *context.T, call rpc.ServerCall, i0 string) error {
	return s.impl.Add(ctx, call, i0)
}

func (s implFortuneServerStub) Get(ctx *context.T, call rpc.ServerCall) (string, error) {
	return s.impl.Get(ctx, call)
}

func (s implFortuneServerStub) ParameterizedGet(ctx *context.T, call rpc.ServerCall) (map[string]string, error) {
	return s.impl.ParameterizedGet(ctx, call)
}

func (s implFortuneServerStub) StreamingGet(ctx *context.T, call *FortuneStreamingGetServerCallStub) (int32, error) {
	return s.impl.StreamingGet(ctx, call)
}

func (s implFortuneServerStub) MultipleGet(ctx *context.T, call rpc.ServerCall) (string, string, error) {
	return s.impl.MultipleGet(ctx, call)
}

func (s implFortuneServerStub) MultipleStreamingGet(ctx *context.T, call *FortuneMultipleStreamingGetServerCallStub) (int32, int32, error) {
	return s.impl.MultipleStreamingGet(ctx, call)
}

func (s implFortuneServerStub) GetComplexError(ctx *context.T, call rpc.ServerCall) error {
	return s.impl.GetComplexError(ctx, call)
}

func (s implFortuneServerStub) NoTags(ctx *context.T, call rpc.ServerCall) error {
	return s.impl.NoTags(ctx, call)
}

func (s implFortuneServerStub) TestServerCall(ctx *context.T, call rpc.ServerCall) error {
	return s.impl.TestServerCall(ctx, call)
}

func (s implFortuneServerStub) GetServerThread(ctx *context.T, call rpc.ServerCall) (string, error) {
	return s.impl.GetServerThread(ctx, call)
}

func (s implFortuneServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implFortuneServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{FortuneDesc}
}

// FortuneDesc describes the Fortune interface.
var FortuneDesc rpc.InterfaceDesc = descFortune

// descFortune hides the desc to keep godoc clean.
var descFortune = rpc.InterfaceDesc{
	Name:    "Fortune",
	PkgPath: "v.io/x/jni/test/fortune",
	Doc:     "// Fortune allows clients to Get and Add fortune strings.",
	Methods: []rpc.MethodDesc{
		{
			Name: "Add",
			Doc:  "// Add stores a fortune in the set used by Get.",
			InArgs: []rpc.ArgDesc{
				{"Fortune", ``}, // string
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Write"))},
		},
		{
			Name: "Get",
			Doc:  "// Get returns the last-added fortune.",
			OutArgs: []rpc.ArgDesc{
				{"Fortune", ``}, // string
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Read"))},
		},
		{
			Name: "ParameterizedGet",
			Doc:  "// ParameterizedGet returns the last-added fortune as a map (which is a parameterized\n// type in Java).",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // map[string]string
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Read"))},
		},
		{
			Name: "StreamingGet",
			Doc:  "// StreamingGet returns a stream that can be used to obtain fortunes, and returns the\n// total number of items sent.",
			OutArgs: []rpc.ArgDesc{
				{"total", ``}, // int32
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Read"))},
		},
		{
			Name: "MultipleGet",
			Doc:  "// MultipleGet returns the same fortune twice.",
			OutArgs: []rpc.ArgDesc{
				{"Fortune", ``}, // string
				{"Another", ``}, // string
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Read"))},
		},
		{
			Name: "MultipleStreamingGet",
			Doc:  "// MultipleStreamingGet returns a stream that can be used to obtain fortunes, and returns\n// the total number of items sent, twice.",
			OutArgs: []rpc.ArgDesc{
				{"total", ``},   // int32
				{"another", ``}, // int32
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Read"))},
		},
		{
			Name: "GetComplexError",
			Doc:  "// GetComplexError returns (always!) ErrComplex.",
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Read"))},
		},
		{
			Name: "NoTags",
			Doc:  "// NoTags is a method without tags.",
		},
		{
			Name: "TestServerCall",
			Doc:  "// TestServerCall is a method used for testing that the server receives a\n// correct ServerCall.",
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Read"))},
		},
		{
			Name: "GetServerThread",
			Doc:  "// GetServerThread returns a name of the server thread that executes this method.",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // string
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Read"))},
		},
	},
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

// FortuneStreamingGetServerCall represents the context passed to Fortune.StreamingGet.
type FortuneStreamingGetServerCall interface {
	rpc.ServerCall
	FortuneStreamingGetServerStream
}

// FortuneStreamingGetServerCallStub is a wrapper that converts rpc.StreamServerCall into
// a typesafe stub that implements FortuneStreamingGetServerCall.
type FortuneStreamingGetServerCallStub struct {
	rpc.StreamServerCall
	valRecv bool
	errRecv error
}

// Init initializes FortuneStreamingGetServerCallStub from rpc.StreamServerCall.
func (s *FortuneStreamingGetServerCallStub) Init(call rpc.StreamServerCall) {
	s.StreamServerCall = call
}

// RecvStream returns the receiver side of the Fortune.StreamingGet server stream.
func (s *FortuneStreamingGetServerCallStub) RecvStream() interface {
	Advance() bool
	Value() bool
	Err() error
} {
	return implFortuneStreamingGetServerCallRecv{s}
}

type implFortuneStreamingGetServerCallRecv struct {
	s *FortuneStreamingGetServerCallStub
}

func (s implFortuneStreamingGetServerCallRecv) Advance() bool {
	s.s.errRecv = s.s.Recv(&s.s.valRecv)
	return s.s.errRecv == nil
}
func (s implFortuneStreamingGetServerCallRecv) Value() bool {
	return s.s.valRecv
}
func (s implFortuneStreamingGetServerCallRecv) Err() error {
	if s.s.errRecv == io.EOF {
		return nil
	}
	return s.s.errRecv
}

// SendStream returns the send side of the Fortune.StreamingGet server stream.
func (s *FortuneStreamingGetServerCallStub) SendStream() interface {
	Send(item string) error
} {
	return implFortuneStreamingGetServerCallSend{s}
}

type implFortuneStreamingGetServerCallSend struct {
	s *FortuneStreamingGetServerCallStub
}

func (s implFortuneStreamingGetServerCallSend) Send(item string) error {
	return s.s.Send(item)
}

// FortuneMultipleStreamingGetServerStream is the server stream for Fortune.MultipleStreamingGet.
type FortuneMultipleStreamingGetServerStream interface {
	// RecvStream returns the receiver side of the Fortune.MultipleStreamingGet server stream.
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
	// SendStream returns the send side of the Fortune.MultipleStreamingGet server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item string) error
	}
}

// FortuneMultipleStreamingGetServerCall represents the context passed to Fortune.MultipleStreamingGet.
type FortuneMultipleStreamingGetServerCall interface {
	rpc.ServerCall
	FortuneMultipleStreamingGetServerStream
}

// FortuneMultipleStreamingGetServerCallStub is a wrapper that converts rpc.StreamServerCall into
// a typesafe stub that implements FortuneMultipleStreamingGetServerCall.
type FortuneMultipleStreamingGetServerCallStub struct {
	rpc.StreamServerCall
	valRecv bool
	errRecv error
}

// Init initializes FortuneMultipleStreamingGetServerCallStub from rpc.StreamServerCall.
func (s *FortuneMultipleStreamingGetServerCallStub) Init(call rpc.StreamServerCall) {
	s.StreamServerCall = call
}

// RecvStream returns the receiver side of the Fortune.MultipleStreamingGet server stream.
func (s *FortuneMultipleStreamingGetServerCallStub) RecvStream() interface {
	Advance() bool
	Value() bool
	Err() error
} {
	return implFortuneMultipleStreamingGetServerCallRecv{s}
}

type implFortuneMultipleStreamingGetServerCallRecv struct {
	s *FortuneMultipleStreamingGetServerCallStub
}

func (s implFortuneMultipleStreamingGetServerCallRecv) Advance() bool {
	s.s.errRecv = s.s.Recv(&s.s.valRecv)
	return s.s.errRecv == nil
}
func (s implFortuneMultipleStreamingGetServerCallRecv) Value() bool {
	return s.s.valRecv
}
func (s implFortuneMultipleStreamingGetServerCallRecv) Err() error {
	if s.s.errRecv == io.EOF {
		return nil
	}
	return s.s.errRecv
}

// SendStream returns the send side of the Fortune.MultipleStreamingGet server stream.
func (s *FortuneMultipleStreamingGetServerCallStub) SendStream() interface {
	Send(item string) error
} {
	return implFortuneMultipleStreamingGetServerCallSend{s}
}

type implFortuneMultipleStreamingGetServerCallSend struct {
	s *FortuneMultipleStreamingGetServerCallStub
}

func (s implFortuneMultipleStreamingGetServerCallSend) Send(item string) error {
	return s.s.Send(item)
}
