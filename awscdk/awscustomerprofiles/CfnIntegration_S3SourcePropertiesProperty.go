package awscustomerprofiles


// The properties that are applied when Amazon S3 is being used as the flow source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3SourcePropertiesProperty := &S3SourcePropertiesProperty{
//   	BucketName: jsii.String("bucketName"),
//
//   	// the properties below are optional
//   	BucketPrefix: jsii.String("bucketPrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-s3sourceproperties.html
//
type CfnIntegration_S3SourcePropertiesProperty struct {
	// The Amazon S3 bucket name where the source files are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-s3sourceproperties.html#cfn-customerprofiles-integration-s3sourceproperties-bucketname
	//
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The object key for the Amazon S3 bucket in which the source files are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-customerprofiles-integration-s3sourceproperties.html#cfn-customerprofiles-integration-s3sourceproperties-bucketprefix
	//
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
}

