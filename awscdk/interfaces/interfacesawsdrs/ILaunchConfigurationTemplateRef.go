package interfacesawsdrs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdrs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LaunchConfigurationTemplate.
// Experimental.
type ILaunchConfigurationTemplateRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a LaunchConfigurationTemplate resource.
	// Experimental.
	LaunchConfigurationTemplateRef() *LaunchConfigurationTemplateReference
}

// The jsii proxy for ILaunchConfigurationTemplateRef
type jsiiProxy_ILaunchConfigurationTemplateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ILaunchConfigurationTemplateRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ILaunchConfigurationTemplateRef) LaunchConfigurationTemplateRef() *LaunchConfigurationTemplateReference {
	var returns *LaunchConfigurationTemplateReference
	_jsii_.Get(
		j,
		"launchConfigurationTemplateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILaunchConfigurationTemplateRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILaunchConfigurationTemplateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

