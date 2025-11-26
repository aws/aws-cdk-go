package previewawspersonalizemixins


// Properties for CfnSchemaPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSchemaMixinProps := &CfnSchemaMixinProps{
//   	Domain: jsii.String("domain"),
//   	Name: jsii.String("name"),
//   	Schema: jsii.String("schema"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-schema.html
//
type CfnSchemaMixinProps struct {
	// The domain of a schema that you created for a dataset in a Domain dataset group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-schema.html#cfn-personalize-schema-domain
	//
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The name of the schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-schema.html#cfn-personalize-schema-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-schema.html#cfn-personalize-schema-schema
	//
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

