package awskinesisfirehose


// The `ProcessingConfiguration` property configures data processing for an Amazon Kinesis Data Firehose delivery stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   processingConfigurationProperty := &ProcessingConfigurationProperty{
//   	Enabled: jsii.Boolean(false),
//   	Processors: []interface{}{
//   		&ProcessorProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Parameters: []interface{}{
//   				&ProcessorParameterProperty{
//   					ParameterName: jsii.String("parameterName"),
//   					ParameterValue: jsii.String("parameterValue"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-processingconfiguration.html
//
type CfnDeliveryStream_ProcessingConfigurationProperty struct {
	// Indicates whether data processing is enabled (true) or disabled (false).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-processingconfiguration.html#cfn-kinesisfirehose-deliverystream-processingconfiguration-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The data processors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-processingconfiguration.html#cfn-kinesisfirehose-deliverystream-processingconfiguration-processors
	//
	Processors interface{} `field:"optional" json:"processors" yaml:"processors"`
}

