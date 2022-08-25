package awss3


// Specifies how data related to the storage class analysis for an Amazon S3 bucket should be exported.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataExportProperty := &dataExportProperty{
//   	destination: &destinationProperty{
//   		bucketArn: jsii.String("bucketArn"),
//   		format: jsii.String("format"),
//
//   		// the properties below are optional
//   		bucketAccountId: jsii.String("bucketAccountId"),
//   		prefix: jsii.String("prefix"),
//   	},
//   	outputSchemaVersion: jsii.String("outputSchemaVersion"),
//   }
//
type CfnBucket_DataExportProperty struct {
	// The place to store the data for an analysis.
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// The version of the output schema to use when exporting data.
	//
	// Must be `V_1` .
	OutputSchemaVersion *string `field:"required" json:"outputSchemaVersion" yaml:"outputSchemaVersion"`
}

