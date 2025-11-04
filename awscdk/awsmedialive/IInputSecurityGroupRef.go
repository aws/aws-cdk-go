package awsmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InputSecurityGroup.
// Experimental.
type IInputSecurityGroupRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a InputSecurityGroup resource.
	// Experimental.
	InputSecurityGroupRef() *InputSecurityGroupReference
}

// The jsii proxy for IInputSecurityGroupRef
type jsiiProxy_IInputSecurityGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IInputSecurityGroupRef) InputSecurityGroupRef() *InputSecurityGroupReference {
	var returns *InputSecurityGroupReference
	_jsii_.Get(
		j,
		"inputSecurityGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInputSecurityGroupRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInputSecurityGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

