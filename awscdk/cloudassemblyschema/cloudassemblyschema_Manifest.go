package cloudassemblyschema

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Protocol utility class.
// Experimental.
type Manifest interface {
}

// The jsii proxy struct for Manifest
type jsiiProxy_Manifest struct {
	_ byte // padding
}

// Deprecated.
// Deprecated: use `loadAssemblyManifest()`.
func Manifest_Load(filePath *string) *AssemblyManifest {
	_init_.Initialize()

	var returns *AssemblyManifest

	_jsii_.StaticInvoke(
		"monocdk.cloud_assembly_schema.Manifest",
		"load",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Load and validates the cloud assembly manifest from file.
// Experimental.
func Manifest_LoadAssemblyManifest(filePath *string, options *LoadManifestOptions) *AssemblyManifest {
	_init_.Initialize()

	var returns *AssemblyManifest

	_jsii_.StaticInvoke(
		"monocdk.cloud_assembly_schema.Manifest",
		"loadAssemblyManifest",
		[]interface{}{filePath, options},
		&returns,
	)

	return returns
}

// Load and validates the asset manifest from file.
// Experimental.
func Manifest_LoadAssetManifest(filePath *string) *AssetManifest {
	_init_.Initialize()

	var returns *AssetManifest

	_jsii_.StaticInvoke(
		"monocdk.cloud_assembly_schema.Manifest",
		"loadAssetManifest",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Load and validates the integ manifest from file.
// Experimental.
func Manifest_LoadIntegManifest(filePath *string) *IntegManifest {
	_init_.Initialize()

	var returns *IntegManifest

	_jsii_.StaticInvoke(
		"monocdk.cloud_assembly_schema.Manifest",
		"loadIntegManifest",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Deprecated.
// Deprecated: use `saveAssemblyManifest()`.
func Manifest_Save(manifest *AssemblyManifest, filePath *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.cloud_assembly_schema.Manifest",
		"save",
		[]interface{}{manifest, filePath},
	)
}

// Validates and saves the cloud assembly manifest to file.
// Experimental.
func Manifest_SaveAssemblyManifest(manifest *AssemblyManifest, filePath *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.cloud_assembly_schema.Manifest",
		"saveAssemblyManifest",
		[]interface{}{manifest, filePath},
	)
}

// Validates and saves the asset manifest to file.
// Experimental.
func Manifest_SaveAssetManifest(manifest *AssetManifest, filePath *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.cloud_assembly_schema.Manifest",
		"saveAssetManifest",
		[]interface{}{manifest, filePath},
	)
}

// Validates and saves the integ manifest to file.
// Experimental.
func Manifest_SaveIntegManifest(manifest *IntegManifest, filePath *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.cloud_assembly_schema.Manifest",
		"saveIntegManifest",
		[]interface{}{manifest, filePath},
	)
}

// Fetch the current schema version number.
// Experimental.
func Manifest_Version() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"monocdk.cloud_assembly_schema.Manifest",
		"version",
		nil, // no parameters
		&returns,
	)

	return returns
}

