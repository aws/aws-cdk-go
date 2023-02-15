package awsapigateway


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integrationResponse := &integrationResponse{
//   	statusCode: jsii.String("statusCode"),
//
//   	// the properties below are optional
//   	contentHandling: awscdk.Aws_apigateway.contentHandling_CONVERT_TO_BINARY,
//   	responseParameters: map[string]*string{
//   		"responseParametersKey": jsii.String("responseParameters"),
//   	},
//   	responseTemplates: map[string]*string{
//   		"responseTemplatesKey": jsii.String("responseTemplates"),
//   	},
//   	selectionPattern: jsii.String("selectionPattern"),
//   }
//
type IntegrationResponse struct {
	// The status code that API Gateway uses to map the integration response to a MethodResponse status code.
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
	// Specifies how to handle request payload content type conversions.
	ContentHandling ContentHandling `field:"optional" json:"contentHandling" yaml:"contentHandling"`
	// The response parameters from the backend response that API Gateway sends to the method response.
	//
	// Use the destination as the key and the source as the value:
	//
	// - The destination must be an existing response parameter in the
	//    MethodResponse property.
	// - The source must be an existing method request parameter or a static
	//    value. You must enclose static values in single quotation marks and
	//    pre-encode these values based on the destination specified in the
	//    request.
	// See: http://docs.aws.amazon.com/apigateway/latest/developerguide/request-response-data-mappings.html
	//
	ResponseParameters *map[string]*string `field:"optional" json:"responseParameters" yaml:"responseParameters"`
	// The templates that are used to transform the integration response body.
	//
	// Specify templates as key-value pairs, with a content type as the key and
	// a template as the value.
	// See: http://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html
	//
	ResponseTemplates *map[string]*string `field:"optional" json:"responseTemplates" yaml:"responseTemplates"`
	// Specifies the regular expression (regex) pattern used to choose an integration response based on the response from the back end.
	//
	// For example, if the success response returns nothing and the error response returns some string, you
	// could use the ``.+`` regex to match error response. However, make sure that the error response does not contain any
	// newline (``\n``) character in such cases. If the back end is an AWS Lambda function, the AWS Lambda function error
	// header is matched. For all other HTTP and AWS back ends, the HTTP status code is matched.
	// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-integration-settings-integration-response.html
	//
	SelectionPattern *string `field:"optional" json:"selectionPattern" yaml:"selectionPattern"`
}

