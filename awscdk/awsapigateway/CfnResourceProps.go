package awsapigateway


// Properties for defining a `CfnResource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceProps := &CfnResourceProps{
//   	ParentId: jsii.String("parentId"),
//   	PathPart: jsii.String("pathPart"),
//   	RestApiId: jsii.String("restApiId"),
//   }
//
type CfnResourceProps struct {
	// The parent resource's identifier.
	ParentId *string `field:"required" json:"parentId" yaml:"parentId"`
	// The last path segment for this resource.
	PathPart *string `field:"required" json:"pathPart" yaml:"pathPart"`
	// The string identifier of the associated RestApi.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
}

