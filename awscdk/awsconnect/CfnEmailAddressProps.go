package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnEmailAddress`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEmailAddressProps := &CfnEmailAddressProps{
//   	EmailAddress: jsii.String("emailAddress"),
//   	InstanceArn: jsii.String("instanceArn"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-emailaddress.html
//
type CfnEmailAddressProps struct {
	// The email address, including the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-emailaddress.html#cfn-connect-emailaddress-emailaddress
	//
	EmailAddress *string `field:"required" json:"emailAddress" yaml:"emailAddress"`
	// The Amazon Resource Name (ARN) of the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-emailaddress.html#cfn-connect-emailaddress-instancearn
	//
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The description of the email address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-emailaddress.html#cfn-connect-emailaddress-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name of email address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-emailaddress.html#cfn-connect-emailaddress-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-emailaddress.html#cfn-connect-emailaddress-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

