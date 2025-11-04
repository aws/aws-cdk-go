package awsarczonalshift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsarczonalshift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AutoshiftObserverNotificationStatus.
// Experimental.
type IAutoshiftObserverNotificationStatusRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a AutoshiftObserverNotificationStatus resource.
	// Experimental.
	AutoshiftObserverNotificationStatusRef() *AutoshiftObserverNotificationStatusReference
}

// The jsii proxy for IAutoshiftObserverNotificationStatusRef
type jsiiProxy_IAutoshiftObserverNotificationStatusRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IAutoshiftObserverNotificationStatusRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

