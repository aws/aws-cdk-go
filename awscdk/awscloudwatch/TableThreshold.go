package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Thresholds for highlighting cells in TableWidget.
//
// Example:
//   var dashboard dashboard
//
//
//   dashboard.AddWidgets(cloudwatch.NewTableWidget(&TableWidgetProps{
//   	// ...
//
//   	Thresholds: []tableThreshold{
//   		cloudwatch.*tableThreshold_Above(jsii.Number(1000), cloudwatch.Color_RED()),
//   		cloudwatch.*tableThreshold_Between(jsii.Number(500), jsii.Number(1000), cloudwatch.Color_ORANGE()),
//   		cloudwatch.*tableThreshold_Below(jsii.Number(500), cloudwatch.Color_GREEN()),
//   	},
//   }))
//
type TableThreshold interface {
	ToJson() interface{}
}

// The jsii proxy struct for TableThreshold
type jsiiProxy_TableThreshold struct {
	_ byte // padding
}

// A threshold for highlighting and coloring cells above the specified value.
func TableThreshold_Above(value *float64, color *string) TableThreshold {
	_init_.Initialize()

	if err := validateTableThreshold_AboveParameters(value); err != nil {
		panic(err)
	}
	var returns TableThreshold

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.TableThreshold",
		"above",
		[]interface{}{value, color},
		&returns,
	)

	return returns
}

// A threshold for highlighting and coloring cells below the specified value.
func TableThreshold_Below(value *float64, color *string) TableThreshold {
	_init_.Initialize()

	if err := validateTableThreshold_BelowParameters(value); err != nil {
		panic(err)
	}
	var returns TableThreshold

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.TableThreshold",
		"below",
		[]interface{}{value, color},
		&returns,
	)

	return returns
}

// A threshold for highlighting and coloring cells within the specified values.
func TableThreshold_Between(lowerBound *float64, upperBound *float64, color *string) TableThreshold {
	_init_.Initialize()

	if err := validateTableThreshold_BetweenParameters(lowerBound, upperBound); err != nil {
		panic(err)
	}
	var returns TableThreshold

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_cloudwatch.TableThreshold",
		"between",
		[]interface{}{lowerBound, upperBound, color},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TableThreshold) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

