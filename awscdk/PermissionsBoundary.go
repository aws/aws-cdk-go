package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Apply a permissions boundary to all IAM Roles and Users within a specific scope.
//
// A permissions boundary is typically applied at the `Stage` scope.
// This allows setting different permissions boundaries per Stage. For
// example, you may _not_ apply a boundary to the `Dev` stage which deploys
// to a personal dev account, but you _do_ apply the default boundary to the
// `Prod` stage.
//
// It is possible to apply different permissions boundaries to different scopes
// within your app. In this case the most specifically applied one wins
//
// Example:
//   // no permissions boundary for dev stage
//   // no permissions boundary for dev stage
//   awscdk.NewStage(app, jsii.String("DevStage"))
//
//   // default boundary for prod stage
//   prodStage := awscdk.NewStage(app, jsii.String("ProdStage"), &StageProps{
//   	PermissionsBoundary: awscdk.PermissionsBoundary_FromName(jsii.String("prod-pb")),
//   })
//
//   // overriding the pb applied for this stack
//   // overriding the pb applied for this stack
//   awscdk.Newstack(prodStage, jsii.String("ProdStack1"), &stackProps{
//   	PermissionsBoundary: awscdk.PermissionsBoundary_*FromName(jsii.String("stack-pb")),
//   })
//
//   // will inherit the permissions boundary from the stage
//   // will inherit the permissions boundary from the stage
//   awscdk.Newstack(prodStage, jsii.String("ProdStack2"))
//
type PermissionsBoundary interface {
}

// The jsii proxy struct for PermissionsBoundary
type jsiiProxy_PermissionsBoundary struct {
	_ byte // padding
}

// Apply a permissions boundary with the given ARN to all IAM Roles and Users created within a scope.
//
// The arn can include placeholders for the partition, region, qualifier, and account
// These placeholders will be replaced with the actual values if available. This requires
// that the Stack has the environment specified, it does not work with environment
// agnostic stacks.
//
// - '${AWS::Partition}'
// - '${AWS::Region}'
// - '${AWS::AccountId}'
// - '${Qualifier}'.
//
// Example:
//   awscdk.NewStage(app, jsii.String("ProdStage"), &StageProps{
//   	PermissionsBoundary: awscdk.PermissionsBoundary_FromArn(jsii.String("arn:aws:iam::${AWS::AccountId}:policy/my-custom-permissions-boundary")),
//   })
//
func PermissionsBoundary_FromArn(arn *string) PermissionsBoundary {
	_init_.Initialize()

	if err := validatePermissionsBoundary_FromArnParameters(arn); err != nil {
		panic(err)
	}
	var returns PermissionsBoundary

	_jsii_.StaticInvoke(
		"aws-cdk-lib.PermissionsBoundary",
		"fromArn",
		[]interface{}{arn},
		&returns,
	)

	return returns
}

// Apply a permissions boundary with the given name to all IAM Roles and Users created within a scope.
//
// The name can include placeholders for the partition, region, qualifier, and account
// These placeholders will be replaced with the actual values if available. This requires
// that the Stack has the environment specified, it does not work with environment
// agnostic stacks.
//
// - '${AWS::Partition}'
// - '${AWS::Region}'
// - '${AWS::AccountId}'
// - '${Qualifier}'.
//
// Example:
//   awscdk.NewStage(app, jsii.String("ProdStage"), &StageProps{
//   	PermissionsBoundary: awscdk.PermissionsBoundary_FromName(jsii.String("my-custom-permissions-boundary")),
//   })
//
func PermissionsBoundary_FromName(name *string) PermissionsBoundary {
	_init_.Initialize()

	if err := validatePermissionsBoundary_FromNameParameters(name); err != nil {
		panic(err)
	}
	var returns PermissionsBoundary

	_jsii_.StaticInvoke(
		"aws-cdk-lib.PermissionsBoundary",
		"fromName",
		[]interface{}{name},
		&returns,
	)

	return returns
}

