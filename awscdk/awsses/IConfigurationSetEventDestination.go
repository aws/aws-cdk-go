package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsses"
	"github.com/aws/constructs-go/constructs/v10"
)

// A configuration set event destination.
type IConfigurationSetEventDestination interface {
	interfacesawsses.IConfigurationSetEventDestinationRef
	awscdk.IResource
	// The ID of the configuration set event destination.
	ConfigurationSetEventDestinationId() *string
}

// The jsii proxy for IConfigurationSetEventDestination
type jsiiProxy_IConfigurationSetEventDestination struct {
	internal.Type__interfacesawssesIConfigurationSetEventDestinationRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IConfigurationSetEventDestination) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IConfigurationSetEventDestination) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IConfigurationSetEventDestination) ConfigurationSetEventDestinationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationSetEventDestinationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfigurationSetEventDestination) ConfigurationSetEventDestinationRef() *interfacesawsses.ConfigurationSetEventDestinationReference {
	var returns *interfacesawsses.ConfigurationSetEventDestinationReference
	_jsii_.Get(
		j,
		"configurationSetEventDestinationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfigurationSetEventDestination) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfigurationSetEventDestination) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConfigurationSetEventDestination) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

