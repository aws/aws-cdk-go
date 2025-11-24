package mixinsawsbcmdataexports


// The compression type, file format, and overwrite preference for the data export.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3OutputConfigurationsProperty := &S3OutputConfigurationsProperty{
//   	Compression: jsii.String("compression"),
//   	Format: jsii.String("format"),
//   	OutputType: jsii.String("outputType"),
//   	Overwrite: jsii.String("overwrite"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-s3outputconfigurations.html
//
type CfnExportPropsMixin_S3OutputConfigurationsProperty struct {
	// The compression type for the data export.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-s3outputconfigurations.html#cfn-bcmdataexports-export-s3outputconfigurations-compression
	//
	Compression *string `field:"optional" json:"compression" yaml:"compression"`
	// The file format for the data export.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-s3outputconfigurations.html#cfn-bcmdataexports-export-s3outputconfigurations-format
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
	// The output type for the data export.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-s3outputconfigurations.html#cfn-bcmdataexports-export-s3outputconfigurations-outputtype
	//
	OutputType *string `field:"optional" json:"outputType" yaml:"outputType"`
	// The rule to follow when generating a version of the data export file.
	//
	// You have the choice to overwrite the previous version or to be delivered in addition to the previous versions. Overwriting exports can save on Amazon S3 storage costs. Creating new export versions allows you to track the changes in cost and usage data over time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-s3outputconfigurations.html#cfn-bcmdataexports-export-s3outputconfigurations-overwrite
	//
	Overwrite *string `field:"optional" json:"overwrite" yaml:"overwrite"`
}

