package awsbudgets


// Properties for defining a `CfnBudgetsAction`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBudgetsActionProps := &CfnBudgetsActionProps{
//   	ActionThreshold: &ActionThresholdProperty{
//   		Type: jsii.String("type"),
//   		Value: jsii.Number(123),
//   	},
//   	ActionType: jsii.String("actionType"),
//   	BudgetName: jsii.String("budgetName"),
//   	Definition: &DefinitionProperty{
//   		IamActionDefinition: &IamActionDefinitionProperty{
//   			PolicyArn: jsii.String("policyArn"),
//
//   			// the properties below are optional
//   			Groups: []*string{
//   				jsii.String("groups"),
//   			},
//   			Roles: []*string{
//   				jsii.String("roles"),
//   			},
//   			Users: []*string{
//   				jsii.String("users"),
//   			},
//   		},
//   		ScpActionDefinition: &ScpActionDefinitionProperty{
//   			PolicyId: jsii.String("policyId"),
//   			TargetIds: []*string{
//   				jsii.String("targetIds"),
//   			},
//   		},
//   		SsmActionDefinition: &SsmActionDefinitionProperty{
//   			InstanceIds: []*string{
//   				jsii.String("instanceIds"),
//   			},
//   			Region: jsii.String("region"),
//   			Subtype: jsii.String("subtype"),
//   		},
//   	},
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	NotificationType: jsii.String("notificationType"),
//   	Subscribers: []interface{}{
//   		&SubscriberProperty{
//   			Address: jsii.String("address"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	ApprovalModel: jsii.String("approvalModel"),
//   }
//
type CfnBudgetsActionProps struct {
	// The trigger threshold of the action.
	ActionThreshold interface{} `field:"required" json:"actionThreshold" yaml:"actionThreshold"`
	// The type of action.
	//
	// This defines the type of tasks that can be carried out by this action. This field also determines the format for definition.
	ActionType *string `field:"required" json:"actionType" yaml:"actionType"`
	// A string that represents the budget name.
	//
	// ":" and "\" characters aren't allowed.
	BudgetName *string `field:"required" json:"budgetName" yaml:"budgetName"`
	// Specifies all of the type-specific parameters.
	Definition interface{} `field:"required" json:"definition" yaml:"definition"`
	// The role passed for action execution and reversion.
	//
	// Roles and actions must be in the same account.
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The type of a notification.
	NotificationType *string `field:"required" json:"notificationType" yaml:"notificationType"`
	// A list of subscribers.
	Subscribers interface{} `field:"required" json:"subscribers" yaml:"subscribers"`
	// This specifies if the action needs manual or automatic approval.
	ApprovalModel *string `field:"optional" json:"approvalModel" yaml:"approvalModel"`
}

