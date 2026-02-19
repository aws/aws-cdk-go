package interfacesawsodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsodb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CloudExadataInfrastructure.
// Experimental.
type ICloudExadataInfrastructureRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a CloudExadataInfrastructure resource.
	// Experimental.
	CloudExadataInfrastructureRef() *CloudExadataInfrastructureReference
}

// The jsii proxy for ICloudExadataInfrastructureRef
type jsiiProxy_ICloudExadataInfrastructureRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ICloudExadataInfrastructureRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ICloudExadataInfrastructureRef) CloudExadataInfrastructureRef() *CloudExadataInfrastructureReference {
	var returns *CloudExadataInfrastructureReference
	_jsii_.Get(
		j,
		"cloudExadataInfrastructureRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICloudExadataInfrastructureRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
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

