package cxapi

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/cloudassemblyschema"
)

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cx_api "github.com/aws/aws-cdk-go/awscdk/cx_api"
//   assemblyBuildOptions := &assemblyBuildOptions{
//   }
//
type AssemblyBuildOptions struct {
}

// Asset manifest is a description of a set of assets which need to be built and published.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"import awscdk "github.com/aws/aws-cdk-go/awscdk"import cx_api "github.com/aws/aws-cdk-go/awscdk/cx_api"
//
//   var cloudAssembly cloudAssembly
//   assetManifestArtifact := cx_api.NewAssetManifestArtifact(cloudAssembly, jsii.String("name"), &artifactManifest{
//   	type: cloud_assembly_schema.artifactType_NONE,
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
type AssetManifestArtifact interface {
	CloudArtifact
	Assembly() CloudAssembly
	// Name of SSM parameter with bootstrap stack version.
	BootstrapStackVersionSsmParameter() *string
	// Returns all the artifacts that this artifact depends on.
	Dependencies() *[]CloudArtifact
	// The file name of the asset manifest.
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
	// Version of bootstrap stack required to deploy this stack.
	RequiresBootstrapStackVersion() *float64
	// Returns: all the metadata entries of a specific type in this artifact.
	FindMetadataByType(type_ *string) *[]*MetadataEntryResult
}

// The jsii proxy struct for AssetManifestArtifact
type jsiiProxy_AssetManifestArtifact struct {
	jsiiProxy_CloudArtifact
}

func (j *jsiiProxy_AssetManifestArtifact) Assembly() CloudAssembly {
	var returns CloudAssembly
	_jsii_.Get(
		j,
		"assembly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetManifestArtifact) BootstrapStackVersionSsmParameter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapStackVersionSsmParameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetManifestArtifact) Dependencies() *[]CloudArtifact {
	var returns *[]CloudArtifact
	_jsii_.Get(
		j,
		"dependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetManifestArtifact) File() *string {
	var returns *string
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetManifestArtifact) HierarchicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hierarchicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetManifestArtifact) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetManifestArtifact) Manifest() *cloudassemblyschema.ArtifactManifest {
	var returns *cloudassemblyschema.ArtifactManifest
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetManifestArtifact) Messages() *[]*SynthesisMessage {
	var returns *[]*SynthesisMessage
	_jsii_.Get(
		j,
		"messages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetManifestArtifact) RequiresBootstrapStackVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiresBootstrapStackVersion",
		&returns,
	)
	return returns
}


func NewAssetManifestArtifact(assembly CloudAssembly, name *string, artifact *cloudassemblyschema.ArtifactManifest) AssetManifestArtifact {
	_init_.Initialize()

	j := jsiiProxy_AssetManifestArtifact{}

	_jsii_.Create(
		"aws-cdk-lib.cx_api.AssetManifestArtifact",
		[]interface{}{assembly, name, artifact},
		&j,
	)

	return &j
}

func NewAssetManifestArtifact_Override(a AssetManifestArtifact, assembly CloudAssembly, name *string, artifact *cloudassemblyschema.ArtifactManifest) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.cx_api.AssetManifestArtifact",
		[]interface{}{assembly, name, artifact},
		a,
	)
}

