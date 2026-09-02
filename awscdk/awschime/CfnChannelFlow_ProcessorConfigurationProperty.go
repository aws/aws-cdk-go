package awschime


// A processor's metadata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   processorConfigurationProperty := &ProcessorConfigurationProperty{
//   	Lambda: &LambdaConfigurationProperty{
//   		InvocationType: jsii.String("invocationType"),
//   		ResourceArn: jsii.String("resourceArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-channelflow-processorconfiguration.html
//
type CfnChannelFlow_ProcessorConfigurationProperty struct {
	// Stores metadata about a Lambda processor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-channelflow-processorconfiguration.html#cfn-chime-channelflow-processorconfiguration-lambda
	//
	Lambda interface{} `field:"required" json:"lambda" yaml:"lambda"`
}

