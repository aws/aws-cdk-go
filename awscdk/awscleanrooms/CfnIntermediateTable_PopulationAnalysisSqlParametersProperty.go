package awscleanrooms


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   populationAnalysisSqlParametersProperty := &PopulationAnalysisSqlParametersProperty{
//   	AnalysisTemplateArn: jsii.String("analysisTemplateArn"),
//   	QueryString: jsii.String("queryString"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-populationanalysissqlparameters.html
//
type CfnIntermediateTable_PopulationAnalysisSqlParametersProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-populationanalysissqlparameters.html#cfn-cleanrooms-intermediatetable-populationanalysissqlparameters-analysistemplatearn
	//
	AnalysisTemplateArn *string `field:"optional" json:"analysisTemplateArn" yaml:"analysisTemplateArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-populationanalysissqlparameters.html#cfn-cleanrooms-intermediatetable-populationanalysissqlparameters-querystring
	//
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
}

