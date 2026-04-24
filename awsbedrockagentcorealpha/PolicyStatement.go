package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Type-safe builder for creating Cedar authorization policy statements.
//
// This builder provides a fluent API for constructing Cedar policies without
// requiring knowledge of Cedar syntax. It supports:
// - Permit and forbid effects
// - Principal, action, and resource specifications
// - Conditional logic (when/unless clauses)
// - Raw Cedar for advanced cases
//
// The builder generates valid Cedar policy statements that can be used with
// the Policy construct.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//   var engine PolicyEngine
//
//
//   // Example 4: Raw Cedar policy
//   // For advanced Cedar features not supported by the builder
//   // Example 4: Raw Cedar policy
//   // For advanced Cedar features not supported by the builder
//   awsbedrockagentcorealpha.NewPolicy(this, jsii.String("CustomPolicy"), &PolicyProps{
//   	PolicyEngine: engine,
//   	Definition: jsii.String("permit(principal, action, resource) when { context.custom > 10 };"),
//   })
//
//   // Or using fromCedar():
//   // Or using fromCedar():
//   awsbedrockagentcorealpha.NewPolicy(this, jsii.String("ImportedPolicy"), &PolicyProps{
//   	PolicyEngine: engine,
//   	Statement: awsbedrockagentcorealpha.PolicyStatement_FromCedar(jsii.String("forbid(principal, action, resource) when { resource.confidential == true };")),
//   })
//
// Experimental.
type PolicyStatement interface {
	// Apply to all principals (any user, service, or entity).
	//
	// Generates: `principal` in Cedar.
	// Experimental.
	ForAllPrincipals() PolicyStatement
	// Apply to a specific principal entity.
	//
	// Generates: `principal == EntityType::"entityId"` in Cedar.
	// Experimental.
	ForPrincipal(entityType *string, entityId *string) PolicyStatement
	// Apply to principals that are members of a specific group.
	//
	// Generates: `principal in Group::"groupId"` in Cedar.
	// Experimental.
	ForPrincipalInGroup(groupType *string, groupId *string) PolicyStatement
	// Apply to a single specific action.
	//
	// Generates: `action == Action::"name"` in Cedar.
	// Experimental.
	OnAction(action *string) PolicyStatement
	// Apply to specific action(s).
	//
	// Generates: `action == Action::"name"` or `action in [Action::"name1", Action::"name2"]` in Cedar.
	// Experimental.
	OnActions(actions *[]*string) PolicyStatement
	// Apply to all actions (any operation).
	//
	// Generates: `action` in Cedar.
	// Experimental.
	OnAllActions() PolicyStatement
	// Apply to all resources of a specific type.
	//
	// **AWS Requirement**: AWS Bedrock AgentCore Policy service does not allow wildcard
	// resources (`resource`). This method provides type-constrained resources which are
	// required for policy validation to succeed.
	//
	// Generates: `resource is EntityType` in Cedar.
	//
	// Example:
	//   import "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
	//
	//
	//   // Constrain to Gateway resources (default)
	//   awsbedrockagentcorealpha.PolicyStatement_Permit().ForAllPrincipals().OnAllActions().OnAllResources() // → "resource is AgentCore::Gateway"
	//
	//   // Constrain to Runtime resources
	//   awsbedrockagentcorealpha.PolicyStatement_Permit().ForAllPrincipals().OnAllActions().OnAllResources(jsii.String("AgentCore::Runtime"))
	//
	// Experimental.
	OnAllResources(entityType *string) PolicyStatement
	// Apply to a specific resource instance.
	//
	// **AWS Requirement**: When using specific actions (e.g., `action == Action::"Delete"`),
	// you must constrain the resource to a specific instance, not just a type.
	//
	// Generates: `resource == EntityType::"arn"` in Cedar.
	//
	// Example:
	//   import "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
	//   var gatewayArn string
	//
	//
	//   awsbedrockagentcorealpha.PolicyStatement_Forbid().ForAllPrincipals().OnAction(jsii.String("AgentCore::Action::Delete")).OnResource(jsii.String("AgentCore::Gateway"), gatewayArn)
	//
	// Experimental.
	OnResource(entityType *string, entityArn *string) PolicyStatement
	// Apply to all resources of a specific type (explicit method).
	//
	// **AWS Requirement**: Resource type constraints are required by AWS Bedrock
	// AgentCore when using wildcard principals or actions.
	//
	// Generates: `resource is EntityType` in Cedar.
	//
	// Example:
	//   import "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
	//
	//
	//   awsbedrockagentcorealpha.PolicyStatement_Permit().ForAllPrincipals().OnAllActions().OnResourceType(jsii.String("AgentCore::Gateway"))
	//
	// Experimental.
	OnResourceType(entityType *string) PolicyStatement
	// Generate the Cedar policy statement string.
	//
	// Converts the builder state into valid Cedar policy syntax.
	// This is called internally by the Policy construct.
	//
	// Returns: Valid Cedar policy statement.
	// Experimental.
	ToCedar() *string
	// Add unless conditions - policy applies only if these conditions are false.
	//
	// Unless conditions define negative requirements (exclusions).
	// The policy applies when these conditions are NOT met.
	//
	// Returns a ConditionBuilder that you can chain condition methods on.
	// Call done() when finished to return to the PolicyStatement.
	// Experimental.
	Unless() ConditionalPolicyStatement
	// Add when conditions - policy applies only if these conditions are true.
	//
	// When conditions define positive requirements that must be met.
	// Multiple conditions can be combined with AND/OR operators.
	//
	// Returns a ConditionBuilder that you can chain condition methods on.
	// Call done() when finished to return to the PolicyStatement.
	// Experimental.
	When() ConditionalPolicyStatement
}

