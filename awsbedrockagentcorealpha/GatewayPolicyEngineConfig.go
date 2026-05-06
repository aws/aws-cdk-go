package awsbedrockagentcorealpha


// Configuration for associating a policy engine with a gateway.
//
// When configured, the policy engine intercepts all agent requests through this
// gateway and evaluates them against the defined Cedar policies.
// [disable-awslint:prefer-ref-interface].
//
// Example:
//   // Create a Policy engine
//   policyEngine := agentcore.NewPolicyEngine(this, jsii.String("MyPolicyEngine"), &PolicyEngineProps{
//   	PolicyEngineName: jsii.String("my_policy_engine"),
//   	Description: jsii.String("Policy engine for access control"),
//   })
//
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   	PolicyEngineConfiguration: &GatewayPolicyEngineConfig{
//   		PolicyEngine: policyEngine,
//   		Mode: agentcore.PolicyEngineMode_ENFORCE(),
//   	},
//   })
//
//   // Add policy to policy engine
//   policyEngine.AddPolicy(jsii.String("AllowAllActions"), &AddPolicyOptions{
//   	Definition: fmt.Sprintf("\n    permit(\n      principal,\n      action,\n      resource == AgentCore::Gateway::\"%v\"\n    );\n  ", gateway.GatewayArn),
//   	Description: jsii.String("Allow all actions on specific gateway (development)"),
//   	ValidationMode: agentcore.PolicyValidationMode_IGNORE_ALL_FINDINGS(),
//   })
//
//   // you can add multiple policies to the policy engine
//   policyEngine.AddPolicy(jsii.String("SpecificToolPolicy"), &AddPolicyOptions{
//   	Definition: fmt.Sprintf("\n    permit(\n      principal is AgentCore::OAuthUser,\n      action == AgentCore::Action::\"WeatherTool__get_forecast\",\n      resource == AgentCore::Gateway::\"%v\"\n    );\n  ", gateway.*GatewayArn),
//   	Description: jsii.String("Allow specific weather tool access"),
//   	ValidationMode: agentcore.PolicyValidationMode_FAIL_ON_ANY_FINDINGS(),
//   })
//
// Experimental.
type GatewayPolicyEngineConfig struct {
	// The policy engine to associate with this gateway.
	//
	// [disable-awslint:prefer-ref-interface].
	// Experimental.
	PolicyEngine IPolicyEngine `field:"required" json:"policyEngine" yaml:"policyEngine"`
	// The enforcement mode for the policy engine.
	//
	// - `LOG_ONLY`: Evaluates and logs decisions without enforcing them. Use for testing.
	// - `ENFORCE`: Actively allows or denies requests based on Cedar policy evaluation.
	// Default: PolicyEngineMode.LOG_ONLY
	//
	// Experimental.
	Mode PolicyEngineMode `field:"optional" json:"mode" yaml:"mode"`
}

