package awsodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsodb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CloudAutonomousVmCluster.
// Experimental.
type ICloudAutonomousVmClusterRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICloudAutonomousVmClusterRef
type jsiiProxy_ICloudAutonomousVmClusterRef struct {
	internal.Type__constructsIConstruct
}

