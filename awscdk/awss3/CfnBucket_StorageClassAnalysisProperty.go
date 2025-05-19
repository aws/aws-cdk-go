package awss3


// Specifies data related to access patterns to be collected and made available to analyze the tradeoffs between different storage classes for an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   storageClassAnalysisProperty := &StorageClassAnalysisProperty{
//   	DataExport: &DataExportProperty{
//   		Destination: &DestinationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			Format: jsii.String("format"),
//
//   			// the properties below are optional
//   			BucketAccountId: jsii.String("bucketAccountId"),
//   			Prefix: jsii.String("prefix"),
//   		},
//   		OutputSchemaVersion: jsii.String("outputSchemaVersion"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-storageclassanalysis.html
//
type CfnBucket_StorageClassAnalysisProperty struct {
	// Specifies how data related to the storage class analysis for an Amazon S3 bucket should be exported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-storageclassanalysis.html#cfn-s3-bucket-storageclassanalysis-dataexport
	//
	DataExport interface{} `field:"optional" json:"dataExport" yaml:"dataExport"`
}

