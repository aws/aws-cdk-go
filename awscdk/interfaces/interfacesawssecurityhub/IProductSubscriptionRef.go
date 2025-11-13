package interfacesawssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProductSubscription.
// Experimental.
type IProductSubscriptionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ProductSubscription resource.
	// Experimental.
	ProductSubscriptionRef() *ProductSubscriptionReference
}

// The jsii proxy for IProductSubscriptionRef
type jsiiProxy_IProductSubscriptionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IProductSubscriptionRef) ProductSubscriptionRef() *ProductSubscriptionReference {
	var returns *ProductSubscriptionReference
	_jsii_.Get(
		j,
		"productSubscriptionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProductSubscriptionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProductSubscriptionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

