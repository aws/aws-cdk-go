package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWorkloadIdentityApplicationLogsFirehoseProps := &CfnWorkloadIdentityApplicationLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnWorkloadIdentityApplicationLogsOutputFormat.Firehose_PLAIN,
//   	RecordFields: []CfnWorkloadIdentityApplicationLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnWorkloadIdentityApplicationLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnWorkloadIdentityApplicationLogsFirehoseProps struct {
	// Format for log output, options are plain,json,raw.
	// Experimental.
	OutputFormat CfnWorkloadIdentityApplicationLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnWorkloadIdentityApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

