package awslookoutmetrics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslookoutmetrics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Alert.
// Experimental.
type IAlertRef interface {
	constructs.IConstruct
	// A reference to a Alert resource.
	// Experimental.
	AlertRef() *AlertReference
}

// The jsii proxy for IAlertRef
type jsiiProxy_IAlertRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAlertRef) AlertRef() *AlertReference {
	var returns *AlertReference
	_jsii_.Get(
		j,
		"alertRef",
		&returns,
	)
	return returns
}

