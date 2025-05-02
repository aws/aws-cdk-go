package awscloudfront


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterDefinitionProperty := &ParameterDefinitionProperty{
//   	Definition: &DefinitionProperty{
//   		StringSchema: &StringSchemaProperty{
//   			Required: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			Comment: jsii.String("comment"),
//   			DefaultValue: jsii.String("defaultValue"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-parameterdefinition.html
//
type CfnDistribution_ParameterDefinitionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-parameterdefinition.html#cfn-cloudfront-distribution-parameterdefinition-definition
	//
	Definition interface{} `field:"required" json:"definition" yaml:"definition"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-parameterdefinition.html#cfn-cloudfront-distribution-parameterdefinition-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
}

