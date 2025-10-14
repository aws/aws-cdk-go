package awsecr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RepositoryCreationTemplate.
// Experimental.
type IRepositoryCreationTemplateRef interface {
	constructs.IConstruct
	// A reference to a RepositoryCreationTemplate resource.
	// Experimental.
	RepositoryCreationTemplateRef() *RepositoryCreationTemplateReference
}

// The jsii proxy for IRepositoryCreationTemplateRef
type jsiiProxy_IRepositoryCreationTemplateRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRepositoryCreationTemplateRef) RepositoryCreationTemplateRef() *RepositoryCreationTemplateReference {
	var returns *RepositoryCreationTemplateReference
	_jsii_.Get(
		j,
		"repositoryCreationTemplateRef",
		&returns,
	)
	return returns
}

