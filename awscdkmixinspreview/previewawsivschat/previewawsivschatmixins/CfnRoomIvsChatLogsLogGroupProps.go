package previewawsivschatmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRoomIvsChatLogsLogGroupProps := &CfnRoomIvsChatLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnRoomIvsChatLogsOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnRoomIvsChatLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnRoomIvsChatLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

