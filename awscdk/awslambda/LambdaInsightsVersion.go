package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Version of CloudWatch Lambda Insights.
//
// Example:
//   layerArn := "arn:aws:lambda:us-east-1:580247275435:layer:LambdaInsightsExtension:14"
//   lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_18_X(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   	InsightsVersion: lambda.LambdaInsightsVersion_FromInsightVersionArn(layerArn),
//   })
//
type LambdaInsightsVersion interface {
	// The arn of the Lambda Insights extension.
	LayerVersionArn() *string
}

// The jsii proxy struct for LambdaInsightsVersion
type jsiiProxy_LambdaInsightsVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_LambdaInsightsVersion) LayerVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layerVersionArn",
		&returns,
	)
	return returns
}


func NewLambdaInsightsVersion_Override(l LambdaInsightsVersion) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.LambdaInsightsVersion",
		nil, // no parameters
		l,
	)
}

// Use the insights extension associated with the provided ARN.
//
// Make sure the ARN is associated
// with same region as your function.
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Lambda-Insights-extension-versions.html
//
func LambdaInsightsVersion_FromInsightVersionArn(arn *string) LambdaInsightsVersion {
	_init_.Initialize()

	if err := validateLambdaInsightsVersion_FromInsightVersionArnParameters(arn); err != nil {
		panic(err)
	}
	var returns LambdaInsightsVersion

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.LambdaInsightsVersion",
		"fromInsightVersionArn",
		[]interface{}{arn},
		&returns,
	)

	return returns
}

func LambdaInsightsVersion_VERSION_1_0_119_0() LambdaInsightsVersion {
	_init_.Initialize()
	var returns LambdaInsightsVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.LambdaInsightsVersion",
		"VERSION_1_0_119_0",
		&returns,
	)
	return returns
}

func LambdaInsightsVersion_VERSION_1_0_135_0() LambdaInsightsVersion {
	_init_.Initialize()
	var returns LambdaInsightsVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.LambdaInsightsVersion",
		"VERSION_1_0_135_0",
		&returns,
	)
	return returns
}

func LambdaInsightsVersion_VERSION_1_0_143_0() LambdaInsightsVersion {
	_init_.Initialize()
	var returns LambdaInsightsVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.LambdaInsightsVersion",
		"VERSION_1_0_143_0",
		&returns,
	)
	return returns
}

func LambdaInsightsVersion_VERSION_1_0_178_0() LambdaInsightsVersion {
	_init_.Initialize()
	var returns LambdaInsightsVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.LambdaInsightsVersion",
		"VERSION_1_0_178_0",
		&returns,
	)
	return returns
}

func LambdaInsightsVersion_VERSION_1_0_229_0() LambdaInsightsVersion {
	_init_.Initialize()
	var returns LambdaInsightsVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.LambdaInsightsVersion",
		"VERSION_1_0_229_0",
		&returns,
	)
	return returns
}

func LambdaInsightsVersion_VERSION_1_0_54_0() LambdaInsightsVersion {
	_init_.Initialize()
	var returns LambdaInsightsVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.LambdaInsightsVersion",
		"VERSION_1_0_54_0",
		&returns,
	)
	return returns
}

func LambdaInsightsVersion_VERSION_1_0_86_0() LambdaInsightsVersion {
	_init_.Initialize()
	var returns LambdaInsightsVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.LambdaInsightsVersion",
		"VERSION_1_0_86_0",
		&returns,
	)
	return returns
}

func LambdaInsightsVersion_VERSION_1_0_89_0() LambdaInsightsVersion {
	_init_.Initialize()
	var returns LambdaInsightsVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.LambdaInsightsVersion",
		"VERSION_1_0_89_0",
		&returns,
	)
	return returns
}

func LambdaInsightsVersion_VERSION_1_0_98_0() LambdaInsightsVersion {
	_init_.Initialize()
	var returns LambdaInsightsVersion
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.LambdaInsightsVersion",
		"VERSION_1_0_98_0",
		&returns,
	)
	return returns
}

