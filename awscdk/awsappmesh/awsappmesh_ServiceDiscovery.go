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
//   namespace := cloudmap.NewPrivateDnsNamespace(this, jsii.String("test-namespace"), &privateDnsNamespaceProps{
//   	vpc: vpc,
//   	name: jsii.String("domain.local"),
//   })
//   service := namespace.createService(jsii.String("Svc"))
//   node := mesh.addVirtualNode(jsii.String("virtual-node"), &virtualNodeBaseProps{
//   	serviceDiscovery: appmesh.serviceDiscovery.cloudMap(service),
//   	listeners: []virtualNodeListener{
//   		appmesh.*virtualNodeListener.http(&httpVirtualNodeListenerOptions{
//   			port: jsii.Number(8081),
//   			healthCheck: appmesh.healthCheck.http(&httpHealthCheckOptions{
//   				healthyThreshold: jsii.Number(3),
//   				interval: cdk.duration.seconds(jsii.Number(5)),
//   				 // minimum
//   				path: jsii.String("/health-check-path"),
//   				timeout: cdk.*duration.seconds(jsii.Number(2)),
//   				 // minimum
//   				unhealthyThreshold: jsii.Number(2),
//   			}),
//   		}),
//   	},
//   	accessLog: appmesh.accessLog.fromFilePath(jsii.String("/dev/stdout")),
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

