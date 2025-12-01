package awsconnect


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
	// The interaction type of a contact.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformtargetconfiguration.html#cfn-connect-evaluationform-evaluationformtargetconfiguration-contactinteractiontype
	//
	ContactInteractionType *string `field:"required" json:"contactInteractionType" yaml:"contactInteractionType"`
}

