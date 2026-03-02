package previewawstransfermixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnServerTransferLogsLogGroupProps := &CfnServerTransferLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnServerTransferLogsOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnServerTransferLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnServerTransferLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

