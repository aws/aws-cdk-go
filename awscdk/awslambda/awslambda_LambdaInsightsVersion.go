package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Version of CloudWatch Lambda Insights.
//
// Example:
//   layerArn := "arn:aws:lambda:us-east-1:580247275435:layer:LambdaInsightsExtension:14"
//   lambda.NewFunction(this, jsii.String("MyFunction"), &functionProps{
//   	runtime: lambda.runtime_NODEJS_16_X(),
//   	handler: jsii.String("index.handler"),
//   	code: lambda.code.fromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   	insightsVersion: lambda.lambdaInsightsVersion.fromInsightVersionArn(layerArn),
//   })
//
// Experimental.
type LambdaInsightsVersion interface {
	// The arn of the Lambda Insights extension.
	// Experimental.
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


// Experimental.
func NewLambdaInsightsVersion_Override(l LambdaInsightsVersion) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_lambda.LambdaInsightsVersion",
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
// Experimental.
func LambdaInsightsVersion_FromInsightVersionArn(arn *string) LambdaInsightsVersion {
	_init_.Initialize()

	if err := validateLambdaInsightsVersion_FromInsightVersionArnParameters(arn); err != nil {
		panic(err)
	}
	var returns LambdaInsightsVersion

	_jsii_.StaticInvoke(
		"monocdk.aws_lambda.LambdaInsightsVersion",
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
		"monocdk.aws_lambda.LambdaInsightsVersion",
		"VERSION_1_0_119_0",
		&returns,
	)
	return returns
}

func LambdaInsightsVersion_VERSION_1_0_135_0() LambdaInsightsVersion {
	_init_.Initialize()
	var returns LambdaInsightsVersion
	_jsii_.StaticGet(
		"monocdk.aws_lambda.LambdaInsightsVersion",
		"VERSION_1_0_135_0",
		&returns,
	)
	return returns
}

func LambdaInsightsVersion_VERSION_1_0_54_0() LambdaInsightsVersion {
	_init_.Initialize()
	var returns LambdaInsightsVersion
	_jsii_.StaticGet(
		"monocdk.aws_lambda.LambdaInsightsVersion",
		"VERSION_1_0_54_0",
		&returns,
	)
	return returns
}

func LambdaInsightsVersion_VERSION_1_0_86_0() LambdaInsightsVersion {
	_init_.Initialize()
	var returns LambdaInsightsVersion
	_jsii_.StaticGet(
		"monocdk.aws_lambda.LambdaInsightsVersion",
		"VERSION_1_0_86_0",
		&returns,
	)
	return returns
}

func LambdaInsightsVersion_VERSION_1_0_89_0() LambdaInsightsVersion {
	_init_.Initialize()
	var returns LambdaInsightsVersion
	_jsii_.StaticGet(
		"monocdk.aws_lambda.LambdaInsightsVersion",
		"VERSION_1_0_89_0",
		&returns,
	)
	return returns
}

func LambdaInsightsVersion_VERSION_1_0_98_0() LambdaInsightsVersion {
	_init_.Initialize()
	var returns LambdaInsightsVersion
	_jsii_.StaticGet(
		"monocdk.aws_lambda.LambdaInsightsVersion",
		"VERSION_1_0_98_0",
		&returns,
	)
	return returns
}

