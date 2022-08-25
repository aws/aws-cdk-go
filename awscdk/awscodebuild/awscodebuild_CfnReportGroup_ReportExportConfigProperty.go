package awscodebuild


// Information about the location where the run of a report is exported.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   reportExportConfigProperty := &reportExportConfigProperty{
//   	exportConfigType: jsii.String("exportConfigType"),
//
//   	// the properties below are optional
//   	s3Destination: &s3ReportExportConfigProperty{
//   		bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		bucketOwner: jsii.String("bucketOwner"),
//   		encryptionDisabled: jsii.Boolean(false),
//   		encryptionKey: jsii.String("encryptionKey"),
//   		packaging: jsii.String("packaging"),
//   		path: jsii.String("path"),
//   	},
//   }
//
type CfnReportGroup_ReportExportConfigProperty struct {
	// The export configuration type. Valid values are:.
	//
	// - `S3` : The report results are exported to an S3 bucket.
	// - `NO_EXPORT` : The report results are not exported.
	ExportConfigType *string `field:"required" json:"exportConfigType" yaml:"exportConfigType"`
	// A `S3ReportExportConfig` object that contains information about the S3 bucket where the run of a report is exported.
	S3Destination interface{} `field:"optional" json:"s3Destination" yaml:"s3Destination"`
}

