package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/cloudassemblyschema"
)

// Build an asset manifest from assets added to a stack.
//
// This class does not need to be used by app builders; it is only necessary for building Stack Synthesizers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   assetManifestBuilder := cdk.NewAssetManifestBuilder()
//
type AssetManifestBuilder interface {
	// Whether there are any assets registered in the manifest.
	HasAssets() *bool
	// Add a docker asset source and destination to the manifest.
	//
	// sourceHash should be unique for every source.
	AddDockerImageAsset(stack Stack, sourceHash *string, source *cloudassemblyschema.DockerImageSource, dest *cloudassemblyschema.DockerImageDestination) *cloudassemblyschema.DockerImageDestination
	// Add a file asset source and destination to the manifest.
	//
	// sourceHash should be unique for every source.
	AddFileAsset(stack Stack, sourceHash *string, source *cloudassemblyschema.FileSource, dest *cloudassemblyschema.FileDestination) *cloudassemblyschema.FileDestination
	// Add a docker image asset to the manifest with default settings.
	//
	// Derive the region from the stack, use the asset hash as the key, and set the prefix.
	DefaultAddDockerImageAsset(stack Stack, asset *DockerImageAssetSource, target *AssetManifestDockerImageDestination) *cloudassemblyschema.DockerImageDestination
	// Add a file asset to the manifest with default settings.
	//
	// Derive the region from the stack, use the asset hash as the key, copy the
	// file extension over, and set the prefix.
	DefaultAddFileAsset(stack Stack, asset *FileAssetSource, target *AssetManifestFileDestination) *cloudassemblyschema.FileDestination
	// Write the manifest to disk, and add it to the synthesis session.
	//
	// Return the artifact id, which should be added to the `additionalDependencies`
	// field of the stack artifact.
	EmitManifest(stack Stack, session ISynthesisSession, options *cloudassemblyschema.AssetManifestOptions, dependencies *[]*string) *string
}

// The jsii proxy struct for AssetManifestBuilder
type jsiiProxy_AssetManifestBuilder struct {
	_ byte // padding
}

func (j *jsiiProxy_AssetManifestBuilder) HasAssets() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"hasAssets",
		&returns,
	)
	return returns
}


func NewAssetManifestBuilder() AssetManifestBuilder {
	_init_.Initialize()

	j := jsiiProxy_AssetManifestBuilder{}

	_jsii_.Create(
		"aws-cdk-lib.AssetManifestBuilder",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewAssetManifestBuilder_Override(a AssetManifestBuilder) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.AssetManifestBuilder",
		nil, // no parameters
		a,
	)
}

func (a *jsiiProxy_AssetManifestBuilder) AddDockerImageAsset(stack Stack, sourceHash *string, source *cloudassemblyschema.DockerImageSource, dest *cloudassemblyschema.DockerImageDestination) *cloudassemblyschema.DockerImageDestination {
	if err := a.validateAddDockerImageAssetParameters(stack, sourceHash, source, dest); err != nil {
		panic(err)
	}
	var returns *cloudassemblyschema.DockerImageDestination

	_jsii_.Invoke(
		a,
		"addDockerImageAsset",
		[]interface{}{stack, sourceHash, source, dest},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetManifestBuilder) AddFileAsset(stack Stack, sourceHash *string, source *cloudassemblyschema.FileSource, dest *cloudassemblyschema.FileDestination) *cloudassemblyschema.FileDestination {
	if err := a.validateAddFileAssetParameters(stack, sourceHash, source, dest); err != nil {
		panic(err)
	}
	var returns *cloudassemblyschema.FileDestination

	_jsii_.Invoke(
		a,
		"addFileAsset",
		[]interface{}{stack, sourceHash, source, dest},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetManifestBuilder) DefaultAddDockerImageAsset(stack Stack, asset *DockerImageAssetSource, target *AssetManifestDockerImageDestination) *cloudassemblyschema.DockerImageDestination {
	if err := a.validateDefaultAddDockerImageAssetParameters(stack, asset, target); err != nil {
		panic(err)
	}
	var returns *cloudassemblyschema.DockerImageDestination

	_jsii_.Invoke(
		a,
		"defaultAddDockerImageAsset",
		[]interface{}{stack, asset, target},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetManifestBuilder) DefaultAddFileAsset(stack Stack, asset *FileAssetSource, target *AssetManifestFileDestination) *cloudassemblyschema.FileDestination {
	if err := a.validateDefaultAddFileAssetParameters(stack, asset, target); err != nil {
		panic(err)
	}
	var returns *cloudassemblyschema.FileDestination

	_jsii_.Invoke(
		a,
		"defaultAddFileAsset",
		[]interface{}{stack, asset, target},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetManifestBuilder) EmitManifest(stack Stack, session ISynthesisSession, options *cloudassemblyschema.AssetManifestOptions, dependencies *[]*string) *string {
	if err := a.validateEmitManifestParameters(stack, session, options); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"emitManifest",
		[]interface{}{stack, session, options, dependencies},
		&returns,
	)

	return returns
}

