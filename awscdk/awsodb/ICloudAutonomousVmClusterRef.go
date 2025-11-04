package awsodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsodb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CloudAutonomousVmCluster.
// Experimental.
type ICloudAutonomousVmClusterRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a CloudAutonomousVmCluster resource.
	// Experimental.
	CloudAutonomousVmClusterRef() *CloudAutonomousVmClusterReference
}

// The jsii proxy for ICloudAutonomousVmClusterRef
type jsiiProxy_ICloudAutonomousVmClusterRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_ICloudAutonomousVmClusterRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

