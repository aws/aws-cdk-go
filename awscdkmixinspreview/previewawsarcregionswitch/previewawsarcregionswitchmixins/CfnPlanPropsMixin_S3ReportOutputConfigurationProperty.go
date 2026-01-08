package previewawsarcregionswitchmixins


// Configuration for delivering generated reports to an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3ReportOutputConfigurationProperty := &S3ReportOutputConfigurationProperty{
//   	BucketOwner: jsii.String("bucketOwner"),
//   	BucketPath: jsii.String("bucketPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-s3reportoutputconfiguration.html
//
type CfnPlanPropsMixin_S3ReportOutputConfigurationProperty struct {
	// The AWS account ID that owns the S3 bucket.
	//
	// Required to ensure the bucket is still owned by the same expected owner at generation time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-s3reportoutputconfiguration.html#cfn-arcregionswitch-plan-s3reportoutputconfiguration-bucketowner
	//
	BucketOwner *string `field:"optional" json:"bucketOwner" yaml:"bucketOwner"`
	// The S3 bucket name and optional prefix where reports are stored.
	//
	// Format: bucket-name or bucket-name/prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-s3reportoutputconfiguration.html#cfn-arcregionswitch-plan-s3reportoutputconfiguration-bucketpath
	//
	BucketPath *string `field:"optional" json:"bucketPath" yaml:"bucketPath"`
}

