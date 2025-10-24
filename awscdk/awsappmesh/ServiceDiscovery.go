package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicediscovery"
	"github.com/aws/constructs-go/constructs/v10"
)

// Provides the Service Discovery method a VirtualNode uses.
//
// Example:
//   var mesh Mesh
//   vpc := ec2.NewVpc(this, jsii.String("vpc"))
//   namespace := cloudmap.NewPrivateDnsNamespace(this, jsii.String("test-namespace"), &PrivateDnsNamespaceProps{
//   	Vpc: Vpc,
//   	Name: jsii.String("domain.local"),
//   })
//   service := namespace.CreateService(jsii.String("Svc"))
//   node := mesh.addVirtualNode(jsii.String("virtual-node"), &VirtualNodeBaseProps{
//   	ServiceDiscovery: appmesh.ServiceDiscovery_CloudMap(service),
//   	Listeners: []VirtualNodeListener{
//   		appmesh.VirtualNodeListener_Http(&HttpVirtualNodeListenerOptions{
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
type ServiceDiscovery interface {
	// Binds the current object when adding Service Discovery to a VirtualNode.
	Bind(scope constructs.Construct) *ServiceDiscoveryConfig
}

// The jsii proxy struct for ServiceDiscovery
type jsiiProxy_ServiceDiscovery struct {
	_ byte // padding
}

func NewServiceDiscovery_Override(s ServiceDiscovery) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.ServiceDiscovery",
		nil, // no parameters
		s,
	)
}

// Returns Cloud Map based service discovery.
func ServiceDiscovery_CloudMap(service awsservicediscovery.IService, instanceAttributes *map[string]*string, ipPreference IpPreference) ServiceDiscovery {
	_init_.Initialize()

	if err := validateServiceDiscovery_CloudMapParameters(service); err != nil {
		panic(err)
	}
	var returns ServiceDiscovery

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.ServiceDiscovery",
		"cloudMap",
		[]interface{}{service, instanceAttributes, ipPreference},
		&returns,
	)

	return returns
}

// Returns DNS based service discovery.
func ServiceDiscovery_Dns(hostname *string, responseType DnsResponseType, ipPreference IpPreference) ServiceDiscovery {
	_init_.Initialize()

	if err := validateServiceDiscovery_DnsParameters(hostname); err != nil {
		panic(err)
	}
	var returns ServiceDiscovery

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.ServiceDiscovery",
		"dns",
		[]interface{}{hostname, responseType, ipPreference},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceDiscovery) Bind(scope constructs.Construct) *ServiceDiscoveryConfig {
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

