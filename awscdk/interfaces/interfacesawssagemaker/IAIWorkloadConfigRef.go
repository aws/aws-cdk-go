package interfacesawssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AIWorkloadConfig.
// Experimental.
type IAIWorkloadConfigRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AIWorkloadConfig resource.
	// Experimental.
	AiWorkloadConfigRef() *AIWorkloadConfigReference
}

// The jsii proxy for IAIWorkloadConfigRef
type jsiiProxy_IAIWorkloadConfigRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IAIWorkloadConfigRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IAIWorkloadConfigRef) AiWorkloadConfigRef() *AIWorkloadConfigReference {
	var returns *AIWorkloadConfigReference
	_jsii_.Get(
		j,
		"aiWorkloadConfigRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAIWorkloadConfigRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAIWorkloadConfigRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

