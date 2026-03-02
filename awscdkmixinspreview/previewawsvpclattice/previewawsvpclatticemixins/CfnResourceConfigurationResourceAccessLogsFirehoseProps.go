package previewawsvpclatticemixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResourceConfigurationResourceAccessLogsFirehoseProps := &CfnResourceConfigurationResourceAccessLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnResourceConfigurationResourceAccessLogsOutputFormat.Firehose_JSON,
//   }
//
// Experimental.
type CfnResourceConfigurationResourceAccessLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnResourceConfigurationResourceAccessLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

