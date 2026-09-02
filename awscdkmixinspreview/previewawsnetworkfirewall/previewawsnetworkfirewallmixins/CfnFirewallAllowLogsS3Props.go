package previewawsnetworkfirewallmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskms"
)

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var keyRef IKeyRef
//
//   cfnFirewallAllowLogsS3Props := &CfnFirewallAllowLogsS3Props{
//   	EncryptionKey: keyRef,
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnFirewallAllowLogsOutputFormat.S3_PLAIN,
//   	RecordFields: []CfnFirewallAllowLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnFirewallAllowLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnFirewallAllowLogsS3Props struct {
	// Encrpytion key for your delivery bucket.
	// Experimental.
	EncryptionKey interfacesawskms.IKeyRef `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Format for log output, options are plain,json,w3c,parquet.
	// Experimental.
	OutputFormat CfnFirewallAllowLogsOutputFormat_S3 `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnFirewallAllowLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

