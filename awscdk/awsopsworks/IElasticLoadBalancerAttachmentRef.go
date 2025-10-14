package awsopsworks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsopsworks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ElasticLoadBalancerAttachment.
// Experimental.
type IElasticLoadBalancerAttachmentRef interface {
	constructs.IConstruct
	// A reference to a ElasticLoadBalancerAttachment resource.
	// Experimental.
	ElasticLoadBalancerAttachmentRef() *ElasticLoadBalancerAttachmentReference
}

// The jsii proxy for IElasticLoadBalancerAttachmentRef
type jsiiProxy_IElasticLoadBalancerAttachmentRef struct {
	internal.Type__constructsIConstruct
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

