package awsresourceexplorer2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsresourceexplorer2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DefaultViewAssociation.
// Experimental.
type IDefaultViewAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDefaultViewAssociationRef
type jsiiProxy_IDefaultViewAssociationRef struct {
	internal.Type__constructsIConstruct
}

