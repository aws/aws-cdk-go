package awss3assets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An asset represents a local file or directory, which is automatically uploaded to S3 and then can be referenced within a CDK application.
//
// TODO: EXAMPLE
//
// Experimental.
type Asset interface {
	constructs.Construct
	awscdk.IAsset
	AssetHash() *string
	AssetPath() *string
	Bucket() awss3.IBucket
	HttpUrl() *string
	IsFile() *bool
	IsZipArchive() *bool
	Node() constructs.Node
	S3BucketName() *string
	S3ObjectKey() *string
	S3ObjectUrl() *string
	S3Url() *string
	SourceHash() *string
	AddResourceMetadata(resource awscdk.CfnResource, resourceProperty *string)
	GrantRead(grantee awsiam.IGrantable)
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

func (j *jsiiProxy_Asset) S3Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Asset) SourceHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceHash",
		&returns,
	)
	return returns
}


// Experimental.
func NewAsset(scope constructs.Construct, id *string, props *AssetProps) Asset {
	_init_.Initialize()

	j := jsiiProxy_Asset{}

	_jsii_.Create(
		"aws-cdk-lib.aws_s3_assets.Asset",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
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
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Asset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_s3_assets.Asset",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

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
// Experimental.
func (a *jsiiProxy_Asset) AddResourceMetadata(resource awscdk.CfnResource, resourceProperty *string) {
	_jsii_.InvokeVoid(
		a,
		"addResourceMetadata",
		[]interface{}{resource, resourceProperty},
	)
}

// Grants read permissions to the principal on the assets bucket.
// Experimental.
func (a *jsiiProxy_Asset) GrantRead(grantee awsiam.IGrantable) {
	_jsii_.InvokeVoid(
		a,
		"grantRead",
		[]interface{}{grantee},
	)
}

// Returns a string representation of this construct.
// Experimental.
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

// TODO: EXAMPLE
//
// Experimental.
type AssetOptions struct {
	// Glob patterns to exclude from the copy.
	// Experimental.
	Exclude *[]*string `json:"exclude"`
	// A strategy for how to handle symlinks.
	// Deprecated: use `followSymlinks` instead
	Follow assets.FollowMode `json:"follow"`
	// The ignore behavior to use for exclude patterns.
	// Experimental.
	IgnoreMode awscdk.IgnoreMode `json:"ignoreMode"`
	// A strategy for how to handle symlinks.
	// Experimental.
	FollowSymlinks awscdk.SymlinkFollowMode `json:"followSymlinks"`
	// Specify a custom hash for this asset.
	//
	// If `assetHashType` is set it must
	// be set to `AssetHashType.CUSTOM`. For consistency, this custom hash will
	// be SHA256 hashed and encoded as hex. The resulting hash will be the asset
	// hash.
	//
	// NOTE: the hash is used in order to identify a specific revision of the asset, and
	// used for optimizing and caching deployment activities related to this asset such as
	// packaging, uploading to Amazon S3, etc. If you chose to customize the hash, you will
	// need to make sure it is updated every time the asset changes, or otherwise it is
	// possible that some deployments will not be invalidated.
	// Experimental.
	AssetHash *string `json:"assetHash"`
	// Specifies the type of hash to calculate for this asset.
	//
	// If `assetHash` is configured, this option must be `undefined` or
	// `AssetHashType.CUSTOM`.
	// Experimental.
	AssetHashType awscdk.AssetHashType `json:"assetHashType"`
	// Bundle the asset by executing a command in a Docker container or a custom bundling provider.
	//
	// The asset path will be mounted at `/asset-input`. The Docker
	// container is responsible for putting content at `/asset-output`.
	// The content at `/asset-output` will be zipped and used as the
	// final asset.
	// Experimental.
	Bundling *awscdk.BundlingOptions `json:"bundling"`
	// A list of principals that should be able to read this asset from S3.
	//
	// You can use `asset.grantRead(principal)` to grant read permissions later.
	// Experimental.
	Readers *[]awsiam.IGrantable `json:"readers"`
	// Custom hash to use when identifying the specific version of the asset.
	//
	// For consistency,
	// this custom hash will be SHA256 hashed and encoded as hex. The resulting hash will be
	// the asset hash.
	//
	// NOTE: the source hash is used in order to identify a specific revision of the asset,
	// and used for optimizing and caching deployment activities related to this asset such as
	// packaging, uploading to Amazon S3, etc. If you chose to customize the source hash,
	// you will need to make sure it is updated every time the source changes, or otherwise
	// it is possible that some deployments will not be invalidated.
	// Deprecated: see `assetHash` and `assetHashType`
	SourceHash *string `json:"sourceHash"`
}

// TODO: EXAMPLE
//
// Experimental.
type AssetProps struct {
	// Glob patterns to exclude from the copy.
	// Experimental.
	Exclude *[]*string `json:"exclude"`
	// A strategy for how to handle symlinks.
	// Deprecated: use `followSymlinks` instead
	Follow assets.FollowMode `json:"follow"`
	// The ignore behavior to use for exclude patterns.
	// Experimental.
	IgnoreMode awscdk.IgnoreMode `json:"ignoreMode"`
	// A strategy for how to handle symlinks.
	// Experimental.
	FollowSymlinks awscdk.SymlinkFollowMode `json:"followSymlinks"`
	// Specify a custom hash for this asset.
	//
	// If `assetHashType` is set it must
	// be set to `AssetHashType.CUSTOM`. For consistency, this custom hash will
	// be SHA256 hashed and encoded as hex. The resulting hash will be the asset
	// hash.
	//
	// NOTE: the hash is used in order to identify a specific revision of the asset, and
	// used for optimizing and caching deployment activities related to this asset such as
	// packaging, uploading to Amazon S3, etc. If you chose to customize the hash, you will
	// need to make sure it is updated every time the asset changes, or otherwise it is
	// possible that some deployments will not be invalidated.
	// Experimental.
	AssetHash *string `json:"assetHash"`
	// Specifies the type of hash to calculate for this asset.
	//
	// If `assetHash` is configured, this option must be `undefined` or
	// `AssetHashType.CUSTOM`.
	// Experimental.
	AssetHashType awscdk.AssetHashType `json:"assetHashType"`
	// Bundle the asset by executing a command in a Docker container or a custom bundling provider.
	//
	// The asset path will be mounted at `/asset-input`. The Docker
	// container is responsible for putting content at `/asset-output`.
	// The content at `/asset-output` will be zipped and used as the
	// final asset.
	// Experimental.
	Bundling *awscdk.BundlingOptions `json:"bundling"`
	// A list of principals that should be able to read this asset from S3.
	//
	// You can use `asset.grantRead(principal)` to grant read permissions later.
	// Experimental.
	Readers *[]awsiam.IGrantable `json:"readers"`
	// Custom hash to use when identifying the specific version of the asset.
	//
	// For consistency,
	// this custom hash will be SHA256 hashed and encoded as hex. The resulting hash will be
	// the asset hash.
	//
	// NOTE: the source hash is used in order to identify a specific revision of the asset,
	// and used for optimizing and caching deployment activities related to this asset such as
	// packaging, uploading to Amazon S3, etc. If you chose to customize the source hash,
	// you will need to make sure it is updated every time the source changes, or otherwise
	// it is possible that some deployments will not be invalidated.
	// Deprecated: see `assetHash` and `assetHashType`
	SourceHash *string `json:"sourceHash"`
	// The disk location of the asset.
	//
	// The path should refer to one of the following:
	// - A regular file or a .zip file, in which case the file will be uploaded as-is to S3.
	// - A directory, in which case it will be archived into a .zip file and uploaded to S3.
	// Experimental.
	Path *string `json:"path"`
}

