package awsb2bi


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mappingProperty := &MappingProperty{
//   	TemplateLanguage: jsii.String("templateLanguage"),
//
//   	// the properties below are optional
//   	Template: jsii.String("template"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-mapping.html
//
type CfnTransformer_MappingProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-mapping.html#cfn-b2bi-transformer-mapping-templatelanguage
	//
	TemplateLanguage *string `field:"required" json:"templateLanguage" yaml:"templateLanguage"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-mapping.html#cfn-b2bi-transformer-mapping-template
	//
	Template *string `field:"optional" json:"template" yaml:"template"`
}

