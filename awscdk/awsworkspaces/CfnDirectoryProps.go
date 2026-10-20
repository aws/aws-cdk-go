package awsworkspaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDirectory`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDirectoryProps := &CfnDirectoryProps{
//   	ActiveDirectoryConfig: &ActiveDirectoryConfigProperty{
//   		DomainName: jsii.String("domainName"),
//   		ServiceAccountSecretArn: jsii.String("serviceAccountSecretArn"),
//   	},
//   	CertificateBasedAuthProperties: &CertificateBasedAuthPropertiesProperty{
//   		CertificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//   		Status: jsii.String("status"),
//   	},
//   	EnableSelfService: jsii.Boolean(false),
//   	EndpointEncryptionMode: jsii.String("endpointEncryptionMode"),
//   	IdcInstanceArn: jsii.String("idcInstanceArn"),
//   	IpGroupIds: []*string{
//   		jsii.String("ipGroupIds"),
//   	},
//   	MicrosoftEntraConfig: &MicrosoftEntraConfigProperty{
//   		ApplicationConfigSecretArn: jsii.String("applicationConfigSecretArn"),
//   		TenantId: jsii.String("tenantId"),
//   	},
//   	SamlProperties: &SamlPropertiesProperty{
//   		RelayStateParameterName: jsii.String("relayStateParameterName"),
//   		Status: jsii.String("status"),
//   		UserAccessUrl: jsii.String("userAccessUrl"),
//   	},
//   	SelfservicePermissions: &SelfservicePermissionsProperty{
//   		ChangeComputeType: jsii.String("changeComputeType"),
//   		IncreaseVolumeSize: jsii.String("increaseVolumeSize"),
//   		RebuildWorkspace: jsii.String("rebuildWorkspace"),
//   		RestartWorkspace: jsii.String("restartWorkspace"),
//   		SwitchRunningMode: jsii.String("switchRunningMode"),
//   	},
//   	StreamingProperties: &StreamingPropertiesProperty{
//   		GlobalAccelerator: &GlobalAcceleratorForDirectoryProperty{
//   			Mode: jsii.String("mode"),
//
//   			// the properties below are optional
//   			PreferredProtocol: jsii.String("preferredProtocol"),
//   		},
//   		StorageConnectors: []interface{}{
//   			&StorageConnectorProperty{
//   				ConnectorType: jsii.String("connectorType"),
//   				Status: jsii.String("status"),
//   			},
//   		},
//   		StreamingExperiencePreferredProtocol: jsii.String("streamingExperiencePreferredProtocol"),
//   		UserSettings: []interface{}{
//   			&UserSettingProperty{
//   				Action: jsii.String("action"),
//   				Permission: jsii.String("permission"),
//
//   				// the properties below are optional
//   				MaximumLength: jsii.Number(123),
//   			},
//   		},
//   	},
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Tenancy: jsii.String("tenancy"),
//   	UserIdentityType: jsii.String("userIdentityType"),
//   	WorkspaceAccessProperties: &WorkspaceAccessPropertiesProperty{
//   		AccessEndpointConfig: &AccessEndpointConfigProperty{
//   			AccessEndpoints: []interface{}{
//   				&AccessEndpointProperty{
//   					AccessEndpointType: jsii.String("accessEndpointType"),
//   					VpcEndpointId: jsii.String("vpcEndpointId"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			InternetFallbackProtocols: []*string{
//   				jsii.String("internetFallbackProtocols"),
//   			},
//   		},
//   		DeviceTypeAndroid: jsii.String("deviceTypeAndroid"),
//   		DeviceTypeChromeOs: jsii.String("deviceTypeChromeOs"),
//   		DeviceTypeIos: jsii.String("deviceTypeIos"),
//   		DeviceTypeLinux: jsii.String("deviceTypeLinux"),
//   		DeviceTypeOsx: jsii.String("deviceTypeOsx"),
//   		DeviceTypeWeb: jsii.String("deviceTypeWeb"),
//   		DeviceTypeWindows: jsii.String("deviceTypeWindows"),
//   		DeviceTypeWorkSpacesThinClient: jsii.String("deviceTypeWorkSpacesThinClient"),
//   		DeviceTypeZeroClient: jsii.String("deviceTypeZeroClient"),
//   	},
//   	WorkspaceCreationProperties: &DefaultWorkspaceCreationPropertiesProperty{
//   		CustomSecurityGroupId: jsii.String("customSecurityGroupId"),
//   		DefaultOu: jsii.String("defaultOu"),
//   		EnableInternetAccess: jsii.Boolean(false),
//   		EnableMaintenanceMode: jsii.Boolean(false),
//   		InstanceIamRoleArn: jsii.String("instanceIamRoleArn"),
//   		UserEnabledAsLocalAdministrator: jsii.Boolean(false),
//   	},
//   	WorkspaceDirectoryDescription: jsii.String("workspaceDirectoryDescription"),
//   	WorkspaceDirectoryName: jsii.String("workspaceDirectoryName"),
//   	WorkspaceType: jsii.String("workspaceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-directory.html
//
type CfnDirectoryProps struct {
	// Information about the Active Directory config.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-directory.html#cfn-workspaces-directory-activedirectoryconfig
	//
	ActiveDirectoryConfig interface{} `field:"optional" json:"activeDirectoryConfig" yaml:"activeDirectoryConfig"`
	// Describes the properties of the certificate-based authentication you want to use with your WorkSpaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-directory.html#cfn-workspaces-directory-certificatebasedauthproperties
	//
	CertificateBasedAuthProperties interface{} `field:"optional" json:"certificateBasedAuthProperties" yaml:"certificateBasedAuthProperties"`
	// Indicates whether self-service capabilities are enabled or disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-directory.html#cfn-workspaces-directory-enableselfservice
	//
	EnableSelfService interface{} `field:"optional" json:"enableSelfService" yaml:"enableSelfService"`
	// Endpoint encryption mode that allows you to configure the specified directory between Standard TLS and FIPS 140-2 validated mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-directory.html#cfn-workspaces-directory-endpointencryptionmode
	//
	EndpointEncryptionMode *string `field:"optional" json:"endpointEncryptionMode" yaml:"endpointEncryptionMode"`
	// The Amazon Resource Name (ARN) of the identity center instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-directory.html#cfn-workspaces-directory-idcinstancearn
	//
	IdcInstanceArn *string `field:"optional" json:"idcInstanceArn" yaml:"idcInstanceArn"`
	// The identifiers of the IP access control groups associated with the directory.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-directory.html#cfn-workspaces-directory-ipgroupids
	//
	IpGroupIds *[]*string `field:"optional" json:"ipGroupIds" yaml:"ipGroupIds"`
	// Specifies the configurations of the Microsoft Entra.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-directory.html#cfn-workspaces-directory-microsoftentraconfig
	//
	MicrosoftEntraConfig interface{} `field:"optional" json:"microsoftEntraConfig" yaml:"microsoftEntraConfig"`
	// Describes the enablement status, user access URL, and relay state parameter name that are used for configuring federation with an SAML 2.0 identity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-directory.html#cfn-workspaces-directory-samlproperties
	//
	SamlProperties interface{} `field:"optional" json:"samlProperties" yaml:"samlProperties"`
	// Describes the self-service permissions for a directory.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-directory.html#cfn-workspaces-directory-selfservicepermissions
	//
	SelfservicePermissions interface{} `field:"optional" json:"selfservicePermissions" yaml:"selfservicePermissions"`
	// Describes the streaming properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-directory.html#cfn-workspaces-directory-streamingproperties
	//
	StreamingProperties interface{} `field:"optional" json:"streamingProperties" yaml:"streamingProperties"`
	// The identifiers of the subnets for your virtual private cloud (VPC).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-directory.html#cfn-workspaces-directory-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// The tags associated with the directory.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-directory.html#cfn-workspaces-directory-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Indicates whether your WorkSpace directory is dedicated or shared.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-directory.html#cfn-workspaces-directory-tenancy
	//
	Tenancy *string `field:"optional" json:"tenancy" yaml:"tenancy"`
	// The type of identity management the user is using.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-directory.html#cfn-workspaces-directory-useridentitytype
	//
	UserIdentityType *string `field:"optional" json:"userIdentityType" yaml:"userIdentityType"`
	// The device types and operating systems that can be used to access a WorkSpace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-directory.html#cfn-workspaces-directory-workspaceaccessproperties
	//
	WorkspaceAccessProperties interface{} `field:"optional" json:"workspaceAccessProperties" yaml:"workspaceAccessProperties"`
	// The default values that are used to create WorkSpaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-directory.html#cfn-workspaces-directory-workspacecreationproperties
	//
	WorkspaceCreationProperties interface{} `field:"optional" json:"workspaceCreationProperties" yaml:"workspaceCreationProperties"`
	// Description of the directory to register.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-directory.html#cfn-workspaces-directory-workspacedirectorydescription
	//
	WorkspaceDirectoryDescription *string `field:"optional" json:"workspaceDirectoryDescription" yaml:"workspaceDirectoryDescription"`
	// The name of the directory to register.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-directory.html#cfn-workspaces-directory-workspacedirectoryname
	//
	WorkspaceDirectoryName *string `field:"optional" json:"workspaceDirectoryName" yaml:"workspaceDirectoryName"`
	// Indicates whether the directory's WorkSpace type is personal or pools.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-directory.html#cfn-workspaces-directory-workspacetype
	//
	WorkspaceType *string `field:"optional" json:"workspaceType" yaml:"workspaceType"`
}

