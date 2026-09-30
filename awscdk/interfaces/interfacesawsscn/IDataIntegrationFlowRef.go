package interfacesawsscn

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsscn/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataIntegrationFlow.
// Experimental.
type IDataIntegrationFlowRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DataIntegrationFlow resource.
	// Experimental.
	DataIntegrationFlowRef() *DataIntegrationFlowReference
}

// The jsii proxy for IDataIntegrationFlowRef
type jsiiProxy_IDataIntegrationFlowRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDataIntegrationFlowRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IDataIntegrationFlowRef) DataIntegrationFlowRef() *DataIntegrationFlowReference {
	var returns *DataIntegrationFlowReference
	_jsii_.Get(
		j,
		"dataIntegrationFlowRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataIntegrationFlowRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataIntegrationFlowRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

