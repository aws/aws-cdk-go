package cloudassemblyschema

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Protocol utility class.
type Manifest interface {
}

// The jsii proxy struct for Manifest
type jsiiProxy_Manifest struct {
	_ byte // padding
}

// Load and validates the cloud assembly manifest from file.
func Manifest_LoadAssemblyManifest(filePath *string, options *LoadManifestOptions) *AssemblyManifest {
	_init_.Initialize()

	if err := validateManifest_LoadAssemblyManifestParameters(filePath, options); err != nil {
		panic(err)
	}
	var returns *AssemblyManifest

	_jsii_.StaticInvoke(
		"aws-cdk-lib.cloud_assembly_schema.Manifest",
		"loadAssemblyManifest",
		[]interface{}{filePath, options},
		&returns,
	)

	return returns
}

// Load and validates the asset manifest from file.
func Manifest_LoadAssetManifest(filePath *string) *AssetManifest {
	_init_.Initialize()

	if err := validateManifest_LoadAssetManifestParameters(filePath); err != nil {
		panic(err)
	}
	var returns *AssetManifest

	_jsii_.StaticInvoke(
		"aws-cdk-lib.cloud_assembly_schema.Manifest",
		"loadAssetManifest",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Load and validates the integ manifest from file.
func Manifest_LoadIntegManifest(filePath *string) *IntegManifest {
	_init_.Initialize()

	if err := validateManifest_LoadIntegManifestParameters(filePath); err != nil {
		panic(err)
	}
	var returns *IntegManifest

	_jsii_.StaticInvoke(
		"aws-cdk-lib.cloud_assembly_schema.Manifest",
		"loadIntegManifest",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

// Validates and saves the cloud assembly manifest to file.
func Manifest_SaveAssemblyManifest(manifest *AssemblyManifest, filePath *string) {
	_init_.Initialize()

	if err := validateManifest_SaveAssemblyManifestParameters(manifest, filePath); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.cloud_assembly_schema.Manifest",
		"saveAssemblyManifest",
		[]interface{}{manifest, filePath},
	)
}

// Validates and saves the asset manifest to file.
func Manifest_SaveAssetManifest(manifest *AssetManifest, filePath *string) {
	_init_.Initialize()

	if err := validateManifest_SaveAssetManifestParameters(manifest, filePath); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.cloud_assembly_schema.Manifest",
		"saveAssetManifest",
		[]interface{}{manifest, filePath},
	)
}

// Validates and saves the integ manifest to file.
func Manifest_SaveIntegManifest(manifest *IntegManifest, filePath *string) {
	_init_.Initialize()

	if err := validateManifest_SaveIntegManifestParameters(manifest, filePath); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.cloud_assembly_schema.Manifest",
		"saveIntegManifest",
		[]interface{}{manifest, filePath},
	)
}

// Fetch the current schema version number.
func Manifest_Version() *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.cloud_assembly_schema.Manifest",
		"version",
		nil, // no parameters
		&returns,
	)

	return returns
}

