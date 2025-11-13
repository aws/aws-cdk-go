package interfacesawsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataAutomationProject.
// Experimental.
type IDataAutomationProjectRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DataAutomationProject resource.
	// Experimental.
	DataAutomationProjectRef() *DataAutomationProjectReference
}

// The jsii proxy for IDataAutomationProjectRef
type jsiiProxy_IDataAutomationProjectRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IDataAutomationProjectRef) DataAutomationProjectRef() *DataAutomationProjectReference {
	var returns *DataAutomationProjectReference
	_jsii_.Get(
		j,
		"dataAutomationProjectRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataAutomationProjectRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataAutomationProjectRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

