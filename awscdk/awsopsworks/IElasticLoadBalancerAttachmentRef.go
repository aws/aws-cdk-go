package awsopsworks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopsworks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ElasticLoadBalancerAttachment.
// Experimental.
type IElasticLoadBalancerAttachmentRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ElasticLoadBalancerAttachment resource.
	// Experimental.
	ElasticLoadBalancerAttachmentRef() *ElasticLoadBalancerAttachmentReference
}

// The jsii proxy for IElasticLoadBalancerAttachmentRef
type jsiiProxy_IElasticLoadBalancerAttachmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IElasticLoadBalancerAttachmentRef) ElasticLoadBalancerAttachmentRef() *ElasticLoadBalancerAttachmentReference {
	var returns *ElasticLoadBalancerAttachmentReference
	_jsii_.Get(
		j,
		"elasticLoadBalancerAttachmentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IElasticLoadBalancerAttachmentRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IElasticLoadBalancerAttachmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

