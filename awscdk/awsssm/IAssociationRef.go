package awsssm

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Association.
// Experimental.
type IAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAssociationRef
type jsiiProxy_IAssociationRef struct {
	internal.Type__constructsIConstruct
}

