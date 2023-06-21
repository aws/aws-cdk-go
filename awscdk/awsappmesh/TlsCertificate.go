package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a TLS certificate.
//
// Example:
//   // A Virtual Node with listener TLS from an ACM provided certificate
//   var cert certificate
//   var mesh mesh
//
//
//   node := appmesh.NewVirtualNode(this, jsii.String("node"), &VirtualNodeProps{
//   	Mesh: Mesh,
//   	ServiceDiscovery: appmesh.ServiceDiscovery_Dns(jsii.String("node")),
//   	Listeners: []virtualNodeListener{
//   		appmesh.*virtualNodeListener_Grpc(&GrpcVirtualNodeListenerOptions{
//   			Port: jsii.Number(80),
//   			Tls: &ListenerTlsOptions{
//   				Mode: appmesh.TlsMode_STRICT,
//   				Certificate: appmesh.TlsCertificate_Acm(cert),
//   			},
//   		}),
//   	},
//   })
//
//   // A Virtual Gateway with listener TLS from a customer provided file certificate
//   gateway := appmesh.NewVirtualGateway(this, jsii.String("gateway"), &VirtualGatewayProps{
//   	Mesh: Mesh,
//   	Listeners: []virtualGatewayListener{
//   		appmesh.*virtualGatewayListener_Grpc(&GrpcGatewayListenerOptions{
//   			Port: jsii.Number(8080),
//   			Tls: &ListenerTlsOptions{
//   				Mode: appmesh.TlsMode_STRICT,
//   				Certificate: appmesh.TlsCertificate_File(jsii.String("path/to/certChain"), jsii.String("path/to/privateKey")),
//   			},
//   		}),
//   	},
//   	VirtualGatewayName: jsii.String("gateway"),
//   })
//
//   // A Virtual Gateway with listener TLS from a SDS provided certificate
//   gateway2 := appmesh.NewVirtualGateway(this, jsii.String("gateway2"), &VirtualGatewayProps{
//   	Mesh: Mesh,
//   	Listeners: []*virtualGatewayListener{
//   		appmesh.*virtualGatewayListener_Http2(&Http2GatewayListenerOptions{
//   			Port: jsii.Number(8080),
//   			Tls: &ListenerTlsOptions{
//   				Mode: appmesh.TlsMode_STRICT,
//   				Certificate: appmesh.TlsCertificate_Sds(jsii.String("secrete_certificate")),
//   			},
//   		}),
//   	},
//   	VirtualGatewayName: jsii.String("gateway2"),
//   })
//
type TlsCertificate interface {
	// Returns TLS certificate based provider.
	Bind(_scope constructs.Construct) *TlsCertificateConfig
}

// The jsii proxy struct for TlsCertificate
type jsiiProxy_TlsCertificate struct {
	_ byte // padding
}

func NewTlsCertificate_Override(t TlsCertificate) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.TlsCertificate",
		nil, // no parameters
		t,
	)
}

// Returns an ACM TLS Certificate.
func TlsCertificate_Acm(certificate awscertificatemanager.ICertificate) TlsCertificate {
	_init_.Initialize()

	if err := validateTlsCertificate_AcmParameters(certificate); err != nil {
		panic(err)
	}
	var returns TlsCertificate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.TlsCertificate",
		"acm",
		[]interface{}{certificate},
		&returns,
	)

	return returns
}

// Returns an File TLS Certificate.
func TlsCertificate_File(certificateChainPath *string, privateKeyPath *string) MutualTlsCertificate {
	_init_.Initialize()

	if err := validateTlsCertificate_FileParameters(certificateChainPath, privateKeyPath); err != nil {
		panic(err)
	}
	var returns MutualTlsCertificate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.TlsCertificate",
		"file",
		[]interface{}{certificateChainPath, privateKeyPath},
		&returns,
	)

	return returns
}

// Returns an SDS TLS Certificate.
func TlsCertificate_Sds(secretName *string) MutualTlsCertificate {
	_init_.Initialize()

	if err := validateTlsCertificate_SdsParameters(secretName); err != nil {
		panic(err)
	}
	var returns MutualTlsCertificate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.TlsCertificate",
		"sds",
		[]interface{}{secretName},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TlsCertificate) Bind(_scope constructs.Construct) *TlsCertificateConfig {
	if err := t.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *TlsCertificateConfig

	_jsii_.Invoke(
		t,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

