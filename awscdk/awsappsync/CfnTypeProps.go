package awsappsync


// Properties for defining a `CfnType`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTypeProps := &CfnTypeProps{
//   	ApiId: jsii.String("apiId"),
//   	Definition: jsii.String("definition"),
//   	Format: jsii.String("format"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-type.html
//
type CfnTypeProps struct {
	// The API ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-type.html#cfn-appsync-type-apiid
	//
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The type definition, in GraphQL Schema Definition Language (SDL) format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-type.html#cfn-appsync-type-definition
	//
	Definition *string `field:"required" json:"definition" yaml:"definition"`
	// The type format: SDL or JSON.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-type.html#cfn-appsync-type-format
	//
	Format *string `field:"required" json:"format" yaml:"format"`
}

