package previewawssecurityhubmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
)

// Builder for CfnHubLogsMixin to generate SECURITY_FINDING_LOGS for CfnHub.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnHubSecurityFindingLogs := awscdkmixinspreview.Mixins.NewCfnHubSecurityFindingLogs()
//
type CfnHubSecurityFindingLogs interface {
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are CWL
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnHubSecurityFindingLogsDestProps) CfnHubLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnHubSecurityFindingLogsLogGroupProps) CfnHubLogsMixin
}

// The jsii proxy struct for CfnHubSecurityFindingLogs
type jsiiProxy_CfnHubSecurityFindingLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnHubSecurityFindingLogs() CfnHubSecurityFindingLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnHubSecurityFindingLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnHubSecurityFindingLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnHubSecurityFindingLogs_Override(c CfnHubSecurityFindingLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnHubSecurityFindingLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnHubSecurityFindingLogs) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnHubSecurityFindingLogsDestProps) CfnHubLogsMixin {
	if err := c.validateToDestinationParameters(destination, props); err != nil {
		panic(err)
	}
	var returns CfnHubLogsMixin

	_jsii_.Invoke(
		c,
		"toDestination",
		[]interface{}{destination, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHubSecurityFindingLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef, props *CfnHubSecurityFindingLogsLogGroupProps) CfnHubLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup, props); err != nil {
		panic(err)
	}
	var returns CfnHubLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup, props},
		&returns,
	)

	return returns
}

