// The CDK Construct Library for AWS::Cloud9
package awscdkawscloud9alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkawscloud9alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodecommit"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdkawscloud9alpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The class for different repository providers.
// Experimental.
type CloneRepository interface {
	PathComponent() *string
	RepositoryUrl() *string
}

// The jsii proxy struct for CloneRepository
type jsiiProxy_CloneRepository struct {
	_ byte // padding
}

func (j *jsiiProxy_CloneRepository) PathComponent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathComponent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloneRepository) RepositoryUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryUrl",
		&returns,
	)
	return returns
}


// import repository to cloud9 environment from AWS CodeCommit.
// Experimental.
func CloneRepository_FromCodeCommit(repository awscodecommit.IRepository, path *string) CloneRepository {
	_init_.Initialize()

	var returns CloneRepository

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-cloud9-alpha.CloneRepository",
		"fromCodeCommit",
		[]interface{}{repository, path},
		&returns,
	)

	return returns
}

// A Cloud9 Environment with Amazon EC2.
// Experimental.
type Ec2Environment interface {
	awscdk.Resource
	IEc2Environment
	Ec2EnvironmentArn() *string
	Ec2EnvironmentName() *string
	Env() *awscdk.ResourceEnvironment
	EnvironmentId() *string
	IdeUrl() *string
	Node() constructs.Node
	PhysicalName() *string
	Stack() awscdk.Stack
	Vpc() awsec2.IVpc
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
}

// The jsii proxy struct for Ec2Environment
type jsiiProxy_Ec2Environment struct {
	internal.Type__awscdkResource
	jsiiProxy_IEc2Environment
}

func (j *jsiiProxy_Ec2Environment) Ec2EnvironmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2EnvironmentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Environment) Ec2EnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2EnvironmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Environment) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Environment) EnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Environment) IdeUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ideUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Environment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Environment) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Environment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Environment) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Experimental.
func NewEc2Environment(scope constructs.Construct, id *string, props *Ec2EnvironmentProps) Ec2Environment {
	_init_.Initialize()

	j := jsiiProxy_Ec2Environment{}

	_jsii_.Create(
		"@aws-cdk/aws-cloud9-alpha.Ec2Environment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEc2Environment_Override(e Ec2Environment, scope constructs.Construct, id *string, props *Ec2EnvironmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-cloud9-alpha.Ec2Environment",
		[]interface{}{scope, id, props},
		e,
	)
}

// import from EnvironmentEc2Name.
// Experimental.
func Ec2Environment_FromEc2EnvironmentName(scope constructs.Construct, id *string, ec2EnvironmentName *string) IEc2Environment {
	_init_.Initialize()

	var returns IEc2Environment

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-cloud9-alpha.Ec2Environment",
		"fromEc2EnvironmentName",
		[]interface{}{scope, id, ec2EnvironmentName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Ec2Environment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-cloud9-alpha.Ec2Environment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Ec2Environment_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-cloud9-alpha.Ec2Environment",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (e *jsiiProxy_Ec2Environment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		e,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (e *jsiiProxy_Ec2Environment) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (e *jsiiProxy_Ec2Environment) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (e *jsiiProxy_Ec2Environment) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (e *jsiiProxy_Ec2Environment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for Ec2Environment.
// Experimental.
type Ec2EnvironmentProps struct {
	// The VPC that AWS Cloud9 will use to communicate with the Amazon Elastic Compute Cloud (Amazon EC2) instance.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// The AWS CodeCommit repository to be cloned.
	// Experimental.
	ClonedRepositories *[]CloneRepository `json:"clonedRepositories"`
	// Description of the environment.
	// Experimental.
	Description *string `json:"description"`
	// Name of the environment.
	// Experimental.
	Ec2EnvironmentName *string `json:"ec2EnvironmentName"`
	// The type of instance to connect to the environment.
	// Experimental.
	InstanceType awsec2.InstanceType `json:"instanceType"`
	// The subnetSelection of the VPC that AWS Cloud9 will use to communicate with the Amazon EC2 instance.
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection"`
}

// A Cloud9 Environment.
// Experimental.
type IEc2Environment interface {
	awscdk.IResource
	// The arn of the EnvironmentEc2.
	// Experimental.
	Ec2EnvironmentArn() *string
	// The name of the EnvironmentEc2.
	// Experimental.
	Ec2EnvironmentName() *string
}

// The jsii proxy for IEc2Environment
type jsiiProxy_IEc2Environment struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IEc2Environment) Ec2EnvironmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2EnvironmentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEc2Environment) Ec2EnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2EnvironmentName",
		&returns,
	)
	return returns
}

