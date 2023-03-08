package awsbudgets


// Properties for defining a `CfnBudgetsAction`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBudgetsActionProps := &cfnBudgetsActionProps{
//   	actionThreshold: &actionThresholdProperty{
//   		type: jsii.String("type"),
//   		value: jsii.Number(123),
//   	},
//   	actionType: jsii.String("actionType"),
//   	budgetName: jsii.String("budgetName"),
//   	definition: &definitionProperty{
//   		iamActionDefinition: &iamActionDefinitionProperty{
//   			policyArn: jsii.String("policyArn"),
//
//   			// the properties below are optional
//   			groups: []*string{
//   				jsii.String("groups"),
//   			},
//   			roles: []*string{
//   				jsii.String("roles"),
//   			},
//   			users: []*string{
//   				jsii.String("users"),
//   			},
//   		},
//   		scpActionDefinition: &scpActionDefinitionProperty{
//   			policyId: jsii.String("policyId"),
//   			targetIds: []*string{
//   				jsii.String("targetIds"),
//   			},
//   		},
//   		ssmActionDefinition: &ssmActionDefinitionProperty{
//   			instanceIds: []*string{
//   				jsii.String("instanceIds"),
//   			},
//   			region: jsii.String("region"),
//   			subtype: jsii.String("subtype"),
//   		},
//   	},
//   	executionRoleArn: jsii.String("executionRoleArn"),
//   	notificationType: jsii.String("notificationType"),
//   	subscribers: []interface{}{
//   		&subscriberProperty{
//   			address: jsii.String("address"),
//   			type: jsii.String("type"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	approvalModel: jsii.String("approvalModel"),
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

