package appstagingsynthesizeralpha

import (
	_init_ "github.com/aws/aws-cdk-go/appstagingsynthesizeralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/appstagingsynthesizeralpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/cloudassemblyschema"
)

// App Staging Synthesizer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import app_staging_synthesizer_alpha "github.com/aws/aws-cdk-go/appstagingsynthesizeralpha"
//
//   var deploymentIdentities DeploymentIdentities
//   var stagingResourcesFactory IStagingResourcesFactory
//
//   appStagingSynthesizer := app_staging_synthesizer_alpha.AppStagingSynthesizer_CustomFactory(&CustomFactoryOptions{
//   	Factory: stagingResourcesFactory,
//
//   	// the properties below are optional
//   	BootstrapQualifier: jsii.String("bootstrapQualifier"),
//   	DeploymentIdentities: deploymentIdentities,
//   	OncePerEnv: jsii.Boolean(false),
//   })
//
// Experimental.
type AppStagingSynthesizer interface {
	awscdk.StackSynthesizer
	awscdk.IReusableStackSynthesizer
	// The qualifier used to bootstrap this stack.
	// Experimental.
	BootstrapQualifier() *string
	// Retrieve the bound stack.
	//
	// Fails if the stack hasn't been bound yet.
	// Experimental.
	BoundStack() awscdk.Stack
	// The role used to lookup for this stack.
	// Experimental.
	LookupRole() *string
	// Add a CfnRule to the bound stack that checks whether an SSM parameter exceeds a given version.
	//
	// This will modify the template, so must be called before the stack is synthesized.
	// Experimental.
	AddBootstrapVersionRule(requiredVersion *float64, bootstrapStackVersionSsmParameter *string)
	// Implemented for legacy purposes;
	//
	// this will never be called.
	// Experimental.
	AddDockerImageAsset(_asset *awscdk.DockerImageAssetSource) *awscdk.DockerImageAssetLocation
	// Implemented for legacy purposes;
	//
	// this will never be called.
	// Experimental.
	AddFileAsset(_asset *awscdk.FileAssetSource) *awscdk.FileAssetLocation
	// Implemented for legacy purposes;
	//
	// this will never be called.
	// Experimental.
	Bind(_stack awscdk.Stack)
	// Turn a docker asset location into a CloudFormation representation of that location.
	//
	// If any of the fields contain placeholders, the result will be wrapped in a `Fn.sub`.
	// Experimental.
	CloudFormationLocationFromDockerImageAsset(dest *cloudassemblyschema.DockerImageDestination) *awscdk.DockerImageAssetLocation
	// Turn a file asset location into a CloudFormation representation of that location.
	//
	// If any of the fields contain placeholders, the result will be wrapped in a `Fn.sub`.
	// Experimental.
	CloudFormationLocationFromFileAsset(location *cloudassemblyschema.FileDestination) *awscdk.FileAssetLocation
	// Write the CloudFormation stack artifact to the session.
	//
	// Use default settings to add a CloudFormationStackArtifact artifact to
	// the given synthesis session. The Stack artifact will control the settings for the
	// CloudFormation deployment.
	// Experimental.
	EmitArtifact(session awscdk.ISynthesisSession, options *awscdk.SynthesizeStackArtifactOptions)
	// Write the stack artifact to the session.
	//
	// Use default settings to add a CloudFormationStackArtifact artifact to
	// the given synthesis session.
	// Deprecated: Use `emitArtifact` instead.
	EmitStackArtifact(stack awscdk.Stack, session awscdk.ISynthesisSession, options *awscdk.SynthesizeStackArtifactOptions)
	// Returns a version of the synthesizer bound to a stack.
	// Experimental.
	ReusableBind(stack awscdk.Stack) awscdk.IBoundStackSynthesizer
	// Implemented for legacy purposes;
	//
	// this will never be called.
	// Experimental.
	Synthesize(_session awscdk.ISynthesisSession)
	// Have the stack write out its template.
	// Deprecated: Use `synthesizeTemplate` instead.
	SynthesizeStackTemplate(stack awscdk.Stack, session awscdk.ISynthesisSession)
	// Write the stack template to the given session.
	//
	// Return a descriptor that represents the stack template as a file asset
	// source, for adding to an asset manifest (if desired). This can be used to
	// have the asset manifest system (`cdk-assets`) upload the template to S3
	// using the appropriate role, so that afterwards only a CloudFormation
	// deployment is necessary.
	//
	// If the template is uploaded as an asset, the `stackTemplateAssetObjectUrl`
	// property should be set when calling `emitArtifact.`
	//
	// If the template is *NOT* uploaded as an asset first and the template turns
	// out to be >50KB, it will need to be uploaded to S3 anyway. At that point
	// the credentials will be the same identity that is doing the `UpdateStack`
	// call, which may not have the right permissions to write to S3.
	// Experimental.
	SynthesizeTemplate(session awscdk.ISynthesisSession, lookupRoleArn *string, lookupRoleExternalId *string, lookupRoleAdditionalOptions *map[string]interface{}) *awscdk.FileAssetSource
}

