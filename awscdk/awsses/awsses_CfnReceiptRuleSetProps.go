package awsses


// Properties for defining a `CfnReceiptRuleSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReceiptRuleSetProps := &cfnReceiptRuleSetProps{
//   	ruleSetName: jsii.String("ruleSetName"),
//   }
//
type CfnReceiptRuleSetProps struct {
	// The name of the receipt rule set to reorder.
	RuleSetName *string `field:"optional" json:"ruleSetName" yaml:"ruleSetName"`
}

