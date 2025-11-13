package awscleanrooms


// Individual budget parameter configuration that defines specific budget allocation settings for access budgets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   budgetParameterProperty := &BudgetParameterProperty{
//   	Budget: jsii.Number(123),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	AutoRefresh: jsii.String("autoRefresh"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-privacybudgettemplate-budgetparameter.html
//
type CfnPrivacyBudgetTemplate_BudgetParameterProperty struct {
	// The budget allocation amount for this specific parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-privacybudgettemplate-budgetparameter.html#cfn-cleanrooms-privacybudgettemplate-budgetparameter-budget
	//
	Budget *float64 `field:"required" json:"budget" yaml:"budget"`
	// The type of budget parameter being configured.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-privacybudgettemplate-budgetparameter.html#cfn-cleanrooms-privacybudgettemplate-budgetparameter-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Whether this individual budget parameter automatically refreshes when the budget period resets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-privacybudgettemplate-budgetparameter.html#cfn-cleanrooms-privacybudgettemplate-budgetparameter-autorefresh
	//
	AutoRefresh *string `field:"optional" json:"autoRefresh" yaml:"autoRefresh"`
}

