package awslicensemanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnReportGenerator`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnReportGeneratorProps := &CfnReportGeneratorProps{
//   	ReportContext: &ReportContextProperty{
//   		LicenseAssetGroupArns: []*string{
//   			jsii.String("licenseAssetGroupArns"),
//   		},
//   		LicenseConfigurationArns: []*string{
//   			jsii.String("licenseConfigurationArns"),
//   		},
//   		ReportEndDate: jsii.String("reportEndDate"),
//   		ReportStartDate: jsii.String("reportStartDate"),
//   	},
//   	ReportFrequency: &ReportFrequencyProperty{
//   		Period: jsii.String("period"),
//   		Value: jsii.Number(123),
//   	},
//   	ReportGeneratorName: jsii.String("reportGeneratorName"),
//   	ReportType: []*string{
//   		jsii.String("reportType"),
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-reportgenerator.html
//
type CfnReportGeneratorProps struct {
	// Details of the license configurations and asset groups that this generator reports on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-reportgenerator.html#cfn-licensemanager-reportgenerator-reportcontext
	//
	ReportContext interface{} `field:"required" json:"reportContext" yaml:"reportContext"`
	// Details about how frequently reports are generated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-reportgenerator.html#cfn-licensemanager-reportgenerator-reportfrequency
	//
	ReportFrequency interface{} `field:"required" json:"reportFrequency" yaml:"reportFrequency"`
	// Name of the report generator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-reportgenerator.html#cfn-licensemanager-reportgenerator-reportgeneratorname
	//
	ReportGeneratorName *string `field:"required" json:"reportGeneratorName" yaml:"reportGeneratorName"`
	// Type of reports to generate.
	//
	// The report type determines the data reported on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-reportgenerator.html#cfn-licensemanager-reportgenerator-reporttype
	//
	ReportType *[]*string `field:"required" json:"reportType" yaml:"reportType"`
	// Description of the report generator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-reportgenerator.html#cfn-licensemanager-reportgenerator-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-reportgenerator.html#cfn-licensemanager-reportgenerator-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

