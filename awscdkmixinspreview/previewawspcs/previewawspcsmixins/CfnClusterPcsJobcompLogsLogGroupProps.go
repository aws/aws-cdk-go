package previewawspcsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterPcsJobcompLogsLogGroupProps := &CfnClusterPcsJobcompLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnClusterPcsJobcompLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnClusterPcsJobcompLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnClusterPcsJobcompLogsRecordFields_RESOURCE_ID,
//   	},
//   }
//
// Experimental.
type CfnClusterPcsJobcompLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnClusterPcsJobcompLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnClusterPcsJobcompLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

