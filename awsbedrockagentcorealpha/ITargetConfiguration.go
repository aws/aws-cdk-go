package awsbedrockagentcorealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Base interface for target configurations.
// Experimental.
type ITargetConfiguration interface {
	// Binds the configuration to a construct scope Sets up permissions and dependencies.
	// Experimental.
	Bind(scope constructs.Construct, gateway IGateway) *TargetConfigurationConfig
	// The target type.
	// Experimental.
	TargetType() McpTargetType
}

// The jsii proxy for ITargetConfiguration
type jsiiProxy_ITargetConfiguration struct {
	_ byte // padding
}

func (i *jsiiProxy_ITargetConfiguration) Bind(scope constructs.Construct, gateway IGateway) *TargetConfigurationConfig {
	if err := i.validateBindParameters(scope, gateway); err != nil {
		panic(err)
	}
	var returns *TargetConfigurationConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, gateway},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ITargetConfiguration) TargetType() McpTargetType {
	var returns McpTargetType
	_jsii_.Get(
		j,
		"targetType",
		&returns,
	)
	return returns
}

