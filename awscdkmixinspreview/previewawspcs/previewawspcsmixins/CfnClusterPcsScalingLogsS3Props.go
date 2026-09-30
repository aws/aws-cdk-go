package previewawspcsmixins

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
//   cfnClusterPcsScalingLogsS3Props := &CfnClusterPcsScalingLogsS3Props{
//   	EncryptionKey: keyRef,
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnClusterPcsScalingLogsOutputFormat.S3_PLAIN,
//   	RecordFields: []CfnClusterPcsScalingLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnClusterPcsScalingLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnClusterPcsScalingLogsS3Props struct {
	// Encrpytion key for your delivery bucket.
	// Experimental.
	EncryptionKey interfacesawskms.IKeyRef `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Format for log output, options are plain,json,w3c,parquet.
	// Experimental.
	OutputFormat CfnClusterPcsScalingLogsOutputFormat_S3 `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnClusterPcsScalingLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

