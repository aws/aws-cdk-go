package awsodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsodb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CloudExadataInfrastructure.
// Experimental.
type ICloudExadataInfrastructureRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a CloudExadataInfrastructure resource.
	// Experimental.
	CloudExadataInfrastructureRef() *CloudExadataInfrastructureReference
}

// The jsii proxy for ICloudExadataInfrastructureRef
type jsiiProxy_ICloudExadataInfrastructureRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_ICloudExadataInfrastructureRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICloudExadataInfrastructureRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

