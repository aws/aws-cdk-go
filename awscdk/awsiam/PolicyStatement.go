package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a statement in an IAM policy document.
//
// Example:
//   crossAccountRoleArn := "arn:aws:iam::OTHERACCOUNT:role/CrossAccountRoleName" // arn of role deployed in separate account
//
//   callRegion := "us-west-1" // sdk call to be made in specified region (optional)
//
//    // sdk call to be made in specified region (optional)
//   cr.NewAwsCustomResource(this, jsii.String("CrossAccount"), &AwsCustomResourceProps{
//   	OnCreate: &AwsSdkCall{
//   		AssumedRoleArn: crossAccountRoleArn,
//   		Region: callRegion,
//   		 // optional
//   		Service: jsii.String("sts"),
//   		Action: jsii.String("GetCallerIdentity"),
//   		PhysicalResourceId: cr.PhysicalResourceId_Of(jsii.String("id")),
//   	},
//   	Policy: cr.AwsCustomResourcePolicy_FromStatements([]policyStatement{
//   		iam.*policyStatement_FromJson(map[string]*string{
//   			"Effect": jsii.String("Allow"),
//   			"Action": jsii.String("sts:AssumeRole"),
//   			"Resource": crossAccountRoleArn,
//   		}),
//   	}),
//   })
//
type PolicyStatement interface {
	// The Actions added to this statement.
	Actions() *[]*string
	// The conditions added to this statement.
	Conditions() interface{}
	// Whether to allow or deny the actions in this statement Set effect for this statement.
	Effect() Effect
	SetEffect(val Effect)
	// Whether the PolicyStatement has been frozen.
	//
	// The statement object is frozen when `freeze()` is called.
	Frozen() *bool
	// Indicates if this permission has a "Principal" section.
	HasPrincipal() *bool
	// Indicates if this permission has at least one resource associated with it.
	HasResource() *bool
	// The NotActions added to this statement.
	NotActions() *[]*string
	// The NotPrincipals added to this statement.
	NotPrincipals() *[]IPrincipal
	// The NotResources added to this statement.
	NotResources() *[]*string
	// The Principals added to this statement.
	Principals() *[]IPrincipal
	// The Resources added to this statement.
	Resources() *[]*string
	// Statement ID for this statement Set Statement ID for this statement.
	Sid() *string
	SetSid(val *string)
	// Add a `StringEquals` condition that limits to a given account from `sts:ExternalId`.
	//
	// This method can only be called once: subsequent calls will overwrite earlier calls.
	// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user_externalid.html
	//
	AddAccountCondition(accountId *string)
	// Adds an AWS account root user principal to this policy statement.
	AddAccountRootPrincipal()
	// Specify allowed actions into the "Action" section of the policy statement.
	// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_action.html
	//
	AddActions(actions ...*string)
	// Adds a ``"*"`` resource to this statement.
	AddAllResources()
	// Adds all identities in all accounts ("*") to this policy statement.
	AddAnyPrincipal()
	// Specify a principal using the ARN  identifier of the principal.
	//
	// You cannot specify IAM groups and instance profiles as principals.
	AddArnPrincipal(arn *string)
	// Specify AWS account ID as the principal entity to the "Principal" section of a policy statement.
	AddAwsAccountPrincipal(accountId *string)
	// Adds a canonical user ID principal to this policy document.
	AddCanonicalUserPrincipal(canonicalUserId *string)
	// Add a condition to the Policy.
	//
	// If multiple calls are made to add a condition with the same operator and field, only
	// the last one wins. For example:
	//
	// ```ts
	// declare const stmt: iam.PolicyStatement;
	//
	// stmt.addCondition('StringEquals', { 'aws:SomeField': '1' });
	// stmt.addCondition('StringEquals', { 'aws:SomeField': '2' });
	// ```
	//
	// Will end up with the single condition `StringEquals: { 'aws:SomeField': '2' }`.
	//
	// If you meant to add a condition to say that the field can be *either* `1` or `2`, write
	// this:
	//
	// ```ts
	// declare const stmt: iam.PolicyStatement;
	//
	// stmt.addCondition('StringEquals', { 'aws:SomeField': ['1', '2'] });
	// ```.
	AddCondition(key *string, value interface{})
	// Add multiple conditions to the Policy.
	//
	// See the `addCondition` function for a caveat on calling this method multiple times.
	AddConditions(conditions *map[string]interface{})
	// Adds a federated identity provider such as Amazon Cognito to this policy statement.
	AddFederatedPrincipal(federated interface{}, conditions *map[string]interface{})
	// Explicitly allow all actions except the specified list of actions into the "NotAction" section of the policy document.
	// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_notaction.html
	//
	AddNotActions(notActions ...*string)
	// Specify principals that is not allowed or denied access to the "NotPrincipal" section of a policy statement.
	// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_notprincipal.html
	//
	AddNotPrincipals(notPrincipals ...IPrincipal)
	// Specify resources that this policy statement will not apply to in the "NotResource" section of this policy statement.
	//
	// All resources except the specified list will be matched.
	// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_notresource.html
	//
	AddNotResources(arns ...*string)
	// Adds principals to the "Principal" section of a policy statement.
	// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_principal.html
	//
	AddPrincipals(principals ...IPrincipal)
	// Specify resources that this policy statement applies into the "Resource" section of this policy statement.
	// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_resource.html
	//
	AddResources(arns ...*string)
	// Adds a service principal to this policy statement.
	AddServicePrincipal(service *string, opts *ServicePrincipalOpts)
	// Add an `StringEquals` condition that limits to a given account from `aws:SourceAccount`.
	//
	// This method can only be called once: subsequent calls will overwrite earlier calls.
	// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_condition-keys.html#condition-keys-sourceaccount
	//
	AddSourceAccountCondition(accountId *string)
	// Add an `ArnEquals` condition that limits to a given resource arn from `aws:SourceArn`.
	//
	// This method can only be called once: subsequent calls will overwrite earlier calls.
	// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_condition-keys.html#condition-keys-sourcearn
	//
	AddSourceArnCondition(arn *string)
	// Create a new `PolicyStatement` with the same exact properties as this one, except for the overrides.
	Copy(overrides *PolicyStatementProps) PolicyStatement
	// Make the PolicyStatement immutable.
	//
	// After calling this, any of the `addXxx()` methods will throw an exception.
	//
	// Libraries that lazily generate statement bodies can override this method to
	// fill the actual PolicyStatement fields. Be aware that this method may be called
	// multiple times.
	Freeze() PolicyStatement
	// JSON-ify the statement.
	//
	// Used when JSON.stringify() is called
	ToJSON() interface{}
	// JSON-ify the policy statement.
	//
	// Used when JSON.stringify() is called
	ToStatementJson() interface{}
	// String representation of this policy statement.
	ToString() *string
	// Validate that the policy statement satisfies base requirements for a policy.
	//
	// Returns: An array of validation error messages, or an empty array if the statement is valid.
	ValidateForAnyPolicy() *[]*string
	// Validate that the policy statement satisfies all requirements for an identity-based policy.
	//
	// Returns: An array of validation error messages, or an empty array if the statement is valid.
	ValidateForIdentityPolicy() *[]*string
	// Validate that the policy statement satisfies all requirements for a resource-based policy.
	//
	// Returns: An array of validation error messages, or an empty array if the statement is valid.
	ValidateForResourcePolicy() *[]*string
}

