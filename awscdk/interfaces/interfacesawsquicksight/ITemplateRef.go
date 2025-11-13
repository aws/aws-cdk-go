package interfacesawsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Template.
// Experimental.
type ITemplateRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Template resource.
	// Experimental.
	TemplateRef() *TemplateReference
}

// The jsii proxy for ITemplateRef
type jsiiProxy_ITemplateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ITemplateRef) TemplateRef() *TemplateReference {
	var returns *TemplateReference
	_jsii_.Get(
		j,
		"templateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITemplateRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITemplateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

