package awsssm


// Properties for defining a `CfnSession`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSessionProps := &CfnSessionProps{
//   	Reason: jsii.String("reason"),
//   	Target: jsii.String("target"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-session.html
//
type CfnSessionProps struct {
	// The reason for connecting to the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-session.html#cfn-ssm-session-reason
	//
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
	// The managed node to connect to for the session.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-session.html#cfn-ssm-session-target
	//
	Target *string `field:"optional" json:"target" yaml:"target"`
}

