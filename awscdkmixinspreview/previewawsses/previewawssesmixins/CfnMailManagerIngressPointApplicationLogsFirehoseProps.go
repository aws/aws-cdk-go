package previewawssesmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMailManagerIngressPointApplicationLogsFirehoseProps := &CfnMailManagerIngressPointApplicationLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnMailManagerIngressPointApplicationLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnMailManagerIngressPointApplicationLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnMailManagerIngressPointApplicationLogsRecordFields_INGRESS_POINT_TYPE,
//   	},
//   }
//
// Experimental.
type CfnMailManagerIngressPointApplicationLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnMailManagerIngressPointApplicationLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnMailManagerIngressPointApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

