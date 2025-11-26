package previewawskinesisfirehosemixins


// The `ProcessingConfiguration` property configures data processing for an Amazon Kinesis Data Firehose delivery stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   processingConfigurationProperty := &ProcessingConfigurationProperty{
//   	Enabled: jsii.Boolean(false),
//   	Processors: []interface{}{
//   		&ProcessorProperty{
//   			Parameters: []interface{}{
//   				&ProcessorParameterProperty{
//   					ParameterName: jsii.String("parameterName"),
//   					ParameterValue: jsii.String("parameterValue"),
//   				},
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-processingconfiguration.html
//
type CfnDeliveryStreamPropsMixin_ProcessingConfigurationProperty struct {
	// Indicates whether data processing is enabled (true) or disabled (false).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-processingconfiguration.html#cfn-kinesisfirehose-deliverystream-processingconfiguration-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The data processors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-processingconfiguration.html#cfn-kinesisfirehose-deliverystream-processingconfiguration-processors
	//
	Processors interface{} `field:"optional" json:"processors" yaml:"processors"`
}

