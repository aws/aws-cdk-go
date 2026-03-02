package previewawspcsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterPcsJobcompLogsFirehoseProps := &CfnClusterPcsJobcompLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnClusterPcsJobcompLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnClusterPcsJobcompLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnClusterPcsJobcompLogsRecordFields_RESOURCE_ID,
//   	},
//   }
//
// Experimental.
type CfnClusterPcsJobcompLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnClusterPcsJobcompLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnClusterPcsJobcompLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

