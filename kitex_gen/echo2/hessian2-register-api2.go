package echo2

import (
	"fmt"

	"github.com/kitex-contrib/codec-dubbo/pkg/hessian2"
	codec "github.com/kitex-contrib/codec-dubbo/pkg/iface"
	"github.com/pkg/errors"
)

var objectsApi2 = []interface{}{
	&EchoRequest{},
	&EchoResponse{},
}

func init() {
	hessian2.Register(objectsApi2)
}

func GetTestService2IDLAnnotations() map[string][]string {
	return map[string][]string{
		"JavaClassName": {"kitex.echo2.TestService2"},
	}
}

func (p *EchoRequest) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Message)
	if err != nil {
		return err
	}

	err = e.Encode(p.Obj)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoRequest) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Message)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Obj)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoRequest) JavaClassName() string {
	return "kitex.echo2.EchoRequest"
}

func (p *EchoResponse) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Message)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoResponse) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Message)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoResponse) JavaClassName() string {
	return "kitex.echo2.EchoResponse"
}

func (p *TestService2EchoArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestService2EchoArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestService2EchoResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestService2EchoResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}
