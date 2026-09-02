package previewawssecurityhubmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
)

// Builder for CfnHubV2LogsMixin to generate SECURITY_FINDING_LOGS for CfnHubV2.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnHubV2SecurityFindingLogs := awscdkmixinspreview.Mixins.NewCfnHubV2SecurityFindingLogs()
//
type CfnHubV2SecurityFindingLogs interface {
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are CWL
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnHubV2SecurityFindingLogsDestProps) CfnHubV2LogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnHubV2SecurityFindingLogsLogGroupProps) CfnHubV2LogsMixin
}

// The jsii proxy struct for CfnHubV2SecurityFindingLogs
type jsiiProxy_CfnHubV2SecurityFindingLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnHubV2SecurityFindingLogs() CfnHubV2SecurityFindingLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnHubV2SecurityFindingLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnHubV2SecurityFindingLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnHubV2SecurityFindingLogs_Override(c CfnHubV2SecurityFindingLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnHubV2SecurityFindingLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnHubV2SecurityFindingLogs) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnHubV2SecurityFindingLogsDestProps) CfnHubV2LogsMixin {
	if err := c.validateToDestinationParameters(destination, props); err != nil {
		panic(err)
	}
	var returns CfnHubV2LogsMixin

	_jsii_.Invoke(
		c,
		"toDestination",
		[]interface{}{destination, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHubV2SecurityFindingLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnHubV2SecurityFindingLogsLogGroupProps) CfnHubV2LogsMixin {
	if err := c.validateToLogGroupParameters(logGroup, props); err != nil {
		panic(err)
	}
	var returns CfnHubV2LogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup, props},
		&returns,
	)

	return returns
}

