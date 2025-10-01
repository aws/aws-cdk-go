package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for configuring alternate target groups for blue/green deployments.
type IAlternateTarget interface {
	// Bind this configuration to a service.
	//
	// Returns: The configuration to apply to the service.
	Bind(scope constructs.IConstruct) *AlternateTargetConfig
}

// The jsii proxy for IAlternateTarget
type jsiiProxy_IAlternateTarget struct {
	_ byte // padding
}

func (i *jsiiProxy_IAlternateTarget) Bind(scope constructs.IConstruct) *AlternateTargetConfig {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *AlternateTargetConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

