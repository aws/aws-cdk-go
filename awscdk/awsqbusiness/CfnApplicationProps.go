package awsqbusiness

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
//   	DisplayName: jsii.String("displayName"),
//
//   	// the properties below are optional
//   	AttachmentsConfiguration: &AttachmentsConfigurationProperty{
//   		AttachmentsControlMode: jsii.String("attachmentsControlMode"),
//   	},
//   	AutoSubscriptionConfiguration: &AutoSubscriptionConfigurationProperty{
//   		AutoSubscribe: jsii.String("autoSubscribe"),
//
//   		// the properties below are optional
//   		DefaultSubscriptionType: jsii.String("defaultSubscriptionType"),
//   	},
//   	ClientIdsForOidc: []*string{
//   		jsii.String("clientIdsForOidc"),
//   	},
//   	Description: jsii.String("description"),
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	IamIdentityProviderArn: jsii.String("iamIdentityProviderArn"),
//   	IdentityCenterInstanceArn: jsii.String("identityCenterInstanceArn"),
//   	IdentityType: jsii.String("identityType"),
//   	PersonalizationConfiguration: &PersonalizationConfigurationProperty{
//   		PersonalizationControlMode: jsii.String("personalizationControlMode"),
//   	},
//   	QAppsConfiguration: &QAppsConfigurationProperty{
//   		QAppsControlMode: jsii.String("qAppsControlMode"),
//   	},
//   	QuickSightConfiguration: &QuickSightConfigurationProperty{
//   		ClientNamespace: jsii.String("clientNamespace"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-application.html
//
type CfnApplicationProps struct {
	// The name of the Amazon Q Business application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-application.html#cfn-qbusiness-application-displayname
	//
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Configuration information for the file upload during chat feature.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-application.html#cfn-qbusiness-application-attachmentsconfiguration
	//
	AttachmentsConfiguration interface{} `field:"optional" json:"attachmentsConfiguration" yaml:"attachmentsConfiguration"`
	// Subscription configuration information for an Amazon Q Business application using IAM identity federation for user management.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-application.html#cfn-qbusiness-application-autosubscriptionconfiguration
	//
	AutoSubscriptionConfiguration interface{} `field:"optional" json:"autoSubscriptionConfiguration" yaml:"autoSubscriptionConfiguration"`
	// The OIDC client ID for a Amazon Q Business application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-application.html#cfn-qbusiness-application-clientidsforoidc
	//
	ClientIdsForOidc *[]*string `field:"optional" json:"clientIdsForOidc" yaml:"clientIdsForOidc"`
	// A description for the Amazon Q Business application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-application.html#cfn-qbusiness-application-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Provides the identifier of the AWS KMS key used to encrypt data indexed by Amazon Q Business.
	//
	// Amazon Q Business doesn't support asymmetric keys.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-application.html#cfn-qbusiness-application-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The Amazon Resource Name (ARN) of an identity provider being used by an Amazon Q Business application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-application.html#cfn-qbusiness-application-iamidentityproviderarn
	//
	IamIdentityProviderArn *string `field:"optional" json:"iamIdentityProviderArn" yaml:"iamIdentityProviderArn"`
	// The Amazon Resource Name (ARN) of the IAM Identity Center instance you are either creating for—or connecting to—your Amazon Q Business application.
	//
	// *Required* : `Yes`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-application.html#cfn-qbusiness-application-identitycenterinstancearn
	//
	IdentityCenterInstanceArn *string `field:"optional" json:"identityCenterInstanceArn" yaml:"identityCenterInstanceArn"`
	// The authentication type being used by a Amazon Q Business application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-application.html#cfn-qbusiness-application-identitytype
	//
	IdentityType *string `field:"optional" json:"identityType" yaml:"identityType"`
	// Configuration information about chat response personalization.
	//
	// For more information, see [Personalizing chat responses](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/personalizing-chat-responses.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-application.html#cfn-qbusiness-application-personalizationconfiguration
	//
	PersonalizationConfiguration interface{} `field:"optional" json:"personalizationConfiguration" yaml:"personalizationConfiguration"`
	// Configuration information about Amazon Q Apps.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-application.html#cfn-qbusiness-application-qappsconfiguration
	//
	QAppsConfiguration interface{} `field:"optional" json:"qAppsConfiguration" yaml:"qAppsConfiguration"`
	// The Amazon Quick Suite configuration for an Amazon Q Business application that uses Quick Suite as the identity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-application.html#cfn-qbusiness-application-quicksightconfiguration
	//
	QuickSightConfiguration interface{} `field:"optional" json:"quickSightConfiguration" yaml:"quickSightConfiguration"`
	// The Amazon Resource Name (ARN) of an IAM role with permissions to access your Amazon CloudWatch logs and metrics.
	//
	// If this property is not specified, Amazon Q Business will create a [service linked role (SLR)](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/using-service-linked-roles.html#slr-permissions) and use it as the application's role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-application.html#cfn-qbusiness-application-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// A list of key-value pairs that identify or categorize your Amazon Q Business application.
	//
	// You can also use tags to help control access to the application. Tag keys and values can consist of Unicode letters, digits, white space, and any of the following symbols: _ . : / = + -
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-qbusiness-application.html#cfn-qbusiness-application-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

