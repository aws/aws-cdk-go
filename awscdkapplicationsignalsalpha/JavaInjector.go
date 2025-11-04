package awscdkapplicationsignalsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// Java-specific implementation of the SDK injector.
//
// Handles Java agent configuration and environment setup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import applicationsignals_alpha "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha"
//
//   var instrumentationVersion InstrumentationVersion
//
//   javaInjector := applicationsignals_alpha.NewJavaInjector(jsii.String("sharedVolumeName"), instrumentationVersion, []EnvironmentExtension{
//   	&EnvironmentExtension{
//   		Name: jsii.String("name"),
//   		Value: jsii.String("value"),
//   	},
//   })
//
// Experimental.
type JavaInjector interface {
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

// The jsii proxy struct for JavaInjector
type jsiiProxy_JavaInjector struct {
	jsiiProxy_Injector
}

func (j *jsiiProxy_JavaInjector) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaInjector) ContainerPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaInjector) InstrumentationVersion() InstrumentationVersion {
	var returns InstrumentationVersion
	_jsii_.Get(
		j,
		"instrumentationVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaInjector) SharedVolumeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedVolumeName",
		&returns,
	)
	return returns
}


// Experimental.
func NewJavaInjector(sharedVolumeName *string, instrumentationVersion InstrumentationVersion, overrideEnvironments *[]*EnvironmentExtension) JavaInjector {
	_init_.Initialize()

	if err := validateNewJavaInjectorParameters(sharedVolumeName, instrumentationVersion, overrideEnvironments); err != nil {
		panic(err)
	}
	j := jsiiProxy_JavaInjector{}

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.JavaInjector",
		[]interface{}{sharedVolumeName, instrumentationVersion, overrideEnvironments},
		&j,
	)

	return &j
}

// Experimental.
func NewJavaInjector_Override(j JavaInjector, sharedVolumeName *string, instrumentationVersion InstrumentationVersion, overrideEnvironments *[]*EnvironmentExtension) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.JavaInjector",
		[]interface{}{sharedVolumeName, instrumentationVersion, overrideEnvironments},
		j,
	)
}

func (j *jsiiProxy_JavaInjector)SetInstrumentationVersion(val InstrumentationVersion) {
	if err := j.validateSetInstrumentationVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instrumentationVersion",
		val,
	)
}

func (j *jsiiProxy_JavaInjector)SetSharedVolumeName(val *string) {
	if err := j.validateSetSharedVolumeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedVolumeName",
		val,
	)
}

func JavaInjector_DEFAULT_ENVS() *[]*EnvironmentExtension {
	_init_.Initialize()
	var returns *[]*EnvironmentExtension
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.JavaInjector",
		"DEFAULT_ENVS",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JavaInjector) InjectAdditionalEnvironments(envsToInject *map[string]*string, envsFromTaskDef *map[string]*string) {
	if err := j.validateInjectAdditionalEnvironmentsParameters(envsToInject, envsFromTaskDef); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"injectAdditionalEnvironments",
		[]interface{}{envsToInject, envsFromTaskDef},
	)
}

func (j *jsiiProxy_JavaInjector) InjectInitContainer(taskDefinition awsecs.TaskDefinition) awsecs.ContainerDefinition {
	if err := j.validateInjectInitContainerParameters(taskDefinition); err != nil {
		panic(err)
	}
	var returns awsecs.ContainerDefinition

	_jsii_.Invoke(
		j,
		"injectInitContainer",
		[]interface{}{taskDefinition},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JavaInjector) OverrideAdditionalEnvironments(envsToOverride *map[string]*string, envsFromTaskDef *map[string]*string) {
	if err := j.validateOverrideAdditionalEnvironmentsParameters(envsToOverride, envsFromTaskDef); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"overrideAdditionalEnvironments",
		[]interface{}{envsToOverride, envsFromTaskDef},
	)
}

func (j *jsiiProxy_JavaInjector) RenderDefaultContainer(taskDefinition awsecs.TaskDefinition) {
	if err := j.validateRenderDefaultContainerParameters(taskDefinition); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"renderDefaultContainer",
		[]interface{}{taskDefinition},
	)
}

