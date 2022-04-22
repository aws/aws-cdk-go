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
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"type DockerImageAsset awscdk.DockerImageAsset
//   type NetworkMode awscdk.NetworkMode
//
//   asset := NewDockerImageAsset(this, jsii.String("MyBuildImage"), &dockerImageAssetProps{
//   	directory: path.join(__dirname, jsii.String("my-image")),
//   	networkMode: networkMode_HOST(),
//   })
//
type DockerImageAsset interface {
	constructs.Construct
	// A hash of this asset, which is available at construction time.
	//
	// As this is a plain string, it
	// can be used in construct IDs in order to enforce creation of a new resource when the content
	// hash has changed.
	AssetHash() *string
	// The full URI of the image (including a tag).
	//
	// Use this reference to pull
	// the asset.
	ImageUri() *string
	SetImageUri(val *string)
	// The tree node.
	Node() constructs.Node
	// Repository where the image is stored.
	Repository() awsecr.IRepository
	SetRepository(val awsecr.IRepository)
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
	// Returns a string representation of this construct.
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
// Deprecated: use `x instanceof Construct` instead.
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

func (d *jsiiProxy_DockerImageAsset) AddResourceMetadata(resource awscdk.CfnResource, resourceProperty *string) {
	_jsii_.InvokeVoid(
		d,
		"addResourceMetadata",
		[]interface{}{resource, resourceProperty},
	)
}

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
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"type DockerImageAsset awscdk.DockerImageAsset
//
//   asset := NewDockerImageAsset(this, jsii.String("MyBuildImage"), &dockerImageAssetProps{
//   	directory: path.join(__dirname, jsii.String("my-image")),
//   	buildArgs: map[string]*string{
//   		"HTTP_PROXY": jsii.String("http://10.20.30.2:1234"),
//   	},
//   	invalidation: &dockerImageAssetInvalidationOptions{
//   		buildArgs: jsii.Boolean(false),
//   	},
//   })
//
type DockerImageAssetInvalidationOptions struct {
	// Use `buildArgs` while calculating the asset hash.
	BuildArgs *bool `json:"buildArgs" yaml:"buildArgs"`
	// Use `extraHash` while calculating the asset hash.
	ExtraHash *bool `json:"extraHash" yaml:"extraHash"`
	// Use `file` while calculating the asset hash.
	File *bool `json:"file" yaml:"file"`
	// Use `networkMode` while calculating the asset hash.
	NetworkMode *bool `json:"networkMode" yaml:"networkMode"`
	// Use `repositoryName` while calculating the asset hash.
	RepositoryName *bool `json:"repositoryName" yaml:"repositoryName"`
	// Use `target` while calculating the asset hash.
	Target *bool `json:"target" yaml:"target"`
}

// Options for DockerImageAsset.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"import awscdk "github.com/aws/aws-cdk-go/awscdk"import ecr_assets "github.com/aws/aws-cdk-go/awscdk/aws_ecr_assets"
//
//   var networkMode networkMode
//   dockerImageAssetOptions := &dockerImageAssetOptions{
//   	buildArgs: map[string]*string{
//   		"buildArgsKey": jsii.String("buildArgs"),
//   	},
//   	exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	extraHash: jsii.String("extraHash"),
//   	file: jsii.String("file"),
//   	followSymlinks: cdk.symlinkFollowMode_NEVER,
//   	ignoreMode: cdk.ignoreMode_GLOB,
//   	invalidation: &dockerImageAssetInvalidationOptions{
//   		buildArgs: jsii.Boolean(false),
//   		extraHash: jsii.Boolean(false),
//   		file: jsii.Boolean(false),
//   		networkMode: jsii.Boolean(false),
//   		repositoryName: jsii.Boolean(false),
//   		target: jsii.Boolean(false),
//   	},
//   	networkMode: networkMode,
//   	target: jsii.String("target"),
//   }
//
type DockerImageAssetOptions struct {
	// Glob patterns to exclude from the copy.
	Exclude *[]*string `json:"exclude" yaml:"exclude"`
	// A strategy for how to handle symlinks.
	FollowSymlinks awscdk.SymlinkFollowMode `json:"followSymlinks" yaml:"followSymlinks"`
	// The ignore behavior to use for exclude patterns.
	IgnoreMode awscdk.IgnoreMode `json:"ignoreMode" yaml:"ignoreMode"`
	// Extra information to encode into the fingerprint (e.g. build instructions and other inputs).
	ExtraHash *string `json:"extraHash" yaml:"extraHash"`
	// Build args to pass to the `docker build` command.
	//
	// Since Docker build arguments are resolved before deployment, keys and
	// values cannot refer to unresolved tokens (such as `lambda.functionArn` or
	// `queue.queueUrl`).
	BuildArgs *map[string]*string `json:"buildArgs" yaml:"buildArgs"`
	// Path to the Dockerfile (relative to the directory).
	File *string `json:"file" yaml:"file"`
	// Options to control which parameters are used to invalidate the asset hash.
	Invalidation *DockerImageAssetInvalidationOptions `json:"invalidation" yaml:"invalidation"`
	// Networking mode for the RUN commands during build.
	//
	// Support docker API 1.25+.
	NetworkMode NetworkMode `json:"networkMode" yaml:"networkMode"`
	// Docker target to build to.
	Target *string `json:"target" yaml:"target"`
}

// Props for DockerImageAssets.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"type DockerImageAsset awscdk.DockerImageAsset
//
//   asset := NewDockerImageAsset(this, jsii.String("MyBuildImage"), &dockerImageAssetProps{
//   	directory: path.join(__dirname, jsii.String("my-image")),
//   	buildArgs: map[string]*string{
//   		"HTTP_PROXY": jsii.String("http://10.20.30.2:1234"),
//   	},
//   	invalidation: &dockerImageAssetInvalidationOptions{
//   		buildArgs: jsii.Boolean(false),
//   	},
//   })
//
type DockerImageAssetProps struct {
	// Glob patterns to exclude from the copy.
	Exclude *[]*string `json:"exclude" yaml:"exclude"`
	// A strategy for how to handle symlinks.
	FollowSymlinks awscdk.SymlinkFollowMode `json:"followSymlinks" yaml:"followSymlinks"`
	// The ignore behavior to use for exclude patterns.
	IgnoreMode awscdk.IgnoreMode `json:"ignoreMode" yaml:"ignoreMode"`
	// Extra information to encode into the fingerprint (e.g. build instructions and other inputs).
	ExtraHash *string `json:"extraHash" yaml:"extraHash"`
	// Build args to pass to the `docker build` command.
	//
	// Since Docker build arguments are resolved before deployment, keys and
	// values cannot refer to unresolved tokens (such as `lambda.functionArn` or
	// `queue.queueUrl`).
	BuildArgs *map[string]*string `json:"buildArgs" yaml:"buildArgs"`
	// Path to the Dockerfile (relative to the directory).
	File *string `json:"file" yaml:"file"`
	// Options to control which parameters are used to invalidate the asset hash.
	Invalidation *DockerImageAssetInvalidationOptions `json:"invalidation" yaml:"invalidation"`
	// Networking mode for the RUN commands during build.
	//
	// Support docker API 1.25+.
	NetworkMode NetworkMode `json:"networkMode" yaml:"networkMode"`
	// Docker target to build to.
	Target *string `json:"target" yaml:"target"`
	// The directory where the Dockerfile is stored.
	//
	// Any directory inside with a name that matches the CDK output folder (cdk.out by default) will be excluded from the asset
	Directory *string `json:"directory" yaml:"directory"`
}

// networking mode on build time supported by docker.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"type DockerImageAsset awscdk.DockerImageAsset
//   type NetworkMode awscdk.NetworkMode
//
//   asset := NewDockerImageAsset(this, jsii.String("MyBuildImage"), &dockerImageAssetProps{
//   	directory: path.join(__dirname, jsii.String("my-image")),
//   	networkMode: networkMode_HOST(),
//   })
//
type NetworkMode interface {
	// The networking mode to use for docker build.
	Mode() *string
}

// The jsii proxy struct for NetworkMode
type jsiiProxy_NetworkMode struct {
	_ byte // padding
}

func (j *jsiiProxy_NetworkMode) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}


