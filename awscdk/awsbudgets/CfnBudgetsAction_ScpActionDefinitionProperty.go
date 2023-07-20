package awsbudgets


// The service control policies (SCP) action definition details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scpActionDefinitionProperty := &ScpActionDefinitionProperty{
//   	PolicyId: jsii.String("policyId"),
//   	TargetIds: []*string{
//   		jsii.String("targetIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budgetsaction-scpactiondefinition.html
//
type CfnBudgetsAction_ScpActionDefinitionProperty struct {
	// The policy ID attached.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budgetsaction-scpactiondefinition.html#cfn-budgets-budgetsaction-scpactiondefinition-policyid
	//
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
	// A list of target IDs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budgetsaction-scpactiondefinition.html#cfn-budgets-budgetsaction-scpactiondefinition-targetids
	//
	TargetIds *[]*string `field:"required" json:"targetIds" yaml:"targetIds"`
}