// The jsii proxy struct for AppStagingSynthesizer
type jsiiProxy_AppStagingSynthesizer struct {
	internal.Type__awscdkStackSynthesizer
	internal.Type__awscdkIReusableStackSynthesizer
}

func (j *jsiiProxy_AppStagingSynthesizer) BootstrapQualifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapQualifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppStagingSynthesizer) BoundStack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"boundStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppStagingSynthesizer) LookupRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lookupRole",
		&returns,
	)
	return returns
}


// Supply your own stagingStackFactory method for creating an IStagingStack when a stack is bound to the synthesizer.
//
// By default, `oncePerEnv = true`, which means that a new instance of the IStagingStack
// will be created in new environments. Set `oncePerEnv = false` to turn off that behavior.
// Experimental.
func AppStagingSynthesizer_CustomFactory(options *CustomFactoryOptions) AppStagingSynthesizer {
	_init_.Initialize()

	if err := validateAppStagingSynthesizer_CustomFactoryParameters(options); err != nil {
		panic(err)
	}
	var returns AppStagingSynthesizer

	_jsii_.StaticInvoke(
		"@aws-cdk/app-staging-synthesizer-alpha.AppStagingSynthesizer",
		"customFactory",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Use these exact staging resources for every stack that this synthesizer is used for.
// Experimental.
func AppStagingSynthesizer_CustomResources(options *CustomResourcesOptions) AppStagingSynthesizer {
	_init_.Initialize()

	if err := validateAppStagingSynthesizer_CustomResourcesParameters(options); err != nil {
		panic(err)
	}
	var returns AppStagingSynthesizer

	_jsii_.StaticInvoke(
		"@aws-cdk/app-staging-synthesizer-alpha.AppStagingSynthesizer",
		"customResources",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Use the Default Staging Resources, creating a single stack per environment this app is deployed in.
// Experimental.
func AppStagingSynthesizer_DefaultResources(options *DefaultResourcesOptions) AppStagingSynthesizer {
	_init_.Initialize()

	if err := validateAppStagingSynthesizer_DefaultResourcesParameters(options); err != nil {
		panic(err)
	}
	var returns AppStagingSynthesizer

	_jsii_.StaticInvoke(
		"@aws-cdk/app-staging-synthesizer-alpha.AppStagingSynthesizer",
		"defaultResources",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func AppStagingSynthesizer_DEFAULT_CLOUDFORMATION_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/app-staging-synthesizer-alpha.AppStagingSynthesizer",
		"DEFAULT_CLOUDFORMATION_ROLE_ARN",
		&returns,
	)
	return returns
}

func AppStagingSynthesizer_DEFAULT_DEPLOY_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/app-staging-synthesizer-alpha.AppStagingSynthesizer",
		"DEFAULT_DEPLOY_ROLE_ARN",
		&returns,
	)
	return returns
}

func AppStagingSynthesizer_DEFAULT_LOOKUP_ROLE_ARN() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/app-staging-synthesizer-alpha.AppStagingSynthesizer",
		"DEFAULT_LOOKUP_ROLE_ARN",
		&returns,
	)
	return returns
}

func AppStagingSynthesizer_DEFAULT_QUALIFIER() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/app-staging-synthesizer-alpha.AppStagingSynthesizer",
		"DEFAULT_QUALIFIER",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AppStagingSynthesizer) AddBootstrapVersionRule(requiredVersion *float64, bootstrapStackVersionSsmParameter *string) {
	if err := a.validateAddBootstrapVersionRuleParameters(requiredVersion, bootstrapStackVersionSsmParameter); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addBootstrapVersionRule",
		[]interface{}{requiredVersion, bootstrapStackVersionSsmParameter},
	)
}

