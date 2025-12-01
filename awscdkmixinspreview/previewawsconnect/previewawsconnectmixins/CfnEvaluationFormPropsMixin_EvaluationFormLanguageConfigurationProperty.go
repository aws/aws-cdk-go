package previewawsconnectmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   evaluationFormLanguageConfigurationProperty := &EvaluationFormLanguageConfigurationProperty{
//   	FormLanguage: jsii.String("formLanguage"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformlanguageconfiguration.html
//
type CfnEvaluationFormPropsMixin_EvaluationFormLanguageConfigurationProperty struct {
	// The language of the form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformlanguageconfiguration.html#cfn-connect-evaluationform-evaluationformlanguageconfiguration-formlanguage
	//
	FormLanguage *string `field:"optional" json:"formLanguage" yaml:"formLanguage"`
}

