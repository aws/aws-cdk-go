package previewawsec2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRouteServerPeerEventLogsFirehoseProps := &CfnRouteServerPeerEventLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnRouteServerPeerEventLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnRouteServerPeerEventLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnRouteServerPeerEventLogsRecordFields_STATUS,
//   	},
//   }
//
// Experimental.
type CfnRouteServerPeerEventLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnRouteServerPeerEventLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnRouteServerPeerEventLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

