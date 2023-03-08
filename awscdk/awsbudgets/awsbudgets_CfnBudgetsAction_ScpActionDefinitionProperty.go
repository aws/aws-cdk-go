package awsbudgets


// The service control policies (SCP) action definition details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scpActionDefinitionProperty := &scpActionDefinitionProperty{
//   	policyId: jsii.String("policyId"),
//   	targetIds: []*string{
//   		jsii.String("targetIds"),
//   	},
//   }
//
type CfnBudgetsAction_ScpActionDefinitionProperty struct {
	// The policy ID attached.
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
	// A list of target IDs.
	TargetIds *[]*string `field:"required" json:"targetIds" yaml:"targetIds"`
}

