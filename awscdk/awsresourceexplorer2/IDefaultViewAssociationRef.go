package awsresourceexplorer2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsresourceexplorer2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DefaultViewAssociation.
// Experimental.
type IDefaultViewAssociationRef interface {
	constructs.IConstruct
	// A reference to a DefaultViewAssociation resource.
	// Experimental.
	DefaultViewAssociationRef() *DefaultViewAssociationReference
}

// The jsii proxy for IDefaultViewAssociationRef
type jsiiProxy_IDefaultViewAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDefaultViewAssociationRef) DefaultViewAssociationRef() *DefaultViewAssociationReference {
	var returns *DefaultViewAssociationReference
	_jsii_.Get(
		j,
		"defaultViewAssociationRef",
		&returns,
	)
	return returns
}

