package awsssmquicksetup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssmquicksetup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LifecycleAutomation.
// Experimental.
type ILifecycleAutomationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a LifecycleAutomation resource.
	// Experimental.
	LifecycleAutomationRef() *LifecycleAutomationReference
}

// The jsii proxy for ILifecycleAutomationRef
type jsiiProxy_ILifecycleAutomationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ILifecycleAutomationRef) LifecycleAutomationRef() *LifecycleAutomationReference {
	var returns *LifecycleAutomationReference
	_jsii_.Get(
		j,
		"lifecycleAutomationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILifecycleAutomationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILifecycleAutomationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

