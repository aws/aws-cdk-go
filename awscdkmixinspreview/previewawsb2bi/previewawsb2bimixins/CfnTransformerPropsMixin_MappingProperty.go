package previewawsb2bimixins


// Specifies the mapping template for the transformer.
//
// This template is used to map the parsed EDI file using JSONata or XSLT.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mappingProperty := &MappingProperty{
//   	Template: jsii.String("template"),
//   	TemplateLanguage: jsii.String("templateLanguage"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-mapping.html
//
type CfnTransformerPropsMixin_MappingProperty struct {
	// A string that represents the mapping template, in the transformation language specified in `templateLanguage` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-mapping.html#cfn-b2bi-transformer-mapping-template
	//
	Template *string `field:"optional" json:"template" yaml:"template"`
	// The transformation language for the template, either XSLT or JSONATA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-transformer-mapping.html#cfn-b2bi-transformer-mapping-templatelanguage
	//
	TemplateLanguage *string `field:"optional" json:"templateLanguage" yaml:"templateLanguage"`
}

