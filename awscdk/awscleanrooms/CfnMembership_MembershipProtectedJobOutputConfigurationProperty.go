package awscleanrooms


// Contains configurations for protected job results.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   membershipProtectedJobOutputConfigurationProperty := &MembershipProtectedJobOutputConfigurationProperty{
//   	S3: &ProtectedJobS3OutputConfigurationInputProperty{
//   		Bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		KeyPrefix: jsii.String("keyPrefix"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershipprotectedjoboutputconfiguration.html
//
type CfnMembership_MembershipProtectedJobOutputConfigurationProperty struct {
	// Contains the configuration to write the job results to S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershipprotectedjoboutputconfiguration.html#cfn-cleanrooms-membership-membershipprotectedjoboutputconfiguration-s3
	//
	S3 interface{} `field:"required" json:"s3" yaml:"s3"`
}