// The jsii proxy struct for PolicyStatement
type jsiiProxy_PolicyStatement struct {
	_ byte // padding
}

func (j *jsiiProxy_PolicyStatement) Actions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyStatement) Conditions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyStatement) Effect() Effect {
	var returns Effect
	_jsii_.Get(
		j,
		"effect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyStatement) Frozen() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"frozen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyStatement) HasPrincipal() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"hasPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyStatement) HasResource() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"hasResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyStatement) NotActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyStatement) NotPrincipals() *[]IPrincipal {
	var returns *[]IPrincipal
	_jsii_.Get(
		j,
		"notPrincipals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyStatement) NotResources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyStatement) Principals() *[]IPrincipal {
	var returns *[]IPrincipal
	_jsii_.Get(
		j,
		"principals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyStatement) Resources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyStatement) Sid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sid",
		&returns,
	)
	return returns
}


func NewPolicyStatement(props *PolicyStatementProps) PolicyStatement {
	_init_.Initialize()

	if err := validateNewPolicyStatementParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_PolicyStatement{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.PolicyStatement",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewPolicyStatement_Override(p PolicyStatement, props *PolicyStatementProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.PolicyStatement",
		[]interface{}{props},
		p,
	)
}

func (j *jsiiProxy_PolicyStatement)SetEffect(val Effect) {
	if err := j.validateSetEffectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"effect",
		val,
	)
}

func (j *jsiiProxy_PolicyStatement)SetSid(val *string) {
	_jsii_.Set(
		j,
		"sid",
		val,
	)
}

// Creates a new PolicyStatement based on the object provided.
//
// This will accept an object created from the `.toJSON()` call
func PolicyStatement_FromJson(obj interface{}) PolicyStatement {
	_init_.Initialize()

	if err := validatePolicyStatement_FromJsonParameters(obj); err != nil {
		panic(err)
	}
	var returns PolicyStatement

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.PolicyStatement",
		"fromJson",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyStatement) AddAccountCondition(accountId *string) {
	if err := p.validateAddAccountConditionParameters(accountId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addAccountCondition",
		[]interface{}{accountId},
	)
}

