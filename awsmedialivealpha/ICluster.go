package awsmedialivealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmedialive"
	"github.com/aws/aws-cdk-go/awsmedialivealpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a MediaLive Cluster.
// Experimental.
type ICluster interface {
	interfacesawsmedialive.IClusterRef
	awscdk.IResource
	// The ARN of the cluster.
	// Experimental.
	ClusterArn() *string
	// The IDs of channels running on this cluster.
	// Experimental.
	ClusterChannelIds() *[]*string
	// The ID of the cluster.
	// Experimental.
	ClusterId() *string
	// The current state of the cluster.
	// Experimental.
	ClusterState() *string
}

// The jsii proxy for ICluster
type jsiiProxy_ICluster struct {
	internal.Type__interfacesawsmedialiveIClusterRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ICluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_ICluster) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ICluster) ClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) ClusterChannelIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterChannelIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) ClusterState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) ClusterRef() *interfacesawsmedialive.ClusterReference {
	var returns *interfacesawsmedialive.ClusterReference
	_jsii_.Get(
		j,
		"clusterRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

