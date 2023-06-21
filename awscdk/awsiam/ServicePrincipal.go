package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An IAM principal that represents an AWS service (i.e. `sqs.amazonaws.com`).
//
// Example:
//   lambdaRole := iam.NewRole(this, jsii.String("Role"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
//   	Description: jsii.String("Example role..."),
//   })
//
//   stream := kinesis.NewStream(this, jsii.String("MyEncryptedStream"), &StreamProps{
//   	Encryption: kinesis.StreamEncryption_KMS,
//   })
//
//   // give lambda permissions to read stream
//   stream.grantRead(lambdaRole)
//
type ServicePrincipal interface {
	PrincipalBase
	// When this Principal is used in an AssumeRole policy, the action to use.
	AssumeRoleAction() *string
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
	// AWS service (i.e. sqs.amazonaws.com).
	Service() *string
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

// The jsii proxy struct for ServicePrincipal
type jsiiProxy_ServicePrincipal struct {
	jsiiProxy_PrincipalBase
}

func (j *jsiiProxy_ServicePrincipal) AssumeRoleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) GrantPrincipal() IPrincipal {
	var returns IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) PolicyFragment() PrincipalPolicyFragment {
	var returns PrincipalPolicyFragment
	_jsii_.Get(
		j,
		"policyFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) PrincipalAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicePrincipal) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}


// Reference an AWS service, optionally in a given region.
func NewServicePrincipal(service *string, opts *ServicePrincipalOpts) ServicePrincipal {
	_init_.Initialize()

	if err := validateNewServicePrincipalParameters(service, opts); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServicePrincipal{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.ServicePrincipal",
		[]interface{}{service, opts},
		&j,
	)

	return &j
}

// Reference an AWS service, optionally in a given region.
func NewServicePrincipal_Override(s ServicePrincipal, service *string, opts *ServicePrincipalOpts) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.ServicePrincipal",
		[]interface{}{service, opts},
		s,
	)
}

// Return the service principal name based on the region it's used in.
//
// Some service principal names used to be different for different partitions,
// and some were not. This method would return the appropriate region-specific
// service principal name, getting that information from the `region-info`
// module.
//
// These days all service principal names are standardized, and they are all
// of the form `<servicename>.amazonaws.com`.
//
// If the feature flag `@aws-cdk/aws-iam:standardizedServicePrincipals` is set, this
// method will always return its input. If this feature flag is not set, this
// method will perform the legacy behavior, which appends the region-specific
// domain suffix for some select services (for example, it would append `.cn`
// to some service principal names).
//
// Example:
//   principalName := iam.ServicePrincipal_ServicePrincipalName(jsii.String("ec2.amazonaws.com"))
//
func ServicePrincipal_ServicePrincipalName(service *string) *string {
	_init_.Initialize()

	if err := validateServicePrincipal_ServicePrincipalNameParameters(service); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.ServicePrincipal",
		"servicePrincipalName",
		[]interface{}{service},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) AddToAssumeRolePolicy(document PolicyDocument) {
	if err := s.validateAddToAssumeRolePolicyParameters(document); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addToAssumeRolePolicy",
		[]interface{}{document},
	)
}

func (s *jsiiProxy_ServicePrincipal) AddToPolicy(statement PolicyStatement) *bool {
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

func (s *jsiiProxy_ServicePrincipal) AddToPrincipalPolicy(_statement PolicyStatement) *AddToPrincipalPolicyResult {
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

func (s *jsiiProxy_ServicePrincipal) DedupeString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"dedupeString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) ToJSON() *map[string]*[]*string {
	var returns *map[string]*[]*string

	_jsii_.Invoke(
		s,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicePrincipal) WithConditions(conditions *map[string]interface{}) PrincipalBase {
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

func (s *jsiiProxy_ServicePrincipal) WithSessionTags() PrincipalBase {
	var returns PrincipalBase

	_jsii_.Invoke(
		s,
		"withSessionTags",
		nil, // no parameters
		&returns,
	)

	return returns
}

