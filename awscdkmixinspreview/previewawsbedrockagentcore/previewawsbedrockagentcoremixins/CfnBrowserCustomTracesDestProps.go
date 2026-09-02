package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBrowserCustomTracesDestProps := &CfnBrowserCustomTracesDestProps{
//   	RecordFields: []CfnBrowserCustomTracesRecordFields{
//   		awscdkmixinspreview.Mixins.CfnBrowserCustomTracesRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnBrowserCustomTracesDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnBrowserCustomTracesRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

