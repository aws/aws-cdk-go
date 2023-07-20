package awsapigateway


// The `Location` property specifies the location of the Amazon API Gateway API entity that the documentation applies to.
//
// `Location` is a property of the [AWS::ApiGateway::DocumentationPart](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-documentationpart.html) resource.
//
// > For more information about each property, including constraints and valid values, see [DocumentationPart](https://docs.aws.amazon.com/apigateway/latest/api/API_DocumentationPartLocation.html) in the *Amazon API Gateway REST API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationProperty := &LocationProperty{
//   	Method: jsii.String("method"),
//   	Name: jsii.String("name"),
//   	Path: jsii.String("path"),
//   	StatusCode: jsii.String("statusCode"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-documentationpart-location.html
//
type CfnDocumentationPart_LocationProperty struct {
	// The HTTP verb of a method.
	//
	// It is a valid field for the API entity types of `METHOD` , `PATH_PARAMETER` , `QUERY_PARAMETER` , `REQUEST_HEADER` , `REQUEST_BODY` , `RESPONSE` , `RESPONSE_HEADER` , and `RESPONSE_BODY` . The default value is `*` for any method. When an applicable child entity inherits the content of an entity of the same type with more general specifications of the other `location` attributes, the child entity's `method` attribute must match that of the parent entity exactly.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-documentationpart-location.html#cfn-apigateway-documentationpart-location-method
	//
	Method *string `field:"optional" json:"method" yaml:"method"`
	// The name of the targeted API entity.
	//
	// It is a valid and required field for the API entity types of `AUTHORIZER` , `MODEL` , `PATH_PARAMETER` , `QUERY_PARAMETER` , `REQUEST_HEADER` , `REQUEST_BODY` and `RESPONSE_HEADER` . It is an invalid field for any other entity type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-documentationpart-location.html#cfn-apigateway-documentationpart-location-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The URL path of the target.
	//
	// It is a valid field for the API entity types of `RESOURCE` , `METHOD` , `PATH_PARAMETER` , `QUERY_PARAMETER` , `REQUEST_HEADER` , `REQUEST_BODY` , `RESPONSE` , `RESPONSE_HEADER` , and `RESPONSE_BODY` . The default value is `/` for the root resource. When an applicable child entity inherits the content of another entity of the same type with more general specifications of the other `location` attributes, the child entity's `path` attribute must match that of the parent entity as a prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-documentationpart-location.html#cfn-apigateway-documentationpart-location-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The HTTP status code of a response.
	//
	// It is a valid field for the API entity types of `RESPONSE` , `RESPONSE_HEADER` , and `RESPONSE_BODY` . The default value is `*` for any status code. When an applicable child entity inherits the content of an entity of the same type with more general specifications of the other `location` attributes, the child entity's `statusCode` attribute must match that of the parent entity exactly.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-documentationpart-location.html#cfn-apigateway-documentationpart-location-statuscode
	//
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
	// The type of API entity to which the documentation content applies.
	//
	// Valid values are `API` , `AUTHORIZER` , `MODEL` , `RESOURCE` , `METHOD` , `PATH_PARAMETER` , `QUERY_PARAMETER` , `REQUEST_HEADER` , `REQUEST_BODY` , `RESPONSE` , `RESPONSE_HEADER` , and `RESPONSE_BODY` . Content inheritance does not apply to any entity of the `API` , `AUTHORIZER` , `METHOD` , `MODEL` , `REQUEST_BODY` , or `RESOURCE` type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-documentationpart-location.html#cfn-apigateway-documentationpart-location-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

