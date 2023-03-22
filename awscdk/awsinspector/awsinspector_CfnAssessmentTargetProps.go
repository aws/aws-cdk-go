package awsinspector


// Properties for defining a `CfnAssessmentTarget`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAssessmentTargetProps := &cfnAssessmentTargetProps{
//   	assessmentTargetName: jsii.String("assessmentTargetName"),
//   	resourceGroupArn: jsii.String("resourceGroupArn"),
//   }
//
type CfnAssessmentTargetProps struct {
	// The name of the Amazon Inspector assessment target.
	//
	// The name must be unique within the AWS account .
	AssessmentTargetName *string `field:"optional" json:"assessmentTargetName" yaml:"assessmentTargetName"`
	// The ARN that specifies the resource group that is used to create the assessment target.
	//
	// If `resourceGroupArn` is not specified, all EC2 instances in the current AWS account and Region are included in the assessment target.
	ResourceGroupArn *string `field:"optional" json:"resourceGroupArn" yaml:"resourceGroupArn"`
}

