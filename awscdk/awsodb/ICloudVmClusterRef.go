package awsodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsodb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CloudVmCluster.
// Experimental.
type ICloudVmClusterRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICloudVmClusterRef
type jsiiProxy_ICloudVmClusterRef struct {
	internal.Type__constructsIConstruct
}

