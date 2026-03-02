package previewawsqbusinessmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationEventLogsDestProps := &CfnApplicationEventLogsDestProps{
//   	RecordFields: []CfnApplicationEventLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnApplicationEventLogsRecordFields_APPLICATION_ID,
//   	},
//   }
//
// Experimental.
type CfnApplicationEventLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnApplicationEventLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

