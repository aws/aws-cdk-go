package awss3


// Specifies how data related to the storage class analysis for an Amazon S3 bucket should be exported.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataExportProperty := &DataExportProperty{
//   	Destination: &DestinationProperty{
//   		BucketArn: jsii.String("bucketArn"),
//   		Format: jsii.String("format"),
//
//   		// the properties below are optional
//   		BucketAccountId: jsii.String("bucketAccountId"),
//   		Prefix: jsii.String("prefix"),
//   	},
//   	OutputSchemaVersion: jsii.String("outputSchemaVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-dataexport.html
//
type CfnBucket_DataExportProperty struct {
	// The place to store the data for an analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-dataexport.html#cfn-s3-bucket-dataexport-destination
	//
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// The version of the output schema to use when exporting data.
	//
	// Must be `V_1` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-dataexport.html#cfn-s3-bucket-dataexport-outputschemaversion
	//
	OutputSchemaVersion *string `field:"required" json:"outputSchemaVersion" yaml:"outputSchemaVersion"`
}

