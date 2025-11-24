package mixinsawskendra


// Provides the configuration information to connect to ServiceNow as your data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   serviceNowConfigurationProperty := &ServiceNowConfigurationProperty{
//   	AuthenticationType: jsii.String("authenticationType"),
//   	HostUrl: jsii.String("hostUrl"),
//   	KnowledgeArticleConfiguration: &ServiceNowKnowledgeArticleConfigurationProperty{
//   		CrawlAttachments: jsii.Boolean(false),
//   		DocumentDataFieldName: jsii.String("documentDataFieldName"),
//   		DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   		ExcludeAttachmentFilePatterns: []*string{
//   			jsii.String("excludeAttachmentFilePatterns"),
//   		},
//   		FieldMappings: []interface{}{
//   			&DataSourceToIndexFieldMappingProperty{
//   				DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				DateFieldFormat: jsii.String("dateFieldFormat"),
//   				IndexFieldName: jsii.String("indexFieldName"),
//   			},
//   		},
//   		FilterQuery: jsii.String("filterQuery"),
//   		IncludeAttachmentFilePatterns: []*string{
//   			jsii.String("includeAttachmentFilePatterns"),
//   		},
//   	},
//   	SecretArn: jsii.String("secretArn"),
//   	ServiceCatalogConfiguration: &ServiceNowServiceCatalogConfigurationProperty{
//   		CrawlAttachments: jsii.Boolean(false),
//   		DocumentDataFieldName: jsii.String("documentDataFieldName"),
//   		DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   		ExcludeAttachmentFilePatterns: []*string{
//   			jsii.String("excludeAttachmentFilePatterns"),
//   		},
//   		FieldMappings: []interface{}{
//   			&DataSourceToIndexFieldMappingProperty{
//   				DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				DateFieldFormat: jsii.String("dateFieldFormat"),
//   				IndexFieldName: jsii.String("indexFieldName"),
//   			},
//   		},
//   		IncludeAttachmentFilePatterns: []*string{
//   			jsii.String("includeAttachmentFilePatterns"),
//   		},
//   	},
//   	ServiceNowBuildVersion: jsii.String("serviceNowBuildVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowconfiguration.html
//
type CfnDataSourcePropsMixin_ServiceNowConfigurationProperty struct {
	// The type of authentication used to connect to the ServiceNow instance.
	//
	// If you choose `HTTP_BASIC` , Amazon Kendra is authenticated using the user name and password provided in the AWS Secrets Manager secret in the `SecretArn` field. If you choose `OAUTH2` , Amazon Kendra is authenticated using the credentials of client ID, client secret, user name and password.
	//
	// When you use `OAUTH2` authentication, you must generate a token and a client secret using the ServiceNow console. For more information, see [Using a ServiceNow data source](https://docs.aws.amazon.com/kendra/latest/dg/data-source-servicenow.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowconfiguration.html#cfn-kendra-datasource-servicenowconfiguration-authenticationtype
	//
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// The ServiceNow instance that the data source connects to.
	//
	// The host endpoint should look like the following: *{instance}.service-now.com.*
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowconfiguration.html#cfn-kendra-datasource-servicenowconfiguration-hosturl
	//
	HostUrl *string `field:"optional" json:"hostUrl" yaml:"hostUrl"`
	// Configuration information for crawling knowledge articles in the ServiceNow site.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowconfiguration.html#cfn-kendra-datasource-servicenowconfiguration-knowledgearticleconfiguration
	//
	KnowledgeArticleConfiguration interface{} `field:"optional" json:"knowledgeArticleConfiguration" yaml:"knowledgeArticleConfiguration"`
	// The Amazon Resource Name (ARN) of the AWS Secrets Manager secret that contains the user name and password required to connect to the ServiceNow instance.
	//
	// You can also provide OAuth authentication credentials of user name, password, client ID, and client secret. For more information, see [Using a ServiceNow data source](https://docs.aws.amazon.com/kendra/latest/dg/data-source-servicenow.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowconfiguration.html#cfn-kendra-datasource-servicenowconfiguration-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
	// Configuration information for crawling service catalogs in the ServiceNow site.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowconfiguration.html#cfn-kendra-datasource-servicenowconfiguration-servicecatalogconfiguration
	//
	ServiceCatalogConfiguration interface{} `field:"optional" json:"serviceCatalogConfiguration" yaml:"serviceCatalogConfiguration"`
	// The identifier of the release that the ServiceNow host is running.
	//
	// If the host is not running the `LONDON` release, use `OTHERS` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowconfiguration.html#cfn-kendra-datasource-servicenowconfiguration-servicenowbuildversion
	//
	ServiceNowBuildVersion *string `field:"optional" json:"serviceNowBuildVersion" yaml:"serviceNowBuildVersion"`
}

