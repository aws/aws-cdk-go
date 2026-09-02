package awsappsync


// Properties for CfnTypePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnTypeMixinProps := &CfnTypeMixinProps{
//   	ApiId: jsii.String("apiId"),
//   	Definition: jsii.String("definition"),
//   	Format: jsii.String("format"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-type.html
//
type CfnTypeMixinProps struct {
	// The API ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-type.html#cfn-appsync-type-apiid
	//
	ApiId *string `field:"optional" json:"apiId" yaml:"apiId"`
	// The type definition, in GraphQL Schema Definition Language (SDL) format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-type.html#cfn-appsync-type-definition
	//
	Definition *string `field:"optional" json:"definition" yaml:"definition"`
	// The type format: SDL or JSON.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-type.html#cfn-appsync-type-format
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
}

