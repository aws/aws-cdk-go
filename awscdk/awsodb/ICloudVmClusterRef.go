package awsodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsodb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CloudVmCluster.
// Experimental.
type ICloudVmClusterRef interface {
	constructs.IConstruct
	// A reference to a CloudVmCluster resource.
	// Experimental.
	CloudVmClusterRef() *CloudVmClusterReference
}

// The jsii proxy for ICloudVmClusterRef
type jsiiProxy_ICloudVmClusterRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICloudVmClusterRef) CloudVmClusterRef() *CloudVmClusterReference {
	var returns *CloudVmClusterReference
	_jsii_.Get(
		j,
		"cloudVmClusterRef",
		&returns,
	)
	return returns
}

