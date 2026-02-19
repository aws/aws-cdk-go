package interfacesawssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConnectorV2.
// Experimental.
type IConnectorV2Ref interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ConnectorV2 resource.
	// Experimental.
	ConnectorV2Ref() *ConnectorV2Reference
}

// The jsii proxy for IConnectorV2Ref
type jsiiProxy_IConnectorV2Ref struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IConnectorV2Ref) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IConnectorV2Ref) ConnectorV2Ref() *ConnectorV2Reference {
	var returns *ConnectorV2Reference
	_jsii_.Get(
		j,
		"connectorV2Ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnectorV2Ref) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnectorV2Ref) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

