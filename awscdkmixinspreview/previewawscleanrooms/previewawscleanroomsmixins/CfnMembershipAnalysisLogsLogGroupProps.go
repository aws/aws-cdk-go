package previewawscleanroomsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMembershipAnalysisLogsLogGroupProps := &CfnMembershipAnalysisLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnMembershipAnalysisLogsOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnMembershipAnalysisLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnMembershipAnalysisLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

