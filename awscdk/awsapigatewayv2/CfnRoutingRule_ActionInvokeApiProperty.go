package awsapigatewayv2


// Represents an InvokeApi action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionInvokeApiProperty := &ActionInvokeApiProperty{
//   	ApiId: jsii.String("apiId"),
//   	Stage: jsii.String("stage"),
//
//   	// the properties below are optional
//   	StripBasePath: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routingrule-actioninvokeapi.html
//
type CfnRoutingRule_ActionInvokeApiProperty struct {
	// The API identifier of the target API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routingrule-actioninvokeapi.html#cfn-apigatewayv2-routingrule-actioninvokeapi-apiid
	//
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The name of the target stage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routingrule-actioninvokeapi.html#cfn-apigatewayv2-routingrule-actioninvokeapi-stage
	//
	Stage *string `field:"required" json:"stage" yaml:"stage"`
	// The strip base path setting.
	//
	// When true, API Gateway strips the incoming matched base path when forwarding the request to the target API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routingrule-actioninvokeapi.html#cfn-apigatewayv2-routingrule-actioninvokeapi-stripbasepath
	//
	StripBasePath interface{} `field:"optional" json:"stripBasePath" yaml:"stripBasePath"`
}

