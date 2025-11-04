package awscdkapplicationsignalsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// Node.js-specific implementation of the SDK injector. Handles Node.js auto-instrumentation setup and NODE_OPTIONS configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import applicationsignals_alpha "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha"
//
//   var instrumentationVersion InstrumentationVersion
//
//   nodeInjector := applicationsignals_alpha.NewNodeInjector(jsii.String("sharedVolumeName"), instrumentationVersion, []EnvironmentExtension{
//   	&EnvironmentExtension{
//   		Name: jsii.String("name"),
//   		Value: jsii.String("value"),
//   	},
//   })
//
// Experimental.
type NodeInjector interface {
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

// The jsii proxy struct for NodeInjector
type jsiiProxy_NodeInjector struct {
	jsiiProxy_Injector
}

func (j *jsiiProxy_NodeInjector) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeInjector) ContainerPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeInjector) InstrumentationVersion() InstrumentationVersion {
	var returns InstrumentationVersion
	_jsii_.Get(
		j,
		"instrumentationVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeInjector) SharedVolumeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedVolumeName",
		&returns,
	)
	return returns
}


// Experimental.
func NewNodeInjector(sharedVolumeName *string, instrumentationVersion InstrumentationVersion, overrideEnvironments *[]*EnvironmentExtension) NodeInjector {
	_init_.Initialize()

	if err := validateNewNodeInjectorParameters(sharedVolumeName, instrumentationVersion, overrideEnvironments); err != nil {
		panic(err)
	}
	j := jsiiProxy_NodeInjector{}

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.NodeInjector",
		[]interface{}{sharedVolumeName, instrumentationVersion, overrideEnvironments},
		&j,
	)

	return &j
}

// Experimental.
func NewNodeInjector_Override(n NodeInjector, sharedVolumeName *string, instrumentationVersion InstrumentationVersion, overrideEnvironments *[]*EnvironmentExtension) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-applicationsignals-alpha.NodeInjector",
		[]interface{}{sharedVolumeName, instrumentationVersion, overrideEnvironments},
		n,
	)
}

func (j *jsiiProxy_NodeInjector)SetInstrumentationVersion(val InstrumentationVersion) {
	if err := j.validateSetInstrumentationVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instrumentationVersion",
		val,
	)
}

func (j *jsiiProxy_NodeInjector)SetSharedVolumeName(val *string) {
	if err := j.validateSetSharedVolumeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedVolumeName",
		val,
	)
}

func NodeInjector_DEFAULT_ENVS() *[]*EnvironmentExtension {
	_init_.Initialize()
	var returns *[]*EnvironmentExtension
	_jsii_.StaticGet(
		"@aws-cdk/aws-applicationsignals-alpha.NodeInjector",
		"DEFAULT_ENVS",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NodeInjector) InjectAdditionalEnvironments(envsToInject *map[string]*string, envsFromTaskDef *map[string]*string) {
	if err := n.validateInjectAdditionalEnvironmentsParameters(envsToInject, envsFromTaskDef); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"injectAdditionalEnvironments",
		[]interface{}{envsToInject, envsFromTaskDef},
	)
}

func (n *jsiiProxy_NodeInjector) InjectInitContainer(taskDefinition awsecs.TaskDefinition) awsecs.ContainerDefinition {
	if err := n.validateInjectInitContainerParameters(taskDefinition); err != nil {
		panic(err)
	}
	var returns awsecs.ContainerDefinition

	_jsii_.Invoke(
		n,
		"injectInitContainer",
		[]interface{}{taskDefinition},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NodeInjector) OverrideAdditionalEnvironments(envsToOverride *map[string]*string, envsFromTaskDef *map[string]*string) {
	if err := n.validateOverrideAdditionalEnvironmentsParameters(envsToOverride, envsFromTaskDef); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideAdditionalEnvironments",
		[]interface{}{envsToOverride, envsFromTaskDef},
	)
}

func (n *jsiiProxy_NodeInjector) RenderDefaultContainer(taskDefinition awsecs.TaskDefinition) {
	if err := n.validateRenderDefaultContainerParameters(taskDefinition); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"renderDefaultContainer",
		[]interface{}{taskDefinition},
	)
}

