package awscloudtrail

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Util element for InsightSelector.
//
// Example:
//   cloudtrail.NewTrail(this, jsii.String("Insights"), &TrailProps{
//   	InsightTypes: []insightType{
//   		cloudtrail.*insightType_API_CALL_RATE(),
//   		cloudtrail.*insightType_API_ERROR_RATE(),
//   	},
//   })
//
type InsightType interface {
	Value() *string
}

// The jsii proxy struct for InsightType
type jsiiProxy_InsightType struct {
	_ byte // padding
}

func (j *jsiiProxy_InsightType) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NewInsightType(value *string) InsightType {
	_init_.Initialize()

	if err := validateNewInsightTypeParameters(value); err != nil {
		panic(err)
	}
	j := jsiiProxy_InsightType{}

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudtrail.InsightType",
		[]interface{}{value},
		&j,
	)

	return &j
}

func NewInsightType_Override(i InsightType, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_cloudtrail.InsightType",
		[]interface{}{value},
		i,
	)
}

func InsightType_API_CALL_RATE() InsightType {
	_init_.Initialize()
	var returns InsightType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudtrail.InsightType",
		"API_CALL_RATE",
		&returns,
	)
	return returns
}

func InsightType_API_ERROR_RATE() InsightType {
	_init_.Initialize()
	var returns InsightType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_cloudtrail.InsightType",
		"API_ERROR_RATE",
		&returns,
	)
	return returns
}

