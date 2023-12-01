package awss3


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessGrantsLocationConfigurationProperty := &AccessGrantsLocationConfigurationProperty{
//   	S3SubPrefix: jsii.String("s3SubPrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-accessgrant-accessgrantslocationconfiguration.html
//
type CfnAccessGrant_AccessGrantsLocationConfigurationProperty struct {
	// The S3 sub prefix of a registered location in your S3 Access Grants instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-accessgrant-accessgrantslocationconfiguration.html#cfn-s3-accessgrant-accessgrantslocationconfiguration-s3subprefix
	//
	S3SubPrefix *string `field:"required" json:"s3SubPrefix" yaml:"s3SubPrefix"`
}

