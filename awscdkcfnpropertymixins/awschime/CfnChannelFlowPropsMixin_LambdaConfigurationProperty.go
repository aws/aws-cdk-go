package awschime


// Stores metadata about a Lambda processor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   lambdaConfigurationProperty := &LambdaConfigurationProperty{
//   	InvocationType: jsii.String("invocationType"),
//   	ResourceArn: jsii.String("resourceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-channelflow-lambdaconfiguration.html
//
type CfnChannelFlowPropsMixin_LambdaConfigurationProperty struct {
	// Controls how the Lambda function is invoked.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-channelflow-lambdaconfiguration.html#cfn-chime-channelflow-lambdaconfiguration-invocationtype
	//
	InvocationType *string `field:"optional" json:"invocationType" yaml:"invocationType"`
	// The ARN of the Lambda message processing function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-channelflow-lambdaconfiguration.html#cfn-chime-channelflow-lambdaconfiguration-resourcearn
	//
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
}

