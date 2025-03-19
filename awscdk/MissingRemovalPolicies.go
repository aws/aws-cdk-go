package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Manages removal policies for resources without existing policies within a construct scope.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   missingRemovalPolicies := cdk.MissingRemovalPolicies_Of(this)
//
type MissingRemovalPolicies interface {
	// Apply a removal policy only to resources without existing policies within this scope.
	Apply(policy RemovalPolicy, props *RemovalPolicyProps)
	// Apply DESTROY removal policy only to resources without existing policies within this scope.
	Destroy(props *RemovalPolicyProps)
	// Apply RETAIN removal policy only to resources without existing policies within this scope.
	Retain(props *RemovalPolicyProps)
	// Apply RETAIN_ON_UPDATE_OR_DELETE removal policy only to resources without existing policies within this scope.
	RetainOnUpdateOrDelete(props *RemovalPolicyProps)
	// Apply SNAPSHOT removal policy only to resources without existing policies within this scope.
	Snapshot(props *RemovalPolicyProps)
}

// The jsii proxy struct for MissingRemovalPolicies
type jsiiProxy_MissingRemovalPolicies struct {
	_ byte // padding
}

// Returns the missing removal policies API for the given scope.
func MissingRemovalPolicies_Of(scope constructs.IConstruct) MissingRemovalPolicies {
	_init_.Initialize()

	if err := validateMissingRemovalPolicies_OfParameters(scope); err != nil {
		panic(err)
	}
	var returns MissingRemovalPolicies

	_jsii_.StaticInvoke(
		"aws-cdk-lib.MissingRemovalPolicies",
		"of",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MissingRemovalPolicies) Apply(policy RemovalPolicy, props *RemovalPolicyProps) {
	if err := m.validateApplyParameters(policy, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"apply",
		[]interface{}{policy, props},
	)
}

func (m *jsiiProxy_MissingRemovalPolicies) Destroy(props *RemovalPolicyProps) {
	if err := m.validateDestroyParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"destroy",
		[]interface{}{props},
	)
}

func (m *jsiiProxy_MissingRemovalPolicies) Retain(props *RemovalPolicyProps) {
	if err := m.validateRetainParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"retain",
		[]interface{}{props},
	)
}

func (m *jsiiProxy_MissingRemovalPolicies) RetainOnUpdateOrDelete(props *RemovalPolicyProps) {
	if err := m.validateRetainOnUpdateOrDeleteParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"retainOnUpdateOrDelete",
		[]interface{}{props},
	)
}

func (m *jsiiProxy_MissingRemovalPolicies) Snapshot(props *RemovalPolicyProps) {
	if err := m.validateSnapshotParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"snapshot",
		[]interface{}{props},
	)
}

