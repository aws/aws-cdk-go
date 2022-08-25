package awsapigateway


// Properties for defining a `CfnResource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceProps := &cfnResourceProps{
//   	parentId: jsii.String("parentId"),
//   	pathPart: jsii.String("pathPart"),
//   	restApiId: jsii.String("restApiId"),
//   }
//
type CfnResourceProps struct {
	// If you want to create a child resource, the ID of the parent resource.
	//
	// For resources without a parent, specify the `RestApi` root resource ID, such as `{ "Fn::GetAtt": ["MyRestApi", "RootResourceId"] }` .
	ParentId *string `field:"required" json:"parentId" yaml:"parentId"`
	// A path name for the resource.
	PathPart *string `field:"required" json:"pathPart" yaml:"pathPart"`
	// The ID of the [RestApi](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html) resource in which you want to create this resource.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
}

