package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Stages a file or directory from a location on the file system into a staging directory.
//
// This construct can operate in two modes: without and with "bundling". Bundling
// means that an additional command is run to transform the source files into
// a final asset that will be used for deployment.
//
// ## Without bundling
//
// Without the `bundling` property, the asset is just a regular file copy
// operation.
//
// - `exclude` and `ignoreMode` control which files are excluded from the copy operation.
// - `follow` controls how symlinks are handled during the copy operation:
//   the symlink can be copied, the file the symlink points to can be copied,
//   or symlinks pointing outside the source tree can be blocked.
// - The asset hash is calculated based on the files that are copied into the asset,
//   unless the hash is overridden entirely by the `assetHash` property.
//
// The copy-staging operation can be disabled by setting the context
// key `'aws:cdk:asset-staging': false`. This can be useful for local
// development tools like SAM CLI, so that you can update the source files
// directly and immediately see the result in a local emulator without needing
// to run `cdk synth` every time to stage the assets. Note that in this mode,
// since there is no copying, `exclude` and `follow` settings have no effect,
// including `BLOCK_EXTERNAL` which will not be applied.
//
// ## With bundling
//
// With the `bundling` property, the transformation of source files to a final asset
// is done through bundling operation specified in the `bundling` property.
//
// The behavior of the construct depends on the `assetHashType` selected:
//
// - `AssetHashType.SOURCE`: the asset hash will be based on the source files
//   after applying the `exclude`, `ignoreMode` and `follow` settings. That hash
//   will be used to skip the bundling command entirely if an asset with the given
//   hash already exists in the assembly directory. This requires that those
//   settings accurately model the dependencies used by the bundling operation to
//   produce the final asset.
// - `AssetHashType.OUTPUT`: the asset hash will be based on the output of the
//   bundling operation. `exclude`/`ignoreMode` are not used, and no attempt
//   is made to materialize any symlinks that bundling may have written.
// - `AssetHashType.CUSTOM`: the asset hash will be taken from the `assetHash`
//   property. `exclude`/`ignoreMode` are not used, and no attempt is made to
//   materialize any symlinks that bundling may have written.
//
// For all three modes, if `follow` is set to `SymlinkFollowMode.BLOCK_EXTERNAL`,
// any symlinks written by bundling that point outside the output directory
// will be blocked.
//
// For bundled assets, staging can not be disabled for the benefit of local
// development. If you are iterating against a local emulator, you will need to
// run `cdk synth` for every asset change.
//
// ## Other behavior
//
// The `stagedPath` property is a stringified token that represents the location
// of the file or directory after staging. It will be resolved only during the
// "prepare" stage and may be either the original path or the staged path
// depending on the context setting.
//
// The file/directory are staged based on their content hash (fingerprint). This
// means that only if content was changed, copy will happen.
//
// ## Security notes on BLOCK_EXTERNAL symlinks
//
// If a bundled asset is being reused from the `cdk.out` cache, the BLOCK_EXTERNAL
// symlink check is not re-run, and on an asset hash collission (made more likely
// in the case of a chosen `CUSTOM` hash) external symlinks left over from a
// previous unchecked synth can persist. This is unlikely to be a problem in
// practice, but to prevent this from happening to you, delete the `cdk.out`
// directory whenever you switch from non-`BLOCK_EXTERNAL` to `BLOCK_EXTERNAL` follow mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dockerImage DockerImage
//   var localBundling ILocalBundling
//
//   assetStaging := cdk.NewAssetStaging(this, jsii.String("MyAssetStaging"), &AssetStagingProps{
//   	SourcePath: jsii.String("sourcePath"),
//
//   	// the properties below are optional
//   	AssetHash: jsii.String("assetHash"),
//   	AssetHashType: cdk.AssetHashType_SOURCE,
//   	Bundling: &BundlingOptions{
//   		Image: dockerImage,
//
//   		// the properties below are optional
//   		BundlingFileAccess: cdk.BundlingFileAccess_VOLUME_COPY,
//   		Command: []*string{
//   			jsii.String("command"),
//   		},
//   		Entrypoint: []*string{
//   			jsii.String("entrypoint"),
//   		},
//   		Environment: map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   		Local: localBundling,
//   		Network: jsii.String("network"),
//   		OutputType: cdk.BundlingOutput_ARCHIVED,
//   		Platform: jsii.String("platform"),
//   		SecurityOpt: jsii.String("securityOpt"),
//   		User: jsii.String("user"),
//   		Volumes: []DockerVolume{
//   			&DockerVolume{
//   				ContainerPath: jsii.String("containerPath"),
//   				HostPath: jsii.String("hostPath"),
//
//   				// the properties below are optional
//   				Consistency: cdk.DockerVolumeConsistency_CONSISTENT,
//   			},
//   		},
//   		VolumesFrom: []*string{
//   			jsii.String("volumesFrom"),
//   		},
//   		WorkingDirectory: jsii.String("workingDirectory"),
//   	},
//   	Exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	ExtraHash: jsii.String("extraHash"),
//   	Follow: cdk.SymlinkFollowMode_NEVER,
//   	IgnoreMode: cdk.IgnoreMode_GLOB,
//   })
//
type AssetStaging interface {
	constructs.Construct
	// Absolute path to the asset data.
	//
	// If asset staging is disabled, this will just be the source path or
	// a temporary directory used for bundling.
	//
	// If asset staging is enabled it will be the staged path.
	//
	// IMPORTANT: If you are going to call `addFileAsset()`, use
	// `relativeStagedPath()` instead.
	AbsoluteStagedPath() *string
	// A cryptographic hash of the asset.
	AssetHash() *string
	// Whether this asset will ultimately end up in S3 as an archive.
	//
	// It will end up as an archive either if it's a directory (which
	// the CDK will ZIP) or if it's a file that's an archive.
	IsArchive() *bool
	// The tree node.
	Node() constructs.Node
	// How this asset should be packaged.
	Packaging() FileAssetPackaging
	// The absolute path of the asset as it was referenced by the user.
	SourcePath() *string
	// Return the path to the staged asset, relative to the Cloud Assembly (manifest) directory of the given stack.
	//
	// Only returns a relative path if the asset was staged, returns an absolute path if
	// it was not staged.
	//
	// A bundled asset might end up in the outDir and still not count as
	// "staged"; if asset staging is disabled we're technically expected to
	// reference source directories, but we don't have a source directory for the
	// bundled outputs (as the bundle output is written to a temporary
	// directory). Nevertheless, we will still return an absolute path.
	//
	// A non-obvious directory layout may look like this:
	//
	// ```
	//   CLOUD ASSEMBLY ROOT
	//     +-- asset.12345abcdef/
	//     +-- assembly-Stage
	//           +-- MyStack.template.json
	//           +-- MyStack.assets.json <- will contain { "path": "../asset.12345abcdef" }
	// ```.
	RelativeStagedPath(stack Stack) *string
	// Returns a string representation of this construct.
	ToString() *string
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for AssetStaging
type jsiiProxy_AssetStaging struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_AssetStaging) AbsoluteStagedPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"absoluteStagedPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetStaging) AssetHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetStaging) IsArchive() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isArchive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetStaging) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetStaging) Packaging() FileAssetPackaging {
	var returns FileAssetPackaging
	_jsii_.Get(
		j,
		"packaging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetStaging) SourcePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourcePath",
		&returns,
	)
	return returns
}


