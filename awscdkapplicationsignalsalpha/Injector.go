package awscdkapplicationsignalsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// Injector is a base class for all SDK injects to mutate the task definition to inject the ADOT init container and configure the application container with the necessary environment variables.
// Experimental.
type Injector interface {
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

// The jsii proxy struct for Injector
type jsiiProxy_Injector struct {
	_ byte // padding
}

func (j *jsiiProxy_Injector) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Injector) ContainerPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Injector) InstrumentationVersion() InstrumentationVersion {
	var returns InstrumentationVersion
	_jsii_.Get(
		j,
		"instrumentationVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Injector) SharedVolumeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedVolumeName",
		&returns,
	)
	return returns
}


// Experimental.
func NewInjector_Override(i Injector, sharedVolumeName *string, instrumentationVersion InstrumentationVersion, overrideEnvironments *[]*EnvironmentExtension) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.Injector",
		[]interface{}{sharedVolumeName, instrumentationVersion, overrideEnvironments},
		i,
	)
}

func (j *jsiiProxy_Injector)SetInstrumentationVersion(val InstrumentationVersion) {
	if err := j.validateSetInstrumentationVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instrumentationVersion",
		val,
	)
}

func (j *jsiiProxy_Injector)SetSharedVolumeName(val *string) {
	if err := j.validateSetSharedVolumeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedVolumeName",
		val,
	)
}

func Injector_DEFAULT_ENVS() *[]*EnvironmentExtension {
	_init_.Initialize()
	var returns *[]*EnvironmentExtension
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.Injector",
		"DEFAULT_ENVS",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_Injector) InjectAdditionalEnvironments(envsToInject *map[string]*string, envsFromTaskDef *map[string]*string) {
	if err := i.validateInjectAdditionalEnvironmentsParameters(envsToInject, envsFromTaskDef); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"injectAdditionalEnvironments",
		[]interface{}{envsToInject, envsFromTaskDef},
	)
}

func (i *jsiiProxy_Injector) InjectInitContainer(taskDefinition awsecs.TaskDefinition) awsecs.ContainerDefinition {
	if err := i.validateInjectInitContainerParameters(taskDefinition); err != nil {
		panic(err)
	}
	var returns awsecs.ContainerDefinition

	_jsii_.Invoke(
		i,
		"injectInitContainer",
		[]interface{}{taskDefinition},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_Injector) OverrideAdditionalEnvironments(envsToOverride *map[string]*string, envsFromTaskDef *map[string]*string) {
	if err := i.validateOverrideAdditionalEnvironmentsParameters(envsToOverride, envsFromTaskDef); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideAdditionalEnvironments",
		[]interface{}{envsToOverride, envsFromTaskDef},
	)
}

func (i *jsiiProxy_Injector) RenderDefaultContainer(taskDefinition awsecs.TaskDefinition) {
	if err := i.validateRenderDefaultContainerParameters(taskDefinition); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"renderDefaultContainer",
		[]interface{}{taskDefinition},
	)
}

