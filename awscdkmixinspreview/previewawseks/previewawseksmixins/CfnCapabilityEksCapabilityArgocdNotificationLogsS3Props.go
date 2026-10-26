package previewawseksmixins

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
//   cfnCapabilityEksCapabilityArgocdNotificationLogsS3Props := &CfnCapabilityEksCapabilityArgocdNotificationLogsS3Props{
//   	EncryptionKey: keyRef,
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnCapabilityEksCapabilityArgocdNotificationLogsOutputFormat.S3_PLAIN,
//   	RecordFields: []CfnCapabilityEksCapabilityArgocdNotificationLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnCapabilityEksCapabilityArgocdNotificationLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnCapabilityEksCapabilityArgocdNotificationLogsS3Props struct {
	// Encrpytion key for your delivery bucket.
	// Experimental.
	EncryptionKey interfacesawskms.IKeyRef `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Format for log output, options are plain,json,w3c,parquet.
	// Experimental.
	OutputFormat CfnCapabilityEksCapabilityArgocdNotificationLogsOutputFormat_S3 `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnCapabilityEksCapabilityArgocdNotificationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

