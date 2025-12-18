package awsconnect


// Language configuration for an evaluation form.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationFormLanguageConfigurationProperty := &EvaluationFormLanguageConfigurationProperty{
//   	FormLanguage: jsii.String("formLanguage"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformlanguageconfiguration.html
//
type CfnEvaluationForm_EvaluationFormLanguageConfigurationProperty struct {
	// The language for the evaluation form.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-evaluationformlanguageconfiguration.html#cfn-connect-evaluationform-evaluationformlanguageconfiguration-formlanguage
	//
	FormLanguage *string `field:"optional" json:"formLanguage" yaml:"formLanguage"`
}

