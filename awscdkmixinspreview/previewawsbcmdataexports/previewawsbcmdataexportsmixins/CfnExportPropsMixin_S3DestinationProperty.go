package previewawsbcmdataexportsmixins


// Describes the destination Amazon Simple Storage Service (Amazon S3) bucket name and object keys of a data exports file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3DestinationProperty := &S3DestinationProperty{
//   	S3Bucket: jsii.String("s3Bucket"),
//   	S3OutputConfigurations: &S3OutputConfigurationsProperty{
//   		Compression: jsii.String("compression"),
//   		Format: jsii.String("format"),
//   		OutputType: jsii.String("outputType"),
//   		Overwrite: jsii.String("overwrite"),
//   	},
//   	S3Prefix: jsii.String("s3Prefix"),
//   	S3Region: jsii.String("s3Region"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-s3destination.html
//
type CfnExportPropsMixin_S3DestinationProperty struct {
	// The name of the Amazon S3 bucket used as the destination of a data export file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-s3destination.html#cfn-bcmdataexports-export-s3destination-s3bucket
	//
	S3Bucket *string `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
	// The output configuration for the data export.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-s3destination.html#cfn-bcmdataexports-export-s3destination-s3outputconfigurations
	//
	S3OutputConfigurations interface{} `field:"optional" json:"s3OutputConfigurations" yaml:"s3OutputConfigurations"`
	// The S3 path prefix you want prepended to the name of your data export.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-s3destination.html#cfn-bcmdataexports-export-s3destination-s3prefix
	//
	S3Prefix *string `field:"optional" json:"s3Prefix" yaml:"s3Prefix"`
	// The S3 bucket Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-s3destination.html#cfn-bcmdataexports-export-s3destination-s3region
	//
	S3Region *string `field:"optional" json:"s3Region" yaml:"s3Region"`
}

