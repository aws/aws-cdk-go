package previewawss3mixins

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
//   cfnBucketS3ServerAccessLogsS3Props := &CfnBucketS3ServerAccessLogsS3Props{
//   	EncryptionKey: keyRef,
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnBucketS3ServerAccessLogsOutputFormat.S3_JSON,
//   	RecordFields: []CfnBucketS3ServerAccessLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnBucketS3ServerAccessLogsRecordFields_BUCKET_NAME,
//   	},
//   }
//
// Experimental.
type CfnBucketS3ServerAccessLogsS3Props struct {
	// Encrpytion key for your delivery bucket.
	// Experimental.
	EncryptionKey interfacesawskms.IKeyRef `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Format for log output, options are json,parquet.
	// Experimental.
	OutputFormat CfnBucketS3ServerAccessLogsOutputFormat_S3 `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnBucketS3ServerAccessLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

