package previewawsorganizationsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnOrganizationLogsMixin to generate AUTHENTICATION_LOGS for CfnOrganization.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnOrganizationAuthenticationLogs := awscdkmixinspreview.Mixins.NewCfnOrganizationAuthenticationLogs()
//
type CfnOrganizationAuthenticationLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnOrganizationLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnOrganizationLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef) CfnOrganizationLogsMixin
}

// The jsii proxy struct for CfnOrganizationAuthenticationLogs
type jsiiProxy_CfnOrganizationAuthenticationLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnOrganizationAuthenticationLogs() CfnOrganizationAuthenticationLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnOrganizationAuthenticationLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_organizations.mixins.CfnOrganizationAuthenticationLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnOrganizationAuthenticationLogs_Override(c CfnOrganizationAuthenticationLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_organizations.mixins.CfnOrganizationAuthenticationLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnOrganizationAuthenticationLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnOrganizationLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnOrganizationLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnOrganizationAuthenticationLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnOrganizationLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnOrganizationLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnOrganizationAuthenticationLogs) ToS3(bucket interfacesawss3.IBucketRef) CfnOrganizationLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnOrganizationLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket},
		&returns,
	)

	return returns
}

