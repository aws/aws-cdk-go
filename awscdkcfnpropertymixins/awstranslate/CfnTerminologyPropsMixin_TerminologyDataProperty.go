package awstranslate


// The terminology data for the custom terminology being imported.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   terminologyDataProperty := &TerminologyDataProperty{
//   	Directionality: jsii.String("directionality"),
//   	File: jsii.String("file"),
//   	Format: jsii.String("format"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-translate-terminology-terminologydata.html
//
type CfnTerminologyPropsMixin_TerminologyDataProperty struct {
	// The directionality of the terminology resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-translate-terminology-terminologydata.html#cfn-translate-terminology-terminologydata-directionality
	//
	Directionality *string `field:"optional" json:"directionality" yaml:"directionality"`
	// The file containing the custom terminology data, base64-encoded.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-translate-terminology-terminologydata.html#cfn-translate-terminology-terminologydata-file
	//
	File *string `field:"optional" json:"file" yaml:"file"`
	// The data format of the custom terminology.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-translate-terminology-terminologydata.html#cfn-translate-terminology-terminologydata-format
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
}

