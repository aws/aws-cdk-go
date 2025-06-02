package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Manages removal policies for all resources within a construct scope, overriding any existing policies by default.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   removalPolicies := cdk.RemovalPolicies_Of(this)
//
type RemovalPolicies interface {
	// Apply a removal policy to all resources within this scope, overriding any existing policies.
	Apply(policy RemovalPolicy, props *RemovalPolicyProps)
	// Apply DESTROY removal policy to all resources within this scope.
	Destroy(props *RemovalPolicyProps)
	// Apply RETAIN removal policy to all resources within this scope.
	Retain(props *RemovalPolicyProps)
	// Apply RETAIN_ON_UPDATE_OR_DELETE removal policy to all resources within this scope.
	RetainOnUpdateOrDelete(props *RemovalPolicyProps)
	// Apply SNAPSHOT removal policy to all resources within this scope.
	Snapshot(props *RemovalPolicyProps)
}

// The jsii proxy struct for RemovalPolicies
type jsiiProxy_RemovalPolicies struct {
	_ byte // padding
}

// Returns the removal policies API for the given scope.
func RemovalPolicies_Of(scope constructs.IConstruct) RemovalPolicies {
	_init_.Initialize()

	if err := validateRemovalPolicies_OfParameters(scope); err != nil {
		panic(err)
	}
	var returns RemovalPolicies

	_jsii_.StaticInvoke(
		"aws-cdk-lib.RemovalPolicies",
		"of",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RemovalPolicies) Apply(policy RemovalPolicy, props *RemovalPolicyProps) {
	if err := r.validateApplyParameters(policy, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"apply",
		[]interface{}{policy, props},
	)
}

func (r *jsiiProxy_RemovalPolicies) Destroy(props *RemovalPolicyProps) {
	if err := r.validateDestroyParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"destroy",
		[]interface{}{props},
	)
}

func (r *jsiiProxy_RemovalPolicies) Retain(props *RemovalPolicyProps) {
	if err := r.validateRetainParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"retain",
		[]interface{}{props},
	)
}

func (r *jsiiProxy_RemovalPolicies) RetainOnUpdateOrDelete(props *RemovalPolicyProps) {
	if err := r.validateRetainOnUpdateOrDeleteParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"retainOnUpdateOrDelete",
		[]interface{}{props},
	)
}

func (r *jsiiProxy_RemovalPolicies) Snapshot(props *RemovalPolicyProps) {
	if err := r.validateSnapshotParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"snapshot",
		[]interface{}{props},
	)
}