// The jsii proxy struct for PolicyStatement
type jsiiProxy_PolicyStatement struct {
	_ byte // padding
}

// Create a forbid statement - denies the action if conditions are met.
//
// Forbid statements deny access when their conditions evaluate to true.
// Forbid always takes precedence over permit (explicit deny).
// Experimental.
func PolicyStatement_Forbid() PolicyStatement {
	_init_.Initialize()

	var returns PolicyStatement

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.PolicyStatement",
		"forbid",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Create from raw Cedar policy statement string.
//
// Use this for advanced Cedar features not supported by the builder,
// or when migrating existing Cedar policies.
//
// Validation is deferred to the Policy construct's validationMode setting.
// Experimental.
func PolicyStatement_FromCedar(cedarStatement *string) PolicyStatement {
	_init_.Initialize()

	if err := validatePolicyStatement_FromCedarParameters(cedarStatement); err != nil {
		panic(err)
	}
	var returns PolicyStatement

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.PolicyStatement",
		"fromCedar",
		[]interface{}{cedarStatement},
		&returns,
	)

	return returns
}

// Create a permit statement - allows the action if conditions are met.
//
// Permit statements grant access when their conditions evaluate to true.
// Multiple permit statements can apply; any matching permit allows access.
// Experimental.
func PolicyStatement_Permit() PolicyStatement {
	_init_.Initialize()

	var returns PolicyStatement

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.PolicyStatement",
		"permit",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyStatement) ForAllPrincipals() PolicyStatement {
	var returns PolicyStatement

	_jsii_.Invoke(
		p,
		"forAllPrincipals",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyStatement) ForPrincipal(entityType *string, entityId *string) PolicyStatement {
	if err := p.validateForPrincipalParameters(entityType); err != nil {
		panic(err)
	}
	var returns PolicyStatement

	_jsii_.Invoke(
		p,
		"forPrincipal",
		[]interface{}{entityType, entityId},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyStatement) ForPrincipalInGroup(groupType *string, groupId *string) PolicyStatement {
	if err := p.validateForPrincipalInGroupParameters(groupType, groupId); err != nil {
		panic(err)
	}
	var returns PolicyStatement

	_jsii_.Invoke(
		p,
		"forPrincipalInGroup",
		[]interface{}{groupType, groupId},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyStatement) OnAction(action *string) PolicyStatement {
	if err := p.validateOnActionParameters(action); err != nil {
		panic(err)
	}
	var returns PolicyStatement

	_jsii_.Invoke(
		p,
		"onAction",
		[]interface{}{action},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyStatement) OnActions(actions *[]*string) PolicyStatement {
	if err := p.validateOnActionsParameters(actions); err != nil {
		panic(err)
	}
	var returns PolicyStatement

	_jsii_.Invoke(
		p,
		"onActions",
		[]interface{}{actions},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyStatement) OnAllActions() PolicyStatement {
	var returns PolicyStatement

	_jsii_.Invoke(
		p,
		"onAllActions",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyStatement) OnAllResources(entityType *string) PolicyStatement {
	var returns PolicyStatement

	_jsii_.Invoke(
		p,
		"onAllResources",
		[]interface{}{entityType},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyStatement) OnResource(entityType *string, entityArn *string) PolicyStatement {
	if err := p.validateOnResourceParameters(entityType, entityArn); err != nil {
		panic(err)
	}
	var returns PolicyStatement

	_jsii_.Invoke(
		p,
		"onResource",
		[]interface{}{entityType, entityArn},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyStatement) OnResourceType(entityType *string) PolicyStatement {
	if err := p.validateOnResourceTypeParameters(entityType); err != nil {
		panic(err)
	}
	var returns PolicyStatement

	_jsii_.Invoke(
		p,
		"onResourceType",
		[]interface{}{entityType},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyStatement) ToCedar() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toCedar",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyStatement) Unless() ConditionalPolicyStatement {
	var returns ConditionalPolicyStatement

	_jsii_.Invoke(
		p,
		"unless",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyStatement) When() ConditionalPolicyStatement {
	var returns ConditionalPolicyStatement

	_jsii_.Invoke(
		p,
		"when",
		nil, // no parameters
		&returns,
	)

	return returns
}

