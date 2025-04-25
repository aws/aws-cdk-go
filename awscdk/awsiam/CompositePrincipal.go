package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a principal that has multiple types of principals.
//
// A composite principal cannot
// have conditions. i.e. multiple ServicePrincipals that form a composite principal
//
// Example:
//   var vpc vpc
//
//   role := iam.NewRole(this, jsii.String("RDSDirectoryServicesRole"), &RoleProps{
//   	AssumedBy: iam.NewCompositePrincipal(
//   	iam.NewServicePrincipal(jsii.String("rds.amazonaws.com")),
//   	iam.NewServicePrincipal(jsii.String("directoryservice.rds.amazonaws.com"))),
//   	ManagedPolicies: []iManagedPolicy{
//   		iam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("service-role/AmazonRDSDirectoryServiceAccess")),
//   	},
//   })
//   instance := rds.NewDatabaseInstance(this, jsii.String("Instance"), &DatabaseInstanceProps{
//   	Engine: rds.DatabaseInstanceEngine_Mysql(&MySqlInstanceEngineProps{
//   		Version: rds.MysqlEngineVersion_VER_8_0_19(),
//   	}),
//   	Vpc: Vpc,
//   	Domain: jsii.String("d-????????"),
//   	 // The ID of the domain for the instance to join.
//   	DomainRole: role,
//   })
//
type CompositePrincipal interface {
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
	// Returns the principals that make up the CompositePrincipal.
	Principals() *[]IPrincipal
	// Adds IAM principals to the composite principal.
	//
	// Composite principals cannot have
	// conditions.
	AddPrincipals(principals ...IPrincipal) CompositePrincipal
	// Add the principal to the AssumeRolePolicyDocument.
	//
	// Add the statements to the AssumeRolePolicyDocument necessary to give this principal
	// permissions to assume the given role.
	AddToAssumeRolePolicy(doc PolicyDocument)
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

// The jsii proxy struct for CompositePrincipal
type jsiiProxy_CompositePrincipal struct {
	jsiiProxy_PrincipalBase
}

func (j *jsiiProxy_CompositePrincipal) AssumeRoleAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CompositePrincipal) GrantPrincipal() IPrincipal {
	var returns IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CompositePrincipal) PolicyFragment() PrincipalPolicyFragment {
	var returns PrincipalPolicyFragment
	_jsii_.Get(
		j,
		"policyFragment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CompositePrincipal) PrincipalAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CompositePrincipal) Principals() *[]IPrincipal {
	var returns *[]IPrincipal
	_jsii_.Get(
		j,
		"principals",
		&returns,
	)
	return returns
}


func NewCompositePrincipal(principals ...IPrincipal) CompositePrincipal {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range principals {
		args = append(args, a)
	}

	j := jsiiProxy_CompositePrincipal{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.CompositePrincipal",
		args,
		&j,
	)

	return &j
}

func NewCompositePrincipal_Override(c CompositePrincipal, principals ...IPrincipal) {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range principals {
		args = append(args, a)
	}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.CompositePrincipal",
		args,
		c,
	)
}

func (c *jsiiProxy_CompositePrincipal) AddPrincipals(principals ...IPrincipal) CompositePrincipal {
	args := []interface{}{}
	for _, a := range principals {
		args = append(args, a)
	}

	var returns CompositePrincipal

	_jsii_.Invoke(
		c,
		"addPrincipals",
		args,
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CompositePrincipal) AddToAssumeRolePolicy(doc PolicyDocument) {
	if err := c.validateAddToAssumeRolePolicyParameters(doc); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addToAssumeRolePolicy",
		[]interface{}{doc},
	)
}

func (c *jsiiProxy_CompositePrincipal) AddToPolicy(statement PolicyStatement) *bool {
	if err := c.validateAddToPolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"addToPolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CompositePrincipal) AddToPrincipalPolicy(_statement PolicyStatement) *AddToPrincipalPolicyResult {
	if err := c.validateAddToPrincipalPolicyParameters(_statement); err != nil {
		panic(err)
	}
	var returns *AddToPrincipalPolicyResult

	_jsii_.Invoke(
		c,
		"addToPrincipalPolicy",
		[]interface{}{_statement},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CompositePrincipal) DedupeString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"dedupeString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CompositePrincipal) ToJSON() *map[string]*[]*string {
	var returns *map[string]*[]*string

	_jsii_.Invoke(
		c,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CompositePrincipal) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CompositePrincipal) WithConditions(conditions *map[string]interface{}) PrincipalBase {
	if err := c.validateWithConditionsParameters(conditions); err != nil {
		panic(err)
	}
	var returns PrincipalBase

	_jsii_.Invoke(
		c,
		"withConditions",
		[]interface{}{conditions},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CompositePrincipal) WithSessionTags() PrincipalBase {
	var returns PrincipalBase

	_jsii_.Invoke(
		c,
		"withSessionTags",
		nil, // no parameters
		&returns,
	)

	return returns
}

