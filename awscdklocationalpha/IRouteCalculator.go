package awscdklocationalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdklocationalpha/v2/internal"
)

// A Route Calculator.
// Experimental.
type IRouteCalculator interface {
	awscdk.IResource
	// The Amazon Resource Name (ARN) of the route calculator resource.
	// Experimental.
	RouteCalculatorArn() *string
	// The name of the route calculator.
	// Experimental.
	RouteCalculatorName() *string
}

// The jsii proxy for IRouteCalculator
type jsiiProxy_IRouteCalculator struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IRouteCalculator) RouteCalculatorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeCalculatorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouteCalculator) RouteCalculatorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeCalculatorName",
		&returns,
	)
	return returns
}

