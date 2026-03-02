package previewawsroute53profilesmixins

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
//   cfnProfileRoute53ProfilesResolverQueryLogsS3Props := &CfnProfileRoute53ProfilesResolverQueryLogsS3Props{
//   	EncryptionKey: keyRef,
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat.S3_JSON,
//   }
//
// Experimental.
type CfnProfileRoute53ProfilesResolverQueryLogsS3Props struct {
	// Encrpytion key for your delivery bucket.
	// Experimental.
	EncryptionKey interfacesawskms.IKeyRef `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Format for log output, options are json,plain,w3c,parquet.
	// Experimental.
	OutputFormat CfnProfileRoute53ProfilesResolverQueryLogsOutputFormat_S3 `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

