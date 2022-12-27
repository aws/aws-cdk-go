package cxapi

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/cloudassemblyschema"
)

// Asset manifest is a description of a set of assets which need to be built and published.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cloudAssembly cloudAssembly
//
//   nestedCloudAssemblyArtifact := awscdk.Cx_api.NewNestedCloudAssemblyArtifact(cloudAssembly, jsii.String("name"), &artifactManifest{
//   	type: awscdk.Cloud_assembly_schema.artifactType_NONE,
//
//   	// the properties below are optional
//   	dependencies: []*string{
//   		jsii.String("dependencies"),
//   	},
//   	displayName: jsii.String("displayName"),
//   	environment: jsii.String("environment"),
//   	metadata: map[string][]metadataEntry{
//   		"metadataKey": []*metadataEntry{
//   			&metadataEntry{
//   				"type": jsii.String("type"),
//
//   				// the properties below are optional
//   				"data": jsii.String("data"),
//   				"trace": []*string{
//   					jsii.String("trace"),
//   				},
//   			},
//   		},
//   	},
//   	properties: &awsCloudFormationStackProperties{
//   		templateFile: jsii.String("templateFile"),
//
//   		// the properties below are optional
//   		assumeRoleArn: jsii.String("assumeRoleArn"),
//   		assumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   		bootstrapStackVersionSsmParameter: jsii.String("bootstrapStackVersionSsmParameter"),
//   		cloudFormationExecutionRoleArn: jsii.String("cloudFormationExecutionRoleArn"),
//   		lookupRole: &bootstrapRole{
//   			arn: jsii.String("arn"),
//
//   			// the properties below are optional
//   			assumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   			bootstrapStackVersionSsmParameter: jsii.String("bootstrapStackVersionSsmParameter"),
//   			requiresBootstrapStackVersion: jsii.Number(123),
//   		},
//   		parameters: map[string]*string{
//   			"parametersKey": jsii.String("parameters"),
//   		},
//   		requiresBootstrapStackVersion: jsii.Number(123),
//   		stackName: jsii.String("stackName"),
//   		stackTemplateAssetObjectUrl: jsii.String("stackTemplateAssetObjectUrl"),
//   		tags: map[string]*string{
//   			"tagsKey": jsii.String("tags"),
//   		},
//   		terminationProtection: jsii.Boolean(false),
//   		validateOnSynth: jsii.Boolean(false),
//   	},
//   })
//
type NestedCloudAssemblyArtifact interface {
	CloudArtifact
	Assembly() CloudAssembly
	// Returns all the artifacts that this artifact depends on.
	Dependencies() *[]CloudArtifact
	// The relative directory name of the asset manifest.
	DirectoryName() *string
	// Display name.
	DisplayName() *string
	// Full path to the nested assembly directory.
	FullPath() *string
	// An identifier that shows where this artifact is located in the tree of nested assemblies, based on their manifests.
	//
	// Defaults to the normal
	// id. Should only be used in user interfaces.
	HierarchicalId() *string
	Id() *string
	// The artifact's manifest.
	Manifest() *cloudassemblyschema.ArtifactManifest
	// The set of messages extracted from the artifact's metadata.
	Messages() *[]*SynthesisMessage
	// The nested Assembly.
	NestedAssembly() CloudAssembly
	// Returns: all the metadata entries of a specific type in this artifact.
	FindMetadataByType(type_ *string) *[]*MetadataEntryResult
}

// The jsii proxy struct for NestedCloudAssemblyArtifact
type jsiiProxy_NestedCloudAssemblyArtifact struct {
	jsiiProxy_CloudArtifact
}

func (j *jsiiProxy_NestedCloudAssemblyArtifact) Assembly() CloudAssembly {
	var returns CloudAssembly
	_jsii_.Get(
		j,
		"assembly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedCloudAssemblyArtifact) Dependencies() *[]CloudArtifact {
	var returns *[]CloudArtifact
	_jsii_.Get(
		j,
		"dependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedCloudAssemblyArtifact) DirectoryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedCloudAssemblyArtifact) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedCloudAssemblyArtifact) FullPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedCloudAssemblyArtifact) HierarchicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hierarchicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedCloudAssemblyArtifact) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedCloudAssemblyArtifact) Manifest() *cloudassemblyschema.ArtifactManifest {
	var returns *cloudassemblyschema.ArtifactManifest
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedCloudAssemblyArtifact) Messages() *[]*SynthesisMessage {
	var returns *[]*SynthesisMessage
	_jsii_.Get(
		j,
		"messages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NestedCloudAssemblyArtifact) NestedAssembly() CloudAssembly {
	var returns CloudAssembly
	_jsii_.Get(
		j,
		"nestedAssembly",
		&returns,
	)
	return returns
}


func NewNestedCloudAssemblyArtifact(assembly CloudAssembly, name *string, artifact *cloudassemblyschema.ArtifactManifest) NestedCloudAssemblyArtifact {
	_init_.Initialize()

	if err := validateNewNestedCloudAssemblyArtifactParameters(assembly, name, artifact); err != nil {
		panic(err)
	}
	j := jsiiProxy_NestedCloudAssemblyArtifact{}

	_jsii_.Create(
		"aws-cdk-lib.cx_api.NestedCloudAssemblyArtifact",
		[]interface{}{assembly, name, artifact},
		&j,
	)

	return &j
}

func NewNestedCloudAssemblyArtifact_Override(n NestedCloudAssemblyArtifact, assembly CloudAssembly, name *string, artifact *cloudassemblyschema.ArtifactManifest) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.cx_api.NestedCloudAssemblyArtifact",
		[]interface{}{assembly, name, artifact},
		n,
	)
}

// Returns a subclass of `CloudArtifact` based on the artifact type defined in the artifact manifest.
//
// Returns: the `CloudArtifact` that matches the artifact type or `undefined` if it's an artifact type that is unrecognized by this module.
func NestedCloudAssemblyArtifact_FromManifest(assembly CloudAssembly, id *string, artifact *cloudassemblyschema.ArtifactManifest) CloudArtifact {
	_init_.Initialize()

	if err := validateNestedCloudAssemblyArtifact_FromManifestParameters(assembly, id, artifact); err != nil {
		panic(err)
	}
	var returns CloudArtifact

	_jsii_.StaticInvoke(
		"aws-cdk-lib.cx_api.NestedCloudAssemblyArtifact",
		"fromManifest",
		[]interface{}{assembly, id, artifact},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NestedCloudAssemblyArtifact) FindMetadataByType(type_ *string) *[]*MetadataEntryResult {
	if err := n.validateFindMetadataByTypeParameters(type_); err != nil {
		panic(err)
	}
	var returns *[]*MetadataEntryResult

	_jsii_.Invoke(
		n,
		"findMetadataByType",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

