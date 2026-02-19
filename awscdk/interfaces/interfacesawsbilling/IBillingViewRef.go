package interfacesawsbilling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbilling/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a BillingView.
// Experimental.
type IBillingViewRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a BillingView resource.
	// Experimental.
	BillingViewRef() *BillingViewReference
}

// The jsii proxy for IBillingViewRef
type jsiiProxy_IBillingViewRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IBillingViewRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IBillingViewRef) BillingViewRef() *BillingViewReference {
	var returns *BillingViewReference
	_jsii_.Get(
		j,
		"billingViewRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBillingViewRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBillingViewRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