func (p *jsiiProxy_PolicyStatement) AddAccountRootPrincipal() {
	_jsii_.InvokeVoid(
		p,
		"addAccountRootPrincipal",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyStatement) AddActions(actions ...*string) {
	args := []interface{}{}
	for _, a := range actions {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		p,
		"addActions",
		args,
	)
}

func (p *jsiiProxy_PolicyStatement) AddAllResources() {
	_jsii_.InvokeVoid(
		p,
		"addAllResources",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyStatement) AddAnyPrincipal() {
	_jsii_.InvokeVoid(
		p,
		"addAnyPrincipal",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PolicyStatement) AddArnPrincipal(arn *string) {
	if err := p.validateAddArnPrincipalParameters(arn); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addArnPrincipal",
		[]interface{}{arn},
	)
}

func (p *jsiiProxy_PolicyStatement) AddAwsAccountPrincipal(accountId *string) {
	if err := p.validateAddAwsAccountPrincipalParameters(accountId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addAwsAccountPrincipal",
		[]interface{}{accountId},
	)
}

func (p *jsiiProxy_PolicyStatement) AddCanonicalUserPrincipal(canonicalUserId *string) {
	if err := p.validateAddCanonicalUserPrincipalParameters(canonicalUserId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addCanonicalUserPrincipal",
		[]interface{}{canonicalUserId},
	)
}

func (p *jsiiProxy_PolicyStatement) AddCondition(key *string, value interface{}) {
	if err := p.validateAddConditionParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addCondition",
		[]interface{}{key, value},
	)
}

func (p *jsiiProxy_PolicyStatement) AddConditions(conditions *map[string]interface{}) {
	if err := p.validateAddConditionsParameters(conditions); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addConditions",
		[]interface{}{conditions},
	)
}

func (p *jsiiProxy_PolicyStatement) AddFederatedPrincipal(federated interface{}, conditions *map[string]interface{}) {
	if err := p.validateAddFederatedPrincipalParameters(federated, conditions); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addFederatedPrincipal",
		[]interface{}{federated, conditions},
	)
}

func (p *jsiiProxy_PolicyStatement) AddNotActions(notActions ...*string) {
	args := []interface{}{}
	for _, a := range notActions {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		p,
		"addNotActions",
		args,
	)
}

func (p *jsiiProxy_PolicyStatement) AddNotPrincipals(notPrincipals ...IPrincipal) {
	args := []interface{}{}
	for _, a := range notPrincipals {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		p,
		"addNotPrincipals",
		args,
	)
}

func (p *jsiiProxy_PolicyStatement) AddNotResources(arns ...*string) {
	args := []interface{}{}
	for _, a := range arns {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		p,
		"addNotResources",
		args,
	)
}

func (p *jsiiProxy_PolicyStatement) AddPrincipals(principals ...IPrincipal) {
	args := []interface{}{}
	for _, a := range principals {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		p,
		"addPrincipals",
		args,
	)
}

func (p *jsiiProxy_PolicyStatement) AddResources(arns ...*string) {
	args := []interface{}{}
	for _, a := range arns {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		p,
		"addResources",
		args,
	)
}

func (p *jsiiProxy_PolicyStatement) AddServicePrincipal(service *string, opts *ServicePrincipalOpts) {
	if err := p.validateAddServicePrincipalParameters(service, opts); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addServicePrincipal",
		[]interface{}{service, opts},
	)
}

func (p *jsiiProxy_PolicyStatement) AddSourceAccountCondition(accountId *string) {
	if err := p.validateAddSourceAccountConditionParameters(accountId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addSourceAccountCondition",
		[]interface{}{accountId},
	)
}

func (p *jsiiProxy_PolicyStatement) AddSourceArnCondition(arn *string) {
	if err := p.validateAddSourceArnConditionParameters(arn); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addSourceArnCondition",
		[]interface{}{arn},
	)
}

func (p *jsiiProxy_PolicyStatement) Copy(overrides *PolicyStatementProps) PolicyStatement {
	if err := p.validateCopyParameters(overrides); err != nil {
		panic(err)
	}
	var returns PolicyStatement

	_jsii_.Invoke(
		p,
		"copy",
		[]interface{}{overrides},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyStatement) Freeze() PolicyStatement {
	var returns PolicyStatement

	_jsii_.Invoke(
		p,
		"freeze",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyStatement) ToJSON() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyStatement) ToStatementJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toStatementJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyStatement) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyStatement) ValidateForAnyPolicy() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"validateForAnyPolicy",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyStatement) ValidateForIdentityPolicy() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"validateForIdentityPolicy",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyStatement) ValidateForResourcePolicy() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"validateForResourcePolicy",
		nil, // no parameters
		&returns,
	)

	return returns
}

