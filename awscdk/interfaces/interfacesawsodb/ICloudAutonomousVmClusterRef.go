package interfacesawsodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsodb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CloudAutonomousVmCluster.
// Experimental.
type ICloudAutonomousVmClusterRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a CloudAutonomousVmCluster resource.
	// Experimental.
	CloudAutonomousVmClusterRef() *CloudAutonomousVmClusterReference
}

// The jsii proxy for ICloudAutonomousVmClusterRef
type jsiiProxy_ICloudAutonomousVmClusterRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ICloudAutonomousVmClusterRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ICloudAutonomousVmClusterRef) CloudAutonomousVmClusterRef() *CloudAutonomousVmClusterReference {
	var returns *CloudAutonomousVmClusterReference
	_jsii_.Get(
		j,
		"cloudAutonomousVmClusterRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICloudAutonomousVmClusterRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICloudAutonomousVmClusterRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

