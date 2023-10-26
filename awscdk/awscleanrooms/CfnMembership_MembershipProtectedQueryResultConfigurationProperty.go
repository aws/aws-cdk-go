package awscleanrooms


// Contains configurations for protected query results.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   membershipProtectedQueryResultConfigurationProperty := &MembershipProtectedQueryResultConfigurationProperty{
//   	OutputConfiguration: &MembershipProtectedQueryOutputConfigurationProperty{
//   		S3: &ProtectedQueryS3OutputConfigurationProperty{
//   			Bucket: jsii.String("bucket"),
//   			ResultFormat: jsii.String("resultFormat"),
//
//   			// the properties below are optional
//   			KeyPrefix: jsii.String("keyPrefix"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershipprotectedqueryresultconfiguration.html
//
type CfnMembership_MembershipProtectedQueryResultConfigurationProperty struct {
	// Configuration for protected query results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershipprotectedqueryresultconfiguration.html#cfn-cleanrooms-membership-membershipprotectedqueryresultconfiguration-outputconfiguration
	//
	OutputConfiguration interface{} `field:"required" json:"outputConfiguration" yaml:"outputConfiguration"`
	// The unique ARN for an IAM role that is used by to write protected query results to the result location, given by the member who can receive results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershipprotectedqueryresultconfiguration.html#cfn-cleanrooms-membership-membershipprotectedqueryresultconfiguration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

