package awscloudfront


// The value that you assigned to the parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   definitionProperty := &DefinitionProperty{
//   	StringSchema: &StringSchemaProperty{
//   		Required: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		Comment: jsii.String("comment"),
//   		DefaultValue: jsii.String("defaultValue"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-definition.html
//
type CfnDistribution_DefinitionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-definition.html#cfn-cloudfront-distribution-definition-stringschema
	//
	StringSchema interface{} `field:"optional" json:"stringSchema" yaml:"stringSchema"`
}

