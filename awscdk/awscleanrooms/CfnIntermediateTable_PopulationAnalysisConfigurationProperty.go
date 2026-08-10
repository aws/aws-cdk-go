package awscleanrooms


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   populationAnalysisConfigurationProperty := &PopulationAnalysisConfigurationProperty{
//   	SqlParameters: &PopulationAnalysisSqlParametersProperty{
//   		AnalysisTemplateArn: jsii.String("analysisTemplateArn"),
//   		QueryString: jsii.String("queryString"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-populationanalysisconfiguration.html
//
type CfnIntermediateTable_PopulationAnalysisConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-populationanalysisconfiguration.html#cfn-cleanrooms-intermediatetable-populationanalysisconfiguration-sqlparameters
	//
	SqlParameters interface{} `field:"optional" json:"sqlParameters" yaml:"sqlParameters"`
}

