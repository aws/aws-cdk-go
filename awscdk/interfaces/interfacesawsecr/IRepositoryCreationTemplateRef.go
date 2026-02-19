package interfacesawsecr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsecr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RepositoryCreationTemplate.
// Experimental.
type IRepositoryCreationTemplateRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a RepositoryCreationTemplate resource.
	// Experimental.
	RepositoryCreationTemplateRef() *RepositoryCreationTemplateReference
}

// The jsii proxy for IRepositoryCreationTemplateRef
type jsiiProxy_IRepositoryCreationTemplateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IRepositoryCreationTemplateRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_IRepositoryCreationTemplateRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
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

