package mixinsawsdatabrew


// Configuration for data quality validation.
//
// Used to select the Rulesets and Validation Mode to be used in the profile job. When ValidationConfiguration is null, the profile job will run without data quality validation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   validationConfigurationProperty := &ValidationConfigurationProperty{
//   	RulesetArn: jsii.String("rulesetArn"),
//   	ValidationMode: jsii.String("validationMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-validationconfiguration.html
//
type CfnJobPropsMixin_ValidationConfigurationProperty struct {
	// The Amazon Resource Name (ARN) for the ruleset to be validated in the profile job.
	//
	// The TargetArn of the selected ruleset should be the same as the Amazon Resource Name (ARN) of the dataset that is associated with the profile job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-validationconfiguration.html#cfn-databrew-job-validationconfiguration-rulesetarn
	//
	RulesetArn *string `field:"optional" json:"rulesetArn" yaml:"rulesetArn"`
	// Mode of data quality validation.
	//
	// Default mode is “CHECK_ALL” which verifies all rules defined in the selected ruleset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-validationconfiguration.html#cfn-databrew-job-validationconfiguration-validationmode
	//
	ValidationMode *string `field:"optional" json:"validationMode" yaml:"validationMode"`
}

