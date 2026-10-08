package awsbedrockagentcore


// An SNS topic that receives lifecycle hook events asynchronously.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   harnessHookSnsTargetProperty := &HarnessHookSnsTargetProperty{
//   	Arn: jsii.String("arn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesshooksnstarget.html
//
type CfnHarness_HarnessHookSnsTargetProperty struct {
	// The ARN of the SNS topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesshooksnstarget.html#cfn-bedrockagentcore-harness-harnesshooksnstarget-arn
	//
	Arn *string `field:"required" json:"arn" yaml:"arn"`
}

