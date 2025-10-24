package awscdkpipestargetsalpha

import (
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
)

// API Gateway REST API target properties.
//
// Example:
//   var sourceQueue Queue
//
//
//   fn := lambda.NewFunction(this, jsii.String("MyFunc"), &FunctionProps{
//   	Handler: jsii.String("index.handler"),
//   	Runtime: lambda.Runtime_NODEJS_LATEST(),
//   	Code: lambda.Code_FromInline(jsii.String("exports.handler = e => {}")),
//   })
//
//   restApi := api.NewLambdaRestApi(this, jsii.String("MyRestAPI"), &LambdaRestApiProps{
//   	Handler: fn,
//   })
//   apiTarget := targets.NewApiGatewayTarget(restApi, &ApiGatewayTargetParameters{
//   	InputTransformation: pipes.InputTransformation_FromObject(map[string]interface{}{
//   		"body": jsii.String("👀"),
//   	}),
//   })
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
//   	Target: apiTarget,
//   })
//
// Experimental.
type ApiGatewayTargetParameters struct {
	// The headers to send as part of the request invoking the API Gateway REST API.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargethttpparameters.html#cfn-pipes-pipe-pipetargethttpparameters-headerparameters
	//
	// Default: - none.
	//
	// Experimental.
	HeaderParameters *map[string]*string `field:"optional" json:"headerParameters" yaml:"headerParameters"`
	// The input transformation to apply to the message before sending it to the target.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-inputtemplate
	//
	// Default: - none.
	//
	// Experimental.
	InputTransformation awscdkpipesalpha.IInputTransformation `field:"optional" json:"inputTransformation" yaml:"inputTransformation"`
	// The method for API Gateway resource.
	// Default: '*' - ANY.
	//
	// Experimental.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// The path for the API Gateway resource.
	// Default: '/'.
	//
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The path parameter values used to populate the API Gateway REST API path wildcards ("*").
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargethttpparameters.html#cfn-pipes-pipe-pipetargethttpparameters-pathparametervalues
	//
	// Default: - none.
	//
	// Experimental.
	PathParameterValues *[]*string `field:"optional" json:"pathParameterValues" yaml:"pathParameterValues"`
	// The query string keys/values that need to be sent as part of request invoking the API Gateway REST API.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargethttpparameters.html#cfn-pipes-pipe-pipetargethttpparameters-querystringparameters
	//
	// Default: - none.
	//
	// Experimental.
	QueryStringParameters *map[string]*string `field:"optional" json:"queryStringParameters" yaml:"queryStringParameters"`
	// The deployment stage for the API Gateway resource.
	// Default: - the value of `deploymentStage.stageName` of target API Gateway resource.
	//
	// Experimental.
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
}

