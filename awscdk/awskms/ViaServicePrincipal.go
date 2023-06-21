package awskms

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms/internal"
)

// A principal to allow access to a key if it's being used through another AWS service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var principal iPrincipal
//
//   viaServicePrincipal := awscdk.Aws_kms.NewViaServicePrincipal(jsii.String("serviceName"), principal)
//
type ViaServicePrincipal interface {
	awsiam.PrincipalBase
	// When this Principal is used in an AssumeRole policy, the action to use.
	AssumeRoleAction() *string
	// The principal to grant permissions to.
	GrantPrincipal() awsiam.IPrincipal
	// Return the policy fragment that identifies this principal in a Policy.
	PolicyFragment() awsiam.PrincipalPolicyFragment
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
	AddToAssumeRolePolicy(document awsiam.PolicyDocument)
	// Add to the policy of this principal.
	AddToPolicy(statement awsiam.PolicyStatement) *bool
	// Add to the policy of this principal.
	AddToPrincipalPolicy(_statement awsiam.PolicyStatement) *awsiam.AddToPrincipalPolicyResult
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
	WithConditions(conditions *map[string]interface{}) awsiam.PrincipalBase
	// Returns a new principal using this principal as the base, with session tags enabled.
	//
	// Returns: a new SessionTagsPrincipal object.
	WithSessionTags() awsiam.PrincipalBase
}

// The jsii proxy struct for ViaServicePrincipal
type jsiiProxy_ViaServicePrincipal struct {
	internal.Type__awsiamPrincipalBase
}

func (j *jsiiProxy_ViaServicePrincipal) AssumeRoleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ViaServicePrincipal) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ViaServicePrincipal) PolicyFragment() awsiam.PrincipalPolicyFragment {
	var returns awsiam.PrincipalPolicyFragment
	_jsii_.Get(
		j,
		"policyFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ViaServicePrincipal) PrincipalAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalAccount",
		&returns,
	)
	return returns
}


func NewViaServicePrincipal(serviceName *string, basePrincipal awsiam.IPrincipal) ViaServicePrincipal {
	_init_.Initialize()

	if err := validateNewViaServicePrincipalParameters(serviceName); err != nil {
		panic(err)
	}
	j := jsiiProxy_ViaServicePrincipal{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kms.ViaServicePrincipal",
		[]interface{}{serviceName, basePrincipal},
		&j,
	)

	return &j
}

func NewViaServicePrincipal_Override(v ViaServicePrincipal, serviceName *string, basePrincipal awsiam.IPrincipal) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kms.ViaServicePrincipal",
		[]interface{}{serviceName, basePrincipal},
		v,
	)
}

func (v *jsiiProxy_ViaServicePrincipal) AddToAssumeRolePolicy(document awsiam.PolicyDocument) {
	if err := v.validateAddToAssumeRolePolicyParameters(document); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addToAssumeRolePolicy",
		[]interface{}{document},
	)
}

func (v *jsiiProxy_ViaServicePrincipal) AddToPolicy(statement awsiam.PolicyStatement) *bool {
	if err := v.validateAddToPolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		v,
		"addToPolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ViaServicePrincipal) AddToPrincipalPolicy(_statement awsiam.PolicyStatement) *awsiam.AddToPrincipalPolicyResult {
	if err := v.validateAddToPrincipalPolicyParameters(_statement); err != nil {
		panic(err)
	}
	var returns *awsiam.AddToPrincipalPolicyResult

	_jsii_.Invoke(
		v,
		"addToPrincipalPolicy",
		[]interface{}{_statement},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ViaServicePrincipal) DedupeString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"dedupeString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ViaServicePrincipal) ToJSON() *map[string]*[]*string {
	var returns *map[string]*[]*string

	_jsii_.Invoke(
		v,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ViaServicePrincipal) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ViaServicePrincipal) WithConditions(conditions *map[string]interface{}) awsiam.PrincipalBase {
	if err := v.validateWithConditionsParameters(conditions); err != nil {
		panic(err)
	}
	var returns awsiam.PrincipalBase

	_jsii_.Invoke(
		v,
		"withConditions",
		[]interface{}{conditions},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_ViaServicePrincipal) WithSessionTags() awsiam.PrincipalBase {
	var returns awsiam.PrincipalBase

	_jsii_.Invoke(
		v,
		"withSessionTags",
		nil, // no parameters
		&returns,
	)

	return returns
}

