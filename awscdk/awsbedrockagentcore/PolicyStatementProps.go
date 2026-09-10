package awsbedrockagentcore


// Properties for a policy statement.
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
type PolicyStatementProps struct {
	// The action the statement applies to.
	Action PolicyAction `field:"required" json:"action" yaml:"action"`
	// Whether the statement permits or forbids the action.
	Effect PolicyEffect `field:"required" json:"effect" yaml:"effect"`
	// The principal the statement applies to.
	Principal PolicyPrincipal `field:"required" json:"principal" yaml:"principal"`
	// The resource the statement applies to.
	Resource PolicyResource `field:"required" json:"resource" yaml:"resource"`
	// Conditions that must not hold for the statement to apply.
	// Default: - no exclusions.
	//
	Unless *[]PolicyCondition `field:"optional" json:"unless" yaml:"unless"`
	// Conditions that must all hold for the statement to apply.
	//
	// Use `PolicyCondition.anyOf()` for a member that only needs one of several
	// conditions to hold.
	// Default: - the statement applies whenever its principal, action and resource match.
	//
	When *[]PolicyCondition `field:"optional" json:"when" yaml:"when"`
}

