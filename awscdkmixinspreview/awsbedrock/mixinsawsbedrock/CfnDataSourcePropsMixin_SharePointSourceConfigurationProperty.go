package mixinsawsbedrock


// The endpoint information to connect to your SharePoint data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sharePointSourceConfigurationProperty := &SharePointSourceConfigurationProperty{
//   	AuthType: jsii.String("authType"),
//   	CredentialsSecretArn: jsii.String("credentialsSecretArn"),
//   	Domain: jsii.String("domain"),
//   	HostType: jsii.String("hostType"),
//   	SiteUrls: []*string{
//   		jsii.String("siteUrls"),
//   	},
//   	TenantId: jsii.String("tenantId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-sharepointsourceconfiguration.html
//
type CfnDataSourcePropsMixin_SharePointSourceConfigurationProperty struct {
	// The supported authentication type to authenticate and connect to your SharePoint site/sites.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-sharepointsourceconfiguration.html#cfn-bedrock-datasource-sharepointsourceconfiguration-authtype
	//
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// The Amazon Resource Name of an AWS Secrets Manager secret that stores your authentication credentials for your SharePoint site/sites.
	//
	// For more information on the key-value pairs that must be included in your secret, depending on your authentication type, see [SharePoint connection configuration](https://docs.aws.amazon.com/bedrock/latest/userguide/sharepoint-data-source-connector.html#configuration-sharepoint-connector) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-sharepointsourceconfiguration.html#cfn-bedrock-datasource-sharepointsourceconfiguration-credentialssecretarn
	//
	CredentialsSecretArn *string `field:"optional" json:"credentialsSecretArn" yaml:"credentialsSecretArn"`
	// The domain of your SharePoint instance or site URL/URLs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-sharepointsourceconfiguration.html#cfn-bedrock-datasource-sharepointsourceconfiguration-domain
	//
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The supported host type, whether online/cloud or server/on-premises.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-sharepointsourceconfiguration.html#cfn-bedrock-datasource-sharepointsourceconfiguration-hosttype
	//
	HostType *string `field:"optional" json:"hostType" yaml:"hostType"`
	// A list of one or more SharePoint site URLs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-sharepointsourceconfiguration.html#cfn-bedrock-datasource-sharepointsourceconfiguration-siteurls
	//
	SiteUrls *[]*string `field:"optional" json:"siteUrls" yaml:"siteUrls"`
	// The identifier of your Microsoft 365 tenant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-sharepointsourceconfiguration.html#cfn-bedrock-datasource-sharepointsourceconfiguration-tenantid
	//
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

