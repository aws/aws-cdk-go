package previewawssesmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMailManagerIngressPointApplicationLogsLogGroupProps := &CfnMailManagerIngressPointApplicationLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnMailManagerIngressPointApplicationLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnMailManagerIngressPointApplicationLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnMailManagerIngressPointApplicationLogsRecordFields_INGRESS_POINT_TYPE,
//   	},
//   }
//
// Experimental.
type CfnMailManagerIngressPointApplicationLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnMailManagerIngressPointApplicationLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnMailManagerIngressPointApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

