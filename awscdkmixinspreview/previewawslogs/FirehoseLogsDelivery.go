package previewawslogs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Delivers vended logs to a Firehose Delivery Stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var deliveryStreamRef IDeliveryStreamRef
//
//   firehoseLogsDelivery := awscdkmixinspreview.Aws_logs.NewFirehoseLogsDelivery(deliveryStreamRef)
//
// Experimental.
type FirehoseLogsDelivery interface {
	ILogsDelivery
	// Binds the Firehose destination to a delivery source and creates a delivery connection between them.
	// Experimental.
	Bind(scope constructs.IConstruct, deliverySource interfacesawslogs.IDeliverySourceRef, sourceResourceArn *string) ILogsDeliveryConfig
}

// The jsii proxy struct for FirehoseLogsDelivery
type jsiiProxy_FirehoseLogsDelivery struct {
	jsiiProxy_ILogsDelivery
}

// Creates a new Firehose delivery.
// Experimental.
func NewFirehoseLogsDelivery(stream interfacesawskinesisfirehose.IDeliveryStreamRef) FirehoseLogsDelivery {
	_init_.Initialize()

	if err := validateNewFirehoseLogsDeliveryParameters(stream); err != nil {
		panic(err)
	}
	j := jsiiProxy_FirehoseLogsDelivery{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_logs.FirehoseLogsDelivery",
		[]interface{}{stream},
		&j,
	)

	return &j
}

// Creates a new Firehose delivery.
// Experimental.
func NewFirehoseLogsDelivery_Override(f FirehoseLogsDelivery, stream interfacesawskinesisfirehose.IDeliveryStreamRef) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_logs.FirehoseLogsDelivery",
		[]interface{}{stream},
		f,
	)
}

func (f *jsiiProxy_FirehoseLogsDelivery) Bind(scope constructs.IConstruct, deliverySource interfacesawslogs.IDeliverySourceRef, sourceResourceArn *string) ILogsDeliveryConfig {
	if err := f.validateBindParameters(scope, deliverySource, sourceResourceArn); err != nil {
		panic(err)
	}
	var returns ILogsDeliveryConfig

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{scope, deliverySource, sourceResourceArn},
		&returns,
	)

	return returns
}

