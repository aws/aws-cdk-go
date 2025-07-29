package awsapigatewayv2


// Represents a routing rule action.
//
// The only supported action is `invokeApi` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &ActionProperty{
//   	InvokeApi: &ActionInvokeApiProperty{
//   		ApiId: jsii.String("apiId"),
//   		Stage: jsii.String("stage"),
//
//   		// the properties below are optional
//   		StripBasePath: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routingrule-action.html
//
type CfnRoutingRule_ActionProperty struct {
	// Represents an InvokeApi action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-routingrule-action.html#cfn-apigatewayv2-routingrule-action-invokeapi
	//
	InvokeApi interface{} `field:"required" json:"invokeApi" yaml:"invokeApi"`
}

