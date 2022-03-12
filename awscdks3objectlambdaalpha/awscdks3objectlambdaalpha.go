// The CDK Construct Library for AWS::S3ObjectLambda
package awscdks3objectlambdaalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdks3objectlambdaalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdks3objectlambdaalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An S3 object lambda access point for intercepting and transforming `GetObject` requests.
//
// TODO: EXAMPLE
//
// Experimental.
type AccessPoint interface {
	awscdk.Resource
	IAccessPoint
	AccessPointArn() *string
	AccessPointCreationDate() *string
	AccessPointName() *string
	DomainName() *string
	Env() *awscdk.ResourceEnvironment
	Node() constructs.Node
	PhysicalName() *string
	RegionalDomainName() *string
	Stack() awscdk.Stack
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	ToString() *string
	VirtualHostedUrlForObject(key *string, options *awss3.VirtualHostedStyleUrlOptions) *string
}

// The jsii proxy struct for AccessPoint
type jsiiProxy_AccessPoint struct {
	internal.Type__awscdkResource
	jsiiProxy_IAccessPoint
}

func (j *jsiiProxy_AccessPoint) AccessPointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPoint) AccessPointCreationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPointCreationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPoint) AccessPointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPointName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPoint) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPoint) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPoint) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPoint) RegionalDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessPoint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewAccessPoint(scope constructs.Construct, id *string, props *AccessPointProps) AccessPoint {
	_init_.Initialize()

	j := jsiiProxy_AccessPoint{}

	_jsii_.Create(
		"@aws-cdk/aws-s3objectlambda-alpha.AccessPoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAccessPoint_Override(a AccessPoint, scope constructs.Construct, id *string, props *AccessPointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-s3objectlambda-alpha.AccessPoint",
		[]interface{}{scope, id, props},
		a,
	)
}

// Reference an existing AccessPoint defined outside of the CDK code.
// Experimental.
func AccessPoint_FromAccessPointAttributes(scope constructs.Construct, id *string, attrs *AccessPointAttributes) IAccessPoint {
	_init_.Initialize()

	var returns IAccessPoint

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-s3objectlambda-alpha.AccessPoint",
		"fromAccessPointAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func AccessPoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-s3objectlambda-alpha.AccessPoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func AccessPoint_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-s3objectlambda-alpha.AccessPoint",
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
func (a *jsiiProxy_AccessPoint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Experimental.
func (a *jsiiProxy_AccessPoint) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		a,
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
func (a *jsiiProxy_AccessPoint) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		a,
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
func (a *jsiiProxy_AccessPoint) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (a *jsiiProxy_AccessPoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Implement the {@link IAccessPoint.virtualHostedUrlForObject} method.
// Experimental.
func (a *jsiiProxy_AccessPoint) VirtualHostedUrlForObject(key *string, options *awss3.VirtualHostedStyleUrlOptions) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"virtualHostedUrlForObject",
		[]interface{}{key, options},
		&returns,
	)

	return returns
}

// The access point resource attributes.
//
// TODO: EXAMPLE
//
// Experimental.
type AccessPointAttributes struct {
	// The ARN of the access point.
	// Experimental.
	AccessPointArn *string `json:"accessPointArn" yaml:"accessPointArn"`
	// The creation data of the access point.
	// Experimental.
	AccessPointCreationDate *string `json:"accessPointCreationDate" yaml:"accessPointCreationDate"`
}

// The S3 object lambda access point configuration.
//
// TODO: EXAMPLE
//
// Experimental.
type AccessPointProps struct {
	// The bucket to which this access point belongs.
	// Experimental.
	Bucket awss3.IBucket `json:"bucket" yaml:"bucket"`
	// The Lambda function used to transform objects.
	// Experimental.
	Handler awslambda.IFunction `json:"handler" yaml:"handler"`
	// The name of the S3 object lambda access point.
	// Experimental.
	AccessPointName *string `json:"accessPointName" yaml:"accessPointName"`
	// Whether CloudWatch metrics are enabled for the access point.
	// Experimental.
	CloudWatchMetricsEnabled *bool `json:"cloudWatchMetricsEnabled" yaml:"cloudWatchMetricsEnabled"`
	// Additional JSON that provides supplemental data passed to the Lambda function on every request.
	// Experimental.
	Payload *map[string]interface{} `json:"payload" yaml:"payload"`
	// Whether the Lambda function can process `GetObject-PartNumber` requests.
	// Experimental.
	SupportsGetObjectPartNumber *bool `json:"supportsGetObjectPartNumber" yaml:"supportsGetObjectPartNumber"`
	// Whether the Lambda function can process `GetObject-Range` requests.
	// Experimental.
	SupportsGetObjectRange *bool `json:"supportsGetObjectRange" yaml:"supportsGetObjectRange"`
}

// The interface that represents the AccessPoint resource.
// Experimental.
type IAccessPoint interface {
	awscdk.IResource
	// The virtual hosted-style URL of an S3 object through this access point.
	//
	// Specify `regional: false` at the options for non-regional URL.
	//
	// Returns: an ObjectS3Url token
	// Experimental.
	VirtualHostedUrlForObject(key *string, options *awss3.VirtualHostedStyleUrlOptions) *string
	// The ARN of the access point.
	// Experimental.
	AccessPointArn() *string
	// The creation data of the access point.
	// Experimental.
	AccessPointCreationDate() *string
	// The IPv4 DNS name of the access point.
	// Experimental.
	DomainName() *string
	// The regional domain name of the access point.
	// Experimental.
	RegionalDomainName() *string
}

// The jsii proxy for IAccessPoint
type jsiiProxy_IAccessPoint struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IAccessPoint) VirtualHostedUrlForObject(key *string, options *awss3.VirtualHostedStyleUrlOptions) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"virtualHostedUrlForObject",
		[]interface{}{key, options},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IAccessPoint) AccessPointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessPoint) AccessPointCreationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPointCreationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessPoint) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccessPoint) RegionalDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalDomainName",
		&returns,
	)
	return returns
}

