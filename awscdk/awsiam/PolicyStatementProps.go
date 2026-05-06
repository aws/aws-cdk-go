package awsiam


// Interface for creating a policy statement.
//
// Example:
//   executionRole := iam.NewRole(this, jsii.String("EvaluationRole"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("bedrock-agentcore.amazonaws.com")),
//   	Description: jsii.String("Custom role for online evaluation"),
//   })
//
//   // Add required permissions
//   executionRole.AddToPolicy(iam.NewPolicyStatement(&PolicyStatementProps{
//   	Actions: []*string{
//   		jsii.String("logs:DescribeLogGroups"),
//   		jsii.String("logs:GetQueryResults"),
//   		jsii.String("logs:StartQuery"),
//   	},
//   	Resources: []*string{
//   		jsii.String("arn:aws:logs:*:*:log-group:/aws/bedrock-agentcore/*"),
//   	},
//   }))
//
//   evaluation := agentcore.NewOnlineEvaluationConfig(this, jsii.String("CustomRoleEval"), &OnlineEvaluationConfigProps{
//   	OnlineEvaluationConfigName: jsii.String("custom_role_evaluation"),
//   	Evaluators: []EvaluatorReference{
//   		agentcore.EvaluatorReference_Builtin(agentcore.BuiltinEvaluator_HELPFULNESS()),
//   	},
//   	DataSource: agentcore.DataSourceConfig_FromCloudWatchLogs(&CloudWatchLogsDataSourceConfig{
//   		LogGroupNames: []*string{
//   			jsii.String("/aws/bedrock-agentcore/my-agent"),
//   		},
//   		ServiceNames: []*string{
//   			jsii.String("my-agent.default"),
//   		},
//   	}),
//   	ExecutionRole: executionRole,
//   })
//
type PolicyStatementProps struct {
	// List of actions to add to the statement.
	// Default: - no actions.
	//
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// Conditions to add to the statement.
	// Default: - no condition.
	//
	Conditions *map[string]interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// Whether to allow or deny the actions in this statement.
	// Default: Effect.ALLOW
	//
	Effect Effect `field:"optional" json:"effect" yaml:"effect"`
	// List of not actions to add to the statement.
	// Default: - no not-actions.
	//
	NotActions *[]*string `field:"optional" json:"notActions" yaml:"notActions"`
	// List of not principals to add to the statement.
	// Default: - no not principals.
	//
	NotPrincipals *[]IPrincipal `field:"optional" json:"notPrincipals" yaml:"notPrincipals"`
	// NotResource ARNs to add to the statement.
	// Default: - no not-resources.
	//
	NotResources *[]*string `field:"optional" json:"notResources" yaml:"notResources"`
	// List of principals to add to the statement.
	// Default: - no principals.
	//
	Principals *[]IPrincipal `field:"optional" json:"principals" yaml:"principals"`
	// Resource ARNs to add to the statement.
	// Default: - no resources.
	//
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// The Sid (statement ID) is an optional identifier that you provide for the policy statement.
	//
	// You can assign a Sid value to each statement in a
	// statement array. In services that let you specify an ID element, such as
	// SQS and SNS, the Sid value is just a sub-ID of the policy document's ID. In
	// IAM, the Sid value must be unique within a JSON policy.
	// Default: - no sid.
	//
	Sid *string `field:"optional" json:"sid" yaml:"sid"`
}

