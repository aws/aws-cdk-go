package awsapigateway


// Properties for defining a `CfnGatewayResponse`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGatewayResponseProps := &CfnGatewayResponseProps{
//   	ResponseType: jsii.String("responseType"),
//   	RestApiId: jsii.String("restApiId"),
//
//   	// the properties below are optional
//   	ResponseParameters: map[string]*string{
//   		"responseParametersKey": jsii.String("responseParameters"),
//   	},
//   	ResponseTemplates: map[string]*string{
//   		"responseTemplatesKey": jsii.String("responseTemplates"),
//   	},
//   	StatusCode: jsii.String("statusCode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-gatewayresponse.html
//
type CfnGatewayResponseProps struct {
	// The response type of the associated GatewayResponse.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-gatewayresponse.html#cfn-apigateway-gatewayresponse-responsetype
	//
	ResponseType *string `field:"required" json:"responseType" yaml:"responseType"`
	// The string identifier of the associated RestApi.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-gatewayresponse.html#cfn-apigateway-gatewayresponse-restapiid
	//
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// Response parameters (paths, query strings and headers) of the GatewayResponse as a string-to-string map of key-value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-gatewayresponse.html#cfn-apigateway-gatewayresponse-responseparameters
	//
	ResponseParameters interface{} `field:"optional" json:"responseParameters" yaml:"responseParameters"`
	// Response templates of the GatewayResponse as a string-to-string map of key-value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-gatewayresponse.html#cfn-apigateway-gatewayresponse-responsetemplates
	//
	ResponseTemplates interface{} `field:"optional" json:"responseTemplates" yaml:"responseTemplates"`
	// The HTTP status code for this GatewayResponse.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-gatewayresponse.html#cfn-apigateway-gatewayresponse-statuscode
	//
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
}

