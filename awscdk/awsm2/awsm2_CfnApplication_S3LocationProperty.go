package awsm2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationProperty := &s3LocationProperty{
//   	s3Location: jsii.String("s3Location"),
//   }
//
type CfnApplication_S3LocationProperty struct {
	// `CfnApplication.S3LocationProperty.S3Location`.
	S3Location *string `field:"required" json:"s3Location" yaml:"s3Location"`
}

