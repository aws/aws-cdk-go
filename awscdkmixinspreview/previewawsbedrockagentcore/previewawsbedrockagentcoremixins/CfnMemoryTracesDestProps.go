package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMemoryTracesDestProps := &CfnMemoryTracesDestProps{
//   	RecordFields: []tRACE{
//   		awscdkmixinspreview.Mixins.CfnMemoryTracesRecordFields_*tRACE,
//   	},
//   }
//
// Experimental.
type CfnMemoryTracesDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnMemoryTracesRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

