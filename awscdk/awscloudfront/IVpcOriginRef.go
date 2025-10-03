package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VpcOrigin.
// Experimental.
type IVpcOriginRef interface {
	constructs.IConstruct
}

// The jsii proxy for IVpcOriginRef
type jsiiProxy_IVpcOriginRef struct {
	internal.Type__constructsIConstruct
}

