package awsopsworks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopsworks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ElasticLoadBalancerAttachment.
// Experimental.
type IElasticLoadBalancerAttachmentRef interface {
	constructs.IConstruct
}

// The jsii proxy for IElasticLoadBalancerAttachmentRef
type jsiiProxy_IElasticLoadBalancerAttachmentRef struct {
	internal.Type__constructsIConstruct
}

