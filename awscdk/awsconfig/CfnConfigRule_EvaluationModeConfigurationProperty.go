package awsconfig


// Evaluation mode for the AWS Config rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluationModeConfigurationProperty := &EvaluationModeConfigurationProperty{
//   	Mode: jsii.String("mode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-evaluationmodeconfiguration.html
//
type CfnConfigRule_EvaluationModeConfigurationProperty struct {
	// Mode of evaluation of AWS Config rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-evaluationmodeconfiguration.html#cfn-config-configrule-evaluationmodeconfiguration-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

