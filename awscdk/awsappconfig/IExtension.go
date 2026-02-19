package awsappconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappconfig/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappconfig"
	"github.com/aws/constructs-go/constructs/v10"
)

type IExtension interface {
	interfacesawsappconfig.IExtensionRef
	awscdk.IResource
	// The actions for the extension.
	Actions() *[]Action
	// The description of the extension.
	Description() *string
	// The Amazon Resource Name (ARN) of the extension.
	ExtensionArn() *string
	// The ID of the extension.
	ExtensionId() *string
	// The version number of the extension.
	ExtensionVersionNumber() *float64
	// The latest version number of the extension.
	LatestVersionNumber() *float64
	// The name of the extension.
	Name() *string
	// The parameters of the extension.
	Parameters() *[]Parameter
}

// The jsii proxy for IExtension
type jsiiProxy_IExtension struct {
	internal.Type__interfacesawsappconfigIExtensionRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IExtension) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IExtension) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IExtension) Actions() *[]Action {
	var returns *[]Action
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExtension) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExtension) ExtensionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExtension) ExtensionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExtension) ExtensionVersionNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"extensionVersionNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExtension) LatestVersionNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"latestVersionNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExtension) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExtension) Parameters() *[]Parameter {
	var returns *[]Parameter
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExtension) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExtension) ExtensionRef() *interfacesawsappconfig.ExtensionReference {
	var returns *interfacesawsappconfig.ExtensionReference
	_jsii_.Get(
		j,
		"extensionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExtension) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExtension) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

