package interfacesawsgroundstation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsgroundstation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DataflowEndpointGroupV2.
// Experimental.
type IDataflowEndpointGroupV2Ref interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DataflowEndpointGroupV2 resource.
	// Experimental.
	DataflowEndpointGroupV2Ref() *DataflowEndpointGroupV2Reference
}

// The jsii proxy for IDataflowEndpointGroupV2Ref
type jsiiProxy_IDataflowEndpointGroupV2Ref struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IDataflowEndpointGroupV2Ref) DataflowEndpointGroupV2Ref() *DataflowEndpointGroupV2Reference {
	var returns *DataflowEndpointGroupV2Reference
	_jsii_.Get(
		j,
		"dataflowEndpointGroupV2Ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataflowEndpointGroupV2Ref) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDataflowEndpointGroupV2Ref) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

