package previewawscloudfrontmixins

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
//   cfnDistributionConnectionLogsS3Props := &CfnDistributionConnectionLogsS3Props{
//   	EncryptionKey: keyRef,
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnDistributionConnectionLogsOutputFormat.S3_JSON,
//   	RecordFields: []CfnDistributionConnectionLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnDistributionConnectionLogsRecordFields_CONNECTIONSTATUS,
//   	},
//   }
//
// Experimental.
type CfnDistributionConnectionLogsS3Props struct {
	// Encrpytion key for your delivery bucket.
	// Experimental.
	EncryptionKey interfacesawskms.IKeyRef `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Format for log output, options are json,plain,w3c,parquet.
	// Experimental.
	OutputFormat CfnDistributionConnectionLogsOutputFormat_S3 `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnDistributionConnectionLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

