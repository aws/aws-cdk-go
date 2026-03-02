package previewawsec2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVPNConnectionConnectionLogsFirehoseProps := &CfnVPNConnectionConnectionLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnVPNConnectionConnectionLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnVPNConnectionConnectionLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnVPNConnectionConnectionLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnVPNConnectionConnectionLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnVPNConnectionConnectionLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnVPNConnectionConnectionLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

