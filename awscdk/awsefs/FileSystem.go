package awsefs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsefs/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// The Elastic File System implementation of IFileSystem.
//
// It creates a new, empty file system in Amazon Elastic File System (Amazon EFS).
// It also creates mount target (AWS::EFS::MountTarget) implicitly to mount the
// EFS file system on an Amazon Elastic Compute Cloud (Amazon EC2) instance or another resource.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   role := iam.NewRole(this, jsii.String("ClientRole"), &RoleProps{
//   	AssumedBy: iam.NewAnyPrincipal(),
//   })
//   fileSystem := efs.NewFileSystem(this, jsii.String("MyEfsFileSystem"), &FileSystemProps{
//   	Vpc: ec2.NewVpc(this, jsii.String("VPC")),
//   	AllowAnonymousAccess: jsii.Boolean(true),
//   })
//
//   fileSystem.grantRead(role)
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html
//
type FileSystem interface {
	awscdk.Resource
	IFileSystem
	// The security groups/rules used to allow network connections to the file system.
	Connections() awsec2.Connections
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The ARN of the file system.
	FileSystemArn() *string
	// The ID of the file system, assigned by Amazon EFS.
	FileSystemId() *string
	// Dependable that can be depended upon to ensure the mount targets of the filesystem are ready.
	MountTargetsAvailable() constructs.IDependable
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// create access point from this filesystem.
	AddAccessPoint(id *string, accessPointOptions *AccessPointOptions) AccessPoint
	// Adds a statement to the resource policy associated with this file system.
	//
	// A resource policy will be automatically created upon the first call to `addToResourcePolicy`.
	//
	// Note that this does not work with imported file systems.
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Grant the actions defined in actions to the given grantee on this File System resource.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant read permissions for this file system to an IAM principal.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Grant read and write permissions for this file system to an IAM principal.
	GrantReadWrite(grantee awsiam.IGrantable) awsiam.Grant
	// As root user, grant read and write permissions for this file system to an IAM principal.
	GrantRootAccess(grantee awsiam.IGrantable) awsiam.Grant
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for FileSystem
type jsiiProxy_FileSystem struct {
	internal.Type__awscdkResource
	jsiiProxy_IFileSystem
}

func (j *jsiiProxy_FileSystem) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSystem) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSystem) FileSystemArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSystem) FileSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSystem) MountTargetsAvailable() constructs.IDependable {
	var returns constructs.IDependable
	_jsii_.Get(
		j,
		"mountTargetsAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSystem) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSystem) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FileSystem) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Constructor for creating a new EFS FileSystem.
func NewFileSystem(scope constructs.Construct, id *string, props *FileSystemProps) FileSystem {
	_init_.Initialize()

	if err := validateNewFileSystemParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_FileSystem{}

	_jsii_.Create(
		"aws-cdk-lib.aws_efs.FileSystem",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructor for creating a new EFS FileSystem.
func NewFileSystem_Override(f FileSystem, scope constructs.Construct, id *string, props *FileSystemProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_efs.FileSystem",
		[]interface{}{scope, id, props},
		f,
	)
}

// Import an existing File System from the given properties.
func FileSystem_FromFileSystemAttributes(scope constructs.Construct, id *string, attrs *FileSystemAttributes) IFileSystem {
	_init_.Initialize()

	if err := validateFileSystem_FromFileSystemAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IFileSystem

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_efs.FileSystem",
		"fromFileSystemAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func FileSystem_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFileSystem_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_efs.FileSystem",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func FileSystem_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateFileSystem_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_efs.FileSystem",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func FileSystem_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateFileSystem_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_efs.FileSystem",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func FileSystem_DEFAULT_PORT() *float64 {
	_init_.Initialize()
	var returns *float64
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_efs.FileSystem",
		"DEFAULT_PORT",
		&returns,
	)
	return returns
}

func FileSystem_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_efs.FileSystem",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FileSystem) AddAccessPoint(id *string, accessPointOptions *AccessPointOptions) AccessPoint {
	if err := f.validateAddAccessPointParameters(id, accessPointOptions); err != nil {
		panic(err)
	}
	var returns AccessPoint

	_jsii_.Invoke(
		f,
		"addAccessPoint",
		[]interface{}{id, accessPointOptions},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSystem) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	if err := f.validateAddToResourcePolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		f,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSystem) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := f.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (f *jsiiProxy_FileSystem) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSystem) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := f.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSystem) GetResourceNameAttribute(nameAttr *string) *string {
	if err := f.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSystem) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := f.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		f,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSystem) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	if err := f.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		f,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSystem) GrantReadWrite(grantee awsiam.IGrantable) awsiam.Grant {
	if err := f.validateGrantReadWriteParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		f,
		"grantReadWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSystem) GrantRootAccess(grantee awsiam.IGrantable) awsiam.Grant {
	if err := f.validateGrantRootAccessParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		f,
		"grantRootAccess",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FileSystem) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

