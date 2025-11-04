package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a HoursOfOperation.
// Experimental.
type IHoursOfOperationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a HoursOfOperation resource.
	// Experimental.
	HoursOfOperationRef() *HoursOfOperationReference
}

// The jsii proxy for IHoursOfOperationRef
type jsiiProxy_IHoursOfOperationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IHoursOfOperationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHoursOfOperationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

