package previewawsiotfleetwisemixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCampaignIotFleetwiseLogsLogGroupProps := &CfnCampaignIotFleetwiseLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnCampaignIotFleetwiseLogsOutputFormat.LogGroup_PLAIN,
//   }
//
// Experimental.
type CfnCampaignIotFleetwiseLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnCampaignIotFleetwiseLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

