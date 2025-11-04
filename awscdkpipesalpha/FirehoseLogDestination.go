package awscdkpipesalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkpipesalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
)

// Firehose stream for delivery of pipe logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import pipes_alpha "github.com/aws/aws-cdk-go/awscdkpipesalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var deliveryStream DeliveryStream
//
//   firehoseLogDestination := pipes_alpha.NewFirehoseLogDestination(deliveryStream)
//
// Experimental.
type FirehoseLogDestination interface {
	ILogDestination
	// Bind the log destination to the pipe.
	// Experimental.
	Bind(pipe IPipe) *LogDestinationConfig
	// Grant the pipe role to push to the log destination.
	// Experimental.
	GrantPush(grantee awsiam.IRole)
}

// The jsii proxy struct for FirehoseLogDestination
type jsiiProxy_FirehoseLogDestination struct {
	jsiiProxy_ILogDestination
}

// Experimental.
func NewFirehoseLogDestination(deliveryStream awskinesisfirehose.IDeliveryStream) FirehoseLogDestination {
	_init_.Initialize()

	if err := validateNewFirehoseLogDestinationParameters(deliveryStream); err != nil {
		panic(err)
	}
	j := jsiiProxy_FirehoseLogDestination{}

	_jsii_.Create(
		"@aws-cdk/aws-pipes-alpha.FirehoseLogDestination",
		[]interface{}{deliveryStream},
		&j,
	)

	return &j
}

// Experimental.
func NewFirehoseLogDestination_Override(f FirehoseLogDestination, deliveryStream awskinesisfirehose.IDeliveryStream) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-pipes-alpha.FirehoseLogDestination",
		[]interface{}{deliveryStream},
		f,
	)
}

func (f *jsiiProxy_FirehoseLogDestination) Bind(pipe IPipe) *LogDestinationConfig {
	if err := f.validateBindParameters(pipe); err != nil {
		panic(err)
	}
	var returns *LogDestinationConfig

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{pipe},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FirehoseLogDestination) GrantPush(grantee awsiam.IRole) {
	if err := f.validateGrantPushParameters(grantee); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"grantPush",
		[]interface{}{grantee},
	)
}

