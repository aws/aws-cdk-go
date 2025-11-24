package mixinsawscleanrooms


// Individual budget parameter configuration that defines specific budget allocation settings for access budgets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   budgetParameterProperty := &BudgetParameterProperty{
//   	AutoRefresh: jsii.String("autoRefresh"),
//   	Budget: jsii.Number(123),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-privacybudgettemplate-budgetparameter.html
//
type CfnPrivacyBudgetTemplatePropsMixin_BudgetParameterProperty struct {
	// Whether this individual budget parameter automatically refreshes when the budget period resets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-privacybudgettemplate-budgetparameter.html#cfn-cleanrooms-privacybudgettemplate-budgetparameter-autorefresh
	//
	AutoRefresh *string `field:"optional" json:"autoRefresh" yaml:"autoRefresh"`
	// The budget allocation amount for this specific parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-privacybudgettemplate-budgetparameter.html#cfn-cleanrooms-privacybudgettemplate-budgetparameter-budget
	//
	Budget *float64 `field:"optional" json:"budget" yaml:"budget"`
	// The type of budget parameter being configured.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-privacybudgettemplate-budgetparameter.html#cfn-cleanrooms-privacybudgettemplate-budgetparameter-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

