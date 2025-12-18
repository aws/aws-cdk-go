package previewawsconnectmixins


// Configuration that specifies the target for an evaluation form.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   evaluationFormTargetConfigurationProperty := &EvaluationFormTargetConfigurationProperty{
//   	ContactInteractionType: jsii.String("contactInteractionType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformtargetconfiguration.html
//
type CfnEvaluationFormPropsMixin_EvaluationFormTargetConfigurationProperty struct {
	// The contact interaction type for this evaluation form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformtargetconfiguration.html#cfn-connect-evaluationform-evaluationformtargetconfiguration-contactinteractiontype
	//
	ContactInteractionType *string `field:"optional" json:"contactInteractionType" yaml:"contactInteractionType"`
}

