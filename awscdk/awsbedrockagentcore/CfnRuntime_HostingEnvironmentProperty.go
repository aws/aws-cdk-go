package awsbedrockagentcore


// An upstream workload identified by the ARN of its hosting environment (for example a Gateway or Runtime ARN).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hostingEnvironmentProperty := &HostingEnvironmentProperty{
//   	Arn: jsii.String("arn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-hostingenvironment.html
//
type CfnRuntime_HostingEnvironmentProperty struct {
	// The ARN of the bedrock-agentcore hosting environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-hostingenvironment.html#cfn-bedrockagentcore-runtime-hostingenvironment-arn
	//
	Arn *string `field:"required" json:"arn" yaml:"arn"`
}

