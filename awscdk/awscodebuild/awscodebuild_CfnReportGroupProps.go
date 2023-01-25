package awscodebuild

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnReportGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReportGroupProps := &cfnReportGroupProps{
//   	exportConfig: &reportExportConfigProperty{
//   		exportConfigType: jsii.String("exportConfigType"),
//
//   		// the properties below are optional
//   		s3Destination: &s3ReportExportConfigProperty{
//   			bucket: jsii.String("bucket"),
//
//   			// the properties below are optional
//   			bucketOwner: jsii.String("bucketOwner"),
//   			encryptionDisabled: jsii.Boolean(false),
//   			encryptionKey: jsii.String("encryptionKey"),
//   			packaging: jsii.String("packaging"),
//   			path: jsii.String("path"),
//   		},
//   	},
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	deleteReports: jsii.Boolean(false),
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnReportGroupProps struct {
	// Information about the destination where the raw data of this `ReportGroup` is exported.
	ExportConfig interface{} `field:"required" json:"exportConfig" yaml:"exportConfig"`
	// The type of the `ReportGroup` . This can be one of the following values:.
	//
	// - **CODE_COVERAGE** - The report group contains code coverage reports.
	// - **TEST** - The report group contains test reports.
	Type *string `field:"required" json:"type" yaml:"type"`
	// When deleting a report group, specifies if reports within the report group should be deleted.
	//
	// - **true** - Deletes any reports that belong to the report group before deleting the report group.
	// - **false** - You must delete any reports in the report group. This is the default value. If you delete a report group that contains one or more reports, an exception is thrown.
	DeleteReports interface{} `field:"optional" json:"deleteReports" yaml:"deleteReports"`
	// The name of the `ReportGroup` .
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of tag key and value pairs associated with this report group.
	//
	// These tags are available for use by AWS services that support AWS CodeBuild report group tags.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

