package previewawsconfigmixins


// The configuration object for AWS Config rule evaluation mode.
//
// The supported valid values are Detective or Proactive.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   evaluationModeConfigurationProperty := &EvaluationModeConfigurationProperty{
//   	Mode: jsii.String("mode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-evaluationmodeconfiguration.html
//
type CfnConfigRulePropsMixin_EvaluationModeConfigurationProperty struct {
	// The mode of an evaluation.
	//
	// The valid values are Detective or Proactive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-evaluationmodeconfiguration.html#cfn-config-configrule-evaluationmodeconfiguration-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

