package awsecr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RepositoryCreationTemplate.
// Experimental.
type IRepositoryCreationTemplateRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a RepositoryCreationTemplate resource.
	// Experimental.
	RepositoryCreationTemplateRef() *RepositoryCreationTemplateReference
}

// The jsii proxy for IRepositoryCreationTemplateRef
type jsiiProxy_IRepositoryCreationTemplateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IRepositoryCreationTemplateRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRepositoryCreationTemplateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

