package awskendra


// Provides the configuration information to connect to ServiceNow as your data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceNowConfigurationProperty := &ServiceNowConfigurationProperty{
//   	HostUrl: jsii.String("hostUrl"),
//   	SecretArn: jsii.String("secretArn"),
//   	ServiceNowBuildVersion: jsii.String("serviceNowBuildVersion"),
//
//   	// the properties below are optional
//   	AuthenticationType: jsii.String("authenticationType"),
//   	KnowledgeArticleConfiguration: &ServiceNowKnowledgeArticleConfigurationProperty{
//   		DocumentDataFieldName: jsii.String("documentDataFieldName"),
//
//   		// the properties below are optional
//   		CrawlAttachments: jsii.Boolean(false),
//   		DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   		ExcludeAttachmentFilePatterns: []*string{
//   			jsii.String("excludeAttachmentFilePatterns"),
//   		},
//   		FieldMappings: []interface{}{
//   			&DataSourceToIndexFieldMappingProperty{
//   				DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				IndexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				DateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   		FilterQuery: jsii.String("filterQuery"),
//   		IncludeAttachmentFilePatterns: []*string{
//   			jsii.String("includeAttachmentFilePatterns"),
//   		},
//   	},
//   	ServiceCatalogConfiguration: &ServiceNowServiceCatalogConfigurationProperty{
//   		DocumentDataFieldName: jsii.String("documentDataFieldName"),
//
//   		// the properties below are optional
//   		CrawlAttachments: jsii.Boolean(false),
//   		DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   		ExcludeAttachmentFilePatterns: []*string{
//   			jsii.String("excludeAttachmentFilePatterns"),
//   		},
//   		FieldMappings: []interface{}{
//   			&DataSourceToIndexFieldMappingProperty{
//   				DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				IndexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				DateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   		IncludeAttachmentFilePatterns: []*string{
//   			jsii.String("includeAttachmentFilePatterns"),
//   		},
//   	},
//   }
//
type CfnDataSource_ServiceNowConfigurationProperty struct {
	// The ServiceNow instance that the data source connects to.
	//
	// The host endpoint should look like the following: *{instance}.service-now.com.*
	HostUrl *string `field:"required" json:"hostUrl" yaml:"hostUrl"`
	// The Amazon Resource Name (ARN) of the AWS Secrets Manager secret that contains the user name and password required to connect to the ServiceNow instance.
	//
	// You can also provide OAuth authentication credentials of user name, password, client ID, and client secret. For more information, see [Using a ServiceNow data source](https://docs.aws.amazon.com/kendra/latest/dg/data-source-servicenow.html) .
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// The identifier of the release that the ServiceNow host is running.
	//
	// If the host is not running the `LONDON` release, use `OTHERS` .
	ServiceNowBuildVersion *string `field:"required" json:"serviceNowBuildVersion" yaml:"serviceNowBuildVersion"`
	// The type of authentication used to connect to the ServiceNow instance.
	//
	// If you choose `HTTP_BASIC` , Amazon Kendra is authenticated using the user name and password provided in the AWS Secrets Manager secret in the `SecretArn` field. If you choose `OAUTH2` , Amazon Kendra is authenticated using the credentials of client ID, client secret, user name and password.
	//
	// When you use `OAUTH2` authentication, you must generate a token and a client secret using the ServiceNow console. For more information, see [Using a ServiceNow data source](https://docs.aws.amazon.com/kendra/latest/dg/data-source-servicenow.html) .
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// Configuration information for crawling knowledge articles in the ServiceNow site.
	KnowledgeArticleConfiguration interface{} `field:"optional" json:"knowledgeArticleConfiguration" yaml:"knowledgeArticleConfiguration"`
	// Configuration information for crawling service catalogs in the ServiceNow site.
	ServiceCatalogConfiguration interface{} `field:"optional" json:"serviceCatalogConfiguration" yaml:"serviceCatalogConfiguration"`
}

