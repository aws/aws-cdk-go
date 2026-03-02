package previewawsivschatmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRoomIvsChatLogsFirehoseProps := &CfnRoomIvsChatLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnRoomIvsChatLogsOutputFormat.Firehose_JSON,
//   }
//
// Experimental.
type CfnRoomIvsChatLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnRoomIvsChatLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

