package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A principal for use in resources that need to have a role but it's unknown.
//
// Some resources have roles associated with them which they assume, such as
// Lambda Functions, CodeBuild projects, StepFunctions machines, etc.
//
// When those resources are imported, their actual roles are not always
// imported with them. When that happens, we use an instance of this class
// instead, which will add user warnings when statements are attempted to be
// added to it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import constructs "github.com/aws/constructs-go/constructs"
//
//   var construct construct
//
//   unknownPrincipal := awscdk.Aws_iam.NewUnknownPrincipal(&UnknownPrincipalProps{
//   	Resource: construct,
//   })
//
type UnknownPrincipal interface {
	IPrincipal
	// When this Principal is used in an AssumeRole policy, the action to use.
	AssumeRoleAction() *string
	// The principal to grant permissions to.
	GrantPrincipal() IPrincipal
	// Return the policy fragment that identifies this principal in a Policy.
	PolicyFragment() PrincipalPolicyFragment
	// Add to the policy of this principal.
	AddToPolicy(statement PolicyStatement) *bool
	// Add to the policy of this principal.
	AddToPrincipalPolicy(statement PolicyStatement) *AddToPrincipalPolicyResult
}

// The jsii proxy struct for UnknownPrincipal
type jsiiProxy_UnknownPrincipal struct {
	jsiiProxy_IPrincipal
}

func (j *jsiiProxy_UnknownPrincipal) AssumeRoleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UnknownPrincipal) GrantPrincipal() IPrincipal {
	var returns IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UnknownPrincipal) PolicyFragment() PrincipalPolicyFragment {
	var returns PrincipalPolicyFragment
	_jsii_.Get(
		j,
		"policyFragment",
		&returns,
	)
	return returns
}


func NewUnknownPrincipal(props *UnknownPrincipalProps) UnknownPrincipal {
	_init_.Initialize()

	if err := validateNewUnknownPrincipalParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_UnknownPrincipal{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.UnknownPrincipal",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewUnknownPrincipal_Override(u UnknownPrincipal, props *UnknownPrincipalProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.UnknownPrincipal",
		[]interface{}{props},
		u,
	)
}

func (u *jsiiProxy_UnknownPrincipal) AddToPolicy(statement PolicyStatement) *bool {
	if err := u.validateAddToPolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		u,
		"addToPolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UnknownPrincipal) AddToPrincipalPolicy(statement PolicyStatement) *AddToPrincipalPolicyResult {
	if err := u.validateAddToPrincipalPolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *AddToPrincipalPolicyResult

	_jsii_.Invoke(
		u,
		"addToPrincipalPolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

