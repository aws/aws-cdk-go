package awsapigateway


// `IntegrationResponse` is a property of the [Amazon API Gateway Method Integration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html) property type that specifies the response that API Gateway sends after a method's backend finishes processing a request.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integrationResponseProperty := &IntegrationResponseProperty{
//   	StatusCode: jsii.String("statusCode"),
//
//   	// the properties below are optional
//   	ContentHandling: jsii.String("contentHandling"),
//   	ResponseParameters: map[string]*string{
//   		"responseParametersKey": jsii.String("responseParameters"),
//   	},
//   	ResponseTemplates: map[string]*string{
//   		"responseTemplatesKey": jsii.String("responseTemplates"),
//   	},
//   	SelectionPattern: jsii.String("selectionPattern"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-method-integrationresponse.html
//
type CfnMethod_IntegrationResponseProperty struct {
	// Specifies the status code that is used to map the integration response to an existing MethodResponse.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-method-integrationresponse.html#cfn-apigateway-method-integrationresponse-statuscode
	//
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
	// Specifies how to handle response payload content type conversions.
	//
	// Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT` , with the following behaviors:
	//
	// If this property is not defined, the response payload will be passed through from the integration response to the method response without modification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-method-integrationresponse.html#cfn-apigateway-method-integrationresponse-contenthandling
	//
	ContentHandling *string `field:"optional" json:"contentHandling" yaml:"contentHandling"`
	// A key-value map specifying response parameters that are passed to the method response from the back end.
	//
	// The key is a method response header parameter name and the mapped value is an integration response header value, a static value enclosed within a pair of single quotes, or a JSON expression from the integration response body. The mapping key must match the pattern of `method.response.header.{name}` , where `name` is a valid and unique header name. The mapped non-static value must match the pattern of `integration.response.header.{name}` or `integration.response.body.{JSON-expression}` , where `name` is a valid and unique response header name and `JSON-expression` is a valid JSON expression without the `$` prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-method-integrationresponse.html#cfn-apigateway-method-integrationresponse-responseparameters
	//
	ResponseParameters interface{} `field:"optional" json:"responseParameters" yaml:"responseParameters"`
	// Specifies the templates used to transform the integration response body.
	//
	// Response templates are represented as a key/value map, with a content-type as the key and a template as the value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-method-integrationresponse.html#cfn-apigateway-method-integrationresponse-responsetemplates
	//
	ResponseTemplates interface{} `field:"optional" json:"responseTemplates" yaml:"responseTemplates"`
	// Specifies the regular expression (regex) pattern used to choose an integration response based on the response from the back end.
	//
	// For example, if the success response returns nothing and the error response returns some string, you could use the `.+` regex to match error response. However, make sure that the error response does not contain any newline ( `\n` ) character in such cases. If the back end is an AWS Lambda function, the AWS Lambda function error header is matched. For all other HTTP and AWS back ends, the HTTP status code is matched.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-method-integrationresponse.html#cfn-apigateway-method-integrationresponse-selectionpattern
	//
	SelectionPattern *string `field:"optional" json:"selectionPattern" yaml:"selectionPattern"`
}

