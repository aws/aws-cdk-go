package previewawscloudfrontmixins


// A list of parameter values to add to the resource.
//
// A parameter is specified as a key-value pair. A valid parameter value must exist for any parameter that is marked as required in the multi-tenant distribution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   parameterDefinitionProperty := &ParameterDefinitionProperty{
//   	Definition: &DefinitionProperty{
//   		StringSchema: &StringSchemaProperty{
//   			Comment: jsii.String("comment"),
//   			DefaultValue: jsii.String("defaultValue"),
//   			Required: jsii.Boolean(false),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-parameterdefinition.html
//
type CfnDistributionPropsMixin_ParameterDefinitionProperty struct {
	// The value that you assigned to the parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-parameterdefinition.html#cfn-cloudfront-distribution-parameterdefinition-definition
	//
	Definition interface{} `field:"optional" json:"definition" yaml:"definition"`
	// The name of the parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-parameterdefinition.html#cfn-cloudfront-distribution-parameterdefinition-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

