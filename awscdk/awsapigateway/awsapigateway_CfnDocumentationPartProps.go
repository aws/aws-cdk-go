package awsapigateway


// Properties for defining a `CfnDocumentationPart`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDocumentationPartProps := &cfnDocumentationPartProps{
//   	location: &locationProperty{
//   		method: jsii.String("method"),
//   		name: jsii.String("name"),
//   		path: jsii.String("path"),
//   		statusCode: jsii.String("statusCode"),
//   		type: jsii.String("type"),
//   	},
//   	properties: jsii.String("properties"),
//   	restApiId: jsii.String("restApiId"),
//   }
//
type CfnDocumentationPartProps struct {
	// The location of the targeted API entity of the to-be-created documentation part.
	Location interface{} `field:"required" json:"location" yaml:"location"`
	// The new documentation content map of the targeted API entity.
	//
	// Enclosed key-value pairs are API-specific, but only OpenAPI-compliant key-value pairs can be exported and, hence, published.
	Properties *string `field:"required" json:"properties" yaml:"properties"`
	// The string identifier of the associated RestApi.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
}

