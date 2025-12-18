package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsacmpca"
	"github.com/aws/constructs-go/constructs/v10"
)

// Defines the TLS Validation Context Trust.
//
// Example:
//   var mesh Mesh
//   var service Service
//
//
//   node := appmesh.NewVirtualNode(this, jsii.String("node"), &VirtualNodeProps{
//   	Mesh: Mesh,
//   	ServiceDiscovery: appmesh.ServiceDiscovery_CloudMap(service),
//   	Listeners: []VirtualNodeListener{
//   		appmesh.VirtualNodeListener_Http(&HttpVirtualNodeListenerOptions{
//   			Port: jsii.Number(8080),
//   			HealthCheck: appmesh.HealthCheck_Http(&HttpHealthCheckOptions{
//   				HealthyThreshold: jsii.Number(3),
//   				Interval: awscdk.Duration_Seconds(jsii.Number(5)),
//   				Path: jsii.String("/ping"),
//   				Timeout: awscdk.Duration_*Seconds(jsii.Number(2)),
//   				UnhealthyThreshold: jsii.Number(2),
//   			}),
//   			Timeout: &HttpTimeout{
//   				Idle: awscdk.Duration_*Seconds(jsii.Number(5)),
//   			},
//   		}),
//   	},
//   	BackendDefaults: &BackendDefaults{
//   		TlsClientPolicy: &TlsClientPolicy{
//   			Validation: &TlsValidation{
//   				Trust: appmesh.TlsValidationTrust_File(jsii.String("/keys/local_cert_chain.pem")),
//   			},
//   		},
//   	},
//   	AccessLog: appmesh.AccessLog_FromFilePath(jsii.String("/dev/stdout")),
//   })
//
//   cdk.Tags_Of(node).Add(jsii.String("Environment"), jsii.String("Dev"))
//
type TlsValidationTrust interface {
	// Returns Trust context based on trust type.
	Bind(scope constructs.Construct) *TlsValidationTrustConfig
}

// The jsii proxy struct for TlsValidationTrust
type jsiiProxy_TlsValidationTrust struct {
	_ byte // padding
}

func NewTlsValidationTrust_Override(t TlsValidationTrust) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.TlsValidationTrust",
		nil, // no parameters
		t,
	)
}

// TLS Validation Context Trust for ACM Private Certificate Authority (CA).
func TlsValidationTrust_Acm(certificateAuthorities *[]interfacesawsacmpca.ICertificateAuthorityRef) TlsValidationTrust {
	_init_.Initialize()

	if err := validateTlsValidationTrust_AcmParameters(certificateAuthorities); err != nil {
		panic(err)
	}
	var returns TlsValidationTrust

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.TlsValidationTrust",
		"acm",
		[]interface{}{certificateAuthorities},
		&returns,
	)

	return returns
}

// Tells envoy where to fetch the validation context from.
func TlsValidationTrust_File(certificateChain *string) MutualTlsValidationTrust {
	_init_.Initialize()

	if err := validateTlsValidationTrust_FileParameters(certificateChain); err != nil {
		panic(err)
	}
	var returns MutualTlsValidationTrust

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.TlsValidationTrust",
		"file",
		[]interface{}{certificateChain},
		&returns,
	)

	return returns
}

// TLS Validation Context Trust for Envoy' service discovery service.
func TlsValidationTrust_Sds(secretName *string) MutualTlsValidationTrust {
	_init_.Initialize()

	if err := validateTlsValidationTrust_SdsParameters(secretName); err != nil {
		panic(err)
	}
	var returns MutualTlsValidationTrust

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.TlsValidationTrust",
		"sds",
		[]interface{}{secretName},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TlsValidationTrust) Bind(scope constructs.Construct) *TlsValidationTrustConfig {
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

