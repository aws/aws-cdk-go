package awsses


// Construction properties for a ReceiptRuleSet.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import ses "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   app := cdk.NewApp()
//   stack := cdk.NewStack(app, jsii.String("Stack"))
//   awscdk.CustomResourceConfig_Of(app).AddLogRetentionLifetime(logs.RetentionDays_TEN_YEARS)
//   awscdk.CustomResourceConfig_Of(app).AddRemovalPolicy(cdk.RemovalPolicy_DESTROY)
//
//   ses.NewReceiptRuleSet(app, jsii.String("RuleSet"), &ReceiptRuleSetProps{
//   	DropSpam: jsii.Boolean(true),
//   })
//
type ReceiptRuleSetProps struct {
	// Whether to add a first rule to stop processing messages that have at least one spam indicator.
	// Default: false.
	//
	DropSpam *bool `field:"optional" json:"dropSpam" yaml:"dropSpam"`
	// The name for the receipt rule set.
	// Default: - A CloudFormation generated name.
	//
	ReceiptRuleSetName *string `field:"optional" json:"receiptRuleSetName" yaml:"receiptRuleSetName"`
	// The list of rules to add to this rule set.
	//
	// Rules are added in the same
	// order as they appear in the list.
	// Default: - No rules are added to the rule set.
	//
	Rules *[]*ReceiptRuleOptions `field:"optional" json:"rules" yaml:"rules"`
}

