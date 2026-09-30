package awschime


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kinesisDataStreamSinkConfigurationProperty := &KinesisDataStreamSinkConfigurationProperty{
//   	InsightsTarget: jsii.String("insightsTarget"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-kinesisdatastreamsinkconfiguration.html
//
type CfnMediaInsightsPipelineConfiguration_KinesisDataStreamSinkConfigurationProperty struct {
	// The ARN of the Kinesis Data Stream sink.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-kinesisdatastreamsinkconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-kinesisdatastreamsinkconfiguration-insightstarget
	//
	InsightsTarget *string `field:"optional" json:"insightsTarget" yaml:"insightsTarget"`
}

