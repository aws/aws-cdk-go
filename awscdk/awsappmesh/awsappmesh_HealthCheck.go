package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// Contains static factory methods for creating health checks for different protocols.
//
// Example:
//   var mesh mesh
//   vpc := ec2.NewVpc(this, jsii.String("vpc"))
//   namespace := cloudmap.NewPrivateDnsNamespace(this, jsii.String("test-namespace"), &PrivateDnsNamespaceProps{
//   	Vpc: Vpc,
//   	Name: jsii.String("domain.local"),
//   })
//   service := namespace.CreateService(jsii.String("Svc"))
//   node := mesh.addVirtualNode(jsii.String("virtual-node"), &VirtualNodeBaseProps{
//   	ServiceDiscovery: appmesh.ServiceDiscovery_CloudMap(service),
//   	Listeners: []virtualNodeListener{
//   		appmesh.*virtualNodeListener_Http(&HttpVirtualNodeListenerOptions{
//   			Port: jsii.Number(8081),
//   			HealthCheck: appmesh.HealthCheck_Http(&HttpHealthCheckOptions{
//   				HealthyThreshold: jsii.Number(3),
//   				Interval: cdk.Duration_Seconds(jsii.Number(5)),
//   				 // minimum
//   				Path: jsii.String("/health-check-path"),
//   				Timeout: cdk.Duration_*Seconds(jsii.Number(2)),
//   				 // minimum
//   				UnhealthyThreshold: jsii.Number(2),
//   			}),
//   		}),
//   	},
//   	AccessLog: appmesh.AccessLog_FromFilePath(jsii.String("/dev/stdout")),
//   })
//
// Experimental.
type HealthCheck interface {
	// Called when the AccessLog type is initialized.
	//
	// Can be used to enforce
	// mutual exclusivity with future properties.
	// Experimental.
	Bind(scope awscdk.Construct, options *HealthCheckBindOptions) *HealthCheckConfig
}

// The jsii proxy struct for HealthCheck
type jsiiProxy_HealthCheck struct {
	_ byte // padding
}

// Experimental.
func NewHealthCheck_Override(h HealthCheck) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.HealthCheck",
		nil, // no parameters
		h,
	)
}

// Construct a GRPC health check.
// Experimental.
func HealthCheck_Grpc(options *GrpcHealthCheckOptions) HealthCheck {
	_init_.Initialize()

	if err := validateHealthCheck_GrpcParameters(options); err != nil {
		panic(err)
	}
	var returns HealthCheck

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HealthCheck",
		"grpc",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Construct a HTTP health check.
// Experimental.
func HealthCheck_Http(options *HttpHealthCheckOptions) HealthCheck {
	_init_.Initialize()

	if err := validateHealthCheck_HttpParameters(options); err != nil {
		panic(err)
	}
	var returns HealthCheck

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HealthCheck",
		"http",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Construct a HTTP2 health check.
// Experimental.
func HealthCheck_Http2(options *HttpHealthCheckOptions) HealthCheck {
	_init_.Initialize()

	if err := validateHealthCheck_Http2Parameters(options); err != nil {
		panic(err)
	}
	var returns HealthCheck

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HealthCheck",
		"http2",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Construct a TCP health check.
// Experimental.
func HealthCheck_Tcp(options *TcpHealthCheckOptions) HealthCheck {
	_init_.Initialize()

	if err := validateHealthCheck_TcpParameters(options); err != nil {
		panic(err)
	}
	var returns HealthCheck

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HealthCheck",
		"tcp",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthCheck) Bind(scope awscdk.Construct, options *HealthCheckBindOptions) *HealthCheckConfig {
	if err := h.validateBindParameters(scope, options); err != nil {
		panic(err)
	}
	var returns *HealthCheckConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

