package previewawsstepfunctionsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
)

// Builder for CfnStateMachineLogsMixin to generate STANDARD_LOGS for CfnStateMachine.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStateMachineStandardLogs := awscdkmixinspreview.Mixins.NewCfnStateMachineStandardLogs()
//
type CfnStateMachineStandardLogs interface {
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are CWL
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef) CfnStateMachineLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnStateMachineLogsMixin
}

// The jsii proxy struct for CfnStateMachineStandardLogs
type jsiiProxy_CfnStateMachineStandardLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnStateMachineStandardLogs() CfnStateMachineStandardLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnStateMachineStandardLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineStandardLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnStateMachineStandardLogs_Override(c CfnStateMachineStandardLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineStandardLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnStateMachineStandardLogs) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef) CfnStateMachineLogsMixin {
	if err := c.validateToDestinationParameters(destination); err != nil {
		panic(err)
	}
	var returns CfnStateMachineLogsMixin

	_jsii_.Invoke(
		c,
		"toDestination",
		[]interface{}{destination},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStateMachineStandardLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnStateMachineLogsMixin {
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

