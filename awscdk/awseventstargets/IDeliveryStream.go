package awseventstargets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets/internal"
)

// Represents an Amazon Data Firehose delivery stream.
// Deprecated: Use aws_kinesisfirehose.IDeliveryStream
type IDeliveryStream interface {
	awscdk.IResource
	// The ARN of the delivery stream.
	// Deprecated: Use aws_kinesisfirehose.IDeliveryStream
	DeliveryStreamArn() *string
	// The name of the delivery stream.
	// Deprecated: Use aws_kinesisfirehose.IDeliveryStream
	DeliveryStreamName() *string
}

// The jsii proxy for IDeliveryStream
type jsiiProxy_IDeliveryStream struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IDeliveryStream) DeliveryStreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDeliveryStream) DeliveryStreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamName",
		&returns,
	)
	return returns
}

