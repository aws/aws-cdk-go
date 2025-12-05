package cxapi

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/cloudassemblyschema"
	"github.com/aws/aws-cdk-go/awscdk/v2/cxapi/internal"
)

// Represents a deployable cloud application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudAssembly := awscdk.Cx_api.NewCloudAssembly(jsii.String("directory"), &LoadManifestOptions{
//   	SkipEnumCheck: jsii.Boolean(false),
//   	SkipVersionCheck: jsii.Boolean(false),
//   	TopoSort: jsii.Boolean(false),
//   })
//
type CloudAssembly interface {
	cloudassemblyschema.ICloudAssembly
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
	internal.Type__cloudassemblyschemaICloudAssembly
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

	if err := validateNewCloudAssemblyParameters(directory, loadOptions); err != nil {
		panic(err)
	}
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

// Cleans up any temporary assembly directories that got created in this process.
//
// If a Cloud Assembly is emitted to a temporary directory, its directory gets
// added to a list. This function iterates over that list and deletes each
// directory in it, to free up disk space.
//
// This function will normally be called automatically during Node process
// exit and so you don't need to call this. However, some test environments do
// not properly trigger Node's `exit` event. Notably: Jest does not trigger
// the `exit` event (<https://github.com/jestjs/jest/issues/10927>).
//
// ## Cleaning up temporary directories in jest
//
// For Jest, you have to make sure this function is called at the end of the
// test suite instead:
//
// ```js
// import { CloudAssembly } from 'aws-cdk-lib/cx-api';
//
// afterAll(CloudAssembly.cleanupTemporaryDirectories);
// ```
//
// Alternatively, you can use the `setupFilesAfterEnv` feature and use a
// provided helper script to automatically inject the above into every
// test file, so you don't have to do it by hand.
//
// ```
// $ npx jest --setupFilesAfterEnv aws-cdk-lib/testhelpers/jest-autoclean
// ```
//
// Or put the following into `jest.config.js`:
//
// ```js
// module.exports = {
//   // ...
//   setupFilesAfterEnv: ['aws-cdk-lib/testhelpers/jest-cleanup'],
// };
// ```.
func CloudAssembly_CleanupTemporaryDirectories() {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.cx_api.CloudAssembly",
		"cleanupTemporaryDirectories",
		nil, // no parameters
	)
}

// Return whether the given object is a CloudAssembly.
//
// We do attribute detection since we can't reliably use 'instanceof'.
func CloudAssembly_IsCloudAssembly(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudAssembly_IsCloudAssemblyParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.cx_api.CloudAssembly",
		"isCloudAssembly",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAssembly) GetNestedAssembly(artifactId *string) CloudAssembly {
	if err := c.validateGetNestedAssemblyParameters(artifactId); err != nil {
		panic(err)
	}
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
	if err := c.validateGetNestedAssemblyArtifactParameters(artifactId); err != nil {
		panic(err)
	}
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
	if err := c.validateGetStackArtifactParameters(artifactId); err != nil {
		panic(err)
	}
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
	if err := c.validateGetStackByNameParameters(stackName); err != nil {
		panic(err)
	}
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
	if err := c.validateTryGetArtifactParameters(id); err != nil {
		panic(err)
	}
	var returns CloudArtifact

	_jsii_.Invoke(
		c,
		"tryGetArtifact",
		[]interface{}{id},
		&returns,
	)

	return returns
}

