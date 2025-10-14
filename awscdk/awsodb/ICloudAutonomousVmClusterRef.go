package awsodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsodb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CloudAutonomousVmCluster.
// Experimental.
type ICloudAutonomousVmClusterRef interface {
	constructs.IConstruct
	// A reference to a CloudAutonomousVmCluster resource.
	// Experimental.
	CloudAutonomousVmClusterRef() *CloudAutonomousVmClusterReference
}

// The jsii proxy for ICloudAutonomousVmClusterRef
type jsiiProxy_ICloudAutonomousVmClusterRef struct {
	internal.Type__constructsIConstruct
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

