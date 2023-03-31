package awspersonalize


// Properties for defining a `CfnSchema`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSchemaProps := &cfnSchemaProps{
//   	name: jsii.String("name"),
//   	schema: jsii.String("schema"),
//
//   	// the properties below are optional
//   	domain: jsii.String("domain"),
//   }
//
type CfnSchemaProps struct {
	// The name of the schema.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The schema.
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// The domain of a schema that you created for a dataset in a Domain dataset group.
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
}

