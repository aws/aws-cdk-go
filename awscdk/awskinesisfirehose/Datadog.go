package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// A Datadog destination for data from a Kinesis Data Firehose delivery stream.
//
// Example:
//   import secretsmanager "github.com/aws/aws-cdk-go/awscdk"
//
//   var apiKey Secret
//
//   datadogDestination := firehose.NewDatadog(&DatadogProps{
//   	ApiKey: ApiKey,
//   	Endpoint: firehose.DatadogEndpoint_LOGS_US1(),
//   })
//
type Datadog interface {
	HttpEndpoint
	// Binds this destination to the Amazon Data Firehose delivery stream.
	//
	// Implementers should use this method to bind resources to the stack and initialize values using the provided stream.
	Bind(scope constructs.Construct, options *DestinationBindOptions) *DestinationConfig
}

// The jsii proxy struct for Datadog
type jsiiProxy_Datadog struct {
	jsiiProxy_HttpEndpoint
}

func NewDatadog(props *DatadogProps) Datadog {
	_init_.Initialize()

	if err := validateNewDatadogParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Datadog{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisfirehose.Datadog",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewDatadog_Override(d Datadog, props *DatadogProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisfirehose.Datadog",
		[]interface{}{props},
		d,
	)
}

func (d *jsiiProxy_Datadog) Bind(scope constructs.Construct, options *DestinationBindOptions) *DestinationConfig {
	if err := d.validateBindParameters(scope, options); err != nil {
		panic(err)
	}
	var returns *DestinationConfig

	_jsii_.Invoke(
		d,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

