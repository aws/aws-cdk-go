package awsgrafana


// Properties for defining a `CfnWorkspace`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWorkspaceProps := &cfnWorkspaceProps{
//   	accountAccessType: jsii.String("accountAccessType"),
//   	authenticationProviders: []*string{
//   		jsii.String("authenticationProviders"),
//   	},
//   	clientToken: jsii.String("clientToken"),
//   	dataSources: []*string{
//   		jsii.String("dataSources"),
//   	},
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	notificationDestinations: []*string{
//   		jsii.String("notificationDestinations"),
//   	},
//   	organizationalUnits: []*string{
//   		jsii.String("organizationalUnits"),
//   	},
//   	organizationRoleName: jsii.String("organizationRoleName"),
//   	permissionType: jsii.String("permissionType"),
//   	roleArn: jsii.String("roleArn"),
//   	samlConfiguration: &samlConfigurationProperty{
//   		idpMetadata: &idpMetadataProperty{
//   			url: jsii.String("url"),
//   			xml: jsii.String("xml"),
//   		},
//
//   		// the properties below are optional
//   		allowedOrganizations: []*string{
//   			jsii.String("allowedOrganizations"),
//   		},
//   		assertionAttributes: &assertionAttributesProperty{
//   			email: jsii.String("email"),
//   			groups: jsii.String("groups"),
//   			login: jsii.String("login"),
//   			name: jsii.String("name"),
//   			org: jsii.String("org"),
//   			role: jsii.String("role"),
//   		},
//   		loginValidityDuration: jsii.Number(123),
//   		roleValues: &roleValuesProperty{
//   			admin: []*string{
//   				jsii.String("admin"),
//   			},
//   			editor: []*string{
//   				jsii.String("editor"),
//   			},
//   		},
//   	},
//   	stackSetName: jsii.String("stackSetName"),
//   	vpcConfiguration: &vpcConfigurationProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
type CfnWorkspaceProps struct {
	// Specifies whether the workspace can access AWS resources in this AWS account only, or whether it can also access AWS resources in other accounts in the same organization.
	//
	// If this is `ORGANIZATION` , the `workspaceOrganizationalUnits` parameter specifies which organizational units the workspace can access.
	AccountAccessType *string `field:"optional" json:"accountAccessType" yaml:"accountAccessType"`
	// Specifies whether this workspace uses SAML 2.0, AWS IAM Identity Center (successor to AWS Single Sign-On) , or both to authenticate users for using the Grafana console within a workspace. For more information, see [User authentication in Amazon Managed Grafana](https://docs.aws.amazon.com/grafana/latest/userguide/authentication-in-AMG.html) .
	AuthenticationProviders *[]*string `field:"optional" json:"authenticationProviders" yaml:"authenticationProviders"`
	// A unique, case-sensitive, user-provided identifier to ensure the idempotency of the request.
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// Specifies the AWS data sources that have been configured to have IAM roles and permissions created to allow Amazon Managed Grafana to read data from these sources.
	//
	// This list is only used when the workspace was created through the AWS console, and the `permissionType` is `SERVICE_MANAGED` .
	DataSources *[]*string `field:"optional" json:"dataSources" yaml:"dataSources"`
	// The user-defined description of the workspace.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the workspace.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The AWS notification channels that Amazon Managed Grafana can automatically create IAM roles and permissions for, to allow Amazon Managed Grafana to use these channels.
	NotificationDestinations *[]*string `field:"optional" json:"notificationDestinations" yaml:"notificationDestinations"`
	// Specifies the organizational units that this workspace is allowed to use data sources from, if this workspace is in an account that is part of an organization.
	OrganizationalUnits *[]*string `field:"optional" json:"organizationalUnits" yaml:"organizationalUnits"`
	// The name of the IAM role that is used to access resources through Organizations .
	OrganizationRoleName *string `field:"optional" json:"organizationRoleName" yaml:"organizationRoleName"`
	// If this is `SERVICE_MANAGED` , and the workplace was created through the Amazon Managed Grafana console, then Amazon Managed Grafana automatically creates the IAM roles and provisions the permissions that the workspace needs to use AWS data sources and notification channels.
	//
	// If this is `CUSTOMER_MANAGED` , you must manage those roles and permissions yourself.
	//
	// If you are working with a workspace in a member account of an organization and that account is not a delegated administrator account, and you want the workspace to access data sources in other AWS accounts in the organization, this parameter must be set to `CUSTOMER_MANAGED` .
	//
	// For more information about converting between customer and service managed, see [Managing permissions for data sources and notification channels](https://docs.aws.amazon.com/grafana/latest/userguide/AMG-datasource-and-notification.html) . For more information about the roles and permissions that must be managed for customer managed workspaces, see [Amazon Managed Grafana permissions and policies for AWS data sources and notification channels](https://docs.aws.amazon.com/grafana/latest/userguide/AMG-manage-permissions.html)
	PermissionType *string `field:"optional" json:"permissionType" yaml:"permissionType"`
	// The IAM role that grants permissions to the AWS resources that the workspace will view data from.
	//
	// This role must already exist.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// If the workspace uses SAML, use this structure to map SAML assertion attributes to workspace user information and define which groups in the assertion attribute are to have the `Admin` and `Editor` roles in the workspace.
	SamlConfiguration interface{} `field:"optional" json:"samlConfiguration" yaml:"samlConfiguration"`
	// The name of the AWS CloudFormation stack set that is used to generate IAM roles to be used for this workspace.
	StackSetName *string `field:"optional" json:"stackSetName" yaml:"stackSetName"`
	// The configuration for connecting to data sources in a private VPC ( Amazon Virtual Private Cloud ).
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

