package awscleanrooms


// Contains configurations for protected query results.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   membershipProtectedQueryOutputConfigurationProperty := &MembershipProtectedQueryOutputConfigurationProperty{
//   	S3: &ProtectedQueryS3OutputConfigurationProperty{
//   		Bucket: jsii.String("bucket"),
//   		ResultFormat: jsii.String("resultFormat"),
//
//   		// the properties below are optional
//   		KeyPrefix: jsii.String("keyPrefix"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershipprotectedqueryoutputconfiguration.html
//
type CfnMembership_MembershipProtectedQueryOutputConfigurationProperty struct {
	// Required configuration for a protected query with an `S3` output type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershipprotectedqueryoutputconfiguration.html#cfn-cleanrooms-membership-membershipprotectedqueryoutputconfiguration-s3
	//
	S3 interface{} `field:"required" json:"s3" yaml:"s3"`
}

