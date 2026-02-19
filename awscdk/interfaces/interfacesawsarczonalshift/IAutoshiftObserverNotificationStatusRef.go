package interfacesawsarczonalshift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsarczonalshift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AutoshiftObserverNotificationStatus.
// Experimental.
type IAutoshiftObserverNotificationStatusRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AutoshiftObserverNotificationStatus resource.
	// Experimental.
	AutoshiftObserverNotificationStatusRef() *AutoshiftObserverNotificationStatusReference
}

// The jsii proxy for IAutoshiftObserverNotificationStatusRef
type jsiiProxy_IAutoshiftObserverNotificationStatusRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IAutoshiftObserverNotificationStatusRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IAutoshiftObserverNotificationStatusRef) AutoshiftObserverNotificationStatusRef() *AutoshiftObserverNotificationStatusReference {
	var returns *AutoshiftObserverNotificationStatusReference
	_jsii_.Get(
		j,
		"autoshiftObserverNotificationStatusRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAutoshiftObserverNotificationStatusRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAutoshiftObserverNotificationStatusRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