func (a *jsiiProxy_AppStagingSynthesizer) AddDockerImageAsset(_asset *awscdk.DockerImageAssetSource) *awscdk.DockerImageAssetLocation {
	if err := a.validateAddDockerImageAssetParameters(_asset); err != nil {
		panic(err)
	}
	var returns *awscdk.DockerImageAssetLocation

	_jsii_.Invoke(
		a,
		"addDockerImageAsset",
		[]interface{}{_asset},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppStagingSynthesizer) AddFileAsset(_asset *awscdk.FileAssetSource) *awscdk.FileAssetLocation {
	if err := a.validateAddFileAssetParameters(_asset); err != nil {
		panic(err)
	}
	var returns *awscdk.FileAssetLocation

	_jsii_.Invoke(
		a,
		"addFileAsset",
		[]interface{}{_asset},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppStagingSynthesizer) Bind(_stack awscdk.Stack) {
	if err := a.validateBindParameters(_stack); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"bind",
		[]interface{}{_stack},
	)
}

func (a *jsiiProxy_AppStagingSynthesizer) CloudFormationLocationFromDockerImageAsset(dest *cloudassemblyschema.DockerImageDestination) *awscdk.DockerImageAssetLocation {
	if err := a.validateCloudFormationLocationFromDockerImageAssetParameters(dest); err != nil {
		panic(err)
	}
	var returns *awscdk.DockerImageAssetLocation

	_jsii_.Invoke(
		a,
		"cloudFormationLocationFromDockerImageAsset",
		[]interface{}{dest},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppStagingSynthesizer) CloudFormationLocationFromFileAsset(location *cloudassemblyschema.FileDestination) *awscdk.FileAssetLocation {
	if err := a.validateCloudFormationLocationFromFileAssetParameters(location); err != nil {
		panic(err)
	}
	var returns *awscdk.FileAssetLocation

	_jsii_.Invoke(
		a,
		"cloudFormationLocationFromFileAsset",
		[]interface{}{location},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppStagingSynthesizer) EmitArtifact(session awscdk.ISynthesisSession, options *awscdk.SynthesizeStackArtifactOptions) {
	if err := a.validateEmitArtifactParameters(session, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"emitArtifact",
		[]interface{}{session, options},
	)
}

func (a *jsiiProxy_AppStagingSynthesizer) EmitStackArtifact(stack awscdk.Stack, session awscdk.ISynthesisSession, options *awscdk.SynthesizeStackArtifactOptions) {
	if err := a.validateEmitStackArtifactParameters(stack, session, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"emitStackArtifact",
		[]interface{}{stack, session, options},
	)
}

func (a *jsiiProxy_AppStagingSynthesizer) ReusableBind(stack awscdk.Stack) awscdk.IBoundStackSynthesizer {
	if err := a.validateReusableBindParameters(stack); err != nil {
		panic(err)
	}
	var returns awscdk.IBoundStackSynthesizer

	_jsii_.Invoke(
		a,
		"reusableBind",
		[]interface{}{stack},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppStagingSynthesizer) Synthesize(_session awscdk.ISynthesisSession) {
	if err := a.validateSynthesizeParameters(_session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		[]interface{}{_session},
	)
}

func (a *jsiiProxy_AppStagingSynthesizer) SynthesizeStackTemplate(stack awscdk.Stack, session awscdk.ISynthesisSession) {
	if err := a.validateSynthesizeStackTemplateParameters(stack, session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"synthesizeStackTemplate",
		[]interface{}{stack, session},
	)
}

func (a *jsiiProxy_AppStagingSynthesizer) SynthesizeTemplate(session awscdk.ISynthesisSession, lookupRoleArn *string, lookupRoleExternalId *string, lookupRoleAdditionalOptions *map[string]interface{}) *awscdk.FileAssetSource {
	if err := a.validateSynthesizeTemplateParameters(session); err != nil {
		panic(err)
	}
	var returns *awscdk.FileAssetSource

	_jsii_.Invoke(
		a,
		"synthesizeTemplate",
		[]interface{}{session, lookupRoleArn, lookupRoleExternalId, lookupRoleAdditionalOptions},
		&returns,
	)

	return returns
}

