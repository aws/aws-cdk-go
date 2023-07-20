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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html
//
type CfnResourceProps struct {
	// The parent resource's identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html#cfn-apigateway-resource-parentid
	//
	ParentId *string `field:"required" json:"parentId" yaml:"parentId"`
	// The last path segment for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html#cfn-apigateway-resource-pathpart
	//
	PathPart *string `field:"required" json:"pathPart" yaml:"pathPart"`
	// The string identifier of the associated RestApi.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html#cfn-apigateway-resource-restapiid
	//
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
}

