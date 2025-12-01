package awsecr


// Properties for defining a `CfnPullTimeUpdateExclusion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPullTimeUpdateExclusionProps := &CfnPullTimeUpdateExclusionProps{
//   	PrincipalArn: jsii.String("principalArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-pulltimeupdateexclusion.html
//
type CfnPullTimeUpdateExclusionProps struct {
	// The ARN of the IAM principal to remove from the pull time update exclusion list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-pulltimeupdateexclusion.html#cfn-ecr-pulltimeupdateexclusion-principalarn
	//
	PrincipalArn *string `field:"required" json:"principalArn" yaml:"principalArn"`
}

