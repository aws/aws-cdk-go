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
	// The version identifier of the to-be-updated documentation version.
	DocumentationVersion *string `field:"required" json:"documentationVersion" yaml:"documentationVersion"`
	// The string identifier of the associated RestApi.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// A description about the new documentation snapshot.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

