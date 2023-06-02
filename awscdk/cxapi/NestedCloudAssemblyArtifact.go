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
//   nestedCloudAssemblyArtifact := awscdk.Cx_api.NewNestedCloudAssemblyArtifact(cloudAssembly, jsii.String("name"), &ArtifactManifest{
//   	Type: awscdk.Cloud_assembly_schema.ArtifactType_NONE,
//
//   	// the properties below are optional
//   	Dependencies: []*string{
//   		jsii.String("dependencies"),
//   	},
//   	DisplayName: jsii.String("displayName"),
//   	Environment: jsii.String("environment"),
//   	Metadata: map[string][]metadataEntry{
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
//   	Properties: &AwsCloudFormationStackProperties{
//   		TemplateFile: jsii.String("templateFile"),
//
//   		// the properties below are optional
//   		AssumeRoleArn: jsii.String("assumeRoleArn"),
//   		AssumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   		BootstrapStackVersionSsmParameter: jsii.String("bootstrapStackVersionSsmParameter"),
//   		CloudFormationExecutionRoleArn: jsii.String("cloudFormationExecutionRoleArn"),
//   		LookupRole: &BootstrapRole{
//   			Arn: jsii.String("arn"),
//
//   			// the properties below are optional
//   			AssumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   			BootstrapStackVersionSsmParameter: jsii.String("bootstrapStackVersionSsmParameter"),
//   			RequiresBootstrapStackVersion: jsii.Number(123),
//   		},
//   		Parameters: map[string]*string{
//   			"parametersKey": jsii.String("parameters"),
//   		},
//   		RequiresBootstrapStackVersion: jsii.Number(123),
//   		StackName: jsii.String("stackName"),
//   		StackTemplateAssetObjectUrl: jsii.String("stackTemplateAssetObjectUrl"),
//   		Tags: map[string]*string{
//   			"tagsKey": jsii.String("tags"),
//   		},
//   		TerminationProtection: jsii.Boolean(false),
//   		ValidateOnSynth: jsii.Boolean(false),
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

// Checks if `art` is an instance of this class.
//
// Use this method instead of `instanceof` to properly detect `NestedCloudAssemblyArtifact`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `cx-api` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `NestedCloudAssemblyArtifact` in each copy of the `cx-api` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `cx-api`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
func NestedCloudAssemblyArtifact_IsNestedCloudAssemblyArtifact(art interface{}) *bool {
	_init_.Initialize()

	if err := validateNestedCloudAssemblyArtifact_IsNestedCloudAssemblyArtifactParameters(art); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.cx_api.NestedCloudAssemblyArtifact",
		"isNestedCloudAssemblyArtifact",
		[]interface{}{art},
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

