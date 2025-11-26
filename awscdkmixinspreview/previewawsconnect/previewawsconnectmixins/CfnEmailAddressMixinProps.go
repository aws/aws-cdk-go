package previewawsconnectmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnEmailAddressPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEmailAddressMixinProps := &CfnEmailAddressMixinProps{
//   	AliasConfigurations: []interface{}{
//   		&AliasConfigurationProperty{
//   			EmailAddressArn: jsii.String("emailAddressArn"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	EmailAddress: jsii.String("emailAddress"),
//   	InstanceArn: jsii.String("instanceArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-emailaddress.html
//
type CfnEmailAddressMixinProps struct {
	// A list of alias configurations for this email address, showing which email addresses forward to this primary address.
	//
	// Each configuration contains the email address ID of an alias that forwards emails to this address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-emailaddress.html#cfn-connect-emailaddress-aliasconfigurations
	//
	AliasConfigurations interface{} `field:"optional" json:"aliasConfigurations" yaml:"aliasConfigurations"`
	// The description of the email address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-emailaddress.html#cfn-connect-emailaddress-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name of email address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-emailaddress.html#cfn-connect-emailaddress-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The email address, including the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-emailaddress.html#cfn-connect-emailaddress-emailaddress
	//
	EmailAddress *string `field:"optional" json:"emailAddress" yaml:"emailAddress"`
	// The Amazon Resource Name (ARN) of the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-emailaddress.html#cfn-connect-emailaddress-instancearn
	//
	InstanceArn *string `field:"optional" json:"instanceArn" yaml:"instanceArn"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-emailaddress.html#cfn-connect-emailaddress-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