// Returns a subclass of `CloudArtifact` based on the artifact type defined in the artifact manifest.
//
// Returns: the `CloudArtifact` that matches the artifact type or `undefined` if it's an artifact type that is unrecognized by this module.
func AssetManifestArtifact_FromManifest(assembly CloudAssembly, id *string, artifact *cloudassemblyschema.ArtifactManifest) CloudArtifact {
	_init_.Initialize()

	var returns CloudArtifact

	_jsii_.StaticInvoke(
		"aws-cdk-lib.cx_api.AssetManifestArtifact",
		"fromManifest",
		[]interface{}{assembly, id, artifact},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetManifestArtifact) FindMetadataByType(type_ *string) *[]*MetadataEntryResult {
	var returns *[]*MetadataEntryResult

	_jsii_.Invoke(
		a,
		"findMetadataByType",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

// Artifact properties for CloudFormation stacks.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cx_api "github.com/aws/aws-cdk-go/awscdk/cx_api"
//   awsCloudFormationStackProperties := &awsCloudFormationStackProperties{
//   	templateFile: jsii.String("templateFile"),
//
//   	// the properties below are optional
//   	parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	stackName: jsii.String("stackName"),
//   	terminationProtection: jsii.Boolean(false),
//   }
//
type AwsCloudFormationStackProperties struct {
	// A file relative to the assembly root which contains the CloudFormation template for this stack.
	TemplateFile *string `json:"templateFile" yaml:"templateFile"`
	// Values for CloudFormation stack parameters that should be passed when the stack is deployed.
	Parameters *map[string]*string `json:"parameters" yaml:"parameters"`
	// The name to use for the CloudFormation stack.
	StackName *string `json:"stackName" yaml:"stackName"`
	// Whether to enable termination protection for this stack.
	TerminationProtection *bool `json:"terminationProtection" yaml:"terminationProtection"`
}

// Represents an artifact within a cloud assembly.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"import awscdk "github.com/aws/aws-cdk-go/awscdk"import cx_api "github.com/aws/aws-cdk-go/awscdk/cx_api"
//
//   var cloudAssembly cloudAssembly
//   cloudArtifact := cx_api.cloudArtifact.fromManifest(cloudAssembly, jsii.String("MyCloudArtifact"), &artifactManifest{
//   	type: cloud_assembly_schema.artifactType_NONE,
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
type CloudArtifact interface {
	Assembly() CloudAssembly
	// Returns all the artifacts that this artifact depends on.
	Dependencies() *[]CloudArtifact
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

// The jsii proxy struct for CloudArtifact
type jsiiProxy_CloudArtifact struct {
	_ byte // padding
}

func (j *jsiiProxy_CloudArtifact) Assembly() CloudAssembly {
	var returns CloudAssembly
	_jsii_.Get(
		j,
		"assembly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudArtifact) Dependencies() *[]CloudArtifact {
	var returns *[]CloudArtifact
	_jsii_.Get(
		j,
		"dependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudArtifact) HierarchicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hierarchicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudArtifact) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudArtifact) Manifest() *cloudassemblyschema.ArtifactManifest {
	var returns *cloudassemblyschema.ArtifactManifest
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudArtifact) Messages() *[]*SynthesisMessage {
	var returns *[]*SynthesisMessage
	_jsii_.Get(
		j,
		"messages",
		&returns,
	)
	return returns
}


func NewCloudArtifact(assembly CloudAssembly, id *string, manifest *cloudassemblyschema.ArtifactManifest) CloudArtifact {
	_init_.Initialize()

	j := jsiiProxy_CloudArtifact{}

	_jsii_.Create(
		"aws-cdk-lib.cx_api.CloudArtifact",
		[]interface{}{assembly, id, manifest},
		&j,
	)

	return &j
}

func NewCloudArtifact_Override(c CloudArtifact, assembly CloudAssembly, id *string, manifest *cloudassemblyschema.ArtifactManifest) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.cx_api.CloudArtifact",
		[]interface{}{assembly, id, manifest},
		c,
	)
}

// Returns a subclass of `CloudArtifact` based on the artifact type defined in the artifact manifest.
//
// Returns: the `CloudArtifact` that matches the artifact type or `undefined` if it's an artifact type that is unrecognized by this module.
func CloudArtifact_FromManifest(assembly CloudAssembly, id *string, artifact *cloudassemblyschema.ArtifactManifest) CloudArtifact {
	_init_.Initialize()

	var returns CloudArtifact

	_jsii_.StaticInvoke(
		"aws-cdk-lib.cx_api.CloudArtifact",
		"fromManifest",
		[]interface{}{assembly, id, artifact},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudArtifact) FindMetadataByType(type_ *string) *[]*MetadataEntryResult {
	var returns *[]*MetadataEntryResult

	_jsii_.Invoke(
		c,
		"findMetadataByType",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

// Represents a deployable cloud application.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cx_api "github.com/aws/aws-cdk-go/awscdk/cx_api"
//   cloudAssembly := cx_api.NewCloudAssembly(jsii.String("directory"), &loadManifestOptions{
//   	skipEnumCheck: jsii.Boolean(false),
//   	skipVersionCheck: jsii.Boolean(false),
//   })
//
type CloudAssembly interface {
	// All artifacts included in this assembly.
	Artifacts() *[]CloudArtifact
	// The root directory of the cloud assembly.
	Directory() *string
	// The raw assembly manifest.
	Manifest() *cloudassemblyschema.AssemblyManifest
	// The nested assembly artifacts in this assembly.
	NestedAssemblies() *[]NestedCloudAssemblyArtifact
	// Runtime information such as module versions used to synthesize this assembly.
	Runtime() *cloudassemblyschema.RuntimeInfo
	// Returns: all the CloudFormation stack artifacts that are included in this assembly.
	Stacks() *[]CloudFormationStackArtifact
	// Returns all the stacks, including the ones in nested assemblies.
	StacksRecursively() *[]CloudFormationStackArtifact
	// The schema version of the assembly manifest.
	Version() *string
	// Returns a nested assembly.
	GetNestedAssembly(artifactId *string) CloudAssembly
	// Returns a nested assembly artifact.
	GetNestedAssemblyArtifact(artifactId *string) NestedCloudAssemblyArtifact
	// Returns a CloudFormation stack artifact from this assembly.
	//
	// Returns: a `CloudFormationStackArtifact` object.
	GetStackArtifact(artifactId *string) CloudFormationStackArtifact
	// Returns a CloudFormation stack artifact from this assembly.
	//
	// Will only search the current assembly.
	//
	// Returns: a `CloudFormationStackArtifact` object.
	GetStackByName(stackName *string) CloudFormationStackArtifact
	// Returns the tree metadata artifact from this assembly.
	//
	// Returns: a `TreeCloudArtifact` object if there is one defined in the manifest, `undefined` otherwise.
	Tree() TreeCloudArtifact
	// Attempts to find an artifact with a specific identity.
	//
	// Returns: A `CloudArtifact` object or `undefined` if the artifact does not exist in this assembly.
	TryGetArtifact(id *string) CloudArtifact
}

// The jsii proxy struct for CloudAssembly
type jsiiProxy_CloudAssembly struct {
	_ byte // padding
}

func (j *jsiiProxy_CloudAssembly) Artifacts() *[]CloudArtifact {
	var returns *[]CloudArtifact
	_jsii_.Get(
		j,
		"artifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssembly) Directory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssembly) Manifest() *cloudassemblyschema.AssemblyManifest {
	var returns *cloudassemblyschema.AssemblyManifest
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssembly) NestedAssemblies() *[]NestedCloudAssemblyArtifact {
	var returns *[]NestedCloudAssemblyArtifact
	_jsii_.Get(
		j,
		"nestedAssemblies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssembly) Runtime() *cloudassemblyschema.RuntimeInfo {
	var returns *cloudassemblyschema.RuntimeInfo
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssembly) Stacks() *[]CloudFormationStackArtifact {
	var returns *[]CloudFormationStackArtifact
	_jsii_.Get(
		j,
		"stacks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssembly) StacksRecursively() *[]CloudFormationStackArtifact {
	var returns *[]CloudFormationStackArtifact
	_jsii_.Get(
		j,
		"stacksRecursively",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssembly) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Reads a cloud assembly from the specified directory.
func NewCloudAssembly(directory *string, loadOptions *cloudassemblyschema.LoadManifestOptions) CloudAssembly {
	_init_.Initialize()

	j := jsiiProxy_CloudAssembly{}

	_jsii_.Create(
		"aws-cdk-lib.cx_api.CloudAssembly",
		[]interface{}{directory, loadOptions},
		&j,
	)

	return &j
}

// Reads a cloud assembly from the specified directory.
func NewCloudAssembly_Override(c CloudAssembly, directory *string, loadOptions *cloudassemblyschema.LoadManifestOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.cx_api.CloudAssembly",
		[]interface{}{directory, loadOptions},
		c,
	)
}

func (c *jsiiProxy_CloudAssembly) GetNestedAssembly(artifactId *string) CloudAssembly {
	var returns CloudAssembly

	_jsii_.Invoke(
		c,
		"getNestedAssembly",
		[]interface{}{artifactId},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAssembly) GetNestedAssemblyArtifact(artifactId *string) NestedCloudAssemblyArtifact {
	var returns NestedCloudAssemblyArtifact

	_jsii_.Invoke(
		c,
		"getNestedAssemblyArtifact",
		[]interface{}{artifactId},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAssembly) GetStackArtifact(artifactId *string) CloudFormationStackArtifact {
	var returns CloudFormationStackArtifact

	_jsii_.Invoke(
		c,
		"getStackArtifact",
		[]interface{}{artifactId},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAssembly) GetStackByName(stackName *string) CloudFormationStackArtifact {
	var returns CloudFormationStackArtifact

	_jsii_.Invoke(
		c,
		"getStackByName",
		[]interface{}{stackName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAssembly) Tree() TreeCloudArtifact {
	var returns TreeCloudArtifact

	_jsii_.Invoke(
		c,
		"tree",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAssembly) TryGetArtifact(id *string) CloudArtifact {
	var returns CloudArtifact

	_jsii_.Invoke(
		c,
		"tryGetArtifact",
		[]interface{}{id},
		&returns,
	)

	return returns
}

// Can be used to build a cloud assembly.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cx_api "github.com/aws/aws-cdk-go/awscdk/cx_api"
//
//   var cloudAssemblyBuilder cloudAssemblyBuilder
//   cloudAssemblyBuilder := cx_api.NewCloudAssemblyBuilder(jsii.String("outdir"), &cloudAssemblyBuilderProps{
//   	assetOutdir: jsii.String("assetOutdir"),
//   	parentBuilder: cloudAssemblyBuilder,
//   })
//
type CloudAssemblyBuilder interface {
	// The directory where assets of this Cloud Assembly should be stored.
	AssetOutdir() *string
	// The root directory of the resulting cloud assembly.
	Outdir() *string
	// Adds an artifact into the cloud assembly.
	AddArtifact(id *string, manifest *cloudassemblyschema.ArtifactManifest)
	// Reports that some context is missing in order for this cloud assembly to be fully synthesized.
	AddMissing(missing *cloudassemblyschema.MissingContext)
	// Finalizes the cloud assembly into the output directory returns a `CloudAssembly` object that can be used to inspect the assembly.
	BuildAssembly(options *AssemblyBuildOptions) CloudAssembly
	// Creates a nested cloud assembly.
	CreateNestedAssembly(artifactId *string, displayName *string) CloudAssemblyBuilder
}

// The jsii proxy struct for CloudAssemblyBuilder
type jsiiProxy_CloudAssemblyBuilder struct {
	_ byte // padding
}

func (j *jsiiProxy_CloudAssemblyBuilder) AssetOutdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetOutdir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAssemblyBuilder) Outdir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdir",
		&returns,
	)
	return returns
}


// Initializes a cloud assembly builder.
func NewCloudAssemblyBuilder(outdir *string, props *CloudAssemblyBuilderProps) CloudAssemblyBuilder {
	_init_.Initialize()

	j := jsiiProxy_CloudAssemblyBuilder{}

	_jsii_.Create(
		"aws-cdk-lib.cx_api.CloudAssemblyBuilder",
		[]interface{}{outdir, props},
		&j,
	)

	return &j
}

// Initializes a cloud assembly builder.
func NewCloudAssemblyBuilder_Override(c CloudAssemblyBuilder, outdir *string, props *CloudAssemblyBuilderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.cx_api.CloudAssemblyBuilder",
		[]interface{}{outdir, props},
		c,
	)
}

func (c *jsiiProxy_CloudAssemblyBuilder) AddArtifact(id *string, manifest *cloudassemblyschema.ArtifactManifest) {
	_jsii_.InvokeVoid(
		c,
		"addArtifact",
		[]interface{}{id, manifest},
	)
}

func (c *jsiiProxy_CloudAssemblyBuilder) AddMissing(missing *cloudassemblyschema.MissingContext) {
	_jsii_.InvokeVoid(
		c,
		"addMissing",
		[]interface{}{missing},
	)
}

func (c *jsiiProxy_CloudAssemblyBuilder) BuildAssembly(options *AssemblyBuildOptions) CloudAssembly {
	var returns CloudAssembly

	_jsii_.Invoke(
		c,
		"buildAssembly",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAssemblyBuilder) CreateNestedAssembly(artifactId *string, displayName *string) CloudAssemblyBuilder {
	var returns CloudAssemblyBuilder

	_jsii_.Invoke(
		c,
		"createNestedAssembly",
		[]interface{}{artifactId, displayName},
		&returns,
	)

	return returns
}

// Construction properties for CloudAssemblyBuilder.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cx_api "github.com/aws/aws-cdk-go/awscdk/cx_api"
//
//   var cloudAssemblyBuilder cloudAssemblyBuilder
//   cloudAssemblyBuilderProps := &cloudAssemblyBuilderProps{
//   	assetOutdir: jsii.String("assetOutdir"),
//   	parentBuilder: cloudAssemblyBuilder,
//   }
//
type CloudAssemblyBuilderProps struct {
	// Use the given asset output directory.
	AssetOutdir *string `json:"assetOutdir" yaml:"assetOutdir"`
	// If this builder is for a nested assembly, the parent assembly builder.
	ParentBuilder CloudAssemblyBuilder `json:"parentBuilder" yaml:"parentBuilder"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"import awscdk "github.com/aws/aws-cdk-go/awscdk"import cx_api "github.com/aws/aws-cdk-go/awscdk/cx_api"
//
//   var cloudAssembly cloudAssembly
//   cloudFormationStackArtifact := cx_api.NewCloudFormationStackArtifact(cloudAssembly, jsii.String("artifactId"), &artifactManifest{
//   	type: cloud_assembly_schema.artifactType_NONE,
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
type CloudFormationStackArtifact interface {
	CloudArtifact
	Assembly() CloudAssembly
	// Any assets associated with this stack.
	Assets() *[]interface{}
	// The role that needs to be assumed to deploy the stack.
	AssumeRoleArn() *string
	// External ID to use when assuming role for cloudformation deployments.
	AssumeRoleExternalId() *string
	// Name of SSM parameter with bootstrap stack version.
	BootstrapStackVersionSsmParameter() *string
	// The role that is passed to CloudFormation to execute the change set.
	CloudFormationExecutionRoleArn() *string
	// Returns all the artifacts that this artifact depends on.
	Dependencies() *[]CloudArtifact
	// A string that represents this stack.
	//
	// Should only be used in user
	// interfaces. If the stackName has not been set explicitly, or has been set
	// to artifactId, it will return the hierarchicalId of the stack. Otherwise,
	// it will return something like "<hierarchicalId> (<stackName>)".
	DisplayName() *string
	// The environment into which to deploy this artifact.
	Environment() *Environment
	// An identifier that shows where this artifact is located in the tree of nested assemblies, based on their manifests.
	//
	// Defaults to the normal
	// id. Should only be used in user interfaces.
	HierarchicalId() *string
	Id() *string
	// The role to use to look up values from the target AWS account.
	LookupRole() *cloudassemblyschema.BootstrapRole
	// The artifact's manifest.
	Manifest() *cloudassemblyschema.ArtifactManifest
	// The set of messages extracted from the artifact's metadata.
	Messages() *[]*SynthesisMessage
	// The original name as defined in the CDK app.
	OriginalName() *string
	// CloudFormation parameters to pass to the stack.
	Parameters() *map[string]*string
	// Version of bootstrap stack required to deploy this stack.
	RequiresBootstrapStackVersion() *float64
	// The physical name of this stack.
	StackName() *string
	// If the stack template has already been included in the asset manifest, its asset URL.
	StackTemplateAssetObjectUrl() *string
	// CloudFormation tags to pass to the stack.
	Tags() *map[string]*string
	// The CloudFormation template for this stack.
	Template() interface{}
	// The file name of the template.
	TemplateFile() *string
	// Full path to the template file.
	TemplateFullPath() *string
	// Whether termination protection is enabled for this stack.
	TerminationProtection() *bool
	// Whether this stack should be validated by the CLI after synthesis.
	ValidateOnSynth() *bool
	// Returns: all the metadata entries of a specific type in this artifact.
	FindMetadataByType(type_ *string) *[]*MetadataEntryResult
}

// The jsii proxy struct for CloudFormationStackArtifact
type jsiiProxy_CloudFormationStackArtifact struct {
	jsiiProxy_CloudArtifact
}

func (j *jsiiProxy_CloudFormationStackArtifact) Assembly() CloudAssembly {
	var returns CloudAssembly
	_jsii_.Get(
		j,
		"assembly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) Assets() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"assets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) AssumeRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) AssumeRoleExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assumeRoleExternalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) BootstrapStackVersionSsmParameter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapStackVersionSsmParameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) CloudFormationExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudFormationExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) Dependencies() *[]CloudArtifact {
	var returns *[]CloudArtifact
	_jsii_.Get(
		j,
		"dependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) Environment() *Environment {
	var returns *Environment
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) HierarchicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hierarchicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) LookupRole() *cloudassemblyschema.BootstrapRole {
	var returns *cloudassemblyschema.BootstrapRole
	_jsii_.Get(
		j,
		"lookupRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) Manifest() *cloudassemblyschema.ArtifactManifest {
	var returns *cloudassemblyschema.ArtifactManifest
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) Messages() *[]*SynthesisMessage {
	var returns *[]*SynthesisMessage
	_jsii_.Get(
		j,
		"messages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) OriginalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) RequiresBootstrapStackVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiresBootstrapStackVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) StackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) StackTemplateAssetObjectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackTemplateAssetObjectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) Template() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"template",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) TemplateFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) TemplateFullPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateFullPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) TerminationProtection() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"terminationProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFormationStackArtifact) ValidateOnSynth() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"validateOnSynth",
		&returns,
	)
	return returns
}


func NewCloudFormationStackArtifact(assembly CloudAssembly, artifactId *string, artifact *cloudassemblyschema.ArtifactManifest) CloudFormationStackArtifact {
	_init_.Initialize()

	j := jsiiProxy_CloudFormationStackArtifact{}

	_jsii_.Create(
		"aws-cdk-lib.cx_api.CloudFormationStackArtifact",
		[]interface{}{assembly, artifactId, artifact},
		&j,
	)

	return &j
}

func NewCloudFormationStackArtifact_Override(c CloudFormationStackArtifact, assembly CloudAssembly, artifactId *string, artifact *cloudassemblyschema.ArtifactManifest) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.cx_api.CloudFormationStackArtifact",
		[]interface{}{assembly, artifactId, artifact},
		c,
	)
}

// Returns a subclass of `CloudArtifact` based on the artifact type defined in the artifact manifest.
//
// Returns: the `CloudArtifact` that matches the artifact type or `undefined` if it's an artifact type that is unrecognized by this module.
func CloudFormationStackArtifact_FromManifest(assembly CloudAssembly, id *string, artifact *cloudassemblyschema.ArtifactManifest) CloudArtifact {
	_init_.Initialize()

	var returns CloudArtifact

	_jsii_.StaticInvoke(
		"aws-cdk-lib.cx_api.CloudFormationStackArtifact",
		"fromManifest",
		[]interface{}{assembly, id, artifact},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFormationStackArtifact) FindMetadataByType(type_ *string) *[]*MetadataEntryResult {
	var returns *[]*MetadataEntryResult

	_jsii_.Invoke(
		c,
		"findMetadataByType",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

// Query to hosted zone context provider.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cx_api "github.com/aws/aws-cdk-go/awscdk/cx_api"
//   endpointServiceAvailabilityZonesContextQuery := &endpointServiceAvailabilityZonesContextQuery{
//   	account: jsii.String("account"),
//   	region: jsii.String("region"),
//   	serviceName: jsii.String("serviceName"),
//   }
//
type EndpointServiceAvailabilityZonesContextQuery struct {
	// Query account.
	Account *string `json:"account" yaml:"account"`
	// Query region.
	Region *string `json:"region" yaml:"region"`
	// Query service name.
	ServiceName *string `json:"serviceName" yaml:"serviceName"`
}

// Models an AWS execution environment, for use within the CDK toolkit.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cx_api "github.com/aws/aws-cdk-go/awscdk/cx_api"
//   environment := &environment{
//   	account: jsii.String("account"),
//   	name: jsii.String("name"),
//   	region: jsii.String("region"),
//   }
//
type Environment struct {
	// The AWS account this environment deploys into.
	Account *string `json:"account" yaml:"account"`
	// The arbitrary name of this environment (user-set, or at least user-meaningful).
	Name *string `json:"name" yaml:"name"`
	// The AWS region name where this environment deploys into.
	Region *string `json:"region" yaml:"region"`
}

// Return the appropriate values for the environment placeholders.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cx_api "github.com/aws/aws-cdk-go/awscdk/cx_api"
//   environmentPlaceholderValues := &environmentPlaceholderValues{
//   	accountId: jsii.String("accountId"),
//   	partition: jsii.String("partition"),
//   	region: jsii.String("region"),
//   }
//
type EnvironmentPlaceholderValues struct {
	// Return the account.
	AccountId *string `json:"accountId" yaml:"accountId"`
	// Return the partition.
	Partition *string `json:"partition" yaml:"partition"`
	// Return the region.
	Region *string `json:"region" yaml:"region"`
}

// Placeholders which can be used manifests.
//
// These can occur both in the Asset Manifest as well as the general
// Cloud Assembly manifest.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cx_api "github.com/aws/aws-cdk-go/awscdk/cx_api"
//   environmentPlaceholders := cx_api.NewEnvironmentPlaceholders()
//
type EnvironmentPlaceholders interface {
}

// The jsii proxy struct for EnvironmentPlaceholders
type jsiiProxy_EnvironmentPlaceholders struct {
	_ byte // padding
}

func NewEnvironmentPlaceholders() EnvironmentPlaceholders {
	_init_.Initialize()

	j := jsiiProxy_EnvironmentPlaceholders{}

	_jsii_.Create(
		"aws-cdk-lib.cx_api.EnvironmentPlaceholders",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewEnvironmentPlaceholders_Override(e EnvironmentPlaceholders) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.cx_api.EnvironmentPlaceholders",
		nil, // no parameters
		e,
	)
}

// Replace the environment placeholders in all strings found in a complex object.
//
// Duplicated between cdk-assets and aws-cdk CLI because we don't have a good single place to put it
// (they're nominally independent tools).
func EnvironmentPlaceholders_Replace(object interface{}, values *EnvironmentPlaceholderValues) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"aws-cdk-lib.cx_api.EnvironmentPlaceholders",
		"replace",
		[]interface{}{object, values},
		&returns,
	)

	return returns
}

// Like 'replace', but asynchronous.
func EnvironmentPlaceholders_ReplaceAsync(object interface{}, provider IEnvironmentPlaceholderProvider) interface{} {
	_init_.Initialize()

	var returns interface{}

	_jsii_.StaticInvoke(
		"aws-cdk-lib.cx_api.EnvironmentPlaceholders",
		"replaceAsync",
		[]interface{}{object, provider},
		&returns,
	)

	return returns
}

func EnvironmentPlaceholders_CURRENT_ACCOUNT() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.cx_api.EnvironmentPlaceholders",
		"CURRENT_ACCOUNT",
		&returns,
	)
	return returns
}

func EnvironmentPlaceholders_CURRENT_PARTITION() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.cx_api.EnvironmentPlaceholders",
		"CURRENT_PARTITION",
		&returns,
	)
	return returns
}

func EnvironmentPlaceholders_CURRENT_REGION() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.cx_api.EnvironmentPlaceholders",
		"CURRENT_REGION",
		&returns,
	)
	return returns
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cx_api "github.com/aws/aws-cdk-go/awscdk/cx_api"
//   environmentUtils := cx_api.NewEnvironmentUtils()
//
type EnvironmentUtils interface {
}

// The jsii proxy struct for EnvironmentUtils
type jsiiProxy_EnvironmentUtils struct {
	_ byte // padding
}

func NewEnvironmentUtils() EnvironmentUtils {
	_init_.Initialize()

	j := jsiiProxy_EnvironmentUtils{}

	_jsii_.Create(
		"aws-cdk-lib.cx_api.EnvironmentUtils",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewEnvironmentUtils_Override(e EnvironmentUtils) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.cx_api.EnvironmentUtils",
		nil, // no parameters
		e,
	)
}

// Format an environment string from an account and region.
func EnvironmentUtils_Format(account *string, region *string) *string {
	_init_.Initialize()

	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.cx_api.EnvironmentUtils",
		"format",
		[]interface{}{account, region},
		&returns,
	)

	return returns
}

// Build an environment object from an account and region.
func EnvironmentUtils_Make(account *string, region *string) *Environment {
	_init_.Initialize()

	var returns *Environment

	_jsii_.StaticInvoke(
		"aws-cdk-lib.cx_api.EnvironmentUtils",
		"make",
		[]interface{}{account, region},
		&returns,
	)

	return returns
}

func EnvironmentUtils_Parse(environment *string) *Environment {
	_init_.Initialize()

	var returns *Environment

	_jsii_.StaticInvoke(
		"aws-cdk-lib.cx_api.EnvironmentUtils",
		"parse",
		[]interface{}{environment},
		&returns,
	)

	return returns
}

// Return the appropriate values for the environment placeholders.
type IEnvironmentPlaceholderProvider interface {
	// Return the account.
	AccountId() *string
	// Return the partition.
	Partition() *string
	// Return the region.
	Region() *string
}

// The jsii proxy for IEnvironmentPlaceholderProvider
type jsiiProxy_IEnvironmentPlaceholderProvider struct {
	_ byte // padding
}

func (i *jsiiProxy_IEnvironmentPlaceholderProvider) AccountId() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"accountId",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEnvironmentPlaceholderProvider) Partition() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"partition",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IEnvironmentPlaceholderProvider) Region() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"region",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties of a discovered key.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cx_api "github.com/aws/aws-cdk-go/awscdk/cx_api"
//   keyContextResponse := &keyContextResponse{
//   	keyId: jsii.String("keyId"),
//   }
//
type KeyContextResponse struct {
	// Id of the key.
	KeyId *string `json:"keyId" yaml:"keyId"`
}

// Properties of a discovered load balancer.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cx_api "github.com/aws/aws-cdk-go/awscdk/cx_api"
//   loadBalancerContextResponse := &loadBalancerContextResponse{
//   	ipAddressType: cx_api.loadBalancerIpAddressType_IPV4,
//   	loadBalancerArn: jsii.String("loadBalancerArn"),
//   	loadBalancerCanonicalHostedZoneId: jsii.String("loadBalancerCanonicalHostedZoneId"),
//   	loadBalancerDnsName: jsii.String("loadBalancerDnsName"),
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	vpcId: jsii.String("vpcId"),
//   }
//
type LoadBalancerContextResponse struct {
	// Type of IP address.
	IpAddressType LoadBalancerIpAddressType `json:"ipAddressType" yaml:"ipAddressType"`
	// The ARN of the load balancer.
	LoadBalancerArn *string `json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// The hosted zone ID of the load balancer's name.
	LoadBalancerCanonicalHostedZoneId *string `json:"loadBalancerCanonicalHostedZoneId" yaml:"loadBalancerCanonicalHostedZoneId"`
	// Load balancer's DNS name.
	LoadBalancerDnsName *string `json:"loadBalancerDnsName" yaml:"loadBalancerDnsName"`
	// Load balancer's security groups.
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
	// Load balancer's VPC.
	VpcId *string `json:"vpcId" yaml:"vpcId"`
}

// Load balancer ip address type.
type LoadBalancerIpAddressType string

const (
	// IPV4 ip address.
	LoadBalancerIpAddressType_IPV4 LoadBalancerIpAddressType = "IPV4"
	// Dual stack address.
	LoadBalancerIpAddressType_DUAL_STACK LoadBalancerIpAddressType = "DUAL_STACK"
)

// Properties of a discovered load balancer listener.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cx_api "github.com/aws/aws-cdk-go/awscdk/cx_api"
//   loadBalancerListenerContextResponse := &loadBalancerListenerContextResponse{
//   	listenerArn: jsii.String("listenerArn"),
//   	listenerPort: jsii.Number(123),
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   }
//
type LoadBalancerListenerContextResponse struct {
	// The ARN of the listener.
	ListenerArn *string `json:"listenerArn" yaml:"listenerArn"`
	// The port the listener is listening on.
	ListenerPort *float64 `json:"listenerPort" yaml:"listenerPort"`
	// The security groups of the load balancer.
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cx_api "github.com/aws/aws-cdk-go/awscdk/cx_api"
//   metadataEntryResult := &metadataEntryResult{
//   	path: jsii.String("path"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	data: jsii.String("data"),
//   	trace: []*string{
//   		jsii.String("trace"),
//   	},
//   }
//
type MetadataEntryResult struct {
	// The type of the metadata entry.
	Type *string `json:"type" yaml:"type"`
	// The data.
	Data interface{} `json:"data" yaml:"data"`
	// A stack trace for when the entry was created.
	Trace *[]*string `json:"trace" yaml:"trace"`
	// The path in which this entry was defined.
	Path *string `json:"path" yaml:"path"`
}

// Asset manifest is a description of a set of assets which need to be built and published.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"import awscdk "github.com/aws/aws-cdk-go/awscdk"import cx_api "github.com/aws/aws-cdk-go/awscdk/cx_api"
//
//   var cloudAssembly cloudAssembly
//   nestedCloudAssemblyArtifact := cx_api.NewNestedCloudAssemblyArtifact(cloudAssembly, jsii.String("name"), &artifactManifest{
//   	type: cloud_assembly_schema.artifactType_NONE,
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
	var returns *[]*MetadataEntryResult

	_jsii_.Invoke(
		n,
		"findMetadataByType",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

// Properties of a discovered SecurityGroup.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cx_api "github.com/aws/aws-cdk-go/awscdk/cx_api"
//   securityGroupContextResponse := &securityGroupContextResponse{
//   	allowAllOutbound: jsii.Boolean(false),
//   	securityGroupId: jsii.String("securityGroupId"),
//   }
//
type SecurityGroupContextResponse struct {
	// Whether the security group allows all outbound traffic.
	//
	// This will be true
	// when the security group has all-protocol egress permissions to access both
	// `0.0.0.0/0` and `::/0`.
	AllowAllOutbound *bool `json:"allowAllOutbound" yaml:"allowAllOutbound"`
	// The security group's id.
	SecurityGroupId *string `json:"securityGroupId" yaml:"securityGroupId"`
}

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cx_api "github.com/aws/aws-cdk-go/awscdk/cx_api"
//   synthesisMessage := &synthesisMessage{
//   	entry: &metadataEntry{
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		data: jsii.String("data"),
//   		trace: []*string{
//   			jsii.String("trace"),
//   		},
//   	},
//   	id: jsii.String("id"),
//   	level: cx_api.synthesisMessageLevel_INFO,
//   }
//
type SynthesisMessage struct {
	Entry *cloudassemblyschema.MetadataEntry `json:"entry" yaml:"entry"`
	Id *string `json:"id" yaml:"id"`
	Level SynthesisMessageLevel `json:"level" yaml:"level"`
}

type SynthesisMessageLevel string

const (
	SynthesisMessageLevel_INFO SynthesisMessageLevel = "INFO"
	SynthesisMessageLevel_WARNING SynthesisMessageLevel = "WARNING"
	SynthesisMessageLevel_ERROR SynthesisMessageLevel = "ERROR"
)

// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloud_assembly_schema "github.com/aws/aws-cdk-go/awscdk/cloud_assembly_schema"import awscdk "github.com/aws/aws-cdk-go/awscdk"import cx_api "github.com/aws/aws-cdk-go/awscdk/cx_api"
//
//   var cloudAssembly cloudAssembly
//   treeCloudArtifact := cx_api.NewTreeCloudArtifact(cloudAssembly, jsii.String("name"), &artifactManifest{
//   	type: cloud_assembly_schema.artifactType_NONE,
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

	var returns CloudArtifact

	_jsii_.StaticInvoke(
		"aws-cdk-lib.cx_api.TreeCloudArtifact",
		"fromManifest",
		[]interface{}{assembly, id, artifact},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TreeCloudArtifact) FindMetadataByType(type_ *string) *[]*MetadataEntryResult {
	var returns *[]*MetadataEntryResult

	_jsii_.Invoke(
		t,
		"findMetadataByType",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

// Properties of a discovered VPC.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cx_api "github.com/aws/aws-cdk-go/awscdk/cx_api"
//   vpcContextResponse := &vpcContextResponse{
//   	availabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	vpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	isolatedSubnetIds: []*string{
//   		jsii.String("isolatedSubnetIds"),
//   	},
//   	isolatedSubnetNames: []*string{
//   		jsii.String("isolatedSubnetNames"),
//   	},
//   	isolatedSubnetRouteTableIds: []*string{
//   		jsii.String("isolatedSubnetRouteTableIds"),
//   	},
//   	privateSubnetIds: []*string{
//   		jsii.String("privateSubnetIds"),
//   	},
//   	privateSubnetNames: []*string{
//   		jsii.String("privateSubnetNames"),
//   	},
//   	privateSubnetRouteTableIds: []*string{
//   		jsii.String("privateSubnetRouteTableIds"),
//   	},
//   	publicSubnetIds: []*string{
//   		jsii.String("publicSubnetIds"),
//   	},
//   	publicSubnetNames: []*string{
//   		jsii.String("publicSubnetNames"),
//   	},
//   	publicSubnetRouteTableIds: []*string{
//   		jsii.String("publicSubnetRouteTableIds"),
//   	},
//   	subnetGroups: []vpcSubnetGroup{
//   		&vpcSubnetGroup{
//   			name: jsii.String("name"),
//   			subnets: []vpcSubnet{
//   				&vpcSubnet{
//   					availabilityZone: jsii.String("availabilityZone"),
//   					routeTableId: jsii.String("routeTableId"),
//   					subnetId: jsii.String("subnetId"),
//
//   					// the properties below are optional
//   					cidr: jsii.String("cidr"),
//   				},
//   			},
//   			type: cx_api.vpcSubnetGroupType_PUBLIC,
//   		},
//   	},
//   	vpcCidrBlock: jsii.String("vpcCidrBlock"),
//   	vpnGatewayId: jsii.String("vpnGatewayId"),
//   }
//
type VpcContextResponse struct {
	// AZs.
	AvailabilityZones *[]*string `json:"availabilityZones" yaml:"availabilityZones"`
	// VPC id.
	VpcId *string `json:"vpcId" yaml:"vpcId"`
	// IDs of all isolated subnets.
	//
	// Element count: #(availabilityZones) · #(isolatedGroups).
	IsolatedSubnetIds *[]*string `json:"isolatedSubnetIds" yaml:"isolatedSubnetIds"`
	// Name of isolated subnet groups.
	//
	// Element count: #(isolatedGroups).
	IsolatedSubnetNames *[]*string `json:"isolatedSubnetNames" yaml:"isolatedSubnetNames"`
	// Route Table IDs of isolated subnet groups.
	//
	// Element count: #(availabilityZones) · #(isolatedGroups).
	IsolatedSubnetRouteTableIds *[]*string `json:"isolatedSubnetRouteTableIds" yaml:"isolatedSubnetRouteTableIds"`
	// IDs of all private subnets.
	//
	// Element count: #(availabilityZones) · #(privateGroups).
	PrivateSubnetIds *[]*string `json:"privateSubnetIds" yaml:"privateSubnetIds"`
	// Name of private subnet groups.
	//
	// Element count: #(privateGroups).
	PrivateSubnetNames *[]*string `json:"privateSubnetNames" yaml:"privateSubnetNames"`
	// Route Table IDs of private subnet groups.
	//
	// Element count: #(availabilityZones) · #(privateGroups).
	PrivateSubnetRouteTableIds *[]*string `json:"privateSubnetRouteTableIds" yaml:"privateSubnetRouteTableIds"`
	// IDs of all public subnets.
	//
	// Element count: #(availabilityZones) · #(publicGroups).
	PublicSubnetIds *[]*string `json:"publicSubnetIds" yaml:"publicSubnetIds"`
	// Name of public subnet groups.
	//
	// Element count: #(publicGroups).
	PublicSubnetNames *[]*string `json:"publicSubnetNames" yaml:"publicSubnetNames"`
	// Route Table IDs of public subnet groups.
	//
	// Element count: #(availabilityZones) · #(publicGroups).
	PublicSubnetRouteTableIds *[]*string `json:"publicSubnetRouteTableIds" yaml:"publicSubnetRouteTableIds"`
	// The subnet groups discovered for the given VPC.
	//
	// Unlike the above properties, this will include asymmetric subnets,
	// if the VPC has any.
	// This property will only be populated if {@link VpcContextQuery.returnAsymmetricSubnets}
	// is true.
	SubnetGroups *[]*VpcSubnetGroup `json:"subnetGroups" yaml:"subnetGroups"`
	// VPC cidr.
	VpcCidrBlock *string `json:"vpcCidrBlock" yaml:"vpcCidrBlock"`
	// The VPN gateway ID.
	VpnGatewayId *string `json:"vpnGatewayId" yaml:"vpnGatewayId"`
}

// A subnet representation that the VPC provider uses.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cx_api "github.com/aws/aws-cdk-go/awscdk/cx_api"
//   vpcSubnet := &vpcSubnet{
//   	availabilityZone: jsii.String("availabilityZone"),
//   	routeTableId: jsii.String("routeTableId"),
//   	subnetId: jsii.String("subnetId"),
//
//   	// the properties below are optional
//   	cidr: jsii.String("cidr"),
//   }
//
type VpcSubnet struct {
	// The code of the availability zone this subnet is in (for example, 'us-west-2a').
	AvailabilityZone *string `json:"availabilityZone" yaml:"availabilityZone"`
	// The identifier of the route table for this subnet.
	RouteTableId *string `json:"routeTableId" yaml:"routeTableId"`
	// The identifier of the subnet.
	SubnetId *string `json:"subnetId" yaml:"subnetId"`
	// CIDR range of the subnet.
	Cidr *string `json:"cidr" yaml:"cidr"`
}

// A group of subnets returned by the VPC provider.
//
// The included subnets do NOT have to be symmetric!
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import cx_api "github.com/aws/aws-cdk-go/awscdk/cx_api"
//   vpcSubnetGroup := &vpcSubnetGroup{
//   	name: jsii.String("name"),
//   	subnets: []vpcSubnet{
//   		&vpcSubnet{
//   			availabilityZone: jsii.String("availabilityZone"),
//   			routeTableId: jsii.String("routeTableId"),
//   			subnetId: jsii.String("subnetId"),
//
//   			// the properties below are optional
//   			cidr: jsii.String("cidr"),
//   		},
//   	},
//   	type: cx_api.vpcSubnetGroupType_PUBLIC,
//   }
//
type VpcSubnetGroup struct {
	// The name of the subnet group, determined by looking at the tags of of the subnets that belong to it.
	Name *string `json:"name" yaml:"name"`
	// The subnets that are part of this group.
	//
	// There is no condition that the subnets have to be symmetric
	// in the group.
	Subnets *[]*VpcSubnet `json:"subnets" yaml:"subnets"`
	// The type of the subnet group.
	Type VpcSubnetGroupType `json:"type" yaml:"type"`
}

// The type of subnet group.
//
// Same as SubnetType in the @aws-cdk/aws-ec2 package,
// but we can't use that because of cyclical dependencies.
type VpcSubnetGroupType string

const (
	// Public subnet group type.
	VpcSubnetGroupType_PUBLIC VpcSubnetGroupType = "PUBLIC"
	// Private subnet group type.
	VpcSubnetGroupType_PRIVATE VpcSubnetGroupType = "PRIVATE"
	// Isolated subnet group type.
	VpcSubnetGroupType_ISOLATED VpcSubnetGroupType = "ISOLATED"
)

