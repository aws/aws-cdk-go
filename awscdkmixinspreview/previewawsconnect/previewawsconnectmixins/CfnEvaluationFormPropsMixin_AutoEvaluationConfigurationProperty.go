package previewawsconnectmixins


// Configuration information about automated evaluations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   autoEvaluationConfigurationProperty := &AutoEvaluationConfigurationProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-autoevaluationconfiguration.html
//
type CfnEvaluationFormPropsMixin_AutoEvaluationConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-evaluationform-autoevaluationconfiguration.html#cfn-connect-evaluationform-autoevaluationconfiguration-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

