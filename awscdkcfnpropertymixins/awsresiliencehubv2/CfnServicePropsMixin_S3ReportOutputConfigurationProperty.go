package awsresiliencehubv2


// S3 configuration for report output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   s3ReportOutputConfigurationProperty := &S3ReportOutputConfigurationProperty{
//   	BucketOwner: jsii.String("bucketOwner"),
//   	BucketPath: jsii.String("bucketPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-s3reportoutputconfiguration.html
//
type CfnServicePropsMixin_S3ReportOutputConfigurationProperty struct {
	// Account ID of the bucket owner.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-s3reportoutputconfiguration.html#cfn-resiliencehubv2-service-s3reportoutputconfiguration-bucketowner
	//
	BucketOwner *string `field:"optional" json:"bucketOwner" yaml:"bucketOwner"`
	// S3 bucket path where reports will be written.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-s3reportoutputconfiguration.html#cfn-resiliencehubv2-service-s3reportoutputconfiguration-bucketpath
	//
	BucketPath *string `field:"optional" json:"bucketPath" yaml:"bucketPath"`
}

