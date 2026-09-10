package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An attribute of the request that a condition can compare against.
//
// Cedar exposes three sources of attributes. Which one you want depends on where the
// value comes from: the caller, the thing being accessed, or the request environment.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   awscdk.PolicyAttribute_Principal(jsii.String("department")) // principal.department
//   awscdk.PolicyAttribute_Resource(jsii.String("confidential")) // resource.confidential
//   awscdk.PolicyAttribute_Context(jsii.String("sourceIp"))
//
type PolicyAttribute interface {
}

// The jsii proxy struct for PolicyAttribute
type jsiiProxy_PolicyAttribute struct {
	_ byte // padding
}

// An attribute of the request context, meaning the request environment rather than either entity.
//
// For example `sourceIp`, `environment` or `timestamp`.
func PolicyAttribute_Context(attribute *string) PolicyAttribute {
	_init_.Initialize()

	if err := validatePolicyAttribute_ContextParameters(attribute); err != nil {
		panic(err)
	}
	var returns PolicyAttribute

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyAttribute",
		"context",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

// An attribute of the principal, meaning the authenticated user or service making the request.
//
// For example `username`, `department` or `groups`.
func PolicyAttribute_Principal(attribute *string) PolicyAttribute {
	_init_.Initialize()

	if err := validatePolicyAttribute_PrincipalParameters(attribute); err != nil {
		panic(err)
	}
	var returns PolicyAttribute

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyAttribute",
		"principal",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

// An attribute of the resource being accessed.
//
// For example `owner` or
// `classification`.
func PolicyAttribute_Resource(attribute *string) PolicyAttribute {
	_init_.Initialize()

	if err := validatePolicyAttribute_ResourceParameters(attribute); err != nil {
		panic(err)
	}
	var returns PolicyAttribute

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.PolicyAttribute",
		"resource",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

