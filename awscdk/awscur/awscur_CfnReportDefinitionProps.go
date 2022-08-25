package awscur


// Properties for defining a `CfnReportDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReportDefinitionProps := &cfnReportDefinitionProps{
//   	compression: jsii.String("compression"),
//   	format: jsii.String("format"),
//   	refreshClosedReports: jsii.Boolean(false),
//   	reportName: jsii.String("reportName"),
//   	reportVersioning: jsii.String("reportVersioning"),
//   	s3Bucket: jsii.String("s3Bucket"),
//   	s3Prefix: jsii.String("s3Prefix"),
//   	s3Region: jsii.String("s3Region"),
//   	timeUnit: jsii.String("timeUnit"),
//
//   	// the properties below are optional
//   	additionalArtifacts: []*string{
//   		jsii.String("additionalArtifacts"),
//   	},
//   	additionalSchemaElements: []*string{
//   		jsii.String("additionalSchemaElements"),
//   	},
//   	billingViewArn: jsii.String("billingViewArn"),
//   }
//
type CfnReportDefinitionProps struct {
	// The compression format that Amazon Web Services uses for the report.
	Compression *string `field:"required" json:"compression" yaml:"compression"`
	// The format that Amazon Web Services saves the report in.
	Format *string `field:"required" json:"format" yaml:"format"`
	// Whether you want AWS to update your reports after they have been finalized if AWS detects charges related to previous months.
	//
	// These charges can include refunds, credits, or support fees.
	RefreshClosedReports interface{} `field:"required" json:"refreshClosedReports" yaml:"refreshClosedReports"`
	// The name of the report that you want to create.
	//
	// The name must be unique, is case sensitive, and can't include spaces.
	ReportName *string `field:"required" json:"reportName" yaml:"reportName"`
	// Whether you want AWS to overwrite the previous version of each report or to deliver the report in addition to the previous versions.
	ReportVersioning *string `field:"required" json:"reportVersioning" yaml:"reportVersioning"`
	// The S3 bucket where Amazon Web Services delivers the report.
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// The prefix that Amazon Web Services adds to the report name when Amazon Web Services delivers the report.
	//
	// Your prefix can't include spaces.
	S3Prefix *string `field:"required" json:"s3Prefix" yaml:"s3Prefix"`
	// The Region of the S3 bucket that Amazon Web Services delivers the report into.
	S3Region *string `field:"required" json:"s3Region" yaml:"s3Region"`
	// The granularity of the line items in the report.
	TimeUnit *string `field:"required" json:"timeUnit" yaml:"timeUnit"`
	// A list of manifests that you want AWS to create for this report.
	AdditionalArtifacts *[]*string `field:"optional" json:"additionalArtifacts" yaml:"additionalArtifacts"`
	// A list of strings that indicate additional content that AWS includes in the report, such as individual resource IDs.
	AdditionalSchemaElements *[]*string `field:"optional" json:"additionalSchemaElements" yaml:"additionalSchemaElements"`
	// The Amazon Resource Name (ARN) of the billing view.
	//
	// You can get this value by using the billing view service public APIs.
	BillingViewArn *string `field:"optional" json:"billingViewArn" yaml:"billingViewArn"`
}

