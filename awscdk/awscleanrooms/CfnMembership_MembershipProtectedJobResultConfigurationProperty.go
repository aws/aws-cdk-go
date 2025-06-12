package awscleanrooms


// Contains configurations for protected job results.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   membershipProtectedJobResultConfigurationProperty := &MembershipProtectedJobResultConfigurationProperty{
//   	OutputConfiguration: &MembershipProtectedJobOutputConfigurationProperty{
//   		S3: &ProtectedJobS3OutputConfigurationInputProperty{
//   			Bucket: jsii.String("bucket"),
//
//   			// the properties below are optional
//   			KeyPrefix: jsii.String("keyPrefix"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershipprotectedjobresultconfiguration.html
//
type CfnMembership_MembershipProtectedJobResultConfigurationProperty struct {
	// The output configuration for a protected job result.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershipprotectedjobresultconfiguration.html#cfn-cleanrooms-membership-membershipprotectedjobresultconfiguration-outputconfiguration
	//
	OutputConfiguration interface{} `field:"required" json:"outputConfiguration" yaml:"outputConfiguration"`
	// The unique ARN for an IAM role that is used by AWS Clean Rooms to write protected job results to the result location, given by the member who can receive results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershipprotectedjobresultconfiguration.html#cfn-cleanrooms-membership-membershipprotectedjobresultconfiguration-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

