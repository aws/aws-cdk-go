package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an Aurora cluster instance This can be either a provisioned instance or a serverless v2 instance.
type IClusterInstance interface {
	// Create the database instance within the provided cluster.
	Bind(scope constructs.Construct, cluster IDatabaseCluster, options *ClusterInstanceBindOptions) IAuroraClusterInstance
}

// The jsii proxy for IClusterInstance
type jsiiProxy_IClusterInstance struct {
	_ byte // padding
}

func (i *jsiiProxy_IClusterInstance) Bind(scope constructs.Construct, cluster IDatabaseCluster, options *ClusterInstanceBindOptions) IAuroraClusterInstance {
	if err := i.validateBindParameters(scope, cluster, options); err != nil {
		panic(err)
	}
	var returns IAuroraClusterInstance

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, cluster, options},
		&returns,
	)

	return returns
}

