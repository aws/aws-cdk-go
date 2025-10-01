package awscdkapplicationsignalsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// Linux-specific implementation of the .NET SDK injector. Handles CoreCLR profiler setup and paths for Linux environments.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import applicationsignals_alpha "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cpuArchitecture cpuArchitecture
//   var instrumentationVersion instrumentationVersion
//
//   dotNetLinuxInjector := applicationsignals_alpha.NewDotNetLinuxInjector(jsii.String("sharedVolumeName"), instrumentationVersion, cpuArchitecture, []environmentExtension{
//   	&environmentExtension{
//   		Name: jsii.String("name"),
//   		Value: jsii.String("value"),
//   	},
//   })
//
// Experimental.
type DotNetLinuxInjector interface {
	DotNetInjector
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
	OverrideAdditionalEnvironments(_envsToOverride *map[string]*string, _envsFromTaskDef *map[string]*string)
	// Render the application container for SDK instrumentation.
	// Experimental.
	RenderDefaultContainer(taskDefinition awsecs.TaskDefinition)
}

// The jsii proxy struct for DotNetLinuxInjector
type jsiiProxy_DotNetLinuxInjector struct {
	jsiiProxy_DotNetInjector
}

func (j *jsiiProxy_DotNetLinuxInjector) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DotNetLinuxInjector) ContainerPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DotNetLinuxInjector) InstrumentationVersion() InstrumentationVersion {
	var returns InstrumentationVersion
	_jsii_.Get(
		j,
		"instrumentationVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DotNetLinuxInjector) SharedVolumeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedVolumeName",
		&returns,
	)
	return returns
}


// Experimental.
func NewDotNetLinuxInjector(sharedVolumeName *string, instrumentationVersion InstrumentationVersion, cpuArch awsecs.CpuArchitecture, overrideEnvironments *[]*EnvironmentExtension) DotNetLinuxInjector {
	_init_.Initialize()

	if err := validateNewDotNetLinuxInjectorParameters(sharedVolumeName, instrumentationVersion, cpuArch, overrideEnvironments); err != nil {
		panic(err)
	}
	j := jsiiProxy_DotNetLinuxInjector{}

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.DotNetLinuxInjector",
		[]interface{}{sharedVolumeName, instrumentationVersion, cpuArch, overrideEnvironments},
		&j,
	)

	return &j
}

// Experimental.
func NewDotNetLinuxInjector_Override(d DotNetLinuxInjector, sharedVolumeName *string, instrumentationVersion InstrumentationVersion, cpuArch awsecs.CpuArchitecture, overrideEnvironments *[]*EnvironmentExtension) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.DotNetLinuxInjector",
		[]interface{}{sharedVolumeName, instrumentationVersion, cpuArch, overrideEnvironments},
		d,
	)
}

func (j *jsiiProxy_DotNetLinuxInjector)SetInstrumentationVersion(val InstrumentationVersion) {
	if err := j.validateSetInstrumentationVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instrumentationVersion",
		val,
	)
}

func (j *jsiiProxy_DotNetLinuxInjector)SetSharedVolumeName(val *string) {
	if err := j.validateSetSharedVolumeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedVolumeName",
		val,
	)
}

func DotNetLinuxInjector_DEFAULT_ENVS() *[]*EnvironmentExtension {
	_init_.Initialize()
	var returns *[]*EnvironmentExtension
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.DotNetLinuxInjector",
		"DEFAULT_ENVS",
		&returns,
	)
	return returns
}

func DotNetLinuxInjector_DOTNET_COMMON_ENVS() *[]*EnvironmentExtension {
	_init_.Initialize()
	var returns *[]*EnvironmentExtension
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.DotNetLinuxInjector",
		"DOTNET_COMMON_ENVS",
		&returns,
	)
	return returns
}

func DotNetLinuxInjector_DOTNET_LINUX_ENVS() *[]*EnvironmentExtension {
	_init_.Initialize()
	var returns *[]*EnvironmentExtension
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.DotNetLinuxInjector",
		"DOTNET_LINUX_ENVS",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DotNetLinuxInjector) InjectAdditionalEnvironments(envsToInject *map[string]*string, envsFromTaskDef *map[string]*string) {
	if err := d.validateInjectAdditionalEnvironmentsParameters(envsToInject, envsFromTaskDef); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"injectAdditionalEnvironments",
		[]interface{}{envsToInject, envsFromTaskDef},
	)
}

func (d *jsiiProxy_DotNetLinuxInjector) InjectInitContainer(taskDefinition awsecs.TaskDefinition) awsecs.ContainerDefinition {
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

func (d *jsiiProxy_DotNetLinuxInjector) OverrideAdditionalEnvironments(_envsToOverride *map[string]*string, _envsFromTaskDef *map[string]*string) {
	if err := d.validateOverrideAdditionalEnvironmentsParameters(_envsToOverride, _envsFromTaskDef); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideAdditionalEnvironments",
		[]interface{}{_envsToOverride, _envsFromTaskDef},
	)
}

func (d *jsiiProxy_DotNetLinuxInjector) RenderDefaultContainer(taskDefinition awsecs.TaskDefinition) {
	if err := d.validateRenderDefaultContainerParameters(taskDefinition); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"renderDefaultContainer",
		[]interface{}{taskDefinition},
	)
}

