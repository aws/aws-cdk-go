package interfacesawscloud9

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloud9/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EnvironmentEC2.
// Experimental.
type IEnvironmentEC2Ref interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a EnvironmentEC2 resource.
	// Experimental.
	EnvironmentEc2Ref() *EnvironmentEC2Reference
}

// The jsii proxy for IEnvironmentEC2Ref
type jsiiProxy_IEnvironmentEC2Ref struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IEnvironmentEC2Ref) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IEnvironmentEC2Ref) EnvironmentEc2Ref() *EnvironmentEC2Reference {
	var returns *EnvironmentEC2Reference
	_jsii_.Get(
		j,
		"environmentEc2Ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironmentEC2Ref) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironmentEC2Ref) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

