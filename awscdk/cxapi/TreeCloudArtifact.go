package cxapi

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/cloudassemblyschema"
)

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cloudAssembly cloudAssembly
//
//   treeCloudArtifact := awscdk.Cx_api.NewTreeCloudArtifact(cloudAssembly, jsii.String("name"), &ArtifactManifest{
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
type TreeCloudArtifact interface {
	CloudArtifact
	Assembly() CloudAssembly
	// Returns all the artifacts that this artifact depends on.
	Dependencies() *[]CloudArtifact
	File() *string
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
	// Returns: all the metadata entries of a specific type in this artifact.
	FindMetadataByType(type_ *string) *[]*MetadataEntryResult
}

// The jsii proxy struct for TreeCloudArtifact
type jsiiProxy_TreeCloudArtifact struct {
	jsiiProxy_CloudArtifact
}

func (j *jsiiProxy_TreeCloudArtifact) Assembly() CloudAssembly {
	var returns CloudAssembly
	_jsii_.Get(
		j,
		"assembly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TreeCloudArtifact) Dependencies() *[]CloudArtifact {
	var returns *[]CloudArtifact
	_jsii_.Get(
		j,
		"dependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TreeCloudArtifact) File() *string {
	var returns *string
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TreeCloudArtifact) HierarchicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hierarchicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TreeCloudArtifact) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TreeCloudArtifact) Manifest() *cloudassemblyschema.ArtifactManifest {
	var returns *cloudassemblyschema.ArtifactManifest
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TreeCloudArtifact) Messages() *[]*SynthesisMessage {
	var returns *[]*SynthesisMessage
	_jsii_.Get(
		j,
		"messages",
		&returns,
	)
	return returns
}


func NewTreeCloudArtifact(assembly CloudAssembly, name *string, artifact *cloudassemblyschema.ArtifactManifest) TreeCloudArtifact {
	_init_.Initialize()

	if err := validateNewTreeCloudArtifactParameters(assembly, name, artifact); err != nil {
		panic(err)
	}
	j := jsiiProxy_TreeCloudArtifact{}

	_jsii_.Create(
		"aws-cdk-lib.cx_api.TreeCloudArtifact",
		[]interface{}{assembly, name, artifact},
		&j,
	)

	return &j
}

func NewTreeCloudArtifact_Override(t TreeCloudArtifact, assembly CloudAssembly, name *string, artifact *cloudassemblyschema.ArtifactManifest) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.cx_api.TreeCloudArtifact",
		[]interface{}{assembly, name, artifact},
		t,
	)
}

// Returns a subclass of `CloudArtifact` based on the artifact type defined in the artifact manifest.
//
// Returns: the `CloudArtifact` that matches the artifact type or `undefined` if it's an artifact type that is unrecognized by this module.
func TreeCloudArtifact_FromManifest(assembly CloudAssembly, id *string, artifact *cloudassemblyschema.ArtifactManifest) CloudArtifact {
	_init_.Initialize()

	if err := validateTreeCloudArtifact_FromManifestParameters(assembly, id, artifact); err != nil {
		panic(err)
	}
	var returns CloudArtifact

	_jsii_.StaticInvoke(
		"aws-cdk-lib.cx_api.TreeCloudArtifact",
		"fromManifest",
		[]interface{}{assembly, id, artifact},
		&returns,
	)

	return returns
}

// Checks if `art` is an instance of this class.
//
// Use this method instead of `instanceof` to properly detect `TreeCloudArtifact`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `cx-api` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `TreeCloudArtifact` in each copy of the `cx-api` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `cx-api`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
func TreeCloudArtifact_IsTreeCloudArtifact(art interface{}) *bool {
	_init_.Initialize()

	if err := validateTreeCloudArtifact_IsTreeCloudArtifactParameters(art); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.cx_api.TreeCloudArtifact",
		"isTreeCloudArtifact",
		[]interface{}{art},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TreeCloudArtifact) FindMetadataByType(type_ *string) *[]*MetadataEntryResult {
	if err := t.validateFindMetadataByTypeParameters(type_); err != nil {
		panic(err)
	}
	var returns *[]*MetadataEntryResult

	_jsii_.Invoke(
		t,
		"findMetadataByType",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

