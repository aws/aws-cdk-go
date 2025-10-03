package awssigner

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssigner/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SigningProfile.
// Experimental.
type ISigningProfileRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISigningProfileRef
type jsiiProxy_ISigningProfileRef struct {
	internal.Type__constructsIConstruct
}

