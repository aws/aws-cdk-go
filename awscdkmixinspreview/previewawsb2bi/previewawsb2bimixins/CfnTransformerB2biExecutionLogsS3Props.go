package previewawsb2bimixins

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
//   cfnTransformerB2biExecutionLogsS3Props := &CfnTransformerB2biExecutionLogsS3Props{
//   	EncryptionKey: keyRef,
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnTransformerB2biExecutionLogsOutputFormat.S3_JSON,
//   }
//
// Experimental.
type CfnTransformerB2biExecutionLogsS3Props struct {
	// Encrpytion key for your delivery bucket.
	// Experimental.
	EncryptionKey interfacesawskms.IKeyRef `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Format for log output, options are json,plain,w3c,parquet.
	// Experimental.
	OutputFormat CfnTransformerB2biExecutionLogsOutputFormat_S3 `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

