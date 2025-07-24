package awsappconfig

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// This class is meant to be used by AWS AppConfig resources (application, configuration profile, environment) directly.
//
// There is currently no use
// for this class outside of the AWS AppConfig construct implementation. It is
// intended to be used with the resources since there is currently no way to
// inherit from two classes (at least within JSII constraints).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   extensibleBase := awscdk.Aws_appconfig.NewExtensibleBase(this, jsii.String("resourceArn"), jsii.String("resourceName"))
//
type ExtensibleBase interface {
	IExtensible
	// Adds an extension association to the derived resource.
	AddExtension(extension IExtension)
	// Adds an AT_DEPLOYMENT_TICK extension with the provided event destination and also creates an extension association to the derived resource.
	AtDeploymentTick(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an extension defined by the action point and event destination and also creates an extension association to the derived resource.
	On(actionPoint ActionPoint, eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_BAKING extension with the provided event destination and also creates an extension association to the derived resource.
	OnDeploymentBaking(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_COMPLETE extension with the provided event destination and also creates an extension association to the derived resource.
	OnDeploymentComplete(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_ROLLED_BACK extension with the provided event destination and also creates an extension association to the derived resource.
	OnDeploymentRolledBack(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_START extension with the provided event destination and also creates an extension association to the derived resource.
	OnDeploymentStart(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds an ON_DEPLOYMENT_STEP extension with the provided event destination and also creates an extension association to the derived resource.
	OnDeploymentStep(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds a PRE_CREATE_HOSTED_CONFIGURATION_VERSION extension with the provided event destination and also creates an extension association to the derived resource.
	PreCreateHostedConfigurationVersion(eventDestination IEventDestination, options *ExtensionOptions)
	// Adds a PRE_START_DEPLOYMENT extension with the provided event destination and also creates an extension association to the derived resource.
	PreStartDeployment(eventDestination IEventDestination, options *ExtensionOptions)
}

// The jsii proxy struct for ExtensibleBase
type jsiiProxy_ExtensibleBase struct {
	jsiiProxy_IExtensible
}

func NewExtensibleBase(scope constructs.Construct, resourceArn *string, resourceName *string) ExtensibleBase {
	_init_.Initialize()

	if err := validateNewExtensibleBaseParameters(scope, resourceArn); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExtensibleBase{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appconfig.ExtensibleBase",
		[]interface{}{scope, resourceArn, resourceName},
		&j,
	)

	return &j
}

func NewExtensibleBase_Override(e ExtensibleBase, scope constructs.Construct, resourceArn *string, resourceName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appconfig.ExtensibleBase",
		[]interface{}{scope, resourceArn, resourceName},
		e,
	)
}

func (e *jsiiProxy_ExtensibleBase) AddExtension(extension IExtension) {
	if err := e.validateAddExtensionParameters(extension); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addExtension",
		[]interface{}{extension},
	)
}

func (e *jsiiProxy_ExtensibleBase) AtDeploymentTick(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := e.validateAtDeploymentTickParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"atDeploymentTick",
		[]interface{}{eventDestination, options},
	)
}

func (e *jsiiProxy_ExtensibleBase) On(actionPoint ActionPoint, eventDestination IEventDestination, options *ExtensionOptions) {
	if err := e.validateOnParameters(actionPoint, eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"on",
		[]interface{}{actionPoint, eventDestination, options},
	)
}

func (e *jsiiProxy_ExtensibleBase) OnDeploymentBaking(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := e.validateOnDeploymentBakingParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"onDeploymentBaking",
		[]interface{}{eventDestination, options},
	)
}

func (e *jsiiProxy_ExtensibleBase) OnDeploymentComplete(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := e.validateOnDeploymentCompleteParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"onDeploymentComplete",
		[]interface{}{eventDestination, options},
	)
}

func (e *jsiiProxy_ExtensibleBase) OnDeploymentRolledBack(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := e.validateOnDeploymentRolledBackParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"onDeploymentRolledBack",
		[]interface{}{eventDestination, options},
	)
}

func (e *jsiiProxy_ExtensibleBase) OnDeploymentStart(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := e.validateOnDeploymentStartParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"onDeploymentStart",
		[]interface{}{eventDestination, options},
	)
}

func (e *jsiiProxy_ExtensibleBase) OnDeploymentStep(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := e.validateOnDeploymentStepParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"onDeploymentStep",
		[]interface{}{eventDestination, options},
	)
}

func (e *jsiiProxy_ExtensibleBase) PreCreateHostedConfigurationVersion(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := e.validatePreCreateHostedConfigurationVersionParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"preCreateHostedConfigurationVersion",
		[]interface{}{eventDestination, options},
	)
}

func (e *jsiiProxy_ExtensibleBase) PreStartDeployment(eventDestination IEventDestination, options *ExtensionOptions) {
	if err := e.validatePreStartDeploymentParameters(eventDestination, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"preStartDeployment",
		[]interface{}{eventDestination, options},
	)
}

