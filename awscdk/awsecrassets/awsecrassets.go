package awsecrassets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/assets"
	"github.com/aws/aws-cdk-go/awscdk/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/awsecrassets/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// An asset that represents a Docker image.
//
// The image will be created in build time and uploaded to an ECR repository.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   asset := awscdk.NewDockerImageAsset(this, jsii.String("MyBuildImage"), &dockerImageAssetProps{
//   	directory: path.join(__dirname, jsii.String("my-image")),
//   	networkMode: awscdk.NetworkMode_HOST(),
//   })
//
// Experimental.
type DockerImageAsset interface {
	awscdk.Construct
	assets.IAsset
	// A hash of this asset, which is available at construction time.
	//
	// As this is a plain string, it
	// can be used in construct IDs in order to enforce creation of a new resource when the content
	// hash has changed.
	// Experimental.
	AssetHash() *string
	// The full URI of the image (including a tag).
	//
	// Use this reference to pull
	// the asset.
	// Experimental.
	ImageUri() *string
	// Experimental.
	SetImageUri(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Repository where the image is stored.
	// Experimental.
	Repository() awsecr.IRepository
	// Experimental.
	SetRepository(val awsecr.IRepository)
	// A hash of the source of this asset, which is available at construction time.
	//
	// As this is a plain
	// string, it can be used in construct IDs in order to enforce creation of a new resource when
	// the content hash has changed.
	// Deprecated: use assetHash.
	SourceHash() *string
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
	AddResourceMetadata(resource awscdk.CfnResource, resourceProperty *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for DockerImageAsset
type jsiiProxy_DockerImageAsset struct {
	internal.Type__awscdkConstruct
	internal.Type__assetsIAsset
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

func (j *jsiiProxy_DockerImageAsset) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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

func (j *jsiiProxy_DockerImageAsset) SourceHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceHash",
		&returns,
	)
	return returns
}


// Experimental.
func NewDockerImageAsset(scope constructs.Construct, id *string, props *DockerImageAssetProps) DockerImageAsset {
	_init_.Initialize()

	j := jsiiProxy_DockerImageAsset{}

	_jsii_.Create(
		"monocdk.aws_ecr_assets.DockerImageAsset",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDockerImageAsset_Override(d DockerImageAsset, scope constructs.Construct, id *string, props *DockerImageAssetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecr_assets.DockerImageAsset",
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

// Return whether the given object is a Construct.
// Experimental.
func DockerImageAsset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecr_assets.DockerImageAsset",
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

func (d *jsiiProxy_DockerImageAsset) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DockerImageAsset) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (d *jsiiProxy_DockerImageAsset) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DockerImageAsset) Prepare() {
	_jsii_.InvokeVoid(
		d,
		"prepare",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DockerImageAsset) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"synthesize",
		[]interface{}{session},
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

func (d *jsiiProxy_DockerImageAsset) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options to control invalidation of `DockerImageAsset` asset hashes.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   asset := awscdk.NewDockerImageAsset(this, jsii.String("MyBuildImage"), &dockerImageAssetProps{
//   	directory: path.join(__dirname, jsii.String("my-image")),
//   	buildArgs: map[string]*string{
//   		"HTTP_PROXY": jsii.String("http://10.20.30.2:1234"),
//   	},
//   	invalidation: &dockerImageAssetInvalidationOptions{
//   		buildArgs: jsii.Boolean(false),
//   	},
//   })
//
// Experimental.
type DockerImageAssetInvalidationOptions struct {
	// Use `buildArgs` while calculating the asset hash.
	// Experimental.
	BuildArgs *bool `field:"optional" json:"buildArgs" yaml:"buildArgs"`
	// Use `extraHash` while calculating the asset hash.
	// Experimental.
	ExtraHash *bool `field:"optional" json:"extraHash" yaml:"extraHash"`
	// Use `file` while calculating the asset hash.
	// Experimental.
	File *bool `field:"optional" json:"file" yaml:"file"`
	// Use `networkMode` while calculating the asset hash.
	// Experimental.
	NetworkMode *bool `field:"optional" json:"networkMode" yaml:"networkMode"`
	// Use `repositoryName` while calculating the asset hash.
	// Experimental.
	RepositoryName *bool `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// Use `target` while calculating the asset hash.
	// Experimental.
	Target *bool `field:"optional" json:"target" yaml:"target"`
}

// Options for DockerImageAsset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var networkMode networkMode
//
//   dockerImageAssetOptions := &dockerImageAssetOptions{
//   	buildArgs: map[string]*string{
//   		"buildArgsKey": jsii.String("buildArgs"),
//   	},
//   	exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	extraHash: jsii.String("extraHash"),
//   	file: jsii.String("file"),
//   	follow: awscdk.Assets.followMode_NEVER,
//   	followSymlinks: monocdk.symlinkFollowMode_NEVER,
//   	ignoreMode: monocdk.ignoreMode_GLOB,
//   	invalidation: &dockerImageAssetInvalidationOptions{
//   		buildArgs: jsii.Boolean(false),
//   		extraHash: jsii.Boolean(false),
//   		file: jsii.Boolean(false),
//   		networkMode: jsii.Boolean(false),
//   		repositoryName: jsii.Boolean(false),
//   		target: jsii.Boolean(false),
//   	},
//   	networkMode: networkMode,
//   	repositoryName: jsii.String("repositoryName"),
//   	target: jsii.String("target"),
//   }
//
// Experimental.
type DockerImageAssetOptions struct {
	// Glob patterns to exclude from the copy.
	// Experimental.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// A strategy for how to handle symlinks.
	// Deprecated: use `followSymlinks` instead.
	Follow assets.FollowMode `field:"optional" json:"follow" yaml:"follow"`
	// The ignore behavior to use for exclude patterns.
	// Experimental.
	IgnoreMode awscdk.IgnoreMode `field:"optional" json:"ignoreMode" yaml:"ignoreMode"`
	// Extra information to encode into the fingerprint (e.g. build instructions and other inputs).
	// Experimental.
	ExtraHash *string `field:"optional" json:"extraHash" yaml:"extraHash"`
	// A strategy for how to handle symlinks.
	// Experimental.
	FollowSymlinks awscdk.SymlinkFollowMode `field:"optional" json:"followSymlinks" yaml:"followSymlinks"`
	// Build args to pass to the `docker build` command.
	//
	// Since Docker build arguments are resolved before deployment, keys and
	// values cannot refer to unresolved tokens (such as `lambda.functionArn` or
	// `queue.queueUrl`).
	// Experimental.
	BuildArgs *map[string]*string `field:"optional" json:"buildArgs" yaml:"buildArgs"`
	// Path to the Dockerfile (relative to the directory).
	// Experimental.
	File *string `field:"optional" json:"file" yaml:"file"`
	// Options to control which parameters are used to invalidate the asset hash.
	// Experimental.
	Invalidation *DockerImageAssetInvalidationOptions `field:"optional" json:"invalidation" yaml:"invalidation"`
	// Networking mode for the RUN commands during build.
	//
	// Support docker API 1.25+.
	// Experimental.
	NetworkMode NetworkMode `field:"optional" json:"networkMode" yaml:"networkMode"`
	// ECR repository name.
	//
	// Specify this property if you need to statically address the image, e.g.
	// from a Kubernetes Pod. Note, this is only the repository name, without the
	// registry and the tag parts.
	// Deprecated: to control the location of docker image assets, please override
	// `Stack.addDockerImageAsset`. this feature will be removed in future
	// releases.
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// Docker target to build to.
	// Experimental.
	Target *string `field:"optional" json:"target" yaml:"target"`
}

// Props for DockerImageAssets.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   asset := awscdk.NewDockerImageAsset(this, jsii.String("MyBuildImage"), &dockerImageAssetProps{
//   	directory: path.join(__dirname, jsii.String("my-image")),
//   	buildArgs: map[string]*string{
//   		"HTTP_PROXY": jsii.String("http://10.20.30.2:1234"),
//   	},
//   	invalidation: &dockerImageAssetInvalidationOptions{
//   		buildArgs: jsii.Boolean(false),
//   	},
//   })
//
// Experimental.
type DockerImageAssetProps struct {
	// Glob patterns to exclude from the copy.
	// Experimental.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// A strategy for how to handle symlinks.
	// Deprecated: use `followSymlinks` instead.
	Follow assets.FollowMode `field:"optional" json:"follow" yaml:"follow"`
	// The ignore behavior to use for exclude patterns.
	// Experimental.
	IgnoreMode awscdk.IgnoreMode `field:"optional" json:"ignoreMode" yaml:"ignoreMode"`
	// Extra information to encode into the fingerprint (e.g. build instructions and other inputs).
	// Experimental.
	ExtraHash *string `field:"optional" json:"extraHash" yaml:"extraHash"`
	// A strategy for how to handle symlinks.
	// Experimental.
	FollowSymlinks awscdk.SymlinkFollowMode `field:"optional" json:"followSymlinks" yaml:"followSymlinks"`
	// Build args to pass to the `docker build` command.
	//
	// Since Docker build arguments are resolved before deployment, keys and
	// values cannot refer to unresolved tokens (such as `lambda.functionArn` or
	// `queue.queueUrl`).
	// Experimental.
	BuildArgs *map[string]*string `field:"optional" json:"buildArgs" yaml:"buildArgs"`
	// Path to the Dockerfile (relative to the directory).
	// Experimental.
	File *string `field:"optional" json:"file" yaml:"file"`
	// Options to control which parameters are used to invalidate the asset hash.
	// Experimental.
	Invalidation *DockerImageAssetInvalidationOptions `field:"optional" json:"invalidation" yaml:"invalidation"`
	// Networking mode for the RUN commands during build.
	//
	// Support docker API 1.25+.
	// Experimental.
	NetworkMode NetworkMode `field:"optional" json:"networkMode" yaml:"networkMode"`
	// ECR repository name.
	//
	// Specify this property if you need to statically address the image, e.g.
	// from a Kubernetes Pod. Note, this is only the repository name, without the
	// registry and the tag parts.
	// Deprecated: to control the location of docker image assets, please override
	// `Stack.addDockerImageAsset`. this feature will be removed in future
	// releases.
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// Docker target to build to.
	// Experimental.
	Target *string `field:"optional" json:"target" yaml:"target"`
	// The directory where the Dockerfile is stored.
	//
	// Any directory inside with a name that matches the CDK output folder (cdk.out by default) will be excluded from the asset
	// Experimental.
	Directory *string `field:"required" json:"directory" yaml:"directory"`
}

// networking mode on build time supported by docker.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   asset := awscdk.NewDockerImageAsset(this, jsii.String("MyBuildImage"), &dockerImageAssetProps{
//   	directory: path.join(__dirname, jsii.String("my-image")),
//   	networkMode: awscdk.NetworkMode_HOST(),
//   })
//
// Experimental.
type NetworkMode interface {
	// The networking mode to use for docker build.
	// Experimental.
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
// Experimental.
func NetworkMode_Custom(mode *string) NetworkMode {
	_init_.Initialize()

	var returns NetworkMode

	_jsii_.StaticInvoke(
		"monocdk.aws_ecr_assets.NetworkMode",
		"custom",
		[]interface{}{mode},
		&returns,
	)

	return returns
}

// Reuse another container's network stack.
// Experimental.
func NetworkMode_FromContainer(containerId *string) NetworkMode {
	_init_.Initialize()

	var returns NetworkMode

	_jsii_.StaticInvoke(
		"monocdk.aws_ecr_assets.NetworkMode",
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
		"monocdk.aws_ecr_assets.NetworkMode",
		"DEFAULT",
		&returns,
	)
	return returns
}

func NetworkMode_HOST() NetworkMode {
	_init_.Initialize()
	var returns NetworkMode
	_jsii_.StaticGet(
		"monocdk.aws_ecr_assets.NetworkMode",
		"HOST",
		&returns,
	)
	return returns
}

func NetworkMode_NONE() NetworkMode {
	_init_.Initialize()
	var returns NetworkMode
	_jsii_.StaticGet(
		"monocdk.aws_ecr_assets.NetworkMode",
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
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   asset := awscdk.NewTarballImageAsset(this, jsii.String("MyBuildImage"), &tarballImageAssetProps{
//   	tarballFile: jsii.String("local-image.tar"),
//   })
//
// Experimental.
type TarballImageAsset interface {
	awscdk.Construct
	assets.IAsset
	// A hash of this asset, which is available at construction time.
	//
	// As this is a plain string, it
	// can be used in construct IDs in order to enforce creation of a new resource when the content
	// hash has changed.
	// Experimental.
	AssetHash() *string
	// The full URI of the image (including a tag).
	//
	// Use this reference to pull
	// the asset.
	// Experimental.
	ImageUri() *string
	// Experimental.
	SetImageUri(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Repository where the image is stored.
	// Experimental.
	Repository() awsecr.IRepository
	// Experimental.
	SetRepository(val awsecr.IRepository)
	// A hash of the source of this asset, which is available at construction time.
	//
	// As this is a plain
	// string, it can be used in construct IDs in order to enforce creation of a new resource when
	// the content hash has changed.
	// Deprecated: use assetHash.
	SourceHash() *string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for TarballImageAsset
type jsiiProxy_TarballImageAsset struct {
	internal.Type__awscdkConstruct
	internal.Type__assetsIAsset
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

func (j *jsiiProxy_TarballImageAsset) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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

func (j *jsiiProxy_TarballImageAsset) SourceHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceHash",
		&returns,
	)
	return returns
}


// Experimental.
func NewTarballImageAsset(scope constructs.Construct, id *string, props *TarballImageAssetProps) TarballImageAsset {
	_init_.Initialize()

	j := jsiiProxy_TarballImageAsset{}

	_jsii_.Create(
		"monocdk.aws_ecr_assets.TarballImageAsset",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewTarballImageAsset_Override(t TarballImageAsset, scope constructs.Construct, id *string, props *TarballImageAssetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ecr_assets.TarballImageAsset",
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

// Return whether the given object is a Construct.
// Experimental.
func TarballImageAsset_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_ecr_assets.TarballImageAsset",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TarballImageAsset) OnPrepare() {
	_jsii_.InvokeVoid(
		t,
		"onPrepare",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TarballImageAsset) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (t *jsiiProxy_TarballImageAsset) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TarballImageAsset) Prepare() {
	_jsii_.InvokeVoid(
		t,
		"prepare",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TarballImageAsset) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
		"synthesize",
		[]interface{}{session},
	)
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

func (t *jsiiProxy_TarballImageAsset) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Options for TarballImageAsset.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   asset := awscdk.NewTarballImageAsset(this, jsii.String("MyBuildImage"), &tarballImageAssetProps{
//   	tarballFile: jsii.String("local-image.tar"),
//   })
//
// Experimental.
type TarballImageAssetProps struct {
	// Absolute path to the tarball.
	//
	// It is recommended to to use the script running directory (e.g. `__dirname`
	// in Node.js projects or dirname of `__file__` in Python) if your tarball
	// is located as a resource inside your project.
	// Experimental.
	TarballFile *string `field:"required" json:"tarballFile" yaml:"tarballFile"`
}

