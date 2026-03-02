package previewawscleanroomsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMembershipAnalysisLogsFirehoseProps := &CfnMembershipAnalysisLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnMembershipAnalysisLogsOutputFormat.Firehose_JSON,
//   }
//
// Experimental.
type CfnMembershipAnalysisLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnMembershipAnalysisLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

