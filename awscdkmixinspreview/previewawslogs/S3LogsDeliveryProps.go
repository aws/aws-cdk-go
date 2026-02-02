package previewawslogs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskms"
)

// Props for S3LogsDelivery.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var keyRef IKeyRef
//
//   s3LogsDeliveryProps := &S3LogsDeliveryProps{
//   	KmsKey: keyRef,
//   	PermissionsVersion: awscdkmixinspreview.Aws_logs.S3LogsDeliveryPermissionsVersion_V1,
//   }
//
// Experimental.
type S3LogsDeliveryProps struct {
	// KMS key to use for encrypting logs in the S3 bucket.
	//
	// When provided, grants the logs delivery service permissions to use the key.
	// Default: - No encryption key is configured.
	//
	// Experimental.
	KmsKey interfacesawskms.IKeyRef `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The permissions version ('V1' or 'V2') to be used for this delivery.
	//
	// Depending on the source of the logs, different permissions are required.
	// Default: "V2".
	//
	// Experimental.
	PermissionsVersion S3LogsDeliveryPermissionsVersion `field:"optional" json:"permissionsVersion" yaml:"permissionsVersion"`
}

