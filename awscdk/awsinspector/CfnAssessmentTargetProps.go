package awsinspector


// Properties for defining a `CfnAssessmentTarget`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAssessmentTargetProps := &CfnAssessmentTargetProps{
//   	AssessmentTargetName: jsii.String("assessmentTargetName"),
//   	ResourceGroupArn: jsii.String("resourceGroupArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-assessmenttarget.html
//
type CfnAssessmentTargetProps struct {
	// The name of the Amazon Inspector assessment target.
	//
	// The name must be unique within the AWS account .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-assessmenttarget.html#cfn-inspector-assessmenttarget-assessmenttargetname
	//
	AssessmentTargetName *string `field:"optional" json:"assessmentTargetName" yaml:"assessmentTargetName"`
	// The ARN that specifies the resource group that is used to create the assessment target.
	//
	// If `resourceGroupArn` is not specified, all EC2 instances in the current AWS account and Region are included in the assessment target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-assessmenttarget.html#cfn-inspector-assessmenttarget-resourcegrouparn
	//
	ResourceGroupArn *string `field:"optional" json:"resourceGroupArn" yaml:"resourceGroupArn"`
}

