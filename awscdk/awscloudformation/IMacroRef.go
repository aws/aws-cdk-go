package awscloudformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Macro.
// Experimental.
type IMacroRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Macro resource.
	// Experimental.
	MacroRef() *MacroReference
}

// The jsii proxy for IMacroRef
type jsiiProxy_IMacroRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IMacroRef) MacroRef() *MacroReference {
	var returns *MacroReference
	_jsii_.Get(
		j,
		"macroRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMacroRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMacroRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

