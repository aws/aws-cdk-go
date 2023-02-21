// The CDK Construct Library for AWS::S3ObjectLambda
package awscdks3objectlambdaalpha


// The access point resource attributes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import s3objectlambda_alpha "github.com/aws/aws-cdk-go/awscdks3objectlambdaalpha"
//
//   accessPointAttributes := &AccessPointAttributes{
//   	AccessPointArn: jsii.String("accessPointArn"),
//   	AccessPointCreationDate: jsii.String("accessPointCreationDate"),
//   }
//
// Experimental.
type AccessPointAttributes struct {
	// The ARN of the access point.
	// Experimental.
	AccessPointArn *string `field:"required" json:"accessPointArn" yaml:"accessPointArn"`
	// The creation data of the access point.
	// Experimental.
	AccessPointCreationDate *string `field:"required" json:"accessPointCreationDate" yaml:"accessPointCreationDate"`
}

