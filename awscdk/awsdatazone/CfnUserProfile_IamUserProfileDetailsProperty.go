package awsdatazone


// The details of an IAM user profile in Amazon DataZone.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iamUserProfileDetailsProperty := &IamUserProfileDetailsProperty{
//   	Arn: jsii.String("arn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-userprofile-iamuserprofiledetails.html
//
type CfnUserProfile_IamUserProfileDetailsProperty struct {
	// The ARN of an IAM user profile in Amazon DataZone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-userprofile-iamuserprofiledetails.html#cfn-datazone-userprofile-iamuserprofiledetails-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

