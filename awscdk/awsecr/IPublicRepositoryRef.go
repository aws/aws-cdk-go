package awsecr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PublicRepository.
// Experimental.
type IPublicRepositoryRef interface {
	constructs.IConstruct
	// A reference to a PublicRepository resource.
	// Experimental.
	PublicRepositoryRef() *PublicRepositoryReference
}

// The jsii proxy for IPublicRepositoryRef
type jsiiProxy_IPublicRepositoryRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IPublicRepositoryRef) PublicRepositoryRef() *PublicRepositoryReference {
	var returns *PublicRepositoryReference
	_jsii_.Get(
		j,
		"publicRepositoryRef",
		&returns,
	)
	return returns
}

