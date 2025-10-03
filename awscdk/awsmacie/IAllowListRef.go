package awsmacie

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmacie/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AllowList.
// Experimental.
type IAllowListRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAllowListRef
type jsiiProxy_IAllowListRef struct {
	internal.Type__constructsIConstruct
}

