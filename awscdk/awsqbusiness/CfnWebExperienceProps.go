package awsqbusiness

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnWebExperience`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWebExperienceProps := &CfnWebExperienceProps{
//   	ApplicationId: jsii.String("applicationId"),
//
//   	// the properties below are optional
//   	CustomizationConfiguration: &CustomizationConfigurationProperty{
//   		CustomCssUrl: jsii.String("customCssUrl"),
//   		FaviconUrl: jsii.String("faviconUrl"),
//   		FontUrl: jsii.String("fontUrl"),
//   		LogoUrl: jsii.String("logoUrl"),
//   	},
//   	IdentityProviderConfiguration: &IdentityProviderConfigurationProperty{
//   		OpenIdConnectConfiguration: &OpenIDConnectProviderConfigurationProperty{
//   			SecretsArn: jsii.String("secretsArn"),
//   			SecretsRole: jsii.String("secretsRole"),
//   		},
//   		SamlConfiguration: &SamlProviderConfigurationProperty{
//   			AuthenticationUrl: jsii.String("authenticationUrl"),
//   		},
//   	},
//   	Origins: []*string{
//   		jsii.String("origins"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	SamplePromptsControlMode: jsii.String("samplePromptsControlMode"),
//   	Subtitle: jsii.String("subtitle"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Title: jsii.String("title"),
//   	WelcomeMessage: jsii.String("welcomeMessage"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-webexperience.html
//
type CfnWebExperienceProps struct {
	// The identifier of the Amazon Q Business web experience.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-webexperience.html#cfn-qbusiness-webexperience-applicationid
	//
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-webexperience.html#cfn-qbusiness-webexperience-customizationconfiguration
	//
	CustomizationConfiguration interface{} `field:"optional" json:"customizationConfiguration" yaml:"customizationConfiguration"`
	// Provides information about the identity provider (IdP) used to authenticate end users of an Amazon Q Business web experience.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-webexperience.html#cfn-qbusiness-webexperience-identityproviderconfiguration
	//
	IdentityProviderConfiguration interface{} `field:"optional" json:"identityProviderConfiguration" yaml:"identityProviderConfiguration"`
	// Sets the website domain origins that are allowed to embed the Amazon Q Business web experience.
	//
	// The *domain origin* refers to the base URL for accessing a website including the protocol ( `http/https` ), the domain name, and the port number (if specified).
	//
	// > You must only submit a *base URL* and not a full path. For example, `https://docs.aws.amazon.com` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-webexperience.html#cfn-qbusiness-webexperience-origins
	//
	Origins *[]*string `field:"optional" json:"origins" yaml:"origins"`
	// The Amazon Resource Name (ARN) of the service role attached to your web experience.
	//
	// > You must provide this value if you're using IAM Identity Center to manage end user access to your application. If you're using legacy identity management to manage user access, you don't need to provide this value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-webexperience.html#cfn-qbusiness-webexperience-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Determines whether sample prompts are enabled in the web experience for an end user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-webexperience.html#cfn-qbusiness-webexperience-samplepromptscontrolmode
	//
	SamplePromptsControlMode *string `field:"optional" json:"samplePromptsControlMode" yaml:"samplePromptsControlMode"`
	// A subtitle to personalize your Amazon Q Business web experience.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-webexperience.html#cfn-qbusiness-webexperience-subtitle
	//
	Subtitle *string `field:"optional" json:"subtitle" yaml:"subtitle"`
	// A list of key-value pairs that identify or categorize your Amazon Q Business web experience.
	//
	// You can also use tags to help control access to the web experience. Tag keys and values can consist of Unicode letters, digits, white space, and any of the following symbols: _ . : / = + -
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-webexperience.html#cfn-qbusiness-webexperience-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The title for your Amazon Q Business web experience.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-webexperience.html#cfn-qbusiness-webexperience-title
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
	// A message in an Amazon Q Business web experience.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-webexperience.html#cfn-qbusiness-webexperience-welcomemessage
	//
	WelcomeMessage *string `field:"optional" json:"welcomeMessage" yaml:"welcomeMessage"`
}

