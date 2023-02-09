package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An IAM principal that represents an AWS service (i.e. sqs.amazonaws.com).
//
// Example:
//   lambdaRole := iam.NewRole(this, jsii.String("Role"), &roleProps{
//   	assumedBy: iam.NewServicePrincipal(jsii.String("lambda.amazonaws.com")),
//   	description: jsii.String("Example role..."),
//   })
//
//   stream := kinesis.NewStream(this, jsii.String("MyEncryptedStream"), &streamProps{
//   	encryption: kinesis.streamEncryption_KMS,
//   })
//
//   // give lambda permissions to read stream
//   stream.grantRead(lambdaRole)
//
// Experimental.
type ServicePrincipal interface {
	PrincipalBase
	// When this Principal is used in an AssumeRole policy, the action to use.
	// Experimental.
	AssumeRoleAction() *string
	// The principal to grant permissions to.
	// Experimental.
	GrantPrincipal() IPrincipal
	// Return the policy fragment that identifies this principal in a Policy.
	// Experimental.
	PolicyFragment() PrincipalPolicyFragment
	// The AWS account ID of this principal.
	//
	// Can be undefined when the account is not known
	// (for example, for service principals).
	// Can be a Token - in that case,
	// it's assumed to be AWS::AccountId.
	// Experimental.
	PrincipalAccount() *string
	// AWS service (i.e. sqs.amazonaws.com).
	// Experimental.
	Service() *string
	// Add the princpial to the AssumeRolePolicyDocument.
	//
	// Add the statements to the AssumeRolePolicyDocument necessary to give this principal
	// permissions to assume the given role.
	// Experimental.
	AddToAssumeRolePolicy(document PolicyDocument)
	// Add to the policy of this principal.
	// Experimental.
	AddToPolicy(statement PolicyStatement) *bool
	// Add to the policy of this principal.
	// Experimental.
	AddToPrincipalPolicy(_statement PolicyStatement) *AddToPrincipalPolicyResult
	// Return whether or not this principal is equal to the given principal.
	// Experimental.
	DedupeString() *string
	// JSON-ify the principal.
	//
	// Used when JSON.stringify() is called
	// Experimental.
	ToJSON() *map[string]*[]*string
	// Returns a string representation of an object.
	// Experimental.
	ToString() *string
	// Returns a new PrincipalWithConditions using this principal as the base, with the passed conditions added.
	//
	// When there is a value for the same operator and key in both the principal and the
	// conditions parameter, the value from the conditions parameter will be used.
	//
	// Returns: a new PrincipalWithConditions object.
	// Experimental.
	WithConditions(conditions *map[string]interface{}) PrincipalBase
	// Returns a new principal using this principal as the base, with session tags enabled.
	//
	// Returns: a new SessionTagsPrincipal object.
	// Experimental.
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


// Experimental.
func NewServicePrincipal(service *string, opts *ServicePrincipalOpts) ServicePrincipal {
	_init_.Initialize()

	if err := validateNewServicePrincipalParameters(service, opts); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServicePrincipal{}

	_jsii_.Create(
		"monocdk.aws_iam.ServicePrincipal",
		[]interface{}{service, opts},
		&j,
	)

	return &j
}

// Experimental.
func NewServicePrincipal_Override(s ServicePrincipal, service *string, opts *ServicePrincipalOpts) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iam.ServicePrincipal",
		[]interface{}{service, opts},
		s,
	)
}

// Translate the given service principal name based on the region it's used in.
//
// For example, for Chinese regions this may (depending on whether that's necessary
// for the given service principal) append `.cn` to the name.
//
// The `region-info` module is used to obtain this information.
//
// Example:
//   principalName := iam.servicePrincipal.servicePrincipalName(jsii.String("ec2.amazonaws.com"))
//
// Experimental.
func ServicePrincipal_ServicePrincipalName(service *string) *string {
	_init_.Initialize()

	if err := validateServicePrincipal_ServicePrincipalNameParameters(service); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"monocdk.aws_iam.ServicePrincipal",
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

