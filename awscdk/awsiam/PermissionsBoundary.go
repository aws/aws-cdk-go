package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Modify the Permissions Boundaries of Users and Roles in a construct tree.
//
// ```ts
// const policy = iam.ManagedPolicy.fromAwsManagedPolicyName('ReadOnlyAccess');
// iam.PermissionsBoundary.of(this).apply(policy);
// ```.
//
// Example:
//   var project project
//
//   iam.PermissionsBoundary_Of(project).Apply(codebuild.NewUntrustedCodeBoundaryPolicy(this, jsii.String("Boundary")))
//
type PermissionsBoundary interface {
	// Apply the given policy as Permissions Boundary to all Roles and Users in the scope.
	//
	// Will override any Permissions Boundaries configured previously; in case
	// a Permission Boundary is applied in multiple scopes, the Boundary applied
	// closest to the Role wins.
	Apply(boundaryPolicy IManagedPolicy)
	// Remove previously applied Permissions Boundaries.
	Clear()
}

// The jsii proxy struct for PermissionsBoundary
type jsiiProxy_PermissionsBoundary struct {
	_ byte // padding
}

// Access the Permissions Boundaries of a construct tree.
func PermissionsBoundary_Of(scope constructs.IConstruct) PermissionsBoundary {
	_init_.Initialize()

	if err := validatePermissionsBoundary_OfParameters(scope); err != nil {
		panic(err)
	}
	var returns PermissionsBoundary

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.PermissionsBoundary",
		"of",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PermissionsBoundary) Apply(boundaryPolicy IManagedPolicy) {
	if err := p.validateApplyParameters(boundaryPolicy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"apply",
		[]interface{}{boundaryPolicy},
	)
}

func (p *jsiiProxy_PermissionsBoundary) Clear() {
	_jsii_.InvokeVoid(
		p,
		"clear",
		nil, // no parameters
	)
}

