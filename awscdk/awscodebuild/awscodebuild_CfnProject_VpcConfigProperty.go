package awscodebuild


// `VpcConfig` is a property of the [AWS::CodeBuild::Project](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html) resource that enable AWS CodeBuild to access resources in an Amazon VPC. For more information, see [Use AWS CodeBuild with Amazon Virtual Private Cloud](https://docs.aws.amazon.com/codebuild/latest/userguide/vpc-support.html) in the *AWS CodeBuild User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcConfigProperty := &vpcConfigProperty{
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   	vpcId: jsii.String("vpcId"),
//   }
//
type CfnProject_VpcConfigProperty struct {
	// A list of one or more security groups IDs in your Amazon VPC.
	//
	// The maximum count is 5.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// A list of one or more subnet IDs in your Amazon VPC.
	//
	// The maximum count is 16.
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
	// The ID of the Amazon VPC.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

