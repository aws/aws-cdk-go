package awsb2bi

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProfileProps := &CfnProfileProps{
//   	BusinessName: jsii.String("businessName"),
//   	Logging: jsii.String("logging"),
//   	Name: jsii.String("name"),
//   	Phone: jsii.String("phone"),
//
//   	// the properties below are optional
//   	Email: jsii.String("email"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-profile.html
//
type CfnProfileProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-profile.html#cfn-b2bi-profile-businessname
	//
	BusinessName *string `field:"required" json:"businessName" yaml:"businessName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-profile.html#cfn-b2bi-profile-logging
	//
	Logging *string `field:"required" json:"logging" yaml:"logging"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-profile.html#cfn-b2bi-profile-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-profile.html#cfn-b2bi-profile-phone
	//
	Phone *string `field:"required" json:"phone" yaml:"phone"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-profile.html#cfn-b2bi-profile-email
	//
	Email *string `field:"optional" json:"email" yaml:"email"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-profile.html#cfn-b2bi-profile-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

