package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Configuration for Envoy Access logs for mesh endpoints.
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
//   				Interval: awscdk.Duration_Seconds(jsii.Number(5)),
//   				 // minimum
//   				Path: jsii.String("/health-check-path"),
//   				Timeout: awscdk.Duration_*Seconds(jsii.Number(2)),
//   				 // minimum
//   				UnhealthyThreshold: jsii.Number(2),
//   			}),
//   		}),
//   	},
//   	AccessLog: appmesh.AccessLog_FromFilePath(jsii.String("/dev/stdout")),
//   })
//
type AccessLog interface {
	// Called when the AccessLog type is initialized.
	//
	// Can be used to enforce
	// mutual exclusivity with future properties.
	Bind(scope constructs.Construct) *AccessLogConfig
}

// The jsii proxy struct for AccessLog
type jsiiProxy_AccessLog struct {
	_ byte // padding
}

func NewAccessLog_Override(a AccessLog) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.AccessLog",
		nil, // no parameters
		a,
	)
}

// Path to a file to write access logs to.
func AccessLog_FromFilePath(filePath *string) AccessLog {
	_init_.Initialize()

	if err := validateAccessLog_FromFilePathParameters(filePath); err != nil {
		panic(err)
	}
	var returns AccessLog

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.AccessLog",
		"fromFilePath",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessLog) Bind(scope constructs.Construct) *AccessLogConfig {
	if err := a.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *AccessLogConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