// Used to specify a custom networking mode Use this if the networking mode name is not yet supported by the CDK.
func NetworkMode_Custom(mode *string) NetworkMode {
	_init_.Initialize()

	var returns NetworkMode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr_assets.NetworkMode",
		"custom",
		[]interface{}{mode},
		&returns,
	)

	return returns
}

// Reuse another container's network stack.
func NetworkMode_FromContainer(containerId *string) NetworkMode {
	_init_.Initialize()

	var returns NetworkMode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecr_assets.NetworkMode",
		"fromContainer",
		[]interface{}{containerId},
		&returns,
	)

	return returns
}

func NetworkMode_DEFAULT() NetworkMode {
	_init_.Initialize()
	var returns NetworkMode
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecr_assets.NetworkMode",
		"DEFAULT",
		&returns,
	)
	return returns
}

func NetworkMode_HOST() NetworkMode {
	_init_.Initialize()
	var returns NetworkMode
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecr_assets.NetworkMode",
		"HOST",
		&returns,
	)
	return returns
}

func NetworkMode_NONE() NetworkMode {
	_init_.Initialize()
	var returns NetworkMode
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecr_assets.NetworkMode",
		"NONE",
		&returns,
	)
	return returns
}

// An asset that represents a Docker image.
//
// The image will loaded from an existing tarball and uploaded to an ECR repository.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"type TarballImageAsset awscdk.TarballImageAsset
//
//   asset := NewTarballImageAsset(this, jsii.String("MyBuildImage"), &tarballImageAssetProps{
//   	tarballFile: jsii.String("local-image.tar"),
//   })
//
type TarballImageAsset interface {
	constructs.Construct
	// A hash of this asset, which is available at construction time.
	//
	// As this is a plain string, it
	// can be used in construct IDs in order to enforce creation of a new resource when the content
	// hash has changed.
	AssetHash() *string
	// The full URI of the image (including a tag).
	//
	// Use this reference to pull
	// the asset.
	ImageUri() *string
	SetImageUri(val *string)
	// The tree node.
	Node() constructs.Node
	// Repository where the image is stored.
	Repository() awsecr.IRepository
	SetRepository(val awsecr.IRepository)
	// Returns a string representation of this construct.
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
// Deprecated: use `x instanceof Construct` instead.
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
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"type TarballImageAsset awscdk.TarballImageAsset
//
//   asset := NewTarballImageAsset(this, jsii.String("MyBuildImage"), &tarballImageAssetProps{
//   	tarballFile: jsii.String("local-image.tar"),
//   })
//
type TarballImageAssetProps struct {
	// Absolute path to the tarball.
	//
	// It is recommended to to use the script running directory (e.g. `__dirname`
	// in Node.js projects or dirname of `__file__` in Python) if your tarball
	// is located as a resource inside your project.
	TarballFile *string `json:"tarballFile" yaml:"tarballFile"`
}

