package awsworkspaces


// The device types and operating systems that can be used to access a WorkSpace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workspaceAccessPropertiesProperty := &WorkspaceAccessPropertiesProperty{
//   	AccessEndpointConfig: &AccessEndpointConfigProperty{
//   		AccessEndpoints: []interface{}{
//   			&AccessEndpointProperty{
//   				AccessEndpointType: jsii.String("accessEndpointType"),
//   				VpcEndpointId: jsii.String("vpcEndpointId"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		InternetFallbackProtocols: []*string{
//   			jsii.String("internetFallbackProtocols"),
//   		},
//   	},
//   	DeviceTypeAndroid: jsii.String("deviceTypeAndroid"),
//   	DeviceTypeChromeOs: jsii.String("deviceTypeChromeOs"),
//   	DeviceTypeIos: jsii.String("deviceTypeIos"),
//   	DeviceTypeLinux: jsii.String("deviceTypeLinux"),
//   	DeviceTypeOsx: jsii.String("deviceTypeOsx"),
//   	DeviceTypeWeb: jsii.String("deviceTypeWeb"),
//   	DeviceTypeWindows: jsii.String("deviceTypeWindows"),
//   	DeviceTypeWorkSpacesThinClient: jsii.String("deviceTypeWorkSpacesThinClient"),
//   	DeviceTypeZeroClient: jsii.String("deviceTypeZeroClient"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-workspaceaccessproperties.html
//
type CfnDirectory_WorkspaceAccessPropertiesProperty struct {
	// Describes the access endpoint configuration for a WorkSpace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-workspaceaccessproperties.html#cfn-workspaces-directory-workspaceaccessproperties-accessendpointconfig
	//
	AccessEndpointConfig interface{} `field:"optional" json:"accessEndpointConfig" yaml:"accessEndpointConfig"`
	// Indicates whether users can use Android and Android-compatible Chrome OS devices to access their WorkSpaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-workspaceaccessproperties.html#cfn-workspaces-directory-workspaceaccessproperties-devicetypeandroid
	//
	DeviceTypeAndroid *string `field:"optional" json:"deviceTypeAndroid" yaml:"deviceTypeAndroid"`
	// Indicates whether users can use Chromebooks to access their WorkSpaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-workspaceaccessproperties.html#cfn-workspaces-directory-workspaceaccessproperties-devicetypechromeos
	//
	DeviceTypeChromeOs *string `field:"optional" json:"deviceTypeChromeOs" yaml:"deviceTypeChromeOs"`
	// Indicates whether users can use iOS devices to access their WorkSpaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-workspaceaccessproperties.html#cfn-workspaces-directory-workspaceaccessproperties-devicetypeios
	//
	DeviceTypeIos *string `field:"optional" json:"deviceTypeIos" yaml:"deviceTypeIos"`
	// Indicates whether users can use Linux clients to access their WorkSpaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-workspaceaccessproperties.html#cfn-workspaces-directory-workspaceaccessproperties-devicetypelinux
	//
	DeviceTypeLinux *string `field:"optional" json:"deviceTypeLinux" yaml:"deviceTypeLinux"`
	// Indicates whether users can use macOS clients to access their WorkSpaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-workspaceaccessproperties.html#cfn-workspaces-directory-workspaceaccessproperties-devicetypeosx
	//
	DeviceTypeOsx *string `field:"optional" json:"deviceTypeOsx" yaml:"deviceTypeOsx"`
	// Indicates whether users can access their WorkSpaces through a web browser.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-workspaceaccessproperties.html#cfn-workspaces-directory-workspaceaccessproperties-devicetypeweb
	//
	DeviceTypeWeb *string `field:"optional" json:"deviceTypeWeb" yaml:"deviceTypeWeb"`
	// Indicates whether users can use Windows clients to access their WorkSpaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-workspaceaccessproperties.html#cfn-workspaces-directory-workspaceaccessproperties-devicetypewindows
	//
	DeviceTypeWindows *string `field:"optional" json:"deviceTypeWindows" yaml:"deviceTypeWindows"`
	// Indicates whether users can access their WorkSpaces through a WorkSpaces Thin Client.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-workspaceaccessproperties.html#cfn-workspaces-directory-workspaceaccessproperties-devicetypeworkspacesthinclient
	//
	DeviceTypeWorkSpacesThinClient *string `field:"optional" json:"deviceTypeWorkSpacesThinClient" yaml:"deviceTypeWorkSpacesThinClient"`
	// Indicates whether users can use zero client devices to access their WorkSpaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-directory-workspaceaccessproperties.html#cfn-workspaces-directory-workspaceaccessproperties-devicetypezeroclient
	//
	DeviceTypeZeroClient *string `field:"optional" json:"deviceTypeZeroClient" yaml:"deviceTypeZeroClient"`
}

