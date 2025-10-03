package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a HoursOfOperation.
// Experimental.
type IHoursOfOperationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IHoursOfOperationRef
type jsiiProxy_IHoursOfOperationRef struct {
	internal.Type__constructsIConstruct
}

