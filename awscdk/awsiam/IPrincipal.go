package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a logical IAM principal.
//
// An IPrincipal describes a logical entity that can perform AWS API calls
// against sets of resources, optionally under certain conditions.
//
// Examples of simple principals are IAM objects that you create, such
// as Users or Roles.
//
// An example of a more complex principals is a `ServicePrincipal` (such as
// `new ServicePrincipal("sns.amazonaws.com")`, which represents the Simple
// Notifications Service).
//
// A single logical Principal may also map to a set of physical principals.
// For example, `new OrganizationPrincipal('o-1234')` represents all
// identities that are part of the given AWS Organization.
type IPrincipal interface {
	IGrantable
	// Add to the policy of this principal.
	AddToPrincipalPolicy(statement PolicyStatement) *AddToPrincipalPolicyResult
	// When this Principal is used in an AssumeRole policy, the action to use.
	AssumeRoleAction() *string
	// Return the policy fragment that identifies this principal in a Policy.
	PolicyFragment() PrincipalPolicyFragment
	// The AWS account ID of this principal.
	//
	// Can be undefined when the account is not known
	// (for example, for service principals).
	// Can be a Token - in that case,
	// it's assumed to be AWS::AccountId.
	PrincipalAccount() *string
}

// The jsii proxy for IPrincipal
type jsiiProxy_IPrincipal struct {
	jsiiProxy_IGrantable
}

func (i *jsiiProxy_IPrincipal) AddToPrincipalPolicy(statement PolicyStatement) *AddToPrincipalPolicyResult {
	if err := i.validateAddToPrincipalPolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *AddToPrincipalPolicyResult

	_jsii_.Invoke(
		i,
		"addToPrincipalPolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IPrincipal) AssumeRoleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrincipal) PolicyFragment() PrincipalPolicyFragment {
	var returns PrincipalPolicyFragment
	_jsii_.Get(
		j,
		"policyFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrincipal) PrincipalAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalAccount",
		&returns,
	)
	return returns
}

