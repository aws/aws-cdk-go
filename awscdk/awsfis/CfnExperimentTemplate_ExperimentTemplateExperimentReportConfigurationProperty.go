package awsfis


// Describes the report configuration for the experiment template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   experimentTemplateExperimentReportConfigurationProperty := &ExperimentTemplateExperimentReportConfigurationProperty{
//   	DataSources: &DataSourcesProperty{
//   		CloudWatchDashboards: []interface{}{
//   			&CloudWatchDashboardProperty{
//   				DashboardIdentifier: jsii.String("dashboardIdentifier"),
//   			},
//   		},
//   	},
//   	Outputs: &OutputsProperty{
//   		ExperimentReportS3Configuration: &ExperimentReportS3ConfigurationProperty{
//   			BucketName: jsii.String("bucketName"),
//
//   			// the properties below are optional
//   			Prefix: jsii.String("prefix"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	PostExperimentDuration: jsii.String("postExperimentDuration"),
//   	PreExperimentDuration: jsii.String("preExperimentDuration"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateexperimentreportconfiguration.html
//
type CfnExperimentTemplate_ExperimentTemplateExperimentReportConfigurationProperty struct {
	// The data sources for the experiment report.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateexperimentreportconfiguration.html#cfn-fis-experimenttemplate-experimenttemplateexperimentreportconfiguration-datasources
	//
	DataSources interface{} `field:"required" json:"dataSources" yaml:"dataSources"`
	// The output destinations of the experiment report.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateexperimentreportconfiguration.html#cfn-fis-experimenttemplate-experimenttemplateexperimentreportconfiguration-outputs
	//
	Outputs interface{} `field:"required" json:"outputs" yaml:"outputs"`
	// The duration after the experiment end time for the data sources to include in the report.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateexperimentreportconfiguration.html#cfn-fis-experimenttemplate-experimenttemplateexperimentreportconfiguration-postexperimentduration
	//
	PostExperimentDuration *string `field:"optional" json:"postExperimentDuration" yaml:"postExperimentDuration"`
	// The duration before the experiment start time for the data sources to include in the report.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-experimenttemplateexperimentreportconfiguration.html#cfn-fis-experimenttemplate-experimenttemplateexperimentreportconfiguration-preexperimentduration
	//
	PreExperimentDuration *string `field:"optional" json:"preExperimentDuration" yaml:"preExperimentDuration"`
}

