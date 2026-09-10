package awsbedrockagentcore


// Effect of a policy statement, whether it permits or forbids the action.
//
// Example:
//   var policyEngine PolicyEngine
//   var gateway Gateway
//
//
//   // Allow access unless the user is suspended
//   policyWithUnless := agentcore.NewPolicy(this, jsii.String("UnlessPolicy"), &PolicyProps{
//   	PolicyEngine: policyEngine,
//   	PolicyName: jsii.String("unless_suspended"),
//   	Statement: agentcore.NewPolicyStatement(&PolicyStatementProps{
//   		Effect: agentcore.PolicyEffect_PERMIT,
//   		Principal: agentcore.PolicyPrincipal_EntityType(jsii.String("AgentCore::OAuthUser")),
//   		Action: agentcore.PolicyAction_Any(),
//   		Resource: agentcore.PolicyResource_Instance(jsii.String("AgentCore::Gateway"), gateway.GatewayArn),
//   		Unless: []PolicyCondition{
//   			agentcore.PolicyCondition_BooleanEquals(agentcore.PolicyAttribute_Principal(jsii.String("suspended")), jsii.Boolean(true)),
//   		},
//   	}),
//   	Description: jsii.String("Allow all actions unless user is suspended"),
//   	ValidationMode: agentcore.PolicyValidationMode_FAIL_ON_ANY_FINDINGS(),
//   })
//
type PolicyEffect string

const (
	// Permit the action when the statement matches and its conditions hold.
	//
	// Multiple permit statements can apply, and any matching permit grants access.
	PolicyEffect_PERMIT PolicyEffect = "PERMIT"
	// Forbid the action when the statement matches and its conditions hold.
	//
	// A forbid always takes precedence over any permit.
	PolicyEffect_FORBID PolicyEffect = "FORBID"
)

