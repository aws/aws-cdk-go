package awsodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsodb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CloudExadataInfrastructure.
// Experimental.
type ICloudExadataInfrastructureRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICloudExadataInfrastructureRef
type jsiiProxy_ICloudExadataInfrastructureRef struct {
	internal.Type__constructsIConstruct
}

