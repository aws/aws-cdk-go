package awss3assets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An asset represents a local file or directory, which is automatically uploaded to S3 and then can be referenced within a CDK application.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   asset := awscdk.NewAsset(this, jsii.String("BundledAsset"), &AssetProps{
//   	Path: jsii.String("/path/to/asset"),
//   	Bundling: &BundlingOptions{
//   		Image: cdk.DockerImage_FromRegistry(jsii.String("alpine")),
//   		Command: []*string{
//   			jsii.String("command-that-produces-an-archive.sh"),
//   		},
//   		OutputType: cdk.BundlingOutput_NOT_ARCHIVED,
//   	},
//   })
//
type Asset interface {
	constructs.Construct
	awscdk.IAsset
	// A hash of this asset, which is available at construction time.
	//
	// As this is a plain string, it
	// can be used in construct IDs in order to enforce creation of a new resource when the content
	// hash has changed.
	AssetHash() *string
	// The path to the asset, relative to the current Cloud Assembly.
	//
	// If asset staging is disabled, this will just be the original path.
	// If asset staging is enabled it will be the staged path.
	AssetPath() *string
	// The S3 bucket in which this asset resides.
	Bucket() awss3.IBucket
	// Attribute which represents the S3 HTTP URL of this asset.
	//
	// For example, `https://s3.us-west-1.amazonaws.com/bucket/key`
	HttpUrl() *string
	// Indicates if this asset is a single file.
	//
	// Allows constructs to ensure that the
	// correct file type was used.
	IsFile() *bool
	// Indicates if this asset is a zip archive.
	//
	// Allows constructs to ensure that the
	// correct file type was used.
	IsZipArchive() *bool
	// The tree node.
	Node() constructs.Node
	// Attribute that represents the name of the bucket this asset exists in.
	S3BucketName() *string
	// Attribute which represents the S3 object key of this asset.
	S3ObjectKey() *string
	// Attribute which represents the S3 URL of this asset.
	//
	// For example, `s3://bucket/key`.
	S3ObjectUrl() *string
	// Adds CloudFormation template metadata to the specified resource with information that indicates which resource property is mapped to this local asset.
	//
	// This can be used by tools such as SAM CLI to provide local
	// experience such as local invocation and debugging of Lambda functions.
	//
	// Asset metadata will only be included if the stack is synthesized with the
	// "aws:cdk:enable-asset-metadata" context key defined, which is the default
	// behavior when synthesizing via the CDK Toolkit.
	// See: https://github.com/aws/aws-cdk/issues/1432
	//
	AddResourceMetadata(resource awscdk.CfnResource, resourceProperty *string)
	// Grants read permissions to the principal on the assets bucket.
	GrantRead(grantee awsiam.IGrantable)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Asset
type jsiiProxy_Asset struct {
	internal.Type__constructsConstruct
	internal.Type__awscdkIAsset
}

func (j *jsiiProxy_Asset) AssetHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Asset) AssetPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Asset) Bucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Asset) HttpUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Asset) IsFile() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Asset) IsZipArchive() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isZipArchive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Asset) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Asset) S3BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Asset) S3ObjectKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Asset) S3ObjectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectUrl",
		&returns,
	)
	return returns
}


func NewAsset(scope constructs.Construct, id *string, props *AssetProps) Asset {
	_init_.Initialize()

	if err := validateNewAssetParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Asset{}

	_jsii_.Create(
		"aws-cdk-lib.aws_s3_assets.Asset",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewAsset_Override(a Asset, scope constructs.Construct, id *string, props *AssetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_s3_assets.Asset",
		[]interface{}{scope, id, props},
		a,
	)
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
func Asset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAsset_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_assets.Asset",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Asset) AddResourceMetadata(resource awscdk.CfnResource, resourceProperty *string) {
	if err := a.validateAddResourceMetadataParameters(resource, resourceProperty); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addResourceMetadata",
		[]interface{}{resource, resourceProperty},
	)
}

func (a *jsiiProxy_Asset) GrantRead(grantee awsiam.IGrantable) {
	if err := a.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"grantRead",
		[]interface{}{grantee},
	)
}

func (a *jsiiProxy_Asset) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

