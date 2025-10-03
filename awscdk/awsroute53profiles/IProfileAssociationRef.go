package awsroute53profiles

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53profiles/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProfileAssociation.
// Experimental.
type IProfileAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IProfileAssociationRef
type jsiiProxy_IProfileAssociationRef struct {
	internal.Type__constructsIConstruct
}

