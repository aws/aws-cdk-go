package awscleanrooms


// The structure that defines the body of the analysis template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisSourceProperty := &AnalysisSourceProperty{
//   	Text: jsii.String("text"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysissource.html
//
type CfnAnalysisTemplate_AnalysisSourceProperty struct {
	// The query text.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-analysistemplate-analysissource.html#cfn-cleanrooms-analysistemplate-analysissource-text
	//
	Text *string `field:"required" json:"text" yaml:"text"`
}

