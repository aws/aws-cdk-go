package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnUserProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserProfileProps := &cfnUserProfileProps{
//   	domainId: jsii.String("domainId"),
//   	userProfileName: jsii.String("userProfileName"),
//
//   	// the properties below are optional
//   	singleSignOnUserIdentifier: jsii.String("singleSignOnUserIdentifier"),
//   	singleSignOnUserValue: jsii.String("singleSignOnUserValue"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	userSettings: &userSettingsProperty{
//   		executionRole: jsii.String("executionRole"),
//   		jupyterServerAppSettings: &jupyterServerAppSettingsProperty{
//   			defaultResourceSpec: &resourceSpecProperty{
//   				instanceType: jsii.String("instanceType"),
//   				sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   		},
//   		kernelGatewayAppSettings: &kernelGatewayAppSettingsProperty{
//   			customImages: []interface{}{
//   				&customImageProperty{
//   					appImageConfigName: jsii.String("appImageConfigName"),
//   					imageName: jsii.String("imageName"),
//
//   					// the properties below are optional
//   					imageVersionNumber: jsii.Number(123),
//   				},
//   			},
//   			defaultResourceSpec: &resourceSpecProperty{
//   				instanceType: jsii.String("instanceType"),
//   				sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   				sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   			},
//   		},
//   		rStudioServerProAppSettings: &rStudioServerProAppSettingsProperty{
//   			accessStatus: jsii.String("accessStatus"),
//   			userGroup: jsii.String("userGroup"),
//   		},
//   		securityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		sharingSettings: &sharingSettingsProperty{
//   			notebookOutputOption: jsii.String("notebookOutputOption"),
//   			s3KmsKeyId: jsii.String("s3KmsKeyId"),
//   			s3OutputPath: jsii.String("s3OutputPath"),
//   		},
//   	},
//   }
//
type CfnUserProfileProps struct {
	// The domain ID.
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// The user profile name.
	UserProfileName *string `field:"required" json:"userProfileName" yaml:"userProfileName"`
	// A specifier for the type of value specified in SingleSignOnUserValue.
	//
	// Currently, the only supported value is "UserName". If the Domain's AuthMode is SSO, this field is required. If the Domain's AuthMode is not SSO, this field cannot be specified.
	SingleSignOnUserIdentifier *string `field:"optional" json:"singleSignOnUserIdentifier" yaml:"singleSignOnUserIdentifier"`
	// The username of the associated AWS Single Sign-On User for this UserProfile.
	//
	// If the Domain's AuthMode is SSO, this field is required, and must match a valid username of a user in your directory. If the Domain's AuthMode is not SSO, this field cannot be specified.
	SingleSignOnUserValue *string `field:"optional" json:"singleSignOnUserValue" yaml:"singleSignOnUserValue"`
	// An array of key-value pairs to apply to this resource.
	//
	// Tags that you specify for the User Profile are also added to all apps that the User Profile launches.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A collection of settings that apply to users of Amazon SageMaker Studio.
	UserSettings interface{} `field:"optional" json:"userSettings" yaml:"userSettings"`
}

