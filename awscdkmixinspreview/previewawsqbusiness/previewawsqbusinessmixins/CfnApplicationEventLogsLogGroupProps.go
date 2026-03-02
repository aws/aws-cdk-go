package previewawsqbusinessmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationEventLogsLogGroupProps := &CfnApplicationEventLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnApplicationEventLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnApplicationEventLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnApplicationEventLogsRecordFields_APPLICATION_ID,
//   	},
//   }
//
// Experimental.
type CfnApplicationEventLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnApplicationEventLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnApplicationEventLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

