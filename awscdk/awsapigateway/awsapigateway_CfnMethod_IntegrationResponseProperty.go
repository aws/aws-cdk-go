package awsapigateway


// `IntegrationResponse` is a property of the [Amazon API Gateway Method Integration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html) property type that specifies the response that API Gateway sends after a method's backend finishes processing a request.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   integrationResponseProperty := &integrationResponseProperty{
//   	statusCode: jsii.String("statusCode"),
//
//   	// the properties below are optional
//   	contentHandling: jsii.String("contentHandling"),
//   	responseParameters: map[string]*string{
//   		"responseParametersKey": jsii.String("responseParameters"),
//   	},
//   	responseTemplates: map[string]*string{
//   		"responseTemplatesKey": jsii.String("responseTemplates"),
//   	},
//   	selectionPattern: jsii.String("selectionPattern"),
//   }
//
type CfnMethod_IntegrationResponseProperty struct {
	// The status code that API Gateway uses to map the integration response to a [MethodResponse](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-methodresponse.html) status code.
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
	// Specifies how to handle request payload content type conversions. Valid values are:.
	//
	// - `CONVERT_TO_BINARY` : Converts a request payload from a base64-encoded string to a binary blob.
	// - `CONVERT_TO_TEXT` : Converts a request payload from a binary blob to a base64-encoded string.
	//
	// If this property isn't defined, the request payload is passed through from the method request to the integration request without modification.
	ContentHandling *string `field:"optional" json:"contentHandling" yaml:"contentHandling"`
	// The response parameters from the backend response that API Gateway sends to the method response.
	//
	// Specify response parameters as key-value pairs ( [string-to-string mappings](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/mappings-section-structure.html) ).
	//
	// Use the destination as the key and the source as the value:
	//
	// - The destination must be an existing response parameter in the [MethodResponse](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-methodresponse.html) property.
	// - The source must be an existing method request parameter or a static value. You must enclose static values in single quotation marks and pre-encode these values based on the destination specified in the request.
	//
	// For more information about templates, see [API Gateway Mapping Template and Access Logging Variable Reference](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html) in the *API Gateway Developer Guide* .
	ResponseParameters interface{} `field:"optional" json:"responseParameters" yaml:"responseParameters"`
	// The templates that are used to transform the integration response body.
	//
	// Specify templates as key-value pairs (string-to-string mappings), with a content type as the key and a template as the value. For more information, see [API Gateway Mapping Template and Access Logging Variable Reference](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html) in the *API Gateway Developer Guide* .
	ResponseTemplates interface{} `field:"optional" json:"responseTemplates" yaml:"responseTemplates"`
	// A [regular expression](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-regexes.html) that specifies which error strings or status codes from the backend map to the integration response.
	SelectionPattern *string `field:"optional" json:"selectionPattern" yaml:"selectionPattern"`
}

