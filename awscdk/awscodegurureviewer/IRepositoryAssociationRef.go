package awscodegurureviewer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodegurureviewer/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RepositoryAssociation.
// Experimental.
type IRepositoryAssociationRef interface {
	constructs.IConstruct
	// A reference to a RepositoryAssociation resource.
	// Experimental.
	RepositoryAssociationRef() *RepositoryAssociationReference
}

// The jsii proxy for IRepositoryAssociationRef
type jsiiProxy_IRepositoryAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRepositoryAssociationRef) RepositoryAssociationRef() *RepositoryAssociationReference {
	var returns *RepositoryAssociationReference
	_jsii_.Get(
		j,
		"repositoryAssociationRef",
		&returns,
	)
	return returns
}

