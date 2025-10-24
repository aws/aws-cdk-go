package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Configuration for Envoy Access Logging Format for mesh endpoints.
//
// Example:
//   var mesh Mesh
//   var service Service
//
//   node := appmesh.NewVirtualNode(this, jsii.String("node"), &VirtualNodeProps{
//   	Mesh: Mesh,
//   	ServiceDiscovery: appmesh.ServiceDiscovery_CloudMap(service),
//   	Listeners: []VirtualNodeListener{
//   		appmesh.VirtualNodeListener_Http(&HttpVirtualNodeListenerOptions{
//   			Port: jsii.Number(8080),
//   			HealthCheck: appmesh.HealthCheck_Http(&HttpHealthCheckOptions{
//   				HealthyThreshold: jsii.Number(3),
//   				Interval: cdk.Duration_Seconds(jsii.Number(5)),
//   				Path: jsii.String("/ping"),
//   				Timeout: cdk.Duration_*Seconds(jsii.Number(2)),
//   				UnhealthyThreshold: jsii.Number(2),
//   			}),
//   			Timeout: &HttpTimeout{
//   				Idle: cdk.Duration_*Seconds(jsii.Number(5)),
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
//   	AccessLog: appmesh.AccessLog_FromFilePath(jsii.String("/dev/stdout"), appmesh.LoggingFormat_FromJson(map[string]*string{
//   		"testKey1": jsii.String("testValue1"),
//   		"testKey2": jsii.String("testValue2"),
//   	})),
//   })
//
type LoggingFormat interface {
	// Called when the Access Log Format is initialized.
	//
	// Can be used to enforce
	// mutual exclusivity with future properties.
	Bind() *LoggingFormatConfig
}

// The jsii proxy struct for LoggingFormat
type jsiiProxy_LoggingFormat struct {
	_ byte // padding
}

func NewLoggingFormat_Override(l LoggingFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.LoggingFormat",
		nil, // no parameters
		l,
	)
}

// Generate logging format from json key pairs.
func LoggingFormat_FromJson(jsonLoggingFormat *map[string]*string) LoggingFormat {
	_init_.Initialize()

	if err := validateLoggingFormat_FromJsonParameters(jsonLoggingFormat); err != nil {
		panic(err)
	}
	var returns LoggingFormat

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.LoggingFormat",
		"fromJson",
		[]interface{}{jsonLoggingFormat},
		&returns,
	)

	return returns
}

// Generate logging format from text pattern.
func LoggingFormat_FromText(text *string) LoggingFormat {
	_init_.Initialize()

	if err := validateLoggingFormat_FromTextParameters(text); err != nil {
		panic(err)
	}
	var returns LoggingFormat

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.LoggingFormat",
		"fromText",
		[]interface{}{text},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LoggingFormat) Bind() *LoggingFormatConfig {
	var returns *LoggingFormatConfig

	_jsii_.Invoke(
		l,
		"bind",
		nil, // no parameters
		&returns,
	)

	return returns
}

