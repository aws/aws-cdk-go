package awsssmquicksetup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsssmquicksetup/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LifecycleAutomation.
// Experimental.
type ILifecycleAutomationRef interface {
	constructs.IConstruct
	// A reference to a LifecycleAutomation resource.
	// Experimental.
	LifecycleAutomationRef() *LifecycleAutomationReference
}

// The jsii proxy for ILifecycleAutomationRef
type jsiiProxy_ILifecycleAutomationRef struct {
	internal.Type__constructsIConstruct
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

