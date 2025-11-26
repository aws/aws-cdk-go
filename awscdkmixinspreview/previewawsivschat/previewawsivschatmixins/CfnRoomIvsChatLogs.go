package previewawsivschatmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// Builder for CfnRoomLogsMixin to generate IVS_CHAT_LOGS for CfnRoom.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRoomIvsChatLogs := awscdkmixinspreview.Mixins.NewCfnRoomIvsChatLogs()
//
type CfnRoomIvsChatLogs interface {
	// Send logs to a Firehose Delivery Stream.
	ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnRoomLogsMixin
	// Send logs to a CloudWatch Log Group.
	ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnRoomLogsMixin
	// Send logs to an S3 Bucket.
	ToS3(bucket interfacesawss3.IBucketRef) CfnRoomLogsMixin
}

// The jsii proxy struct for CfnRoomIvsChatLogs
type jsiiProxy_CfnRoomIvsChatLogs struct {
	_ byte // padding
}

// Experimental.
func NewCfnRoomIvsChatLogs() CfnRoomIvsChatLogs {
	_init_.Initialize()

	j := jsiiProxy_CfnRoomIvsChatLogs{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ivschat.mixins.CfnRoomIvsChatLogs",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnRoomIvsChatLogs_Override(c CfnRoomIvsChatLogs) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ivschat.mixins.CfnRoomIvsChatLogs",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnRoomIvsChatLogs) ToFirehose(deliveryStream interfacesawskinesisfirehose.IDeliveryStreamRef) CfnRoomLogsMixin {
	if err := c.validateToFirehoseParameters(deliveryStream); err != nil {
		panic(err)
	}
	var returns CfnRoomLogsMixin

	_jsii_.Invoke(
		c,
		"toFirehose",
		[]interface{}{deliveryStream},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRoomIvsChatLogs) ToLogGroup(logGroup interfacesawslogs.ILogGroupRef) CfnRoomLogsMixin {
	if err := c.validateToLogGroupParameters(logGroup); err != nil {
		panic(err)
	}
	var returns CfnRoomLogsMixin

	_jsii_.Invoke(
		c,
		"toLogGroup",
		[]interface{}{logGroup},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRoomIvsChatLogs) ToS3(bucket interfacesawss3.IBucketRef) CfnRoomLogsMixin {
	if err := c.validateToS3Parameters(bucket); err != nil {
		panic(err)
	}
	var returns CfnRoomLogsMixin

	_jsii_.Invoke(
		c,
		"toS3",
		[]interface{}{bucket},
		&returns,
	)

	return returns
}

