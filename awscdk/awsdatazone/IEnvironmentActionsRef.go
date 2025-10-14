package awsdatazone

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EnvironmentActions.
// Experimental.
type IEnvironmentActionsRef interface {
	constructs.IConstruct
	// A reference to a EnvironmentActions resource.
	// Experimental.
	EnvironmentActionsRef() *EnvironmentActionsReference
}

// The jsii proxy for IEnvironmentActionsRef
type jsiiProxy_IEnvironmentActionsRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IEnvironmentActionsRef) EnvironmentActionsRef() *EnvironmentActionsReference {
	var returns *EnvironmentActionsReference
	_jsii_.Get(
		j,
		"environmentActionsRef",
		&returns,
	)
	return returns
}

