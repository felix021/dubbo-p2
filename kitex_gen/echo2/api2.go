// Code generated by thriftgo (0.3.6). DO NOT EDIT.

package echo2

import (
	"context"
	"fmt"
	"github.com/felix021/dubbo-p2/kitex_gen/java"
)

type EchoRequest struct {
	Message string       `thrift:"message,1,required" frugal:"1,required,string" json:"message"`
	Obj     java.Object `thrift:"obj,2,optional" frugal:"2,optional,java.Object" json:"obj,omitempty"`
}

func NewEchoRequest() *EchoRequest {
	return &EchoRequest{}
}

func (p *EchoRequest) InitDefault() {
	*p = EchoRequest{}
}

func (p *EchoRequest) GetMessage() (v string) {
	return p.Message
}

var EchoRequest_Obj_DEFAULT java.Object

func (p *EchoRequest) GetObj() (v java.Object) {
	if !p.IsSetObj() {
		return EchoRequest_Obj_DEFAULT
	}
	return p.Obj
}
func (p *EchoRequest) SetMessage(val string) {
	p.Message = val
}
func (p *EchoRequest) SetObj(val java.Object) {
	p.Obj = val
}

func (p *EchoRequest) IsSetObj() bool {
	return p.Obj != nil
}

func (p *EchoRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoRequest(%+v)", *p)
}

type EchoResponse struct {
	Message string `thrift:"message,1,required" frugal:"1,required,string" json:"message"`
}

func NewEchoResponse() *EchoResponse {
	return &EchoResponse{}
}

func (p *EchoResponse) InitDefault() {
	*p = EchoResponse{}
}

func (p *EchoResponse) GetMessage() (v string) {
	return p.Message
}
func (p *EchoResponse) SetMessage(val string) {
	p.Message = val
}

func (p *EchoResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoResponse(%+v)", *p)
}

type TestService2 interface {
	Echo(ctx context.Context, req *EchoRequest) (r *EchoResponse, err error)
}
