package awsgrafana


// Properties for defining a `CfnWorkspace`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWorkspaceProps := &CfnWorkspaceProps{
//   	AccountAccessType: jsii.String("accountAccessType"),
//   	AuthenticationProviders: []*string{
//   		jsii.String("authenticationProviders"),
//   	},
//   	PermissionType: jsii.String("permissionType"),
//
//   	// the properties below are optional
//   	ClientToken: jsii.String("clientToken"),
//   	DataSources: []*string{
//   		jsii.String("dataSources"),
//   	},
//   	Description: jsii.String("description"),
//   	GrafanaVersion: jsii.String("grafanaVersion"),
//   	Name: jsii.String("name"),
//   	NetworkAccessControl: &NetworkAccessControlProperty{
//   		PrefixListIds: []*string{
//   			jsii.String("prefixListIds"),
//   		},
//   		VpceIds: []*string{
//   			jsii.String("vpceIds"),
//   		},
//   	},
//   	NotificationDestinations: []*string{
//   		jsii.String("notificationDestinations"),
//   	},
//   	OrganizationalUnits: []*string{
//   		jsii.String("organizationalUnits"),
//   	},
//   	OrganizationRoleName: jsii.String("organizationRoleName"),
//   	PluginAdminEnabled: jsii.Boolean(false),
//   	RoleArn: jsii.String("roleArn"),
//   	SamlConfiguration: &SamlConfigurationProperty{
//   		IdpMetadata: &IdpMetadataProperty{
//   			Url: jsii.String("url"),
//   			Xml: jsii.String("xml"),
//   		},
//
//   		// the properties below are optional
//   		AllowedOrganizations: []*string{
//   			jsii.String("allowedOrganizations"),
//   		},
//   		AssertionAttributes: &AssertionAttributesProperty{
//   			Email: jsii.String("email"),
//   			Groups: jsii.String("groups"),
//   			Login: jsii.String("login"),
//   			Name: jsii.String("name"),
//   			Org: jsii.String("org"),
//   			Role: jsii.String("role"),
//   		},
//   		LoginValidityDuration: jsii.Number(123),
//   		RoleValues: &RoleValuesProperty{
//   			Admin: []*string{
//   				jsii.String("admin"),
//   			},
//   			Editor: []*string{
//   				jsii.String("editor"),
//   			},
//   		},
//   	},
//   	StackSetName: jsii.String("stackSetName"),
//   	VpcConfiguration: &VpcConfigurationProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-grafana-workspace.html
//
type CfnWorkspaceProps struct {
	// Specifies whether the workspace can access AWS resources in this AWS account only, or whether it can also access AWS resources in other accounts in the same organization.
	//
	// If this is `ORGANIZATION` , the `OrganizationalUnits` parameter specifies which organizational units the workspace can access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-grafana-workspace.html#cfn-grafana-workspace-accountaccesstype
	//
	AccountAccessType *string `field:"required" json:"accountAccessType" yaml:"accountAccessType"`
	// Specifies whether this workspace uses SAML 2.0, AWS IAM Identity Center , or both to authenticate users for using the Grafana console within a workspace. For more information, see [User authentication in Amazon Managed Grafana](https://docs.aws.amazon.com/grafana/latest/userguide/authentication-in-AMG.html) .
	//
	// *Allowed Values* : `AWS_SSO | SAML`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-grafana-workspace.html#cfn-grafana-workspace-authenticationproviders
	//
	AuthenticationProviders *[]*string `field:"required" json:"authenticationProviders" yaml:"authenticationProviders"`
	// If this is `SERVICE_MANAGED` , and the workplace was created through the Amazon Managed Grafana console, then Amazon Managed Grafana automatically creates the IAM roles and provisions the permissions that the workspace needs to use AWS data sources and notification channels.
	//
	// If this is `CUSTOMER_MANAGED` , you must manage those roles and permissions yourself.
	//
	// If you are working with a workspace in a member account of an organization and that account is not a delegated administrator account, and you want the workspace to access data sources in other AWS accounts in the organization, this parameter must be set to `CUSTOMER_MANAGED` .
	//
	// For more information about converting between customer and service managed, see [Managing permissions for data sources and notification channels](https://docs.aws.amazon.com/grafana/latest/userguide/AMG-datasource-and-notification.html) . For more information about the roles and permissions that must be managed for customer managed workspaces, see [Amazon Managed Grafana permissions and policies for AWS data sources and notification channels](https://docs.aws.amazon.com/grafana/latest/userguide/AMG-manage-permissions.html)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-grafana-workspace.html#cfn-grafana-workspace-permissiontype
	//
	PermissionType *string `field:"required" json:"permissionType" yaml:"permissionType"`
	// A unique, case-sensitive, user-provided identifier to ensure the idempotency of the request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-grafana-workspace.html#cfn-grafana-workspace-clienttoken
	//
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// Specifies the AWS data sources that have been configured to have IAM roles and permissions created to allow Amazon Managed Grafana to read data from these sources.
	//
	// This list is only used when the workspace was created through the AWS console, and the `permissionType` is `SERVICE_MANAGED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-grafana-workspace.html#cfn-grafana-workspace-datasources
	//
	DataSources *[]*string `field:"optional" json:"dataSources" yaml:"dataSources"`
	// The user-defined description of the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-grafana-workspace.html#cfn-grafana-workspace-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies the version of Grafana to support in the workspace.
	//
	// Defaults to the latest version on create (for example, 9.4), or the current version of the workspace on update.
	//
	// Can only be used to upgrade (for example, from 8.4 to 9.4), not downgrade (for example, from 9.4 to 8.4).
	//
	// To know what versions are available to upgrade to for a specific workspace, see the [ListVersions](https://docs.aws.amazon.com/grafana/latest/APIReference/API_ListVersions.html) operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-grafana-workspace.html#cfn-grafana-workspace-grafanaversion
	//
	GrafanaVersion *string `field:"optional" json:"grafanaVersion" yaml:"grafanaVersion"`
	// The name of the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-grafana-workspace.html#cfn-grafana-workspace-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The configuration settings for network access to your workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-grafana-workspace.html#cfn-grafana-workspace-networkaccesscontrol
	//
	NetworkAccessControl interface{} `field:"optional" json:"networkAccessControl" yaml:"networkAccessControl"`
	// The AWS notification channels that Amazon Managed Grafana can automatically create IAM roles and permissions for, to allow Amazon Managed Grafana to use these channels.
	//
	// *AllowedValues* : `SNS`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-grafana-workspace.html#cfn-grafana-workspace-notificationdestinations
	//
	NotificationDestinations *[]*string `field:"optional" json:"notificationDestinations" yaml:"notificationDestinations"`
	// Specifies the organizational units that this workspace is allowed to use data sources from, if this workspace is in an account that is part of an organization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-grafana-workspace.html#cfn-grafana-workspace-organizationalunits
	//
	OrganizationalUnits *[]*string `field:"optional" json:"organizationalUnits" yaml:"organizationalUnits"`
	// The name of the IAM role that is used to access resources through Organizations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-grafana-workspace.html#cfn-grafana-workspace-organizationrolename
	//
	OrganizationRoleName *string `field:"optional" json:"organizationRoleName" yaml:"organizationRoleName"`
	// Whether plugin administration is enabled in the workspace.
	//
	// Setting to `true` allows workspace admins to install, uninstall, and update plugins from within the Grafana workspace.
	//
	// > This option is only valid for workspaces that support Grafana version 9 or newer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-grafana-workspace.html#cfn-grafana-workspace-pluginadminenabled
	//
	PluginAdminEnabled interface{} `field:"optional" json:"pluginAdminEnabled" yaml:"pluginAdminEnabled"`
	// The IAM role that grants permissions to the AWS resources that the workspace will view data from.
	//
	// This role must already exist.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-grafana-workspace.html#cfn-grafana-workspace-rolearn
	//
	RoleArn interface{} `field:"optional" json:"roleArn" yaml:"roleArn"`
	// If the workspace uses SAML, use this structure to map SAML assertion attributes to workspace user information and define which groups in the assertion attribute are to have the `Admin` and `Editor` roles in the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-grafana-workspace.html#cfn-grafana-workspace-samlconfiguration
	//
	SamlConfiguration interface{} `field:"optional" json:"samlConfiguration" yaml:"samlConfiguration"`
	// The name of the AWS CloudFormation stack set that is used to generate IAM roles to be used for this workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-grafana-workspace.html#cfn-grafana-workspace-stacksetname
	//
	StackSetName *string `field:"optional" json:"stackSetName" yaml:"stackSetName"`
	// The configuration settings for an Amazon VPC that contains data sources for your Grafana workspace to connect to.
	//
	// > Connecting to a private VPC is not yet available in the Asia Pacific (Seoul) Region (ap-northeast-2).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-grafana-workspace.html#cfn-grafana-workspace-vpcconfiguration
	//
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

