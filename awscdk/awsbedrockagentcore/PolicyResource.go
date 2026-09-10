package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The resource a policy statement applies to.
//
// AgentCore rejects a policy whose resource is an unconstrained wildcard, so a
// statement must name either a resource type or a specific resource.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   var gatewayArn string
//
//
//   awscdk.PolicyResource_AnyOfType(jsii.String("AgentCore::Gateway")) // resource is AgentCore::Gateway
//   awscdk.PolicyResource_Instance(jsii.String("AgentCore::Gateway"), gatewayArn)
//
type PolicyResource interface {
}

// The jsii proxy struct for PolicyResource
type jsiiProxy_PolicyResource struct {
	_ byte // padding
}

// Any resource of the given entity type.
func PolicyResource_AnyOfType(entityType *string) PolicyResource {
	_init_.Initialize()

	if err := validatePolicyResource_AnyOfTypeParameters(entityType); err != nil {
		panic(err)
	}
	var returns PolicyResource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyResource",
		"anyOfType",
		[]interface{}{entityType},
		&returns,
	)

	return returns
}

// One specific resource.
//
// AgentCore requires a specific resource when the statement names specific
// actions rather than any action.
func PolicyResource_Instance(entityType *string, entityArn *string) PolicyResource {
	_init_.Initialize()

	if err := validatePolicyResource_InstanceParameters(entityType, entityArn); err != nil {
		panic(err)
	}
	var returns PolicyResource

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyResource",
		"instance",
		[]interface{}{entityType, entityArn},
		&returns,
	)

	return returns
}

