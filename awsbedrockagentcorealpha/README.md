# Amazon Bedrock AgentCore Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

> The **Policy** submodule is the only submodule that remains in alpha. All other constructs have graduated to stable in `aws-cdk-lib/aws-bedrockagentcore` and we recommend migrating to the stable versions.

| **Language**                                                                                   | **Package**                             |
| :--------------------------------------------------------------------------------------------- | --------------------------------------- |
| ![Typescript Logo](https://docs.aws.amazon.com/cdk/api/latest/img/typescript32.png) TypeScript | `@aws-cdk/aws-bedrock-agentcore-alpha` |

## Migration to Stable

All constructs except Policy have moved to `aws-cdk-lib/aws-bedrockagentcore`:

```go
// Before
import agentcore "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
```

```go
// After (for all non-Policy constructs)
import agentcore "github.com/aws/aws-cdk-go/awscdk"
```

The following constructs are now in stable:

* **Runtime**: `Runtime`, `RuntimeEndpoint`, `AgentRuntimeArtifact`, `NetworkConfiguration`, `Observability`
* **Gateway**: `Gateway`, `GatewayTarget`, `GatewayAuthorizer`, `GatewayCredentialProvider`, `Interceptor`
* **Tools**: `BrowserCustom`, `CodeInterpreterCustom`
* **Memory**: `Memory`, `MemoryStrategy`
* **Evaluation**: `OnlineEvaluationConfig`, `Evaluator`, `EvaluatorSelector`
* **Identity**: `OAuth2CredentialProvider`, `ApiKeyCredentialProvider`, `WorkloadIdentity`

## What Remains in Alpha

The Policy submodule remains experimental:

* `PolicyEngine`
* `Policy`
* `PolicyStatement`
* `PolicyValidationMode`
* `PolicyEngineMode`

## Policy Engine

A policy engine is a collection of policies that evaluates and authorizes agent tool calls. When associated with a gateway, the policy engine intercepts all agent requests and determines whether to allow or deny each action based on the defined policies.

For more information, see the [Policy in Amazon Bedrock AgentCore documentation](https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/policy.html).

### PolicyEngine Properties

| Name | Type | Required | Description |
|------|------|----------|-------------|
| `policyEngineName` | `string` | No | The name of the policy engine. Valid characters: a-z, A-Z, 0-9, _ (underscore). Must start with a letter, 1-48 characters. If not provided, a unique name will be auto-generated |
| `description` | `string` | No | Optional description for the policy engine (max 4,096 characters). Default: no description |
| `kmsKey` | `IKey` | No | Custom KMS key for encryption. **IMPORTANT**: Once set, cannot be changed (requires replacement). Must be symmetric ENCRYPT_DECRYPT key. If key becomes inaccessible, all authorization decisions will be DENIED. Default: AWS owned key |
| `tags` | `{ [key: string]: string }` | No | Tags for the policy engine (max 50 tags). Default: no tags |

### Understanding Cedar Policies in AgentCore

Policies are constructed using [Cedar language](https://www.cedarpolicy.com/en/tutorial), an open source language for writing and enforcing authorization policies.
Cedar policies in AgentCore follow a specific structure with three main components: **Principal**, **Action**, and **Resource**. Understanding how these components work together is critical for writing effective policies.

#### Policy Structure

Every Cedar policy has this basic structure:

```cedar
permit(              // or forbid
  principal,         // Who is making the request
  action,            // What operation they want to perform
  resource           // What Gateway/tool they want to access
)
when {               // Optional conditions
  // Additional constraints
};
```

Example Policy

```cedar
permit(
  principal,
  action == AgentCore::Action::"ApplicationToolTarget___create_application",
  resource == AgentCore::Gateway::"<gateway-arn>"
) when {
  context.input.coverage_amount <= 1000000
};
```

### Basic PolicyEngine and Policy Creation

Create a policy engine and add policies to it.

#### Policy Engine Mode

When associating a policy engine with a gateway, you can control the enforcement behavior using `PolicyEngineMode`:

* `PolicyEngineMode.LOG_ONLY` (default) — evaluates actions and adds traces but does not enforce decisions. Use this mode for testing and validation before enabling enforcement.
* `PolicyEngineMode.ENFORCE` — actively allows or denies agent operations based on Cedar policy evaluation.

```go
// Create a Policy engine
policyEngine := agentcore.NewPolicyEngine(this, jsii.String("MyPolicyEngine"), &PolicyEngineProps{
	PolicyEngineName: jsii.String("my_policy_engine"),
	Description: jsii.String("Policy engine for access control"),
})

gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
	GatewayName: jsii.String("my-gateway"),
	PolicyEngineConfiguration: &GatewayPolicyEngineConfig{
		PolicyEngine: policyEngine,
		Mode: agentcore.PolicyEngineMode_ENFORCE(),
	},
})

// Add policy to policy engine
policyEngine.AddPolicy(jsii.String("AllowAllActions"), &AddPolicyOptions{
	Definition: fmt.Sprintf("\n    permit(\n      principal,\n      action,\n      resource == AgentCore::Gateway::\"%v\"\n    );\n  ", gateway.GatewayArn),
	Description: jsii.String("Allow all actions on specific gateway (development)"),
	ValidationMode: agentcore.PolicyValidationMode_IGNORE_ALL_FINDINGS(),
})

// you can add multiple policies to the policy engine
policyEngine.AddPolicy(jsii.String("SpecificToolPolicy"), &AddPolicyOptions{
	Definition: fmt.Sprintf("\n    permit(\n      principal is AgentCore::OAuthUser,\n      action == AgentCore::Action::\"WeatherTool__get_forecast\",\n      resource == AgentCore::Gateway::\"%v\"\n    );\n  ", gateway.*GatewayArn),
	Description: jsii.String("Allow specific weather tool access"),
	ValidationMode: agentcore.PolicyValidationMode_FAIL_ON_ANY_FINDINGS(),
})
```

### Type-Safe Policy Builder

For a more type-safe approach, use the `PolicyStatement` builder instead of writing raw Cedar syntax.

```go
gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
	GatewayName: jsii.String("my-gateway"),
})

policyEngine := agentcore.NewPolicyEngine(this, jsii.String("MyPolicyEngine"), &PolicyEngineProps{
	PolicyEngineName: jsii.String("my_policy_engine"),
})

allowAllPolicy := agentcore.NewPolicy(this, jsii.String("AllowAllPolicy"), &PolicyProps{
	PolicyEngine: policyEngine,
	PolicyName: jsii.String("allow_all"),
	Statement: agentcore.PolicyStatement_Permit().ForAllPrincipals().OnAllActions().OnResource(jsii.String("AgentCore::Gateway"), gateway.GatewayArn),
	Description: jsii.String("Allow all actions on specific gateway (development only)"),
	ValidationMode: agentcore.PolicyValidationMode_IGNORE_ALL_FINDINGS(),
})
```

#### Policy with Specific Actions

```go
var policyEngine PolicyEngine
var gateway Gateway


// Allow specific tool actions on specific gateway
// Action names follow pattern: "ToolName__operation"
policyEngine.AddPolicy(jsii.String("SpecificToolPolicy"), &AddPolicyOptions{
	Statement: agentcore.PolicyStatement_Permit().ForPrincipal(jsii.String("AgentCore::OAuthUser::your-client-id")).OnActions([]*string{
		jsii.String("AgentCore::Action::WeatherTool__get_forecast"),
		jsii.String("AgentCore::Action::WeatherTool__get_current"),
	}).OnResource(jsii.String("AgentCore::Gateway"), gateway.GatewayArn),
	Description: jsii.String("Allow specific weather tool operations"),
	ValidationMode: agentcore.PolicyValidationMode_FAIL_ON_ANY_FINDINGS(),
})
```

#### Policy with Conditions

Use `when` clauses to add advanced conditions based on principal tags (from OAuth token) or context:

```go
var policyEngine PolicyEngine
var gateway Gateway


// Policy with when conditions using principal tags
conditionalPolicy := agentcore.NewPolicy(this, jsii.String("ConditionalPolicy"), &PolicyProps{
	PolicyEngine: policyEngine,
	PolicyName: jsii.String("conditional_access"),
	Statement: agentcore.PolicyStatement_Permit().ForPrincipal(jsii.String("AgentCore::OAuthUser")).OnAllActions().OnResource(jsii.String("AgentCore::Gateway"), gateway.GatewayArn).When().PrincipalAttribute(jsii.String("department")).EqualTo(jsii.String("Engineering")).And().ContextAttribute(jsii.String("input.priority")).*EqualTo(jsii.String("high")).Done(),
	Description: jsii.String("Allow engineers for high-priority requests"),
	ValidationMode: agentcore.PolicyValidationMode_FAIL_ON_ANY_FINDINGS(),
})
```

#### Policy with Exclusions (unless)

Use `unless` clauses to exclude specific conditions from a policy. The policy applies when the `unless` conditions are NOT met:

```go
var policyEngine PolicyEngine
var gateway Gateway


// Allow access unless the user is suspended
policyWithUnless := agentcore.NewPolicy(this, jsii.String("UnlessPolicy"), &PolicyProps{
	PolicyEngine: policyEngine,
	PolicyName: jsii.String("unless_suspended"),
	Statement: agentcore.PolicyStatement_Permit().ForPrincipal(jsii.String("AgentCore::OAuthUser")).OnAllActions().OnResource(jsii.String("AgentCore::Gateway"), gateway.GatewayArn).Unless().PrincipalAttribute(jsii.String("suspended")).EqualTo(jsii.Boolean(true)).Done(),
	Description: jsii.String("Allow all actions unless user is suspended"),
	ValidationMode: agentcore.PolicyValidationMode_FAIL_ON_ANY_FINDINGS(),
})
```

You can combine `when` and `unless` clauses in the same policy:

```go
var policyEngine PolicyEngine
var gateway Gateway


// Allow engineers unless they are on probation
policyEngine.AddPolicy(jsii.String("CombinedConditions"), &AddPolicyOptions{
	Statement: agentcore.PolicyStatement_Permit().ForPrincipal(jsii.String("AgentCore::OAuthUser")).OnAllActions().OnResource(jsii.String("AgentCore::Gateway"), gateway.GatewayArn).When().PrincipalAttribute(jsii.String("department")).EqualTo(jsii.String("Engineering")).Done().Unless().*PrincipalAttribute(jsii.String("status")).*EqualTo(jsii.String("probation")).*Done(),
	Description: jsii.String("Allow engineers unless on probation"),
	ValidationMode: agentcore.PolicyValidationMode_FAIL_ON_ANY_FINDINGS(),
})
```

#### Forbid (Deny) Policy

Use `forbid` to explicitly deny access. Forbid policies override permit policies.

```go
var policyEngine PolicyEngine
var gateway Gateway


// Explicitly deny dangerous tool operations
policyEngine.AddPolicy(jsii.String("DenyDangerous"), &AddPolicyOptions{
	Statement: agentcore.PolicyStatement_Forbid().ForAllPrincipals().OnAction(jsii.String("AgentCore::Action::DeleteTool__delete_all")).OnResource(jsii.String("AgentCore::Gateway"), gateway.GatewayArn),
	Description: jsii.String("Forbid delete_all operation for all users"),
	ValidationMode: agentcore.PolicyValidationMode_FAIL_ON_ANY_FINDINGS(),
})
```

#### Raw Cedar for Advanced Cases

For advanced Cedar features not supported by the builder, use raw Cedar strings:

```go
var policyEngine PolicyEngine


// Option 1: Using definition property
advancedPolicy := agentcore.NewPolicy(this, jsii.String("AdvancedPolicy"), &PolicyProps{
	PolicyEngine: policyEngine,
	Definition: jsii.String("permit(principal, action, resource) when { context.custom > 10 };"),
	Description: jsii.String("Advanced policy with custom Cedar logic"),
})

// Option 2: Using fromCedar() with statement property
policyEngine.AddPolicy(jsii.String("CustomPolicy"), &AddPolicyOptions{
	Statement: agentcore.PolicyStatement_FromCedar(jsii.String("forbid(principal, action, resource) when { resource.confidential == true };")),
	Description: jsii.String("Custom policy from Cedar string"),
})
```

**Note**: You must specify **either** `definition` (raw Cedar string) **or** `statement` (PolicyStatement builder), but not both.

#### Accessing Policies on PolicyEngine

You can access the list of policies added to a PolicyEngine using policyEngine.policies.

### PolicyEngine with KMS Encryption

Encrypt policy data with a custom KMS key.

```go
// Create a custom KMS key
policyKey := kms.NewKey(this, jsii.String("PolicyEngineKey"), &KeyProps{
	EnableKeyRotation: jsii.Boolean(true),
	Description: jsii.String("KMS key for policy engine encryption"),
})

// Create policy engine with encryption
policyEngine := agentcore.NewPolicyEngine(this, jsii.String("EncryptedEngine"), &PolicyEngineProps{
	PolicyEngineName: jsii.String("encrypted_engine"),
	Description: jsii.String("Policy engine with KMS encryption"),
	KmsKey: policyKey,
})
```

### Importing Existing PolicyEngine

Import an existing policy engine from its ARN:

```go
importedEngine := agentcore.PolicyEngine_FromPolicyEngineAttributes(this, jsii.String("ImportedEngine"), &PolicyEngineAttributes{
	PolicyEngineArn: jsii.String("policy-engine-arn"),
	KmsKeyArn: jsii.String("kms-arn"),
})

// Use the imported engine
policy := agentcore.NewPolicy(this, jsii.String("PolicyForImportedEngine"), &PolicyProps{
	PolicyEngine: importedEngine,
	Definition: jsii.String("permit(principal, action, resource);"),
})
```

### Importing Existing Policy

Import an existing policy from its ARN:

```go
importedEngine := agentcore.PolicyEngine_FromPolicyEngineAttributes(this, jsii.String("ImportedEngine"), &PolicyEngineAttributes{
	PolicyEngineArn: jsii.String("policy-engine/my-engine-id"),
})

importedPolicy := agentcore.Policy_FromPolicyAttributes(this, jsii.String("ImportedPolicy"), &PolicyAttributes{
	PolicyArn: jsii.String("my-policy-arn"),
	PolicyEngine: importedEngine,
})

// Grant permissions to the imported policy
role := iam.NewRole(this, jsii.String("PolicyRole"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
})

importedPolicy.GrantRead(role)
```

### PolicyEngine IAM Permissions

Grant various levels of access to policy engines:

```go
policyEngine := agentcore.NewPolicyEngine(this, jsii.String("MyEngine"), &PolicyEngineProps{
	PolicyEngineName: jsii.String("my_engine"),
})

lambdaRole := iam.NewRole(this, jsii.String("LambdaRole"), &RoleProps{
	AssumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
})

// Grant read permissions
policyEngine.GrantRead(lambdaRole)

// Grant evaluation permissions
policyEngine.GrantEvaluate(lambdaRole)
```

## Using Policy with Stable Gateway

Since Gateway is now in `aws-cdk-lib/aws-bedrockagentcore` but Policy remains in alpha, use the L1 escape hatch to associate a policy engine with a stable gateway:

> Proper L2 integration will be added in a future update.

```go
import agentcore "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"


// Create policy engine (alpha)
policyEngine := agentcoreAlpha.NewPolicyEngine(this, jsii.String("Engine"), &PolicyEngineProps{
	PolicyEngineName: jsii.String("my_engine"),
})

// Create gateway (stable)
gateway := agentcore.NewGateway(this, jsii.String("Gateway"), &GatewayProps{
	GatewayName: jsii.String("my-gateway"),
})

// Wire policy engine to gateway via the L1 construct
cfnGateway := gateway.Node.defaultChild.(CfnGateway)
cfnGateway.policyEngineConfiguration = &GatewayPolicyEngineConfigurationProperty{
	Arn: policyEngine.PolicyEngineArn,
	Mode: agentcoreAlpha.PolicyEngineMode_ENFORCE().Value,
}

// Grant evaluate permissions to the gateway role
gateway.Role.AddToPrincipalPolicy(iam.NewPolicyStatement(&PolicyStatementProps{
	Actions: []*string{
		jsii.String("bedrock-agentcore:GetPolicyEngine"),
	},
	Resources: []*string{
		policyEngine.*PolicyEngineArn,
	},
}))
gateway.Role.AddToPrincipalPolicy(iam.NewPolicyStatement(&PolicyStatementProps{
	Actions: []*string{
		jsii.String("bedrock-agentcore:AuthorizeAction"),
		jsii.String("bedrock-agentcore:PartiallyAuthorizeActions"),
	},
	Resources: []*string{
		policyEngine.*PolicyEngineArn,
		gateway.GatewayArn,
	},
}))
```
