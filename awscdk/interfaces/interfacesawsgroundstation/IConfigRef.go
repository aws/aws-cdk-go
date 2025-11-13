package interfacesawsgroundstation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsgroundstation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Config.
// Experimental.
type IConfigRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Config resource.
	// Experimental.
	ConfigRef() *ConfigReference
}

// The jsii proxy for IConfigRef
type jsiiProxy_IConfigRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IConfigRef) ConfigRef() *ConfigReference {
	var returns *ConfigReference
	_jsii_.Get(
		j,
		"configRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfigRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfigRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

