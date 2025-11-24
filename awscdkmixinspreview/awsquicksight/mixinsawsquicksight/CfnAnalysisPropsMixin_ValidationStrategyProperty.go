package mixinsawsquicksight


// The option to relax the validation that is required to create and update analyses, dashboards, and templates with definition objects.
//
// When you set this value to `LENIENT` , validation is skipped for specific errors.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   validationStrategyProperty := &ValidationStrategyProperty{
//   	Mode: jsii.String("mode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-validationstrategy.html
//
type CfnAnalysisPropsMixin_ValidationStrategyProperty struct {
	// The mode of validation for the asset to be created or updated.
	//
	// When you set this value to `STRICT` , strict validation for every error is enforced. When you set this value to `LENIENT` , validation is skipped for specific UI errors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-validationstrategy.html#cfn-quicksight-analysis-validationstrategy-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

