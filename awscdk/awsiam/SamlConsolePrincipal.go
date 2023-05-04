package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Principal entity that represents a SAML federated identity provider for programmatic and AWS Management Console access.
//
// Example:
//   provider := iam.NewSamlProvider(this, jsii.String("Provider"), &SamlProviderProps{
//   	MetadataDocument: iam.SamlMetadataDocument_FromFile(jsii.String("/path/to/saml-metadata-document.xml")),
//   })
//   iam.NewRole(this, jsii.String("Role"), &RoleProps{
//   	AssumedBy: iam.NewSamlConsolePrincipal(provider),
//   })
//
type SamlConsolePrincipal interface {
	SamlPrincipal
	// When this Principal is used in an AssumeRole policy, the action to use.
	AssumeRoleAction() *string
	// The conditions under which the policy is in effect.
	// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_condition.html
	//
	Conditions() *map[string]interface{}
	// federated identity provider (i.e. 'cognito-identity.amazonaws.com' for users authenticated through Cognito).
	Federated() *string
	// The principal to grant permissions to.
	GrantPrincipal() IPrincipal
	// Return the policy fragment that identifies this principal in a Policy.
	PolicyFragment() PrincipalPolicyFragment
	// The AWS account ID of this principal.
	//
	// Can be undefined when the account is not known
	// (for example, for service principals).
	// Can be a Token - in that case,
	// it's assumed to be AWS::AccountId.
	PrincipalAccount() *string
	// Add the principal to the AssumeRolePolicyDocument.
	//
	// Add the statements to the AssumeRolePolicyDocument necessary to give this principal
	// permissions to assume the given role.
	AddToAssumeRolePolicy(document PolicyDocument)
	// Add to the policy of this principal.
	AddToPolicy(statement PolicyStatement) *bool
	// Add to the policy of this principal.
	AddToPrincipalPolicy(_statement PolicyStatement) *AddToPrincipalPolicyResult
	// Return whether or not this principal is equal to the given principal.
	DedupeString() *string
	// JSON-ify the principal.
	//
	// Used when JSON.stringify() is called
	ToJSON() *map[string]*[]*string
	// Returns a string representation of an object.
	ToString() *string
	// Returns a new PrincipalWithConditions using this principal as the base, with the passed conditions added.
	//
	// When there is a value for the same operator and key in both the principal and the
	// conditions parameter, the value from the conditions parameter will be used.
	//
	// Returns: a new PrincipalWithConditions object.
	WithConditions(conditions *map[string]interface{}) PrincipalBase
	// Returns a new principal using this principal as the base, with session tags enabled.
	//
	// Returns: a new SessionTagsPrincipal object.
	WithSessionTags() PrincipalBase
}

// The jsii proxy struct for SamlConsolePrincipal
type jsiiProxy_SamlConsolePrincipal struct {
	jsiiProxy_SamlPrincipal
}

func (j *jsiiProxy_SamlConsolePrincipal) AssumeRoleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlConsolePrincipal) Conditions() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlConsolePrincipal) Federated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"federated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlConsolePrincipal) GrantPrincipal() IPrincipal {
	var returns IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlConsolePrincipal) PolicyFragment() PrincipalPolicyFragment {
	var returns PrincipalPolicyFragment
	_jsii_.Get(
		j,
		"policyFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SamlConsolePrincipal) PrincipalAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalAccount",
		&returns,
	)
	return returns
}


func NewSamlConsolePrincipal(samlProvider ISamlProvider, conditions *map[string]interface{}) SamlConsolePrincipal {
	_init_.Initialize()

	if err := validateNewSamlConsolePrincipalParameters(samlProvider); err != nil {
		panic(err)
	}
	j := jsiiProxy_SamlConsolePrincipal{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.SamlConsolePrincipal",
		[]interface{}{samlProvider, conditions},
		&j,
	)

	return &j
}

func NewSamlConsolePrincipal_Override(s SamlConsolePrincipal, samlProvider ISamlProvider, conditions *map[string]interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.SamlConsolePrincipal",
		[]interface{}{samlProvider, conditions},
		s,
	)
}

func (s *jsiiProxy_SamlConsolePrincipal) AddToAssumeRolePolicy(document PolicyDocument) {
	if err := s.validateAddToAssumeRolePolicyParameters(document); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addToAssumeRolePolicy",
		[]interface{}{document},
	)
}

func (s *jsiiProxy_SamlConsolePrincipal) AddToPolicy(statement PolicyStatement) *bool {
	if err := s.validateAddToPolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		s,
		"addToPolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlConsolePrincipal) AddToPrincipalPolicy(_statement PolicyStatement) *AddToPrincipalPolicyResult {
	if err := s.validateAddToPrincipalPolicyParameters(_statement); err != nil {
		panic(err)
	}
	var returns *AddToPrincipalPolicyResult

	_jsii_.Invoke(
		s,
		"addToPrincipalPolicy",
		[]interface{}{_statement},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlConsolePrincipal) DedupeString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"dedupeString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlConsolePrincipal) ToJSON() *map[string]*[]*string {
	var returns *map[string]*[]*string

	_jsii_.Invoke(
		s,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlConsolePrincipal) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlConsolePrincipal) WithConditions(conditions *map[string]interface{}) PrincipalBase {
	if err := s.validateWithConditionsParameters(conditions); err != nil {
		panic(err)
	}
	var returns PrincipalBase

	_jsii_.Invoke(
		s,
		"withConditions",
		[]interface{}{conditions},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SamlConsolePrincipal) WithSessionTags() PrincipalBase {
	var returns PrincipalBase

	_jsii_.Invoke(
		s,
		"withSessionTags",
		nil, // no parameters
		&returns,
	)

	return returns
}

