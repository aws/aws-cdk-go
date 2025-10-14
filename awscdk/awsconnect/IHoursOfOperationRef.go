package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a HoursOfOperation.
// Experimental.
type IHoursOfOperationRef interface {
	constructs.IConstruct
	// A reference to a HoursOfOperation resource.
	// Experimental.
	HoursOfOperationRef() *HoursOfOperationReference
}

// The jsii proxy for IHoursOfOperationRef
type jsiiProxy_IHoursOfOperationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IHoursOfOperationRef) HoursOfOperationRef() *HoursOfOperationReference {
	var returns *HoursOfOperationReference
	_jsii_.Get(
		j,
		"hoursOfOperationRef",
		&returns,
	)
	return returns
}

