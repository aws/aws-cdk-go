package awsapigatewayv2


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routingrule-actioninvokeapi.html#cfn-apigatewayv2-routingrule-actioninvokeapi-apiid
	//
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routingrule-actioninvokeapi.html#cfn-apigatewayv2-routingrule-actioninvokeapi-stage
	//
	Stage *string `field:"required" json:"stage" yaml:"stage"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routingrule-actioninvokeapi.html#cfn-apigatewayv2-routingrule-actioninvokeapi-stripbasepath
	//
	StripBasePath interface{} `field:"optional" json:"stripBasePath" yaml:"stripBasePath"`
}

