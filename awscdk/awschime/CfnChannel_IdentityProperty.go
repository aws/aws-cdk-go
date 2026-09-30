package awschime


// The AppInstanceUser or AppInstanceBot that created the channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   identityProperty := &IdentityProperty{
//   	Arn: jsii.String("arn"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-channel-identity.html
//
type CfnChannel_IdentityProperty struct {
	// The ARN in an identity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-channel-identity.html#cfn-chime-channel-identity-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The name in an identity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-channel-identity.html#cfn-chime-channel-identity-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

