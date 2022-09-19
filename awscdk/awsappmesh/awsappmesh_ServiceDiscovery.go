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
//   // Example automatically generated from non-compiling source. May contain errors.
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