func NewAssetStaging(scope constructs.Construct, id *string, props *AssetStagingProps) AssetStaging {
	_init_.Initialize()

	if err := validateNewAssetStagingParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AssetStaging{}

	_jsii_.Create(
		"aws-cdk-lib.AssetStaging",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewAssetStaging_Override(a AssetStaging, scope constructs.Construct, id *string, props *AssetStagingProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.AssetStaging",
		[]interface{}{scope, id, props},
		a,
	)
}

// Clears the asset hash cache.
func AssetStaging_ClearAssetHashCache() {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.AssetStaging",
		"clearAssetHashCache",
		nil, // no parameters
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
func AssetStaging_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAssetStaging_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.AssetStaging",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AssetStaging_BUNDLING_INPUT_DIR() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.AssetStaging",
		"BUNDLING_INPUT_DIR",
		&returns,
	)
	return returns
}

func AssetStaging_BUNDLING_OUTPUT_DIR() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.AssetStaging",
		"BUNDLING_OUTPUT_DIR",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AssetStaging) RelativeStagedPath(stack Stack) *string {
	if err := a.validateRelativeStagedPathParameters(stack); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"relativeStagedPath",
		[]interface{}{stack},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetStaging) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetStaging) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		a,
		"with",
		args,
		&returns,
	)

	return returns
}

