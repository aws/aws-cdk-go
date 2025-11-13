package awsappstream

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnStack`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStackProps := &CfnStackProps{
//   	AccessEndpoints: []interface{}{
//   		&AccessEndpointProperty{
//   			EndpointType: jsii.String("endpointType"),
//   			VpceId: jsii.String("vpceId"),
//   		},
//   	},
//   	ApplicationSettings: &ApplicationSettingsProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		SettingsGroup: jsii.String("settingsGroup"),
//   	},
//   	AttributesToDelete: []*string{
//   		jsii.String("attributesToDelete"),
//   	},
//   	DeleteStorageConnectors: jsii.Boolean(false),
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	EmbedHostDomains: []*string{
//   		jsii.String("embedHostDomains"),
//   	},
//   	FeedbackUrl: jsii.String("feedbackUrl"),
//   	Name: jsii.String("name"),
//   	RedirectUrl: jsii.String("redirectUrl"),
//   	StorageConnectors: []interface{}{
//   		&StorageConnectorProperty{
//   			ConnectorType: jsii.String("connectorType"),
//
//   			// the properties below are optional
//   			Domains: []*string{
//   				jsii.String("domains"),
//   			},
//   			ResourceIdentifier: jsii.String("resourceIdentifier"),
//   		},
//   	},
//   	StreamingExperienceSettings: &StreamingExperienceSettingsProperty{
//   		PreferredProtocol: jsii.String("preferredProtocol"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UserSettings: []interface{}{
//   		&UserSettingProperty{
//   			Action: jsii.String("action"),
//   			Permission: jsii.String("permission"),
//
//   			// the properties below are optional
//   			MaximumLength: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stack.html
//
type CfnStackProps struct {
	// The list of virtual private cloud (VPC) interface endpoint objects.
	//
	// Users of the stack can connect to WorkSpaces Applications only through the specified endpoints.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stack.html#cfn-appstream-stack-accessendpoints
	//
	AccessEndpoints interface{} `field:"optional" json:"accessEndpoints" yaml:"accessEndpoints"`
	// The persistent application settings for users of the stack.
	//
	// When these settings are enabled, changes that users make to applications and Windows settings are automatically saved after each session and applied to the next session.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stack.html#cfn-appstream-stack-applicationsettings
	//
	ApplicationSettings interface{} `field:"optional" json:"applicationSettings" yaml:"applicationSettings"`
	// The stack attributes to delete.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stack.html#cfn-appstream-stack-attributestodelete
	//
	AttributesToDelete *[]*string `field:"optional" json:"attributesToDelete" yaml:"attributesToDelete"`
	// *This parameter has been deprecated.*.
	//
	// Deletes the storage connectors currently enabled for the stack.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stack.html#cfn-appstream-stack-deletestorageconnectors
	//
	DeleteStorageConnectors interface{} `field:"optional" json:"deleteStorageConnectors" yaml:"deleteStorageConnectors"`
	// The description to display.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stack.html#cfn-appstream-stack-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The stack name to display.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stack.html#cfn-appstream-stack-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The domains where WorkSpaces Applications streaming sessions can be embedded in an iframe.
	//
	// You must approve the domains that you want to host embedded WorkSpaces Applications streaming sessions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stack.html#cfn-appstream-stack-embedhostdomains
	//
	EmbedHostDomains *[]*string `field:"optional" json:"embedHostDomains" yaml:"embedHostDomains"`
	// The URL that users are redirected to after they click the Send Feedback link.
	//
	// If no URL is specified, no Send Feedback link is displayed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stack.html#cfn-appstream-stack-feedbackurl
	//
	FeedbackUrl *string `field:"optional" json:"feedbackUrl" yaml:"feedbackUrl"`
	// The name of the stack.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stack.html#cfn-appstream-stack-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The URL that users are redirected to after their streaming session ends.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stack.html#cfn-appstream-stack-redirecturl
	//
	RedirectUrl *string `field:"optional" json:"redirectUrl" yaml:"redirectUrl"`
	// The storage connectors to enable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stack.html#cfn-appstream-stack-storageconnectors
	//
	StorageConnectors interface{} `field:"optional" json:"storageConnectors" yaml:"storageConnectors"`
	// The streaming protocol that you want your stack to prefer.
	//
	// This can be UDP or TCP. Currently, UDP is only supported in the Windows native client.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stack.html#cfn-appstream-stack-streamingexperiencesettings
	//
	StreamingExperienceSettings interface{} `field:"optional" json:"streamingExperienceSettings" yaml:"streamingExperienceSettings"`
	// An array of key-value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stack.html#cfn-appstream-stack-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The actions that are enabled or disabled for users during their streaming sessions.
	//
	// By default, these actions are enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stack.html#cfn-appstream-stack-usersettings
	//
	UserSettings interface{} `field:"optional" json:"userSettings" yaml:"userSettings"`
}

