package cxapi

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/cloudassemblyschema"
)

// TODO: EXAMPLE
//
type AssemblyBuildOptions struct {
}

// Asset manifest is a description of a set of assets which need to be built and published.
//
// TODO: EXAMPLE
//
type AssetManifestArtifact interface {
	CloudArtifact
	Assembly() CloudAssembly
	BootstrapStackVersionSsmParameter() *string
	Dependencies() *[]CloudArtifact
	File() *string
	HierarchicalId() *string
	Id() *string
	Manifest() *cloudassemblyschema.ArtifactManifest
	Messages() *[]*SynthesisMessage
	RequiresBootstrapStackVersion() *float64
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

// Returns: all the metadata entries of a specific type in this artifact.
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
// TODO: EXAMPLE
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
// TODO: EXAMPLE
//
type CloudArtifact interface {
	Assembly() CloudAssembly
	Dependencies() *[]CloudArtifact
	HierarchicalId() *string
	Id() *string
	Manifest() *cloudassemblyschema.ArtifactManifest
	Messages() *[]*SynthesisMessage
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

// Returns: all the metadata entries of a specific type in this artifact.
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
// TODO: EXAMPLE
//
type CloudAssembly interface {
	Artifacts() *[]CloudArtifact
	Directory() *string
	Manifest() *cloudassemblyschema.AssemblyManifest
	NestedAssemblies() *[]NestedCloudAssemblyArtifact
	Runtime() *cloudassemblyschema.RuntimeInfo
	Stacks() *[]CloudFormationStackArtifact
	StacksRecursively() *[]CloudFormationStackArtifact
	Version() *string
	GetNestedAssembly(artifactId *string) CloudAssembly
	GetNestedAssemblyArtifact(artifactId *string) NestedCloudAssemblyArtifact
	GetStackArtifact(artifactId *string) CloudFormationStackArtifact
	GetStackByName(stackName *string) CloudFormationStackArtifact
	Tree() TreeCloudArtifact
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

// Returns a nested assembly.
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

// Returns a nested assembly artifact.
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

// Returns a CloudFormation stack artifact from this assembly.
//
// Returns: a `CloudFormationStackArtifact` object.
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

// Returns a CloudFormation stack artifact from this assembly.
//
// Will only search the current assembly.
//
// Returns: a `CloudFormationStackArtifact` object.
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

// Returns the tree metadata artifact from this assembly.
//
// Returns: a `TreeCloudArtifact` object if there is one defined in the manifest, `undefined` otherwise.
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

// Attempts to find an artifact with a specific identity.
//
// Returns: A `CloudArtifact` object or `undefined` if the artifact does not exist in this assembly.
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
// TODO: EXAMPLE
//
type CloudAssemblyBuilder interface {
	AssetOutdir() *string
	Outdir() *string
	AddArtifact(id *string, manifest *cloudassemblyschema.ArtifactManifest)
	AddMissing(missing *cloudassemblyschema.MissingContext)
	BuildAssembly(options *AssemblyBuildOptions) CloudAssembly
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

// Adds an artifact into the cloud assembly.
func (c *jsiiProxy_CloudAssemblyBuilder) AddArtifact(id *string, manifest *cloudassemblyschema.ArtifactManifest) {
	_jsii_.InvokeVoid(
		c,
		"addArtifact",
		[]interface{}{id, manifest},
	)
}

// Reports that some context is missing in order for this cloud assembly to be fully synthesized.
func (c *jsiiProxy_CloudAssemblyBuilder) AddMissing(missing *cloudassemblyschema.MissingContext) {
	_jsii_.InvokeVoid(
		c,
		"addMissing",
		[]interface{}{missing},
	)
}

// Finalizes the cloud assembly into the output directory returns a `CloudAssembly` object that can be used to inspect the assembly.
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

// Creates a nested cloud assembly.
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
// TODO: EXAMPLE
//
type CloudAssemblyBuilderProps struct {
	// Use the given asset output directory.
	AssetOutdir *string `json:"assetOutdir" yaml:"assetOutdir"`
	// If this builder is for a nested assembly, the parent assembly builder.
	ParentBuilder CloudAssemblyBuilder `json:"parentBuilder" yaml:"parentBuilder"`
}

// TODO: EXAMPLE
//
type CloudFormationStackArtifact interface {
	CloudArtifact
	Assembly() CloudAssembly
	Assets() *[]interface{}
	AssumeRoleArn() *string
	AssumeRoleExternalId() *string
	BootstrapStackVersionSsmParameter() *string
	CloudFormationExecutionRoleArn() *string
	Dependencies() *[]CloudArtifact
	DisplayName() *string
	Environment() *Environment
	HierarchicalId() *string
	Id() *string
	LookupRole() *cloudassemblyschema.BootstrapRole
	Manifest() *cloudassemblyschema.ArtifactManifest
	Messages() *[]*SynthesisMessage
	OriginalName() *string
	Parameters() *map[string]*string
	RequiresBootstrapStackVersion() *float64
	StackName() *string
	StackTemplateAssetObjectUrl() *string
	Tags() *map[string]*string
	Template() interface{}
	TemplateFile() *string
	TemplateFullPath() *string
	TerminationProtection() *bool
	ValidateOnSynth() *bool
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

// Returns: all the metadata entries of a specific type in this artifact.
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
// TODO: EXAMPLE
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
// TODO: EXAMPLE
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
// TODO: EXAMPLE
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
// TODO: EXAMPLE
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

// TODO: EXAMPLE
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
// TODO: EXAMPLE
//
type KeyContextResponse struct {
	// Id of the key.
	KeyId *string `json:"keyId" yaml:"keyId"`
}

// Properties of a discovered load balancer.
//
// TODO: EXAMPLE
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
	LoadBalancerIpAddressType_IPV4 LoadBalancerIpAddressType = "IPV4"
	LoadBalancerIpAddressType_DUAL_STACK LoadBalancerIpAddressType = "DUAL_STACK"
)

// Properties of a discovered load balancer listener.
//
// TODO: EXAMPLE
//
type LoadBalancerListenerContextResponse struct {
	// The ARN of the listener.
	ListenerArn *string `json:"listenerArn" yaml:"listenerArn"`
	// The port the listener is listening on.
	ListenerPort *float64 `json:"listenerPort" yaml:"listenerPort"`
	// The security groups of the load balancer.
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
}

// TODO: EXAMPLE
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
// TODO: EXAMPLE
//
type NestedCloudAssemblyArtifact interface {
	CloudArtifact
	Assembly() CloudAssembly
	Dependencies() *[]CloudArtifact
	DirectoryName() *string
	DisplayName() *string
	FullPath() *string
	HierarchicalId() *string
	Id() *string
	Manifest() *cloudassemblyschema.ArtifactManifest
	Messages() *[]*SynthesisMessage
	NestedAssembly() CloudAssembly
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

// Returns: all the metadata entries of a specific type in this artifact.
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
// TODO: EXAMPLE
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

// TODO: EXAMPLE
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

// TODO: EXAMPLE
//
type TreeCloudArtifact interface {
	CloudArtifact
	Assembly() CloudAssembly
	Dependencies() *[]CloudArtifact
	File() *string
	HierarchicalId() *string
	Id() *string
	Manifest() *cloudassemblyschema.ArtifactManifest
	Messages() *[]*SynthesisMessage
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

// Returns: all the metadata entries of a specific type in this artifact.
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
// TODO: EXAMPLE
//
type VpcContextResponse struct {
	// AZs.
	AvailabilityZones *[]*string `json:"availabilityZones" yaml:"availabilityZones"`
	// VPC id.
	VpcId *string `json:"vpcId" yaml:"vpcId"`
	// IDs of all isolated subnets.
	//
	// Element count: #(availabilityZones) · #(isolatedGroups)
	IsolatedSubnetIds *[]*string `json:"isolatedSubnetIds" yaml:"isolatedSubnetIds"`
	// Name of isolated subnet groups.
	//
	// Element count: #(isolatedGroups)
	IsolatedSubnetNames *[]*string `json:"isolatedSubnetNames" yaml:"isolatedSubnetNames"`
	// Route Table IDs of isolated subnet groups.
	//
	// Element count: #(availabilityZones) · #(isolatedGroups)
	IsolatedSubnetRouteTableIds *[]*string `json:"isolatedSubnetRouteTableIds" yaml:"isolatedSubnetRouteTableIds"`
	// IDs of all private subnets.
	//
	// Element count: #(availabilityZones) · #(privateGroups)
	PrivateSubnetIds *[]*string `json:"privateSubnetIds" yaml:"privateSubnetIds"`
	// Name of private subnet groups.
	//
	// Element count: #(privateGroups)
	PrivateSubnetNames *[]*string `json:"privateSubnetNames" yaml:"privateSubnetNames"`
	// Route Table IDs of private subnet groups.
	//
	// Element count: #(availabilityZones) · #(privateGroups)
	PrivateSubnetRouteTableIds *[]*string `json:"privateSubnetRouteTableIds" yaml:"privateSubnetRouteTableIds"`
	// IDs of all public subnets.
	//
	// Element count: #(availabilityZones) · #(publicGroups)
	PublicSubnetIds *[]*string `json:"publicSubnetIds" yaml:"publicSubnetIds"`
	// Name of public subnet groups.
	//
	// Element count: #(publicGroups)
	PublicSubnetNames *[]*string `json:"publicSubnetNames" yaml:"publicSubnetNames"`
	// Route Table IDs of public subnet groups.
	//
	// Element count: #(availabilityZones) · #(publicGroups)
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
// TODO: EXAMPLE
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
// TODO: EXAMPLE
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
	VpcSubnetGroupType_PUBLIC VpcSubnetGroupType = "PUBLIC"
	VpcSubnetGroupType_PRIVATE VpcSubnetGroupType = "PRIVATE"
	VpcSubnetGroupType_ISOLATED VpcSubnetGroupType = "ISOLATED"
)

