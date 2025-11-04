package awscdkapplicationsignalsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// Python-specific implementation of the SDK injector.
//
// Handles Python auto-instrumentation setup and PYTHONPATH configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import applicationsignals_alpha "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha"
//
//   var instrumentationVersion InstrumentationVersion
//
//   pythonInjector := applicationsignals_alpha.NewPythonInjector(jsii.String("sharedVolumeName"), instrumentationVersion, []EnvironmentExtension{
//   	&EnvironmentExtension{
//   		Name: jsii.String("name"),
//   		Value: jsii.String("value"),
//   	},
//   })
//
// Experimental.
type PythonInjector interface {
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

// The jsii proxy struct for PythonInjector
type jsiiProxy_PythonInjector struct {
	jsiiProxy_Injector
}

func (j *jsiiProxy_PythonInjector) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonInjector) ContainerPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonInjector) InstrumentationVersion() InstrumentationVersion {
	var returns InstrumentationVersion
	_jsii_.Get(
		j,
		"instrumentationVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PythonInjector) SharedVolumeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedVolumeName",
		&returns,
	)
	return returns
}


// Experimental.
func NewPythonInjector(sharedVolumeName *string, instrumentationVersion InstrumentationVersion, overrideEnvironments *[]*EnvironmentExtension) PythonInjector {
	_init_.Initialize()

	if err := validateNewPythonInjectorParameters(sharedVolumeName, instrumentationVersion, overrideEnvironments); err != nil {
		panic(err)
	}
	j := jsiiProxy_PythonInjector{}

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.PythonInjector",
		[]interface{}{sharedVolumeName, instrumentationVersion, overrideEnvironments},
		&j,
	)

	return &j
}

// Experimental.
func NewPythonInjector_Override(p PythonInjector, sharedVolumeName *string, instrumentationVersion InstrumentationVersion, overrideEnvironments *[]*EnvironmentExtension) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.PythonInjector",
		[]interface{}{sharedVolumeName, instrumentationVersion, overrideEnvironments},
		p,
	)
}

func (j *jsiiProxy_PythonInjector)SetInstrumentationVersion(val InstrumentationVersion) {
	if err := j.validateSetInstrumentationVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instrumentationVersion",
		val,
	)
}

func (j *jsiiProxy_PythonInjector)SetSharedVolumeName(val *string) {
	if err := j.validateSetSharedVolumeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedVolumeName",
		val,
	)
}

func PythonInjector_DEFAULT_ENVS() *[]*EnvironmentExtension {
	_init_.Initialize()
	var returns *[]*EnvironmentExtension
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.PythonInjector",
		"DEFAULT_ENVS",
		&returns,
	)
	return returns
}

func PythonInjector_PYTHON_ENVS() *[]*EnvironmentExtension {
	_init_.Initialize()
	var returns *[]*EnvironmentExtension
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.PythonInjector",
		"PYTHON_ENVS",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PythonInjector) InjectAdditionalEnvironments(envsToInject *map[string]*string, envsFromTaskDef *map[string]*string) {
	if err := p.validateInjectAdditionalEnvironmentsParameters(envsToInject, envsFromTaskDef); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"injectAdditionalEnvironments",
		[]interface{}{envsToInject, envsFromTaskDef},
	)
}

func (p *jsiiProxy_PythonInjector) InjectInitContainer(taskDefinition awsecs.TaskDefinition) awsecs.ContainerDefinition {
	if err := p.validateInjectInitContainerParameters(taskDefinition); err != nil {
		panic(err)
	}
	var returns awsecs.ContainerDefinition

	_jsii_.Invoke(
		p,
		"injectInitContainer",
		[]interface{}{taskDefinition},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PythonInjector) OverrideAdditionalEnvironments(envsToOverride *map[string]*string, envsFromTaskDef *map[string]*string) {
	if err := p.validateOverrideAdditionalEnvironmentsParameters(envsToOverride, envsFromTaskDef); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideAdditionalEnvironments",
		[]interface{}{envsToOverride, envsFromTaskDef},
	)
}

func (p *jsiiProxy_PythonInjector) RenderDefaultContainer(taskDefinition awsecs.TaskDefinition) {
	if err := p.validateRenderDefaultContainerParameters(taskDefinition); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"renderDefaultContainer",
		[]interface{}{taskDefinition},
	)
}

