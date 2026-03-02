package previewawsec2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRouteServerPeerEventLogsLogGroupProps := &CfnRouteServerPeerEventLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnRouteServerPeerEventLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnRouteServerPeerEventLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnRouteServerPeerEventLogsRecordFields_STATUS,
//   	},
//   }
//
// Experimental.
type CfnRouteServerPeerEventLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnRouteServerPeerEventLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnRouteServerPeerEventLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

