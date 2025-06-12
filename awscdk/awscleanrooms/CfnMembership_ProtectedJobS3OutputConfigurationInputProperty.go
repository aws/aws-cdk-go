package awscleanrooms


// Contains input information for protected jobs with an S3 output type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   protectedJobS3OutputConfigurationInputProperty := &ProtectedJobS3OutputConfigurationInputProperty{
//   	Bucket: jsii.String("bucket"),
//
//   	// the properties below are optional
//   	KeyPrefix: jsii.String("keyPrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-protectedjobs3outputconfigurationinput.html
//
type CfnMembership_ProtectedJobS3OutputConfigurationInputProperty struct {
	// The S3 bucket for job output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-protectedjobs3outputconfigurationinput.html#cfn-cleanrooms-membership-protectedjobs3outputconfigurationinput-bucket
	//
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The S3 prefix to unload the protected job results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-protectedjobs3outputconfigurationinput.html#cfn-cleanrooms-membership-protectedjobs3outputconfigurationinput-keyprefix
	//
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
}

