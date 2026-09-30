package awsmacie

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnMemberPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnMemberMixinProps := &CfnMemberMixinProps{
//   	AccountId: jsii.String("accountId"),
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
type CfnMemberMixinProps struct {
	// The AWS account ID for the account to associate with the Amazon Macie administrator account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-member.html#cfn-macie-member-accountid
	//
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
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

