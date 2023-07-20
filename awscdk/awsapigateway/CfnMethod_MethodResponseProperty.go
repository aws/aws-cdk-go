package awsapigateway


// Represents a method response of a given HTTP status code returned to the client.
//
// The method response is passed from the back end through the associated integration response that can be transformed using a mapping template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   methodResponseProperty := &MethodResponseProperty{
//   	StatusCode: jsii.String("statusCode"),
//
//   	// the properties below are optional
//   	ResponseModels: map[string]*string{
//   		"responseModelsKey": jsii.String("responseModels"),
//   	},
//   	ResponseParameters: map[string]interface{}{
//   		"responseParametersKey": jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-method-methodresponse.html
//
type CfnMethod_MethodResponseProperty struct {
	// The method response's status code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-method-methodresponse.html#cfn-apigateway-method-methodresponse-statuscode
	//
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
	// Specifies the Model resources used for the response's content-type.
	//
	// Response models are represented as a key/value map, with a content-type as the key and a Model name as the value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-method-methodresponse.html#cfn-apigateway-method-methodresponse-responsemodels
	//
	ResponseModels interface{} `field:"optional" json:"responseModels" yaml:"responseModels"`
	// A key-value map specifying required or optional response parameters that API Gateway can send back to the caller.
	//
	// A key defines a method response header and the value specifies whether the associated method response header is required or not. The expression of the key must match the pattern `method.response.header.{name}` , where `name` is a valid and unique header name. API Gateway passes certain integration response data to the method response headers specified here according to the mapping you prescribe in the API's IntegrationResponse. The integration response data that can be mapped include an integration response header expressed in `integration.response.header.{name}` , a static value enclosed within a pair of single quotes (e.g., `'application/json'` ), or a JSON expression from the back-end response payload in the form of `integration.response.body.{JSON-expression}` , where `JSON-expression` is a valid JSON expression without the `$` prefix.)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-method-methodresponse.html#cfn-apigateway-method-methodresponse-responseparameters
	//
	ResponseParameters interface{} `field:"optional" json:"responseParameters" yaml:"responseParameters"`
}

