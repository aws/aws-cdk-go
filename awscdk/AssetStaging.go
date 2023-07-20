package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Stages a file or directory from a location on the file system into a staging directory.
//
// This is controlled by the context key 'aws:cdk:asset-staging' and enabled
// by the CLI by default in order to ensure that when the CDK app exists, all
// assets are available for deployment. Otherwise, if an app references assets
// in temporary locations, those will not be available when it exists (see
// https://github.com/aws/aws-cdk/issues/1716).
//
// The `stagedPath` property is a stringified token that represents the location
// of the file or directory after staging. It will be resolved only during the
// "prepare" stage and may be either the original path or the staged path
// depending on the context setting.
//
// The file/directory are staged based on their content hash (fingerprint). This
// means that only if content was changed, copy will happen.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dockerImage dockerImage
//   var localBundling iLocalBundling
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
//   		Volumes: []dockerVolume{
//   			&dockerVolume{
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
	// Whether this asset is an archive (zip or jar).
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

