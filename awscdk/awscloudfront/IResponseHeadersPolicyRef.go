package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResponseHeadersPolicy.
// Experimental.
type IResponseHeadersPolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IResponseHeadersPolicyRef
type jsiiProxy_IResponseHeadersPolicyRef struct {
	internal.Type__constructsIConstruct
}

