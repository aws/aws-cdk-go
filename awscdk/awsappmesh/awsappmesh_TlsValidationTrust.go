package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsacmpca"
)

// Defines the TLS Validation Context Trust.
//
// Example:
//   var mesh mesh
//   var service service
//
//
//   node := appmesh.NewVirtualNode(this, jsii.String("node"), &virtualNodeProps{
//   	mesh: mesh,
//   	serviceDiscovery: appmesh.serviceDiscovery.cloudMap(service),
//   	listeners: []virtualNodeListener{
//   		appmesh.*virtualNodeListener.http(&httpVirtualNodeListenerOptions{
//   			port: jsii.Number(8080),
//   			healthCheck: appmesh.healthCheck.http(&httpHealthCheckOptions{
//   				healthyThreshold: jsii.Number(3),
//   				interval: cdk.duration.seconds(jsii.Number(5)),
//   				path: jsii.String("/ping"),
//   				timeout: cdk.*duration.seconds(jsii.Number(2)),
//   				unhealthyThreshold: jsii.Number(2),
//   			}),
//   			timeout: &httpTimeout{
//   				idle: cdk.*duration.seconds(jsii.Number(5)),
//   			},
//   		}),
//   	},
//   	backendDefaults: &backendDefaults{
//   		tlsClientPolicy: &tlsClientPolicy{
//   			validation: &tlsValidation{
//   				trust: appmesh.tlsValidationTrust.file(jsii.String("/keys/local_cert_chain.pem")),
//   			},
//   		},
//   	},
//   	accessLog: appmesh.accessLog.fromFilePath(jsii.String("/dev/stdout")),
//   })
//
//   cdk.tags.of(node).add(jsii.String("Environment"), jsii.String("Dev"))
//
// Experimental.
type TlsValidationTrust interface {
	// Returns Trust context based on trust type.
	// Experimental.
	Bind(scope awscdk.Construct) *TlsValidationTrustConfig
}

// The jsii proxy struct for TlsValidationTrust
type jsiiProxy_TlsValidationTrust struct {
	_ byte // padding
}

// Experimental.
func NewTlsValidationTrust_Override(t TlsValidationTrust) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.TlsValidationTrust",
		nil, // no parameters
		t,
	)
}

// TLS Validation Context Trust for ACM Private Certificate Authority (CA).
// Experimental.
func TlsValidationTrust_Acm(certificateAuthorities *[]awsacmpca.ICertificateAuthority) TlsValidationTrust {
	_init_.Initialize()

	if err := validateTlsValidationTrust_AcmParameters(certificateAuthorities); err != nil {
		panic(err)
	}
	var returns TlsValidationTrust

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.TlsValidationTrust",
		"acm",
		[]interface{}{certificateAuthorities},
		&returns,
	)

	return returns
}

// Tells envoy where to fetch the validation context from.
// Experimental.
func TlsValidationTrust_File(certificateChain *string) MutualTlsValidationTrust {
	_init_.Initialize()

	if err := validateTlsValidationTrust_FileParameters(certificateChain); err != nil {
		panic(err)
	}
	var returns MutualTlsValidationTrust

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.TlsValidationTrust",
		"file",
		[]interface{}{certificateChain},
		&returns,
	)

	return returns
}

// TLS Validation Context Trust for Envoy' service discovery service.
// Experimental.
func TlsValidationTrust_Sds(secretName *string) MutualTlsValidationTrust {
	_init_.Initialize()

	if err := validateTlsValidationTrust_SdsParameters(secretName); err != nil {
		panic(err)
	}
	var returns MutualTlsValidationTrust

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.TlsValidationTrust",
		"sds",
		[]interface{}{secretName},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TlsValidationTrust) Bind(scope awscdk.Construct) *TlsValidationTrustConfig {
	if err := t.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *TlsValidationTrustConfig

	_jsii_.Invoke(
		t,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

