package awsconnect


// Configuration that specifies the target for an evaluation form.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationFormTargetConfigurationProperty := &EvaluationFormTargetConfigurationProperty{
//   	ContactInteractionType: jsii.String("contactInteractionType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformtargetconfiguration.html
//
type CfnEvaluationForm_EvaluationFormTargetConfigurationProperty struct {
	// The contact interaction type for this evaluation form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformtargetconfiguration.html#cfn-connect-evaluationform-evaluationformtargetconfiguration-contactinteractiontype
	//
	ContactInteractionType *string `field:"required" json:"contactInteractionType" yaml:"contactInteractionType"`
}

