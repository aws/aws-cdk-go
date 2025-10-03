package awsmpa

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmpa/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IdentitySource.
// Experimental.
type IIdentitySourceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IIdentitySourceRef
type jsiiProxy_IIdentitySourceRef struct {
	internal.Type__constructsIConstruct
}

