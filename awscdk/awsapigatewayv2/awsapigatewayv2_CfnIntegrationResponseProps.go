package awsapigatewayv2


// Properties for defining a `CfnIntegrationResponse`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var responseParameters interface{}
//   var responseTemplates interface{}
//
//   cfnIntegrationResponseProps := &cfnIntegrationResponseProps{
//   	apiId: jsii.String("apiId"),
//   	integrationId: jsii.String("integrationId"),
//   	integrationResponseKey: jsii.String("integrationResponseKey"),
//
//   	// the properties below are optional
//   	contentHandlingStrategy: jsii.String("contentHandlingStrategy"),
//   	responseParameters: responseParameters,
//   	responseTemplates: responseTemplates,
//   	templateSelectionExpression: jsii.String("templateSelectionExpression"),
//   }
//
type CfnIntegrationResponseProps struct {
	// The API identifier.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The integration ID.
	IntegrationId *string `field:"required" json:"integrationId" yaml:"integrationId"`
	// The integration response key.
	IntegrationResponseKey *string `field:"required" json:"integrationResponseKey" yaml:"integrationResponseKey"`
	// Supported only for WebSocket APIs.
	//
	// Specifies how to handle response payload content type conversions. Supported values are `CONVERT_TO_BINARY` and `CONVERT_TO_TEXT` , with the following behaviors:
	//
	// `CONVERT_TO_BINARY` : Converts a response payload from a Base64-encoded string to the corresponding binary blob.
	//
	// `CONVERT_TO_TEXT` : Converts a response payload from a binary blob to a Base64-encoded string.
	//
	// If this property is not defined, the response payload will be passed through from the integration response to the route response or method response without modification.
	ContentHandlingStrategy *string `field:"optional" json:"contentHandlingStrategy" yaml:"contentHandlingStrategy"`
	// A key-value map specifying response parameters that are passed to the method response from the backend.
	//
	// The key is a method response header parameter name and the mapped value is an integration response header value, a static value enclosed within a pair of single quotes, or a JSON expression from the integration response body. The mapping key must match the pattern of `method.response.header. *{name}*` , where name is a valid and unique header name. The mapped non-static value must match the pattern of `integration.response.header. *{name}*` or `integration.response.body. *{JSON-expression}*` , where `*{name}*` is a valid and unique response header name and `*{JSON-expression}*` is a valid JSON expression without the `$` prefix.
	ResponseParameters interface{} `field:"optional" json:"responseParameters" yaml:"responseParameters"`
	// The collection of response templates for the integration response as a string-to-string map of key-value pairs.
	//
	// Response templates are represented as a key/value map, with a content-type as the key and a template as the value.
	ResponseTemplates interface{} `field:"optional" json:"responseTemplates" yaml:"responseTemplates"`
	// The template selection expression for the integration response.
	//
	// Supported only for WebSocket APIs.
	TemplateSelectionExpression *string `field:"optional" json:"templateSelectionExpression" yaml:"templateSelectionExpression"`
}

