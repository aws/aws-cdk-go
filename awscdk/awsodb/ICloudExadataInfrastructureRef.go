package awsodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsodb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CloudExadataInfrastructure.
// Experimental.
type ICloudExadataInfrastructureRef interface {
	constructs.IConstruct
	// A reference to a CloudExadataInfrastructure resource.
	// Experimental.
	CloudExadataInfrastructureRef() *CloudExadataInfrastructureReference
}

// The jsii proxy for ICloudExadataInfrastructureRef
type jsiiProxy_ICloudExadataInfrastructureRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICloudExadataInfrastructureRef) CloudExadataInfrastructureRef() *CloudExadataInfrastructureReference {
	var returns *CloudExadataInfrastructureReference
	_jsii_.Get(
		j,
		"cloudExadataInfrastructureRef",
		&returns,
	)
	return returns
}

