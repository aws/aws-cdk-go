package previewawsstepfunctionsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
)

// Builder for CfnStateMachineLogsMixin to generate EXPRESS_LOGS for CfnStateMachine.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStateMachineExpressLogs := awscdkmixinspreview.Mixins.NewCfnStateMachineExpressLogs()
//
type CfnStateMachineExpressLogs interface {
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnStateMachineLogsMixin
}

// The jsii proxy struct for CfnStateMachineExpressLogs
type jsiiProxy_CfnStateMachineExpressLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnStateMachineExpressLogs() CfnStateMachineExpressLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnStateMachineExpressLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineExpressLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnStateMachineExpressLogs_Override(c CfnStateMachineExpressLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineExpressLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnStateMachineExpressLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnStateMachineLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnStateMachineLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

