package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRuntimeTracesDestProps := &CfnRuntimeTracesDestProps{
//   	RecordFields: []tRACE{
//   		awscdkmixinspreview.Mixins.CfnRuntimeTracesRecordFields_*tRACE,
//   	},
//   }
//
// Experimental.
type CfnRuntimeTracesDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnRuntimeTracesRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

