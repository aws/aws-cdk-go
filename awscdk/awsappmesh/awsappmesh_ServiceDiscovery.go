package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsservicediscovery"
)

// Provides the Service Discovery method a VirtualNode uses.
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
type ServiceDiscovery interface {
	// Binds the current object when adding Service Discovery to a VirtualNode.
	// Experimental.
	Bind(scope awscdk.Construct) *ServiceDiscoveryConfig
}

// The jsii proxy struct for ServiceDiscovery
type jsiiProxy_ServiceDiscovery struct {
	_ byte // padding
}

// Experimental.
func NewServiceDiscovery_Override(s ServiceDiscovery) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.ServiceDiscovery",
		nil, // no parameters
		s,
	)
}

// Returns Cloud Map based service discovery.
// Experimental.
func ServiceDiscovery_CloudMap(service awsservicediscovery.IService, instanceAttributes *map[string]*string) ServiceDiscovery {
	_init_.Initialize()

	if err := validateServiceDiscovery_CloudMapParameters(service); err != nil {
		panic(err)
	}
	var returns ServiceDiscovery

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.ServiceDiscovery",
		"cloudMap",
		[]interface{}{service, instanceAttributes},
		&returns,
	)

	return returns
}

// Returns DNS based service discovery.
// Experimental.
func ServiceDiscovery_Dns(hostname *string, responseType DnsResponseType) ServiceDiscovery {
	_init_.Initialize()

	if err := validateServiceDiscovery_DnsParameters(hostname); err != nil {
		panic(err)
	}
	var returns ServiceDiscovery

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.ServiceDiscovery",
		"dns",
		[]interface{}{hostname, responseType},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceDiscovery) Bind(scope awscdk.Construct) *ServiceDiscoveryConfig {
	if err := s.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *ServiceDiscoveryConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

