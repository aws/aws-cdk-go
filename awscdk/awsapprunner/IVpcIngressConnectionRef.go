package awsapprunner

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapprunner/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VpcIngressConnection.
// Experimental.
type IVpcIngressConnectionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IVpcIngressConnectionRef
type jsiiProxy_IVpcIngressConnectionRef struct {
	internal.Type__constructsIConstruct
}

