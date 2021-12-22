package awsecrassets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecrassets/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An asset that represents a Docker image.
//
// The image will be created in build time and uploaded to an ECR repository.
//
// TODO: EXAMPLE
//
type DockerImageAsset interface {
	constructs.Construct
	AssetHash() *string
	ImageUri() *string
	SetImageUri(val *string)
	Node() constructs.Node
	Repository() awsecr.IRepository
	SetRepository(val awsecr.IRepository)
	AddResourceMetadata(resource awscdk.CfnResource, resourceProperty *string)
	ToString() *string
}

// The jsii proxy struct for DockerImageAsset
type jsiiProxy_DockerImageAsset struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_DockerImageAsset) AssetHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageAsset) ImageUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageAsset) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageAsset) Repository() awsecr.IRepository {
	var returns awsecr.IRepository
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}


func NewDockerImageAsset(scope constructs.Construct, id *string, props *DockerImageAssetProps) DockerImageAsset {
	_init_.Initialize()

	j := jsiiProxy_DockerImageAsset{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecr_assets.DockerImageAsset",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDockerImageAsset_Override(d DockerImageAsset, scope constructs.Construct, id *string, props *DockerImageAssetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecr_assets.DockerImageAsset",
		[]interface{}{scope, id, props},
		d,
	)
}

func (j *jsiiProxy_DockerImageAsset) SetImageUri(val *string) {
	_jsii_.Set(
		j,
		"imageUri",
		val,
	)
}

func (j *jsiiProxy_DockerImageAsset) SetRepository(val awsecr.IRepository) {
	_jsii_.Set(
		j,
		"repository",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func DockerImageAsset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr_assets.DockerImageAsset",
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
func (d *jsiiProxy_DockerImageAsset) AddResourceMetadata(resource awscdk.CfnResource, resourceProperty *string) {
	_jsii_.InvokeVoid(
		d,
		"addResourceMetadata",
		[]interface{}{resource, resourceProperty},
	)
}

// Returns a string representation of this construct.
func (d *jsiiProxy_DockerImageAsset) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options to control invalidation of `DockerImageAsset` asset hashes.
//
// TODO: EXAMPLE
//
type DockerImageAssetInvalidationOptions struct {
	// Use `buildArgs` while calculating the asset hash.
	BuildArgs *bool `json:"buildArgs"`
	// Use `extraHash` while calculating the asset hash.
	ExtraHash *bool `json:"extraHash"`
	// Use `file` while calculating the asset hash.
	File *bool `json:"file"`
	// Use `repositoryName` while calculating the asset hash.
	RepositoryName *bool `json:"repositoryName"`
	// Use `target` while calculating the asset hash.
	Target *bool `json:"target"`
}

// Options for DockerImageAsset.
//
// TODO: EXAMPLE
//
type DockerImageAssetOptions struct {
	// Glob patterns to exclude from the copy.
	Exclude *[]*string `json:"exclude"`
	// A strategy for how to handle symlinks.
	FollowSymlinks awscdk.SymlinkFollowMode `json:"followSymlinks"`
	// The ignore behavior to use for exclude patterns.
	IgnoreMode awscdk.IgnoreMode `json:"ignoreMode"`
	// Extra information to encode into the fingerprint (e.g. build instructions and other inputs).
	ExtraHash *string `json:"extraHash"`
	// Build args to pass to the `docker build` command.
	//
	// Since Docker build arguments are resolved before deployment, keys and
	// values cannot refer to unresolved tokens (such as `lambda.functionArn` or
	// `queue.queueUrl`).
	BuildArgs *map[string]*string `json:"buildArgs"`
	// Path to the Dockerfile (relative to the directory).
	File *string `json:"file"`
	// Options to control which parameters are used to invalidate the asset hash.
	Invalidation *DockerImageAssetInvalidationOptions `json:"invalidation"`
	// Docker target to build to.
	Target *string `json:"target"`
}

// Props for DockerImageAssets.
//
// TODO: EXAMPLE
//
type DockerImageAssetProps struct {
	// Glob patterns to exclude from the copy.
	Exclude *[]*string `json:"exclude"`
	// A strategy for how to handle symlinks.
	FollowSymlinks awscdk.SymlinkFollowMode `json:"followSymlinks"`
	// The ignore behavior to use for exclude patterns.
	IgnoreMode awscdk.IgnoreMode `json:"ignoreMode"`
	// Extra information to encode into the fingerprint (e.g. build instructions and other inputs).
	ExtraHash *string `json:"extraHash"`
	// Build args to pass to the `docker build` command.
	//
	// Since Docker build arguments are resolved before deployment, keys and
	// values cannot refer to unresolved tokens (such as `lambda.functionArn` or
	// `queue.queueUrl`).
	BuildArgs *map[string]*string `json:"buildArgs"`
	// Path to the Dockerfile (relative to the directory).
	File *string `json:"file"`
	// Options to control which parameters are used to invalidate the asset hash.
	Invalidation *DockerImageAssetInvalidationOptions `json:"invalidation"`
	// Docker target to build to.
	Target *string `json:"target"`
	// The directory where the Dockerfile is stored.
	//
	// Any directory inside with a name that matches the CDK output folder (cdk.out by default) will be excluded from the asset
	Directory *string `json:"directory"`
}

// An asset that represents a Docker image.
//
// The image will loaded from an existing tarball and uploaded to an ECR repository.
//
// TODO: EXAMPLE
//
type TarballImageAsset interface {
	constructs.Construct
	AssetHash() *string
	ImageUri() *string
	SetImageUri(val *string)
	Node() constructs.Node
	Repository() awsecr.IRepository
	SetRepository(val awsecr.IRepository)
	ToString() *string
}

// The jsii proxy struct for TarballImageAsset
type jsiiProxy_TarballImageAsset struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_TarballImageAsset) AssetHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TarballImageAsset) ImageUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TarballImageAsset) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TarballImageAsset) Repository() awsecr.IRepository {
	var returns awsecr.IRepository
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}


func NewTarballImageAsset(scope constructs.Construct, id *string, props *TarballImageAssetProps) TarballImageAsset {
	_init_.Initialize()

	j := jsiiProxy_TarballImageAsset{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecr_assets.TarballImageAsset",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewTarballImageAsset_Override(t TarballImageAsset, scope constructs.Construct, id *string, props *TarballImageAssetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecr_assets.TarballImageAsset",
		[]interface{}{scope, id, props},
		t,
	)
}

func (j *jsiiProxy_TarballImageAsset) SetImageUri(val *string) {
	_jsii_.Set(
		j,
		"imageUri",
		val,
	)
}

func (j *jsiiProxy_TarballImageAsset) SetRepository(val awsecr.IRepository) {
	_jsii_.Set(
		j,
		"repository",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func TarballImageAsset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr_assets.TarballImageAsset",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
func (t *jsiiProxy_TarballImageAsset) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for TarballImageAsset.
//
// TODO: EXAMPLE
//
type TarballImageAssetProps struct {
	// Absolute path to the tarball.
	//
	// It is recommended to to use the script running directory (e.g. `__dirname`
	// in Node.js projects or dirname of `__file__` in Python) if your tarball
	// is located as a resource inside your project.
	TarballFile *string `json:"tarballFile"`
}

