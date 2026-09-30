package interfacesawsbcmpricingcalculator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbcmpricingcalculator/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WorkloadEstimate.
// Experimental.
type IWorkloadEstimateRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a WorkloadEstimate resource.
	// Experimental.
	WorkloadEstimateRef() *WorkloadEstimateReference
}

// The jsii proxy for IWorkloadEstimateRef
type jsiiProxy_IWorkloadEstimateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IWorkloadEstimateRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IWorkloadEstimateRef) WorkloadEstimateRef() *WorkloadEstimateReference {
	var returns *WorkloadEstimateReference
	_jsii_.Get(
		j,
		"workloadEstimateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkloadEstimateRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkloadEstimateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

