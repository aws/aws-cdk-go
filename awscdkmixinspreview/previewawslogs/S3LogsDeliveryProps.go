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
//   	MandatoryFields: []*string{
//   		jsii.String("mandatoryFields"),
//   	},
//   	OutputFormat: jsii.String("outputFormat"),
//   	PermissionsVersion: awscdkmixinspreview.Aws_logs.S3LogsDeliveryPermissionsVersion_V1,
//   	ProvidedFields: []*string{
//   		jsii.String("providedFields"),
//   	},
//   }
//
// Experimental.
type S3LogsDeliveryProps struct {
	// Any recordFields that a mandatory to be included in a log delivery of a certain log type.
	// Default: - log type has no mandatory fields.
	//
	// Experimental.
	MandatoryFields *[]*string `field:"optional" json:"mandatoryFields" yaml:"mandatoryFields"`
	// RecordFields the user has defined to be used in log delivery.
	// Experimental.
	ProvidedFields *[]*string `field:"optional" json:"providedFields" yaml:"providedFields"`
	// Format of the logs that are sent to the delivery destination specified.
	// Experimental.
	OutputFormat *string `field:"optional" json:"outputFormat" yaml:"outputFormat"`
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

