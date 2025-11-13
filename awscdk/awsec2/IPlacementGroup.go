package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Determines where your instances are placed on the underlying hardware according to the specified PlacementGroupStrategy.
// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html
//
type IPlacementGroup interface {
	interfacesawsec2.IPlacementGroupRef
	awscdk.IResource
	// The number of partitions.
	//
	// Valid only when Strategy is set to PARTITION.
	// Default: 0.
	//
	Partitions() *float64
	// The name of this placement group.
	PlacementGroupName() *string
	// Places instances on distinct hardware.
	//
	// Spread placement groups are recommended for applications
	// that have a small number of critical instances that should be kept separate from each other.
	// Launching instances in a spread level placement group reduces the risk of simultaneous failures
	// that might occur when instances share the same equipment.
	// Spread level placement groups provide access to distinct hardware,
	// and are therefore suitable for mixing instance types or launching instances over time.
	// If you start or launch an instance in a spread placement group and there is insufficient
	// unique hardware to fulfill the request, the request fails. Amazon EC2 makes more distinct hardware
	// available over time, so you can try your request again later.
	// Placement groups can spread instances across racks or hosts.
	// You can use host level spread placement groups only with AWS Outposts.
	// Default: - no spread level.
	//
	SpreadLevel() PlacementGroupSpreadLevel
	// Which strategy to use when launching instances.
	// Default: - `PlacementGroupStrategy.PARTITION` if `partitions` is defined, `CLUSTER` otherwise
	//
	Strategy() PlacementGroupStrategy
}

// The jsii proxy for IPlacementGroup
type jsiiProxy_IPlacementGroup struct {
	internal.Type__interfacesawsec2IPlacementGroupRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IPlacementGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IPlacementGroup) Partitions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPlacementGroup) PlacementGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPlacementGroup) SpreadLevel() PlacementGroupSpreadLevel {
	var returns PlacementGroupSpreadLevel
	_jsii_.Get(
		j,
		"spreadLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPlacementGroup) Strategy() PlacementGroupStrategy {
	var returns PlacementGroupStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPlacementGroup) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPlacementGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPlacementGroup) PlacementGroupRef() *interfacesawsec2.PlacementGroupReference {
	var returns *interfacesawsec2.PlacementGroupReference
	_jsii_.Get(
		j,
		"placementGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPlacementGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

