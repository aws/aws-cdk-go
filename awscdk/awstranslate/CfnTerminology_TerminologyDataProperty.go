package awstranslate


// The terminology data for the custom terminology being imported.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   terminologyDataProperty := &TerminologyDataProperty{
//   	File: jsii.String("file"),
//   	Format: jsii.String("format"),
//
//   	// the properties below are optional
//   	Directionality: jsii.String("directionality"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-translate-terminology-terminologydata.html
//
type CfnTerminology_TerminologyDataProperty struct {
	// The file containing the custom terminology data, base64-encoded.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-translate-terminology-terminologydata.html#cfn-translate-terminology-terminologydata-file
	//
	File *string `field:"required" json:"file" yaml:"file"`
	// The data format of the custom terminology.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-translate-terminology-terminologydata.html#cfn-translate-terminology-terminologydata-format
	//
	Format *string `field:"required" json:"format" yaml:"format"`
	// The directionality of the terminology resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-translate-terminology-terminologydata.html#cfn-translate-terminology-terminologydata-directionality
	//
	Directionality *string `field:"optional" json:"directionality" yaml:"directionality"`
}

