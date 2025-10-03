package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VpcLink.
// Experimental.
type IVpcLinkRef interface {
	constructs.IConstruct
}

// The jsii proxy for IVpcLinkRef
type jsiiProxy_IVpcLinkRef struct {
	internal.Type__constructsIConstruct
}

