package previewawsshieldmixins

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
//   cfnProtectionFlowLogsS3Props := &CfnProtectionFlowLogsS3Props{
//   	EncryptionKey: keyRef,
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnProtectionFlowLogsOutputFormat.S3_JSON,
//   	RecordFields: []CfnProtectionFlowLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnProtectionFlowLogsRecordFields_SRCADDR,
//   	},
//   }
//
// Experimental.
type CfnProtectionFlowLogsS3Props struct {
	// Encrpytion key for your delivery bucket.
	// Experimental.
	EncryptionKey interfacesawskms.IKeyRef `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Format for log output, options are json,plain,w3c,parquet.
	// Experimental.
	OutputFormat CfnProtectionFlowLogsOutputFormat_S3 `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnProtectionFlowLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

