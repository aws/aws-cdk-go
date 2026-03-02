package previewawspcsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterPcsJobcompLogsDestProps := &CfnClusterPcsJobcompLogsDestProps{
//   	RecordFields: []CfnClusterPcsJobcompLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnClusterPcsJobcompLogsRecordFields_RESOURCE_ID,
//   	},
//   }
//
// Experimental.
type CfnClusterPcsJobcompLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnClusterPcsJobcompLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

