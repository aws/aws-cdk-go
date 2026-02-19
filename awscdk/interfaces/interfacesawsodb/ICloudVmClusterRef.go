package interfacesawsodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsodb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CloudVmCluster.
// Experimental.
type ICloudVmClusterRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a CloudVmCluster resource.
	// Experimental.
	CloudVmClusterRef() *CloudVmClusterReference
}

// The jsii proxy for ICloudVmClusterRef
type jsiiProxy_ICloudVmClusterRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ICloudVmClusterRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ICloudVmClusterRef) CloudVmClusterRef() *CloudVmClusterReference {
	var returns *CloudVmClusterReference
	_jsii_.Get(
		j,
		"cloudVmClusterRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICloudVmClusterRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICloudVmClusterRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

