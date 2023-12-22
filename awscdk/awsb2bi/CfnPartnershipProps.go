package awsb2bi

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPartnership`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPartnershipProps := &CfnPartnershipProps{
//   	Email: jsii.String("email"),
//   	Name: jsii.String("name"),
//   	ProfileId: jsii.String("profileId"),
//
//   	// the properties below are optional
//   	Capabilities: []*string{
//   		jsii.String("capabilities"),
//   	},
//   	Phone: jsii.String("phone"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-partnership.html
//
type CfnPartnershipProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-partnership.html#cfn-b2bi-partnership-email
	//
	Email *string `field:"required" json:"email" yaml:"email"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-partnership.html#cfn-b2bi-partnership-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-partnership.html#cfn-b2bi-partnership-profileid
	//
	ProfileId *string `field:"required" json:"profileId" yaml:"profileId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-partnership.html#cfn-b2bi-partnership-capabilities
	//
	Capabilities *[]*string `field:"optional" json:"capabilities" yaml:"capabilities"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-partnership.html#cfn-b2bi-partnership-phone
	//
	Phone *string `field:"optional" json:"phone" yaml:"phone"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-partnership.html#cfn-b2bi-partnership-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

