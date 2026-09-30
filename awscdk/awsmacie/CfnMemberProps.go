package awsmacie

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMember`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMemberProps := &CfnMemberProps{
//   	AccountId: jsii.String("accountId"),
//
//   	// the properties below are optional
//   	Email: jsii.String("email"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-member.html
//
type CfnMemberProps struct {
	// The AWS account ID for the account to associate with the Amazon Macie administrator account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-member.html#cfn-macie-member-accountid
	//
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The email address for the account to associate with the Amazon Macie administrator account.
	//
	// Required by the Amazon Macie CreateMember API at creation time; it is write-only because the service does not return it (it is null when the account is associated through AWS Organizations).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-member.html#cfn-macie-member-email
	//
	Email *string `field:"optional" json:"email" yaml:"email"`
	// The tags to associate with the member account in Amazon Macie.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-member.html#cfn-macie-member-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

