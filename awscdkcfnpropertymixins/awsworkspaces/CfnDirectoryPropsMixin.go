package awsworkspaces

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsworkspaces/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::WorkSpaces::Directory.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnDirectoryPropsMixin := awscdkcfnpropertymixins.Aws_workspaces.NewCfnDirectoryPropsMixin(&CfnDirectoryMixinProps{
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
//   				MaximumLength: jsii.Number(123),
//   				Permission: jsii.String("permission"),
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
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-directory.html
//
type CfnDirectoryPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnDirectoryMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDirectoryPropsMixin
type jsiiProxy_CfnDirectoryPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnDirectoryPropsMixin) Props() *CfnDirectoryMixinProps {
	var returns *CfnDirectoryMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectoryPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::WorkSpaces::Directory`.
func NewCfnDirectoryPropsMixin(props *CfnDirectoryMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnDirectoryPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDirectoryPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDirectoryPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnDirectoryPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::WorkSpaces::Directory`.
func NewCfnDirectoryPropsMixin_Override(c CfnDirectoryPropsMixin, props *CfnDirectoryMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnDirectoryPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnDirectoryPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDirectoryPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnDirectoryPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDirectoryPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_workspaces.CfnDirectoryPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDirectoryPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDirectoryPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

