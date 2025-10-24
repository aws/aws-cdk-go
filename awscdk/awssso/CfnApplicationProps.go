package awssso

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &CfnApplicationProps{
//   	ApplicationProviderArn: jsii.String("applicationProviderArn"),
//   	InstanceArn: jsii.String("instanceArn"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	PortalOptions: &PortalOptionsConfigurationProperty{
//   		SignInOptions: &SignInOptionsProperty{
//   			Origin: jsii.String("origin"),
//
//   			// the properties below are optional
//   			ApplicationUrl: jsii.String("applicationUrl"),
//   		},
//   		Visibility: jsii.String("visibility"),
//   	},
//   	Status: jsii.String("status"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-application.html
//
type CfnApplicationProps struct {
	// The ARN of the application provider for this application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-application.html#cfn-sso-application-applicationproviderarn
	//
	ApplicationProviderArn *string `field:"required" json:"applicationProviderArn" yaml:"applicationProviderArn"`
	// The ARN of the instance of IAM Identity Center that is configured with this application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-application.html#cfn-sso-application-instancearn
	//
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-application.html#cfn-sso-application-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-application.html#cfn-sso-application-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A structure that describes the options for the access portal associated with this application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-application.html#cfn-sso-application-portaloptions
	//
	PortalOptions interface{} `field:"optional" json:"portalOptions" yaml:"portalOptions"`
	// The current status of the application in this instance of IAM Identity Center.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-application.html#cfn-sso-application-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Specifies tags to be attached to the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-application.html#cfn-sso-application-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

