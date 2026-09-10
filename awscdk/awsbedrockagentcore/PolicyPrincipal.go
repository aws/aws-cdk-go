package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The principal a policy statement applies to.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   awscdk.PolicyPrincipal_Any() // principal
//   awscdk.PolicyPrincipal_EntityType(jsii.String("AgentCore::OAuthUser")) // principal is AgentCore::OAuthUser
//   awscdk.PolicyPrincipal_Entity(jsii.String("AgentCore::OAuthUser"), jsii.String("user123")) // principal == AgentCore::OAuthUser::"user123"
//   awscdk.PolicyPrincipal_InGroup(jsii.String("AgentCore::OAuthGroup"), jsii.String("admins"))
//
type PolicyPrincipal interface {
}

// The jsii proxy struct for PolicyPrincipal
type jsiiProxy_PolicyPrincipal struct {
	_ byte // padding
}

// Any principal, whoever the caller is.
func PolicyPrincipal_Any() PolicyPrincipal {
	_init_.Initialize()

	var returns PolicyPrincipal

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyPrincipal",
		"any",
		nil, // no parameters
		&returns,
	)

	return returns
}

// One specific principal.
func PolicyPrincipal_Entity(entityType *string, entityId *string) PolicyPrincipal {
	_init_.Initialize()

	if err := validatePolicyPrincipal_EntityParameters(entityType, entityId); err != nil {
		panic(err)
	}
	var returns PolicyPrincipal

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyPrincipal",
		"entity",
		[]interface{}{entityType, entityId},
		&returns,
	)

	return returns
}

// Any principal of the given entity type.
func PolicyPrincipal_EntityType(entityType *string) PolicyPrincipal {
	_init_.Initialize()

	if err := validatePolicyPrincipal_EntityTypeParameters(entityType); err != nil {
		panic(err)
	}
	var returns PolicyPrincipal

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyPrincipal",
		"entityType",
		[]interface{}{entityType},
		&returns,
	)

	return returns
}

// Any principal that is a member of the given group.
func PolicyPrincipal_InGroup(groupType *string, groupId *string) PolicyPrincipal {
	_init_.Initialize()

	if err := validatePolicyPrincipal_InGroupParameters(groupType, groupId); err != nil {
		panic(err)
	}
	var returns PolicyPrincipal

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyPrincipal",
		"inGroup",
		[]interface{}{groupType, groupId},
		&returns,
	)

	return returns
}

