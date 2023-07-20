package awsapigateway


// Properties for defining a `CfnDocumentationPart`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDocumentationPartProps := &CfnDocumentationPartProps{
//   	Location: &LocationProperty{
//   		Method: jsii.String("method"),
//   		Name: jsii.String("name"),
//   		Path: jsii.String("path"),
//   		StatusCode: jsii.String("statusCode"),
//   		Type: jsii.String("type"),
//   	},
//   	Properties: jsii.String("properties"),
//   	RestApiId: jsii.String("restApiId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-documentationpart.html
//
type CfnDocumentationPartProps struct {
	// The location of the targeted API entity of the to-be-created documentation part.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-documentationpart.html#cfn-apigateway-documentationpart-location
	//
	Location interface{} `field:"required" json:"location" yaml:"location"`
	// The new documentation content map of the targeted API entity.
	//
	// Enclosed key-value pairs are API-specific, but only OpenAPI-compliant key-value pairs can be exported and, hence, published.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-documentationpart.html#cfn-apigateway-documentationpart-properties
	//
	Properties *string `field:"required" json:"properties" yaml:"properties"`
	// The string identifier of the associated RestApi.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-documentationpart.html#cfn-apigateway-documentationpart-restapiid
	//
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
}

