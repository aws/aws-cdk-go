package previewawspcsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterPcsJobcompLogsFirehoseProps := &CfnClusterPcsJobcompLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnClusterPcsJobcompLogsOutputFormat.Firehose_PLAIN,
//   	RecordFields: []CfnClusterPcsJobcompLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnClusterPcsJobcompLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnClusterPcsJobcompLogsFirehoseProps struct {
	// Format for log output, options are plain,json,raw.
	// Experimental.
	OutputFormat CfnClusterPcsJobcompLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnClusterPcsJobcompLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

