package awscontroltower

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscontroltower/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EnabledBaseline.
// Experimental.
type IEnabledBaselineRef interface {
	constructs.IConstruct
	// A reference to a EnabledBaseline resource.
	// Experimental.
	EnabledBaselineRef() *EnabledBaselineReference
}

// The jsii proxy for IEnabledBaselineRef
type jsiiProxy_IEnabledBaselineRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IEnabledBaselineRef) EnabledBaselineRef() *EnabledBaselineReference {
	var returns *EnabledBaselineReference
	_jsii_.Get(
		j,
		"enabledBaselineRef",
		&returns,
	)
	return returns
}

