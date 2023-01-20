package awsapigateway


// Properties for defining a `CfnDocumentationVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDocumentationVersionProps := &cfnDocumentationVersionProps{
//   	documentationVersion: jsii.String("documentationVersion"),
//   	restApiId: jsii.String("restApiId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   }
//
type CfnDocumentationVersionProps struct {
	// The version identifier of the API documentation snapshot.
	DocumentationVersion *string `field:"required" json:"documentationVersion" yaml:"documentationVersion"`
	// The identifier of the API.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// The description of the API documentation snapshot.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

