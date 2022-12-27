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
	// `AWS::Grafana::Workspace.AccountAccessType`.
	AccountAccessType *string `field:"optional" json:"accountAccessType" yaml:"accountAccessType"`
	// `AWS::Grafana::Workspace.AuthenticationProviders`.
	AuthenticationProviders *[]*string `field:"optional" json:"authenticationProviders" yaml:"authenticationProviders"`
	// `AWS::Grafana::Workspace.ClientToken`.
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// `AWS::Grafana::Workspace.DataSources`.
	DataSources *[]*string `field:"optional" json:"dataSources" yaml:"dataSources"`
	// `AWS::Grafana::Workspace.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::Grafana::Workspace.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `AWS::Grafana::Workspace.NotificationDestinations`.
	NotificationDestinations *[]*string `field:"optional" json:"notificationDestinations" yaml:"notificationDestinations"`
	// `AWS::Grafana::Workspace.OrganizationalUnits`.
	OrganizationalUnits *[]*string `field:"optional" json:"organizationalUnits" yaml:"organizationalUnits"`
	// `AWS::Grafana::Workspace.OrganizationRoleName`.
	OrganizationRoleName *string `field:"optional" json:"organizationRoleName" yaml:"organizationRoleName"`
	// `AWS::Grafana::Workspace.PermissionType`.
	PermissionType *string `field:"optional" json:"permissionType" yaml:"permissionType"`
	// `AWS::Grafana::Workspace.RoleArn`.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// `AWS::Grafana::Workspace.SamlConfiguration`.
	SamlConfiguration interface{} `field:"optional" json:"samlConfiguration" yaml:"samlConfiguration"`
	// `AWS::Grafana::Workspace.StackSetName`.
	StackSetName *string `field:"optional" json:"stackSetName" yaml:"stackSetName"`
	// `AWS::Grafana::Workspace.VpcConfiguration`.
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

