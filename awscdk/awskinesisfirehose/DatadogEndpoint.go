package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A Datadog endpoint URL for use with Kinesis Data Firehose.
//
// Use one of the predefined static members for a known Datadog region, or
// `DatadogEndpoint.of(url)` for a custom URL.
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
type DatadogEndpoint interface {
	// The endpoint URL string.
	Url() *string
}

// The jsii proxy struct for DatadogEndpoint
type jsiiProxy_DatadogEndpoint struct {
	_ byte // padding
}

func (j *jsiiProxy_DatadogEndpoint) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}


// Use a custom Datadog endpoint URL.
//
// This is an escape hatch for endpoints not covered by the predefined members (for
// example a new region or a proxy in front of Datadog). The value is passed through
// as-is and is not validated by CDK. It must be a valid HTTPS Datadog intake URL: the
// Firehose API requires the endpoint URL to match `https://`, so an invalid value fails
// at deployment. The caller is responsible for providing a correct URL.
func DatadogEndpoint_Of(url *string) DatadogEndpoint {
	_init_.Initialize()

	if err := validateDatadogEndpoint_OfParameters(url); err != nil {
		panic(err)
	}
	var returns DatadogEndpoint

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesisfirehose.DatadogEndpoint",
		"of",
		[]interface{}{url},
		&returns,
	)

	return returns
}

func DatadogEndpoint_CONFIGURATION_AP1() DatadogEndpoint {
	_init_.Initialize()
	var returns DatadogEndpoint
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.DatadogEndpoint",
		"CONFIGURATION_AP1",
		&returns,
	)
	return returns
}

func DatadogEndpoint_CONFIGURATION_EU() DatadogEndpoint {
	_init_.Initialize()
	var returns DatadogEndpoint
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.DatadogEndpoint",
		"CONFIGURATION_EU",
		&returns,
	)
	return returns
}

func DatadogEndpoint_CONFIGURATION_US_GOV() DatadogEndpoint {
	_init_.Initialize()
	var returns DatadogEndpoint
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.DatadogEndpoint",
		"CONFIGURATION_US_GOV",
		&returns,
	)
	return returns
}

func DatadogEndpoint_CONFIGURATION_US1() DatadogEndpoint {
	_init_.Initialize()
	var returns DatadogEndpoint
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.DatadogEndpoint",
		"CONFIGURATION_US1",
		&returns,
	)
	return returns
}

func DatadogEndpoint_CONFIGURATION_US3() DatadogEndpoint {
	_init_.Initialize()
	var returns DatadogEndpoint
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.DatadogEndpoint",
		"CONFIGURATION_US3",
		&returns,
	)
	return returns
}

func DatadogEndpoint_CONFIGURATION_US5() DatadogEndpoint {
	_init_.Initialize()
	var returns DatadogEndpoint
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.DatadogEndpoint",
		"CONFIGURATION_US5",
		&returns,
	)
	return returns
}

func DatadogEndpoint_LOGS_AP1() DatadogEndpoint {
	_init_.Initialize()
	var returns DatadogEndpoint
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.DatadogEndpoint",
		"LOGS_AP1",
		&returns,
	)
	return returns
}

func DatadogEndpoint_LOGS_EU() DatadogEndpoint {
	_init_.Initialize()
	var returns DatadogEndpoint
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.DatadogEndpoint",
		"LOGS_EU",
		&returns,
	)
	return returns
}

func DatadogEndpoint_LOGS_GOV() DatadogEndpoint {
	_init_.Initialize()
	var returns DatadogEndpoint
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.DatadogEndpoint",
		"LOGS_GOV",
		&returns,
	)
	return returns
}

func DatadogEndpoint_LOGS_US1() DatadogEndpoint {
	_init_.Initialize()
	var returns DatadogEndpoint
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.DatadogEndpoint",
		"LOGS_US1",
		&returns,
	)
	return returns
}

func DatadogEndpoint_LOGS_US3() DatadogEndpoint {
	_init_.Initialize()
	var returns DatadogEndpoint
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.DatadogEndpoint",
		"LOGS_US3",
		&returns,
	)
	return returns
}

func DatadogEndpoint_LOGS_US5() DatadogEndpoint {
	_init_.Initialize()
	var returns DatadogEndpoint
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.DatadogEndpoint",
		"LOGS_US5",
		&returns,
	)
	return returns
}

func DatadogEndpoint_METRICS_AP1() DatadogEndpoint {
	_init_.Initialize()
	var returns DatadogEndpoint
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.DatadogEndpoint",
		"METRICS_AP1",
		&returns,
	)
	return returns
}

func DatadogEndpoint_METRICS_EU() DatadogEndpoint {
	_init_.Initialize()
	var returns DatadogEndpoint
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.DatadogEndpoint",
		"METRICS_EU",
		&returns,
	)
	return returns
}

func DatadogEndpoint_METRICS_US() DatadogEndpoint {
	_init_.Initialize()
	var returns DatadogEndpoint
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.DatadogEndpoint",
		"METRICS_US",
		&returns,
	)
	return returns
}

func DatadogEndpoint_METRICS_US5() DatadogEndpoint {
	_init_.Initialize()
	var returns DatadogEndpoint
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.DatadogEndpoint",
		"METRICS_US5",
		&returns,
	)
	return returns
}

