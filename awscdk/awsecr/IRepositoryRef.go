package awsecr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Repository.
// Experimental.
type IRepositoryRef interface {
	constructs.IConstruct
	// A reference to a Repository resource.
	// Experimental.
	RepositoryRef() *RepositoryReference
}

// The jsii proxy for IRepositoryRef
type jsiiProxy_IRepositoryRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRepositoryRef) RepositoryRef() *RepositoryReference {
	var returns *RepositoryReference
	_jsii_.Get(
		j,
		"repositoryRef",
		&returns,
	)
	return returns
}

