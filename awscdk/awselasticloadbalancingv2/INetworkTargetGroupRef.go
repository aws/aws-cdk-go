package awselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawselasticloadbalancingv2"
)

// Indicates that this resource can be referenced as an NLB TargetGroup.
type INetworkTargetGroupRef interface {
	interfacesawselasticloadbalancingv2.ITargetGroupRef
	// Indicates that this is a Network Target Group.
	//
	// Will always return true, but is necessary to prevent accidental structural
	// equality in TypeScript.
	IsNetworkTargetGroup() *bool
}

// The jsii proxy for INetworkTargetGroupRef
type jsiiProxy_INetworkTargetGroupRef struct {
	internal.Type__interfacesawselasticloadbalancingv2ITargetGroupRef
}

func (j *jsiiProxy_INetworkTargetGroupRef) IsNetworkTargetGroup() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isNetworkTargetGroup",
		&returns,
	)
	return returns
}

