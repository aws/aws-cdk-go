package awsiam

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
)

// A PolicyDocument is a collection of statements.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   myFileSystemPolicy := iam.NewPolicyDocument(&PolicyDocumentProps{
//   	Statements: []policyStatement{
//   		iam.NewPolicyStatement(&PolicyStatementProps{
//   			Actions: []*string{
//   				jsii.String("elasticfilesystem:ClientWrite"),
//   				jsii.String("elasticfilesystem:ClientMount"),
//   			},
//   			Principals: []iPrincipal{
//   				iam.NewAccountRootPrincipal(),
//   			},
//   			Resources: []*string{
//   				jsii.String("*"),
//   			},
//   			Conditions: map[string]interface{}{
//   				"Bool": map[string]*string{
//   					"elasticfilesystem:AccessedViaMountTarget": jsii.String("true"),
//   				},
//   			},
//   		}),
//   	},
//   })
//
//   fileSystem := efs.NewFileSystem(this, jsii.String("MyEfsFileSystem"), &FileSystemProps{
//   	Vpc: ec2.NewVpc(this, jsii.String("VPC")),
//   	FileSystemPolicy: myFileSystemPolicy,
//   })
//
type PolicyDocument interface {
	awscdk.IResolvable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// This may return an array with a single informational element indicating how
	// to get this property populated, if it was skipped for performance reasons.
	CreationStack() *[]*string
	// Whether the policy document contains any statements.
	IsEmpty() *bool
	// The number of statements already added to this policy.
	//
	// Can be used, for example, to generate unique "sid"s within the policy.
	StatementCount() *float64
	// Adds a statement to the policy document.
	AddStatements(statement ...PolicyStatement)
	// Produce the Token's value at resolution time.
	Resolve(context awscdk.IResolveContext) interface{}
	// JSON-ify the document.
	//
	// Used when JSON.stringify() is called
	ToJSON() interface{}
	// Encode the policy document as a string.
	ToString() *string
	// Validate that all policy statements in the policy document satisfies the requirements for any policy.
	//
	// Returns: An array of validation error messages, or an empty array if the document is valid.
	// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#access_policies-json
	//
	ValidateForAnyPolicy() *[]*string
	// Validate that all policy statements in the policy document satisfies the requirements for an identity-based policy.
	//
	// Returns: An array of validation error messages, or an empty array if the document is valid.
	// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#access_policies-json
	//
	ValidateForIdentityPolicy() *[]*string
	// Validate that all policy statements in the policy document satisfies the requirements for a resource-based policy.
	//
	// Returns: An array of validation error messages, or an empty array if the document is valid.
	// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#access_policies-json
	//
	ValidateForResourcePolicy() *[]*string
}

// The jsii proxy struct for PolicyDocument
type jsiiProxy_PolicyDocument struct {
	internal.Type__awscdkIResolvable
}

func (j *jsiiProxy_PolicyDocument) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDocument) IsEmpty() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isEmpty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDocument) StatementCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"statementCount",
		&returns,
	)
	return returns
}


func NewPolicyDocument(props *PolicyDocumentProps) PolicyDocument {
	_init_.Initialize()

	if err := validateNewPolicyDocumentParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_PolicyDocument{}

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.PolicyDocument",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewPolicyDocument_Override(p PolicyDocument, props *PolicyDocumentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_iam.PolicyDocument",
		[]interface{}{props},
		p,
	)
}

// Creates a new PolicyDocument based on the object provided.
//
// This will accept an object created from the `.toJSON()` call
func PolicyDocument_FromJson(obj interface{}) PolicyDocument {
	_init_.Initialize()

	if err := validatePolicyDocument_FromJsonParameters(obj); err != nil {
		panic(err)
	}
	var returns PolicyDocument

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_iam.PolicyDocument",
		"fromJson",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyDocument) AddStatements(statement ...PolicyStatement) {
	args := []interface{}{}
	for _, a := range statement {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		p,
		"addStatements",
		args,
	)
}

func (p *jsiiProxy_PolicyDocument) Resolve(context awscdk.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyDocument) ToJSON() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toJSON",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyDocument) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyDocument) ValidateForAnyPolicy() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"validateForAnyPolicy",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyDocument) ValidateForIdentityPolicy() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"validateForIdentityPolicy",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PolicyDocument) ValidateForResourcePolicy() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"validateForResourcePolicy",
		nil, // no parameters
		&returns,
	)

	return returns
}

