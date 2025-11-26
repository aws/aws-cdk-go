package previewawslogs


// Props for S3LogsDelivery.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3LogsDeliveryProps := &S3LogsDeliveryProps{
//   	PermissionsVersion: awscdkmixinspreview.Aws_logs.S3LogsDeliveryPermissionsVersion_V1,
//   }
//
// Experimental.
type S3LogsDeliveryProps struct {
	// The permissions version ('V1' or 'V2') to be used for this delivery.
	//
	// Depending on the source of the logs, different permissions are required.
	// Default: "V2".
	//
	// Experimental.
	PermissionsVersion S3LogsDeliveryPermissionsVersion `field:"optional" json:"permissionsVersion" yaml:"permissionsVersion"`
}

