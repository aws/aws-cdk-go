package cxapi

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/cloudassemblyschema"
)

// Can be used to build a cloud assembly.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cloudAssemblyBuilder_ cloudAssemblyBuilder
//
//   cloudAssemblyBuilder := awscdk.Cx_api.NewCloudAssemblyBuilder(jsii.String("outdir"), &CloudAssemblyBuilderProps{
//   	AssetOutdir: jsii.String("assetOutdir"),
//   	ParentBuilder: cloudAssemblyBuilder_,
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
	// Delete the cloud assembly directory.
	Delete()
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

	if err := validateNewCloudAssemblyBuilderParameters(props); err != nil {
		panic(err)
	}
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
	if err := c.validateAddArtifactParameters(id, manifest); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addArtifact",
		[]interface{}{id, manifest},
	)
}

func (c *jsiiProxy_CloudAssemblyBuilder) AddMissing(missing *cloudassemblyschema.MissingContext) {
	if err := c.validateAddMissingParameters(missing); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMissing",
		[]interface{}{missing},
	)
}

func (c *jsiiProxy_CloudAssemblyBuilder) BuildAssembly(options *AssemblyBuildOptions) CloudAssembly {
	if err := c.validateBuildAssemblyParameters(options); err != nil {
		panic(err)
	}
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
	if err := c.validateCreateNestedAssemblyParameters(artifactId, displayName); err != nil {
		panic(err)
	}
	var returns CloudAssemblyBuilder

	_jsii_.Invoke(
		c,
		"createNestedAssembly",
		[]interface{}{artifactId, displayName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAssemblyBuilder) Delete() {
	_jsii_.InvokeVoid(
		c,
		"delete",
		nil, // no parameters
	)
}

