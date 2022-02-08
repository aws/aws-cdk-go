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
// TODO: EXAMPLE
//
// Experimental.
type DockerImageAsset interface {
	awscdk.Construct
	assets.IAsset
	AssetHash() *string
	ImageUri() *string
	SetImageUri(val *string)
	Node() awscdk.ConstructNode
	Repository() awsecr.IRepository
	SetRepository(val awsecr.IRepository)
	SourceHash() *string
	AddResourceMetadata(resource awscdk.CfnResource, resourceProperty *string)
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
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
func (d *jsiiProxy_DockerImageAsset) AddResourceMetadata(resource awscdk.CfnResource, resourceProperty *string) {
	_jsii_.InvokeVoid(
		d,
		"addResourceMetadata",
		[]interface{}{resource, resourceProperty},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (d *jsiiProxy_DockerImageAsset) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_DockerImageAsset) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (d *jsiiProxy_DockerImageAsset) Prepare() {
	_jsii_.InvokeVoid(
		d,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_DockerImageAsset) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
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
// TODO: EXAMPLE
//
// Experimental.
type DockerImageAssetInvalidationOptions struct {
	// Use `buildArgs` while calculating the asset hash.
	// Experimental.
	BuildArgs *bool `json:"buildArgs" yaml:"buildArgs"`
	// Use `extraHash` while calculating the asset hash.
	// Experimental.
	ExtraHash *bool `json:"extraHash" yaml:"extraHash"`
	// Use `file` while calculating the asset hash.
	// Experimental.
	File *bool `json:"file" yaml:"file"`
	// Use `networkMode` while calculating the asset hash.
	// Experimental.
	NetworkMode *bool `json:"networkMode" yaml:"networkMode"`
	// Use `repositoryName` while calculating the asset hash.
	// Experimental.
	RepositoryName *bool `json:"repositoryName" yaml:"repositoryName"`
	// Use `target` while calculating the asset hash.
	// Experimental.
	Target *bool `json:"target" yaml:"target"`
}

// Options for DockerImageAsset.
//
// TODO: EXAMPLE
//
// Experimental.
type DockerImageAssetOptions struct {
	// Glob patterns to exclude from the copy.
	// Experimental.
	Exclude *[]*string `json:"exclude" yaml:"exclude"`
	// A strategy for how to handle symlinks.
	// Deprecated: use `followSymlinks` instead
	Follow assets.FollowMode `json:"follow" yaml:"follow"`
	// The ignore behavior to use for exclude patterns.
	// Experimental.
	IgnoreMode awscdk.IgnoreMode `json:"ignoreMode" yaml:"ignoreMode"`
	// Extra information to encode into the fingerprint (e.g. build instructions and other inputs).
	// Experimental.
	ExtraHash *string `json:"extraHash" yaml:"extraHash"`
	// A strategy for how to handle symlinks.
	// Experimental.
	FollowSymlinks awscdk.SymlinkFollowMode `json:"followSymlinks" yaml:"followSymlinks"`
	// Build args to pass to the `docker build` command.
	//
	// Since Docker build arguments are resolved before deployment, keys and
	// values cannot refer to unresolved tokens (such as `lambda.functionArn` or
	// `queue.queueUrl`).
	// Experimental.
	BuildArgs *map[string]*string `json:"buildArgs" yaml:"buildArgs"`
	// Path to the Dockerfile (relative to the directory).
	// Experimental.
	File *string `json:"file" yaml:"file"`
	// Options to control which parameters are used to invalidate the asset hash.
	// Experimental.
	Invalidation *DockerImageAssetInvalidationOptions `json:"invalidation" yaml:"invalidation"`
	// Networking mode for the RUN commands during build.
	//
	// Support docker API 1.25+.
	// Experimental.
	NetworkMode NetworkMode `json:"networkMode" yaml:"networkMode"`
	// ECR repository name.
	//
	// Specify this property if you need to statically address the image, e.g.
	// from a Kubernetes Pod. Note, this is only the repository name, without the
	// registry and the tag parts.
	// Deprecated: to control the location of docker image assets, please override
	// `Stack.addDockerImageAsset`. this feature will be removed in future
	// releases.
	RepositoryName *string `json:"repositoryName" yaml:"repositoryName"`
	// Docker target to build to.
	// Experimental.
	Target *string `json:"target" yaml:"target"`
}

// Props for DockerImageAssets.
//
// TODO: EXAMPLE
//
// Experimental.
type DockerImageAssetProps struct {
	// Glob patterns to exclude from the copy.
	// Experimental.
	Exclude *[]*string `json:"exclude" yaml:"exclude"`
	// A strategy for how to handle symlinks.
	// Deprecated: use `followSymlinks` instead
	Follow assets.FollowMode `json:"follow" yaml:"follow"`
	// The ignore behavior to use for exclude patterns.
	// Experimental.
	IgnoreMode awscdk.IgnoreMode `json:"ignoreMode" yaml:"ignoreMode"`
	// Extra information to encode into the fingerprint (e.g. build instructions and other inputs).
	// Experimental.
	ExtraHash *string `json:"extraHash" yaml:"extraHash"`
	// A strategy for how to handle symlinks.
	// Experimental.
	FollowSymlinks awscdk.SymlinkFollowMode `json:"followSymlinks" yaml:"followSymlinks"`
	// Build args to pass to the `docker build` command.
	//
	// Since Docker build arguments are resolved before deployment, keys and
	// values cannot refer to unresolved tokens (such as `lambda.functionArn` or
	// `queue.queueUrl`).
	// Experimental.
	BuildArgs *map[string]*string `json:"buildArgs" yaml:"buildArgs"`
	// Path to the Dockerfile (relative to the directory).
	// Experimental.
	File *string `json:"file" yaml:"file"`
	// Options to control which parameters are used to invalidate the asset hash.
	// Experimental.
	Invalidation *DockerImageAssetInvalidationOptions `json:"invalidation" yaml:"invalidation"`
	// Networking mode for the RUN commands during build.
	//
	// Support docker API 1.25+.
	// Experimental.
	NetworkMode NetworkMode `json:"networkMode" yaml:"networkMode"`
	// ECR repository name.
	//
	// Specify this property if you need to statically address the image, e.g.
	// from a Kubernetes Pod. Note, this is only the repository name, without the
	// registry and the tag parts.
	// Deprecated: to control the location of docker image assets, please override
	// `Stack.addDockerImageAsset`. this feature will be removed in future
	// releases.
	RepositoryName *string `json:"repositoryName" yaml:"repositoryName"`
	// Docker target to build to.
	// Experimental.
	Target *string `json:"target" yaml:"target"`
	// The directory where the Dockerfile is stored.
	//
	// Any directory inside with a name that matches the CDK output folder (cdk.out by default) will be excluded from the asset
	// Experimental.
	Directory *string `json:"directory" yaml:"directory"`
}

// networking mode on build time supported by docker.
//
// TODO: EXAMPLE
//
// Experimental.
type NetworkMode interface {
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
// TODO: EXAMPLE
//
// Experimental.
type TarballImageAsset interface {
	awscdk.Construct
	assets.IAsset
	AssetHash() *string
	ImageUri() *string
	SetImageUri(val *string)
	Node() awscdk.ConstructNode
	Repository() awsecr.IRepository
	SetRepository(val awsecr.IRepository)
	SourceHash() *string
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (t *jsiiProxy_TarballImageAsset) OnPrepare() {
	_jsii_.InvokeVoid(
		t,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (t *jsiiProxy_TarballImageAsset) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (t *jsiiProxy_TarballImageAsset) Prepare() {
	_jsii_.InvokeVoid(
		t,
		"prepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (t *jsiiProxy_TarballImageAsset) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
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
// TODO: EXAMPLE
//
// Experimental.
type TarballImageAssetProps struct {
	// Absolute path to the tarball.
	//
	// It is recommended to to use the script running directory (e.g. `__dirname`
	// in Node.js projects or dirname of `__file__` in Python) if your tarball
	// is located as a resource inside your project.
	// Experimental.
	TarballFile *string `json:"tarballFile" yaml:"tarballFile"`
}

