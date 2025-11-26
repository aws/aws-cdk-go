package previewawss3mixins


// Specifies how data related to the storage class analysis for an Amazon S3 bucket should be exported.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataExportProperty := &DataExportProperty{
//   	Destination: &DestinationProperty{
//   		BucketAccountId: jsii.String("bucketAccountId"),
//   		BucketArn: jsii.String("bucketArn"),
//   		Format: jsii.String("format"),
//   		Prefix: jsii.String("prefix"),
//   	},
//   	OutputSchemaVersion: jsii.String("outputSchemaVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-dataexport.html
//
type CfnBucketPropsMixin_DataExportProperty struct {
	// The place to store the data for an analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-dataexport.html#cfn-s3-bucket-dataexport-destination
	//
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// The version of the output schema to use when exporting data.
	//
	// Must be `V_1` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-dataexport.html#cfn-s3-bucket-dataexport-outputschemaversion
	//
	OutputSchemaVersion *string `field:"optional" json:"outputSchemaVersion" yaml:"outputSchemaVersion"`
}

