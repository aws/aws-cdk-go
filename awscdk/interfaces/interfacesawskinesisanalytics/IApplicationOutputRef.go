package interfacesawskinesisanalytics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisanalytics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationOutput.
// Experimental.
type IApplicationOutputRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ApplicationOutput resource.
	// Experimental.
	ApplicationOutputRef() *ApplicationOutputReference
}

// The jsii proxy for IApplicationOutputRef
type jsiiProxy_IApplicationOutputRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IApplicationOutputRef) ApplicationOutputRef() *ApplicationOutputReference {
	var returns *ApplicationOutputReference
	_jsii_.Get(
		j,
		"applicationOutputRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationOutputRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationOutputRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

