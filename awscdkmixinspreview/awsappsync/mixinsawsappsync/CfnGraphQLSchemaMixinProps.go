package mixinsawsappsync


// Properties for CfnGraphQLSchemaPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGraphQLSchemaMixinProps := &CfnGraphQLSchemaMixinProps{
//   	ApiId: jsii.String("apiId"),
//   	Definition: jsii.String("definition"),
//   	DefinitionS3Location: jsii.String("definitionS3Location"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlschema.html
//
type CfnGraphQLSchemaMixinProps struct {
	// The AWS AppSync GraphQL API identifier to which you want to apply this schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlschema.html#cfn-appsync-graphqlschema-apiid
	//
	ApiId *string `field:"optional" json:"apiId" yaml:"apiId"`
	// The text representation of a GraphQL schema in SDL format.
	//
	// For more information about using the `Ref` function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlschema.html#cfn-appsync-graphqlschema-definition
	//
	Definition *string `field:"optional" json:"definition" yaml:"definition"`
	// The location of a GraphQL schema file in an Amazon S3 bucket.
	//
	// Use this if you want to provision with the schema living in Amazon S3 rather than embedding it in your CloudFormation template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlschema.html#cfn-appsync-graphqlschema-definitions3location
	//
	DefinitionS3Location *string `field:"optional" json:"definitionS3Location" yaml:"definitionS3Location"`
}

