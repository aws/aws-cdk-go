package awsservicediscovery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicediscovery/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsservicediscovery"
	"github.com/aws/constructs-go/constructs/v10"
)

type IService interface {
	awscdk.IResource
	interfacesawsservicediscovery.IServiceRef
	// The discovery type used by the service.
	DiscoveryType() DiscoveryType
	// The DnsRecordType used by the service.
	DnsRecordType() DnsRecordType
	// The namespace for the Cloudmap Service.
	Namespace() INamespace
	// The Routing Policy used by the service.
	RoutingPolicy() RoutingPolicy
	// The Arn of the namespace that you want to use for DNS configuration.
	ServiceArn() *string
	// The ID of the namespace that you want to use for DNS configuration.
	ServiceId() *string
	// A name for the Cloudmap Service.
	ServiceName() *string
}

// The jsii proxy for IService
type jsiiProxy_IService struct {
	internal.Type__awscdkIResource
	internal.Type__interfacesawsservicediscoveryIServiceRef
}

func (i *jsiiProxy_IService) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IService) DiscoveryType() DiscoveryType {
	var returns DiscoveryType
	_jsii_.Get(
		j,
		"discoveryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IService) DnsRecordType() DnsRecordType {
	var returns DnsRecordType
	_jsii_.Get(
		j,
		"dnsRecordType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IService) Namespace() INamespace {
	var returns INamespace
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IService) RoutingPolicy() RoutingPolicy {
	var returns RoutingPolicy
	_jsii_.Get(
		j,
		"routingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IService) ServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IService) ServiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IService) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IService) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IService) ServiceRef() *interfacesawsservicediscovery.ServiceReference {
	var returns *interfacesawsservicediscovery.ServiceReference
	_jsii_.Get(
		j,
		"serviceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IService) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

