package awscur

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnReportDefinition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReportDefinitionProps := &CfnReportDefinitionProps{
//   	Compression: jsii.String("compression"),
//   	Format: jsii.String("format"),
//   	RefreshClosedReports: jsii.Boolean(false),
//   	ReportName: jsii.String("reportName"),
//   	ReportVersioning: jsii.String("reportVersioning"),
//   	S3Bucket: jsii.String("s3Bucket"),
//   	S3Prefix: jsii.String("s3Prefix"),
//   	S3Region: jsii.String("s3Region"),
//   	TimeUnit: jsii.String("timeUnit"),
//
//   	// the properties below are optional
//   	AdditionalArtifacts: []*string{
//   		jsii.String("additionalArtifacts"),
//   	},
//   	AdditionalSchemaElements: []*string{
//   		jsii.String("additionalSchemaElements"),
//   	},
//   	BillingViewArn: jsii.String("billingViewArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cur-reportdefinition.html
//
type CfnReportDefinitionProps struct {
	// The compression format that Amazon Web Services uses for the report.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cur-reportdefinition.html#cfn-cur-reportdefinition-compression
	//
	Compression *string `field:"required" json:"compression" yaml:"compression"`
	// The format that Amazon Web Services saves the report in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cur-reportdefinition.html#cfn-cur-reportdefinition-format
	//
	Format *string `field:"required" json:"format" yaml:"format"`
	// Whether you want AWS to update your reports after they have been finalized if AWS detects charges related to previous months.
	//
	// These charges can include refunds, credits, or support fees.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cur-reportdefinition.html#cfn-cur-reportdefinition-refreshclosedreports
	//
	RefreshClosedReports interface{} `field:"required" json:"refreshClosedReports" yaml:"refreshClosedReports"`
	// The name of the report that you want to create.
	//
	// The name must be unique, is case sensitive, and can't include spaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cur-reportdefinition.html#cfn-cur-reportdefinition-reportname
	//
	ReportName *string `field:"required" json:"reportName" yaml:"reportName"`
	// Whether you want AWS to overwrite the previous version of each report or to deliver the report in addition to the previous versions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cur-reportdefinition.html#cfn-cur-reportdefinition-reportversioning
	//
	ReportVersioning *string `field:"required" json:"reportVersioning" yaml:"reportVersioning"`
	// The S3 bucket where Amazon Web Services delivers the report.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cur-reportdefinition.html#cfn-cur-reportdefinition-s3bucket
	//
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// The prefix that Amazon Web Services adds to the report name when Amazon Web Services delivers the report.
	//
	// Your prefix can't include spaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cur-reportdefinition.html#cfn-cur-reportdefinition-s3prefix
	//
	S3Prefix *string `field:"required" json:"s3Prefix" yaml:"s3Prefix"`
	// The Region of the S3 bucket that Amazon Web Services delivers the report into.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cur-reportdefinition.html#cfn-cur-reportdefinition-s3region
	//
	S3Region *string `field:"required" json:"s3Region" yaml:"s3Region"`
	// The granularity of the line items in the report.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cur-reportdefinition.html#cfn-cur-reportdefinition-timeunit
	//
	TimeUnit *string `field:"required" json:"timeUnit" yaml:"timeUnit"`
	// A list of manifests that you want AWS to create for this report.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cur-reportdefinition.html#cfn-cur-reportdefinition-additionalartifacts
	//
	AdditionalArtifacts *[]*string `field:"optional" json:"additionalArtifacts" yaml:"additionalArtifacts"`
	// A list of strings that indicate additional content that AWS includes in the report, such as individual resource IDs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cur-reportdefinition.html#cfn-cur-reportdefinition-additionalschemaelements
	//
	AdditionalSchemaElements *[]*string `field:"optional" json:"additionalSchemaElements" yaml:"additionalSchemaElements"`
	// The Amazon Resource Name (ARN) of the billing view.
	//
	// You can get this value by using the billing view service public APIs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cur-reportdefinition.html#cfn-cur-reportdefinition-billingviewarn
	//
	BillingViewArn *string `field:"optional" json:"billingViewArn" yaml:"billingViewArn"`
	// The tags to be assigned to the report definition resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cur-reportdefinition.html#cfn-cur-reportdefinition-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

