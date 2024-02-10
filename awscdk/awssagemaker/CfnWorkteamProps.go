package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnWorkteam`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWorkteamProps := &CfnWorkteamProps{
//   	Description: jsii.String("description"),
//   	MemberDefinitions: []interface{}{
//   		&MemberDefinitionProperty{
//   			CognitoMemberDefinition: &CognitoMemberDefinitionProperty{
//   				CognitoClientId: jsii.String("cognitoClientId"),
//   				CognitoUserGroup: jsii.String("cognitoUserGroup"),
//   				CognitoUserPool: jsii.String("cognitoUserPool"),
//   			},
//   			OidcMemberDefinition: &OidcMemberDefinitionProperty{
//   				OidcGroups: []*string{
//   					jsii.String("oidcGroups"),
//   				},
//   			},
//   		},
//   	},
//   	NotificationConfiguration: &NotificationConfigurationProperty{
//   		NotificationTopicArn: jsii.String("notificationTopicArn"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WorkforceName: jsii.String("workforceName"),
//   	WorkteamName: jsii.String("workteamName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-workteam.html
//
type CfnWorkteamProps struct {
	// A description of the work team.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-workteam.html#cfn-sagemaker-workteam-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of `MemberDefinition` objects that contains objects that identify the workers that make up the work team.
	//
	// Workforces can be created using Amazon Cognito or your own OIDC Identity Provider (IdP). For private workforces created using Amazon Cognito use `CognitoMemberDefinition` . For workforces created using your own OIDC identity provider (IdP) use `OidcMemberDefinition` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-workteam.html#cfn-sagemaker-workteam-memberdefinitions
	//
	MemberDefinitions interface{} `field:"optional" json:"memberDefinitions" yaml:"memberDefinitions"`
	// Configures SNS notifications of available or expiring work items for work teams.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-workteam.html#cfn-sagemaker-workteam-notificationconfiguration
	//
	NotificationConfiguration interface{} `field:"optional" json:"notificationConfiguration" yaml:"notificationConfiguration"`
	// An array of key-value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-workteam.html#cfn-sagemaker-workteam-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the workforce.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-workteam.html#cfn-sagemaker-workteam-workforcename
	//
	WorkforceName *string `field:"optional" json:"workforceName" yaml:"workforceName"`
	// The name of the work team.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-workteam.html#cfn-sagemaker-workteam-workteamname
	//
	WorkteamName *string `field:"optional" json:"workteamName" yaml:"workteamName"`
}

