package previewawscleanroomsmixins


// Contains configurations for protected query results.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   membershipProtectedQueryOutputConfigurationProperty := &MembershipProtectedQueryOutputConfigurationProperty{
//   	S3: &ProtectedQueryS3OutputConfigurationProperty{
//   		Bucket: jsii.String("bucket"),
//   		KeyPrefix: jsii.String("keyPrefix"),
//   		ResultFormat: jsii.String("resultFormat"),
//   		SingleFileOutput: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershipprotectedqueryoutputconfiguration.html
//
type CfnMembershipPropsMixin_MembershipProtectedQueryOutputConfigurationProperty struct {
	// Required configuration for a protected query with an `s3` output type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-membership-membershipprotectedqueryoutputconfiguration.html#cfn-cleanrooms-membership-membershipprotectedqueryoutputconfiguration-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

