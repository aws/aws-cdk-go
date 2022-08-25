package awsappsync


// Properties for defining a `CfnGraphQLSchema`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGraphQLSchemaProps := &cfnGraphQLSchemaProps{
//   	apiId: jsii.String("apiId"),
//
//   	// the properties below are optional
//   	definition: jsii.String("definition"),
//   	definitionS3Location: jsii.String("definitionS3Location"),
//   }
//
type CfnGraphQLSchemaProps struct {
	// The AWS AppSync GraphQL API identifier to which you want to apply this schema.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The text representation of a GraphQL schema in SDL format.
	//
	// For more information about using the `Ref` function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref) .
	Definition *string `field:"optional" json:"definition" yaml:"definition"`
	// The location of a GraphQL schema file in an Amazon S3 bucket.
	//
	// Use this if you want to provision with the schema living in Amazon S3 rather than embedding it in your CloudFormation template.
	DefinitionS3Location *string `field:"optional" json:"definitionS3Location" yaml:"definitionS3Location"`
}

