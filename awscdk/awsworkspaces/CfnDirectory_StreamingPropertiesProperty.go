package awsworkspaces


// Describes the streaming properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   streamingPropertiesProperty := &StreamingPropertiesProperty{
//   	GlobalAccelerator: &GlobalAcceleratorForDirectoryProperty{
//   		Mode: jsii.String("mode"),
//
//   		// the properties below are optional
//   		PreferredProtocol: jsii.String("preferredProtocol"),
//   	},
//   	StorageConnectors: []interface{}{
//   		&StorageConnectorProperty{
//   			ConnectorType: jsii.String("connectorType"),
//   			Status: jsii.String("status"),
//   		},
//   	},
//   	StreamingExperiencePreferredProtocol: jsii.String("streamingExperiencePreferredProtocol"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-streamingproperties.html
//
type CfnDirectory_StreamingPropertiesProperty struct {
	// Describes the Global Accelerator for directory.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-streamingproperties.html#cfn-workspaces-directory-streamingproperties-globalaccelerator
	//
	GlobalAccelerator interface{} `field:"optional" json:"globalAccelerator" yaml:"globalAccelerator"`
	// Indicates the storage connector used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-streamingproperties.html#cfn-workspaces-directory-streamingproperties-storageconnectors
	//
	StorageConnectors interface{} `field:"optional" json:"storageConnectors" yaml:"storageConnectors"`
	// Indicates the type of preferred protocol for the streaming experience.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-streamingproperties.html#cfn-workspaces-directory-streamingproperties-streamingexperiencepreferredprotocol
	//
	StreamingExperiencePreferredProtocol *string `field:"optional" json:"streamingExperiencePreferredProtocol" yaml:"streamingExperiencePreferredProtocol"`
	// Indicates the permission settings associated with the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-streamingproperties.html#cfn-workspaces-directory-streamingproperties-usersettings
	//
	UserSettings interface{} `field:"optional" json:"userSettings" yaml:"userSettings"`
}

