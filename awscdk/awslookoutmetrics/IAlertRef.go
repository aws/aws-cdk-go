package awslookoutmetrics

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslookoutmetrics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Alert.
// Experimental.
type IAlertRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAlertRef
type jsiiProxy_IAlertRef struct {
	internal.Type__constructsIConstruct
}

