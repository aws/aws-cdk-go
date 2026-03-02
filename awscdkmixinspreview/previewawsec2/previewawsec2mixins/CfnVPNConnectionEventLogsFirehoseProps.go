package previewawsec2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVPNConnectionEventLogsFirehoseProps := &CfnVPNConnectionEventLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnVPNConnectionEventLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnVPNConnectionEventLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnVPNConnectionEventLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnVPNConnectionEventLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnVPNConnectionEventLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnVPNConnectionEventLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

