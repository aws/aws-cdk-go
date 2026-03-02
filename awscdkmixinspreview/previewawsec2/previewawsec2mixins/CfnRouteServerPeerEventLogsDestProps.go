package previewawsec2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRouteServerPeerEventLogsDestProps := &CfnRouteServerPeerEventLogsDestProps{
//   	RecordFields: []CfnRouteServerPeerEventLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnRouteServerPeerEventLogsRecordFields_STATUS,
//   	},
//   }
//
// Experimental.
type CfnRouteServerPeerEventLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnRouteServerPeerEventLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

