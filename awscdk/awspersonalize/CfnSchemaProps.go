package awspersonalize


// Properties for defining a `CfnSchema`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSchemaProps := &CfnSchemaProps{
//   	Name: jsii.String("name"),
//   	Schema: jsii.String("schema"),
//
//   	// the properties below are optional
//   	Domain: jsii.String("domain"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-schema.html
//
type CfnSchemaProps struct {
	// The name of the schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-schema.html#cfn-personalize-schema-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-schema.html#cfn-personalize-schema-schema
	//
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// The domain of a schema that you created for a dataset in a Domain dataset group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-schema.html#cfn-personalize-schema-domain
	//
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
}

