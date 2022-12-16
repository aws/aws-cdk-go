package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents the properties needed to define listeners for a VirtualGateway.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var mesh mesh
//
//   certificateAuthorityArn := "arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012"
//
//   gateway := appmesh.NewVirtualGateway(this, jsii.String("gateway"), &virtualGatewayProps{
//   	mesh: mesh,
//   	listeners: []virtualGatewayListener{
//   		appmesh.*virtualGatewayListener.http(&httpGatewayListenerOptions{
//   			port: jsii.Number(443),
//   			healthCheck: appmesh.healthCheck.http(&httpHealthCheckOptions{
//   				interval: cdk.duration.seconds(jsii.Number(10)),
//   			}),
//   		}),
//   	},
//   	backendDefaults: &backendDefaults{
//   		tlsClientPolicy: &tlsClientPolicy{
//   			ports: []*f64{
//   				jsii.Number(8080),
//   				jsii.Number(8081),
//   			},
//   			validation: &tlsValidation{
//   				trust: appmesh.tlsValidationTrust.acm([]iCertificateAuthority{
//   					acmpca.certificateAuthority.fromCertificateAuthorityArn(this, jsii.String("certificate"), certificateAuthorityArn),
//   				}),
//   			},
//   		},
//   	},
//   	accessLog: appmesh.accessLog.fromFilePath(jsii.String("/dev/stdout")),
//   	virtualGatewayName: jsii.String("virtualGateway"),
//   })
//
type VirtualGatewayListener interface {
	// Called when the GatewayListener type is initialized.
	//
	// Can be used to enforce
	// mutual exclusivity.
	Bind(scope constructs.Construct) *VirtualGatewayListenerConfig
}

// The jsii proxy struct for VirtualGatewayListener
type jsiiProxy_VirtualGatewayListener struct {
	_ byte // padding
}

func NewVirtualGatewayListener_Override(v VirtualGatewayListener) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.VirtualGatewayListener",
		nil, // no parameters
		v,
	)
}

// Returns a GRPC Listener for a VirtualGateway.
func VirtualGatewayListener_Grpc(options *GrpcGatewayListenerOptions) VirtualGatewayListener {
	_init_.Initialize()

	if err := validateVirtualGatewayListener_GrpcParameters(options); err != nil {
		panic(err)
	}
	var returns VirtualGatewayListener

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualGatewayListener",
		"grpc",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Returns an HTTP Listener for a VirtualGateway.
func VirtualGatewayListener_Http(options *HttpGatewayListenerOptions) VirtualGatewayListener {
	_init_.Initialize()

	if err := validateVirtualGatewayListener_HttpParameters(options); err != nil {
		panic(err)
	}
	var returns VirtualGatewayListener

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualGatewayListener",
		"http",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Returns an HTTP2 Listener for a VirtualGateway.
func VirtualGatewayListener_Http2(options *Http2GatewayListenerOptions) VirtualGatewayListener {
	_init_.Initialize()

	if err := validateVirtualGatewayListener_Http2Parameters(options); err != nil {
		panic(err)
	}
	var returns VirtualGatewayListener

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.VirtualGatewayListener",
		"http2",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualGatewayListener) Bind(scope constructs.Construct) *VirtualGatewayListenerConfig {
	if err := v.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *VirtualGatewayListenerConfig

	_jsii_.Invoke(
		v,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

