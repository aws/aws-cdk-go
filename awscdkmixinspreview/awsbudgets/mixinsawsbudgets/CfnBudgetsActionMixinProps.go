package mixinsawsbudgets


// Properties for CfnBudgetsActionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBudgetsActionMixinProps := &CfnBudgetsActionMixinProps{
//   	ActionThreshold: &ActionThresholdProperty{
//   		Type: jsii.String("type"),
//   		Value: jsii.Number(123),
//   	},
//   	ActionType: jsii.String("actionType"),
//   	ApprovalModel: jsii.String("approvalModel"),
//   	BudgetName: jsii.String("budgetName"),
//   	Definition: &DefinitionProperty{
//   		IamActionDefinition: &IamActionDefinitionProperty{
//   			Groups: []*string{
//   				jsii.String("groups"),
//   			},
//   			PolicyArn: jsii.String("policyArn"),
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
//   	ResourceTags: []ResourceTagProperty{
//   		&ResourceTagProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Subscribers: []interface{}{
//   		&SubscriberProperty{
//   			Address: jsii.String("address"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budgetsaction.html
//
type CfnBudgetsActionMixinProps struct {
	// The trigger threshold of the action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budgetsaction.html#cfn-budgets-budgetsaction-actionthreshold
	//
	ActionThreshold interface{} `field:"optional" json:"actionThreshold" yaml:"actionThreshold"`
	// The type of action.
	//
	// This defines the type of tasks that can be carried out by this action. This field also determines the format for definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budgetsaction.html#cfn-budgets-budgetsaction-actiontype
	//
	ActionType *string `field:"optional" json:"actionType" yaml:"actionType"`
	// This specifies if the action needs manual or automatic approval.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budgetsaction.html#cfn-budgets-budgetsaction-approvalmodel
	//
	ApprovalModel *string `field:"optional" json:"approvalModel" yaml:"approvalModel"`
	// A string that represents the budget name.
	//
	// ":" and "\" characters aren't allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budgetsaction.html#cfn-budgets-budgetsaction-budgetname
	//
	BudgetName *string `field:"optional" json:"budgetName" yaml:"budgetName"`
	// Specifies all of the type-specific parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budgetsaction.html#cfn-budgets-budgetsaction-definition
	//
	Definition interface{} `field:"optional" json:"definition" yaml:"definition"`
	// The role passed for action execution and reversion.
	//
	// Roles and actions must be in the same account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budgetsaction.html#cfn-budgets-budgetsaction-executionrolearn
	//
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The type of a notification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budgetsaction.html#cfn-budgets-budgetsaction-notificationtype
	//
	NotificationType *string `field:"optional" json:"notificationType" yaml:"notificationType"`
	// An optional list of tags to associate with the specified budget action.
	//
	// Each tag consists of a key and a value, and each key must be unique for the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budgetsaction.html#cfn-budgets-budgetsaction-resourcetags
	//
	ResourceTags *[]*CfnBudgetsActionPropsMixin_ResourceTagProperty `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// A list of subscribers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budgetsaction.html#cfn-budgets-budgetsaction-subscribers
	//
	Subscribers interface{} `field:"optional" json:"subscribers" yaml:"subscribers"`
}

