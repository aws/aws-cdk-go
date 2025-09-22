package awsses


// The result of an analysis can be used in conditions to trigger actions.
//
// Analyses can inspect the email content and report a certain aspect of the email.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisProperty := &AnalysisProperty{
//   	Analyzer: jsii.String("analyzer"),
//   	ResultField: jsii.String("resultField"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-analysis.html
//
type CfnMailManagerRuleSet_AnalysisProperty struct {
	// The Amazon Resource Name (ARN) of an Add On.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-analysis.html#cfn-ses-mailmanagerruleset-analysis-analyzer
	//
	Analyzer *string `field:"required" json:"analyzer" yaml:"analyzer"`
	// The returned value from an Add On.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-analysis.html#cfn-ses-mailmanagerruleset-analysis-resultfield
	//
	ResultField *string `field:"required" json:"resultField" yaml:"resultField"`
}

