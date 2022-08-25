package awsappstream

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnStack`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStackProps := &cfnStackProps{
//   	accessEndpoints: []interface{}{
//   		&accessEndpointProperty{
//   			endpointType: jsii.String("endpointType"),
//   			vpceId: jsii.String("vpceId"),
//   		},
//   	},
//   	applicationSettings: &applicationSettingsProperty{
//   		enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		settingsGroup: jsii.String("settingsGroup"),
//   	},
//   	attributesToDelete: []*string{
//   		jsii.String("attributesToDelete"),
//   	},
//   	deleteStorageConnectors: jsii.Boolean(false),
//   	description: jsii.String("description"),
//   	displayName: jsii.String("displayName"),
//   	embedHostDomains: []*string{
//   		jsii.String("embedHostDomains"),
//   	},
//   	feedbackUrl: jsii.String("feedbackUrl"),
//   	name: jsii.String("name"),
//   	redirectUrl: jsii.String("redirectUrl"),
//   	storageConnectors: []interface{}{
//   		&storageConnectorProperty{
//   			connectorType: jsii.String("connectorType"),
//
//   			// the properties below are optional
//   			domains: []*string{
//   				jsii.String("domains"),
//   			},
//   			resourceIdentifier: jsii.String("resourceIdentifier"),
//   		},
//   	},
//   	streamingExperienceSettings: &streamingExperienceSettingsProperty{
//   		preferredProtocol: jsii.String("preferredProtocol"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	userSettings: []interface{}{
//   		&userSettingProperty{
//   			action: jsii.String("action"),
//   			permission: jsii.String("permission"),
//   		},
//   	},
//   }
//
type CfnStackProps struct {
	// The list of virtual private cloud (VPC) interface endpoint objects.
	//
	// Users of the stack can connect to AppStream 2.0 only through the specified endpoints.
	AccessEndpoints interface{} `field:"optional" json:"accessEndpoints" yaml:"accessEndpoints"`
	// The persistent application settings for users of the stack.
	//
	// When these settings are enabled, changes that users make to applications and Windows settings are automatically saved after each session and applied to the next session.
	ApplicationSettings interface{} `field:"optional" json:"applicationSettings" yaml:"applicationSettings"`
	// The stack attributes to delete.
	AttributesToDelete *[]*string `field:"optional" json:"attributesToDelete" yaml:"attributesToDelete"`
	// *This parameter has been deprecated.*.
	//
	// Deletes the storage connectors currently enabled for the stack.
	DeleteStorageConnectors interface{} `field:"optional" json:"deleteStorageConnectors" yaml:"deleteStorageConnectors"`
	// The description to display.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The stack name to display.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The domains where AppStream 2.0 streaming sessions can be embedded in an iframe. You must approve the domains that you want to host embedded AppStream 2.0 streaming sessions.
	EmbedHostDomains *[]*string `field:"optional" json:"embedHostDomains" yaml:"embedHostDomains"`
	// The URL that users are redirected to after they click the Send Feedback link.
	//
	// If no URL is specified, no Send Feedback link is displayed.
	FeedbackUrl *string `field:"optional" json:"feedbackUrl" yaml:"feedbackUrl"`
	// The name of the stack.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The URL that users are redirected to after their streaming session ends.
	RedirectUrl *string `field:"optional" json:"redirectUrl" yaml:"redirectUrl"`
	// The storage connectors to enable.
	StorageConnectors interface{} `field:"optional" json:"storageConnectors" yaml:"storageConnectors"`
	// `AWS::AppStream::Stack.StreamingExperienceSettings`.
	StreamingExperienceSettings interface{} `field:"optional" json:"streamingExperienceSettings" yaml:"streamingExperienceSettings"`
	// An array of key-value pairs.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The actions that are enabled or disabled for users during their streaming sessions.
	//
	// By default, these actions are enabled.
	UserSettings interface{} `field:"optional" json:"userSettings" yaml:"userSettings"`
}

