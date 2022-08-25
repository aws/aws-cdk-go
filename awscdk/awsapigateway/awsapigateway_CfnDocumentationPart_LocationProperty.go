package awsapigateway


// The `Location` property specifies the location of the Amazon API Gateway API entity that the documentation applies to.
//
// `Location` is a property of the [AWS::ApiGateway::DocumentationPart](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-documentationpart.html) resource.
//
// > For more information about each property, including constraints and valid values, see [DocumentationPart](https://docs.aws.amazon.com/apigateway/api-reference/resource/documentation-part/#location) in the *Amazon API Gateway REST API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationProperty := &locationProperty{
//   	method: jsii.String("method"),
//   	name: jsii.String("name"),
//   	path: jsii.String("path"),
//   	statusCode: jsii.String("statusCode"),
//   	type: jsii.String("type"),
//   }
//
type CfnDocumentationPart_LocationProperty struct {
	// The HTTP verb of a method.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// The name of the targeted API entity.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The URL path of the target.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The HTTP status code of a response.
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
	// The type of API entity that the documentation content applies to.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

