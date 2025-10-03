package awsmacie

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmacie/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomDataIdentifier.
// Experimental.
type ICustomDataIdentifierRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICustomDataIdentifierRef
type jsiiProxy_ICustomDataIdentifierRef struct {
	internal.Type__constructsIConstruct
}

