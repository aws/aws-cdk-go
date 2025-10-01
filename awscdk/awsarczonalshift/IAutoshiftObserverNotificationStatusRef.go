package awsarczonalshift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsarczonalshift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AutoshiftObserverNotificationStatus.
// Experimental.
type IAutoshiftObserverNotificationStatusRef interface {
	constructs.IConstruct
	// A reference to a AutoshiftObserverNotificationStatus resource.
	// Experimental.
	AutoshiftObserverNotificationStatusRef() *AutoshiftObserverNotificationStatusReference
}

// The jsii proxy for IAutoshiftObserverNotificationStatusRef
type jsiiProxy_IAutoshiftObserverNotificationStatusRef struct {
	internal.Type__constructsIConstruct
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

