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
	// The location of the API entity that the documentation applies to.
	Location interface{} `field:"required" json:"location" yaml:"location"`
	// The documentation content map of the targeted API entity.
	Properties *string `field:"required" json:"properties" yaml:"properties"`
	// The identifier of the targeted API entity.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
}

