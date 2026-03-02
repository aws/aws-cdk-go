package previewawsqbusinessmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationEventLogsFirehoseProps := &CfnApplicationEventLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnApplicationEventLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnApplicationEventLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnApplicationEventLogsRecordFields_APPLICATION_ID,
//   	},
//   }
//
// Experimental.
type CfnApplicationEventLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnApplicationEventLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnApplicationEventLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

