package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a schema configuration for Firehose S3 data record format conversion.
//
// Example:
//   var bucket Bucket
//   var schemaGlueTable CfnTable
//
//   s3Destination := firehose.NewS3Bucket(bucket, &S3BucketProps{
//   	DataFormatConversion: &DataFormatConversionProps{
//   		SchemaConfiguration: firehose.SchemaConfiguration_FromCfnTable(schemaGlueTable),
//   		InputFormat: firehose.InputFormat_OPENX_JSON(),
//   		OutputFormat: firehose.OutputFormat_PARQUET(),
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-dataformatconversionconfiguration.html#cfn-kinesisfirehose-deliverystream-dataformatconversionconfiguration-schemaconfiguration
//
type SchemaConfiguration interface {
	// Binds this Schema to the Destination, adding the necessary permissions to the Destination role.
	Bind(scope constructs.Construct, options *SchemaConfigurationBindOptions) *CfnDeliveryStream_SchemaConfigurationProperty
}

// The jsii proxy struct for SchemaConfiguration
type jsiiProxy_SchemaConfiguration struct {
	_ byte // padding
}

// Obtain schema configuration for data record format conversion from an `aws_glue.CfnTable`.
func SchemaConfiguration_FromCfnTable(table awsglue.CfnTable, props *SchemaConfigurationFromCfnTableProps) SchemaConfiguration {
	_init_.Initialize()

	if err := validateSchemaConfiguration_FromCfnTableParameters(table, props); err != nil {
		panic(err)
	}
	var returns SchemaConfiguration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesisfirehose.SchemaConfiguration",
		"fromCfnTable",
		[]interface{}{table, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaConfiguration) Bind(scope constructs.Construct, options *SchemaConfigurationBindOptions) *CfnDeliveryStream_SchemaConfigurationProperty {
	if err := s.validateBindParameters(scope, options); err != nil {
		panic(err)
	}
	var returns *CfnDeliveryStream_SchemaConfigurationProperty

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

