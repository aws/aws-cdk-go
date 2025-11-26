package previewawscloudfrontmixins


// The value that you assigned to the parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   definitionProperty := &DefinitionProperty{
//   	StringSchema: &StringSchemaProperty{
//   		Comment: jsii.String("comment"),
//   		DefaultValue: jsii.String("defaultValue"),
//   		Required: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-definition.html
//
type CfnDistributionPropsMixin_DefinitionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-definition.html#cfn-cloudfront-distribution-definition-stringschema
	//
	StringSchema interface{} `field:"optional" json:"stringSchema" yaml:"stringSchema"`
}

