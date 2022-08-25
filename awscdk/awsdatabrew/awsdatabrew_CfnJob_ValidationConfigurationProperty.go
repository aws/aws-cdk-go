package awsdatabrew


// Configuration for data quality validation.
//
// Used to select the Rulesets and Validation Mode to be used in the profile job. When ValidationConfiguration is null, the profile job will run without data quality validation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   validationConfigurationProperty := &validationConfigurationProperty{
//   	rulesetArn: jsii.String("rulesetArn"),
//
//   	// the properties below are optional
//   	validationMode: jsii.String("validationMode"),
//   }
//
type CfnJob_ValidationConfigurationProperty struct {
	// The Amazon Resource Name (ARN) for the ruleset to be validated in the profile job.
	//
	// The TargetArn of the selected ruleset should be the same as the Amazon Resource Name (ARN) of the dataset that is associated with the profile job.
	RulesetArn *string `field:"required" json:"rulesetArn" yaml:"rulesetArn"`
	// Mode of data quality validation.
	//
	// Default mode is “CHECK_ALL” which verifies all rules defined in the selected ruleset.
	ValidationMode *string `field:"optional" json:"validationMode" yaml:"validationMode"`
}

