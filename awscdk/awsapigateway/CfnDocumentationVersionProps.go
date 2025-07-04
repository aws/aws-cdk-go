package awsapigateway


// Properties for defining a `CfnDocumentationVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDocumentationVersionProps := &CfnDocumentationVersionProps{
//   	DocumentationVersion: jsii.String("documentationVersion"),
//   	RestApiId: jsii.String("restApiId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-documentationversion.html
//
type CfnDocumentationVersionProps struct {
	// The version identifier of the to-be-updated documentation version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-documentationversion.html#cfn-apigateway-documentationversion-documentationversion
	//
	DocumentationVersion *string `field:"required" json:"documentationVersion" yaml:"documentationVersion"`
	// The string identifier of the associated RestApi.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-documentationversion.html#cfn-apigateway-documentationversion-restapiid
	//
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// A description about the new documentation snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-documentationversion.html#cfn-apigateway-documentationversion-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

