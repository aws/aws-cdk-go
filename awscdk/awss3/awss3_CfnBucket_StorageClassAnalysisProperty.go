package awss3


// Specifies data related to access patterns to be collected and made available to analyze the tradeoffs between different storage classes for an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   storageClassAnalysisProperty := &storageClassAnalysisProperty{
//   	dataExport: &dataExportProperty{
//   		destination: &destinationProperty{
//   			bucketArn: jsii.String("bucketArn"),
//   			format: jsii.String("format"),
//
//   			// the properties below are optional
//   			bucketAccountId: jsii.String("bucketAccountId"),
//   			prefix: jsii.String("prefix"),
//   		},
//   		outputSchemaVersion: jsii.String("outputSchemaVersion"),
//   	},
//   }
//
type CfnBucket_StorageClassAnalysisProperty struct {
	// Specifies how data related to the storage class analysis for an Amazon S3 bucket should be exported.
	DataExport interface{} `field:"optional" json:"dataExport" yaml:"dataExport"`
}

