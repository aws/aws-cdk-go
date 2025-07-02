package awscdkapplicationsignalsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// Base class for .NET SDK injectors. Contains common .NET configuration settings used by both Windows and Linux implementations.
// Experimental.
type DotNetInjector interface {
	Injector
	// The command to run the init container.
	// Experimental.
	Command() *[]*string
	// The path to ADOT SDK agent in the init container.
	// Experimental.
	ContainerPath() *string
	// Experimental.
	InstrumentationVersion() InstrumentationVersion
	// Experimental.
	SetInstrumentationVersion(val InstrumentationVersion)
	// Experimental.
	SharedVolumeName() *string
	// Experimental.
	SetSharedVolumeName(val *string)
	// Inject additional environment variables to the application container other than the DEFAULT_ENVS.
	// Experimental.
	InjectAdditionalEnvironments(envsToInject *map[string]*string, envsFromTaskDef *map[string]*string)
	// Inject ADOT SDK agent init container.
	//
	// Returns: The created ContainerDefinition.
	// Experimental.
	InjectInitContainer(taskDefinition awsecs.TaskDefinition) awsecs.ContainerDefinition
	// Override environment variables in the application container.
	// Experimental.
	OverrideAdditionalEnvironments(envsToOverride *map[string]*string, envsFromTaskDef *map[string]*string)
	// Render the application container for SDK instrumentation.
	// Experimental.
	RenderDefaultContainer(taskDefinition awsecs.TaskDefinition)
}

// The jsii proxy struct for DotNetInjector
type jsiiProxy_DotNetInjector struct {
	jsiiProxy_Injector
}

func (j *jsiiProxy_DotNetInjector) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DotNetInjector) ContainerPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DotNetInjector) InstrumentationVersion() InstrumentationVersion {
	var returns InstrumentationVersion
	_jsii_.Get(
		j,
		"instrumentationVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DotNetInjector) SharedVolumeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedVolumeName",
		&returns,
	)
	return returns
}


// Experimental.
func NewDotNetInjector_Override(d DotNetInjector, sharedVolumeName *string, instrumentationVersion InstrumentationVersion, overrideEnvironments *[]*EnvironmentExtension) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.DotNetInjector",
		[]interface{}{sharedVolumeName, instrumentationVersion, overrideEnvironments},
		d,
	)
}

func (j *jsiiProxy_DotNetInjector)SetInstrumentationVersion(val InstrumentationVersion) {
	if err := j.validateSetInstrumentationVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instrumentationVersion",
		val,
	)
}

func (j *jsiiProxy_DotNetInjector)SetSharedVolumeName(val *string) {
	if err := j.validateSetSharedVolumeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedVolumeName",
		val,
	)
}

func DotNetInjector_DEFAULT_ENVS() *[]*EnvironmentExtension {
	_init_.Initialize()
	var returns *[]*EnvironmentExtension
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.DotNetInjector",
		"DEFAULT_ENVS",
		&returns,
	)
	return returns
}

func DotNetInjector_DOTNET_COMMON_ENVS() *[]*EnvironmentExtension {
	_init_.Initialize()
	var returns *[]*EnvironmentExtension
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.DotNetInjector",
		"DOTNET_COMMON_ENVS",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DotNetInjector) InjectAdditionalEnvironments(envsToInject *map[string]*string, envsFromTaskDef *map[string]*string) {
	if err := d.validateInjectAdditionalEnvironmentsParameters(envsToInject, envsFromTaskDef); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"injectAdditionalEnvironments",
		[]interface{}{envsToInject, envsFromTaskDef},
	)
}

func (d *jsiiProxy_DotNetInjector) InjectInitContainer(taskDefinition awsecs.TaskDefinition) awsecs.ContainerDefinition {
	if err := d.validateInjectInitContainerParameters(taskDefinition); err != nil {
		panic(err)
	}
	var returns awsecs.ContainerDefinition

	_jsii_.Invoke(
		d,
		"injectInitContainer",
		[]interface{}{taskDefinition},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DotNetInjector) OverrideAdditionalEnvironments(envsToOverride *map[string]*string, envsFromTaskDef *map[string]*string) {
	if err := d.validateOverrideAdditionalEnvironmentsParameters(envsToOverride, envsFromTaskDef); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideAdditionalEnvironments",
		[]interface{}{envsToOverride, envsFromTaskDef},
	)
}

func (d *jsiiProxy_DotNetInjector) RenderDefaultContainer(taskDefinition awsecs.TaskDefinition) {
	if err := d.validateRenderDefaultContainerParameters(taskDefinition); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"renderDefaultContainer",
		[]interface{}{taskDefinition},
	)
}

