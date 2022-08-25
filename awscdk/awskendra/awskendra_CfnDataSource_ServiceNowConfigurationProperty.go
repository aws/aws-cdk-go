package awskendra


// Provides the configuration information to connect to ServiceNow as your data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceNowConfigurationProperty := &serviceNowConfigurationProperty{
//   	hostUrl: jsii.String("hostUrl"),
//   	secretArn: jsii.String("secretArn"),
//   	serviceNowBuildVersion: jsii.String("serviceNowBuildVersion"),
//
//   	// the properties below are optional
//   	authenticationType: jsii.String("authenticationType"),
//   	knowledgeArticleConfiguration: &serviceNowKnowledgeArticleConfigurationProperty{
//   		documentDataFieldName: jsii.String("documentDataFieldName"),
//
//   		// the properties below are optional
//   		crawlAttachments: jsii.Boolean(false),
//   		documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   		excludeAttachmentFilePatterns: []*string{
//   			jsii.String("excludeAttachmentFilePatterns"),
//   		},
//   		fieldMappings: []interface{}{
//   			&dataSourceToIndexFieldMappingProperty{
//   				dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				indexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				dateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   		filterQuery: jsii.String("filterQuery"),
//   		includeAttachmentFilePatterns: []*string{
//   			jsii.String("includeAttachmentFilePatterns"),
//   		},
//   	},
//   	serviceCatalogConfiguration: &serviceNowServiceCatalogConfigurationProperty{
//   		documentDataFieldName: jsii.String("documentDataFieldName"),
//
//   		// the properties below are optional
//   		crawlAttachments: jsii.Boolean(false),
//   		documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   		excludeAttachmentFilePatterns: []*string{
//   			jsii.String("excludeAttachmentFilePatterns"),
//   		},
//   		fieldMappings: []interface{}{
//   			&dataSourceToIndexFieldMappingProperty{
//   				dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				indexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				dateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   		includeAttachmentFilePatterns: []*string{
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
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// The identifier of the release that the ServiceNow host is running.
	//
	// If the host is not running the `LONDON` release, use `OTHERS` .
	ServiceNowBuildVersion *string `field:"required" json:"serviceNowBuildVersion" yaml:"serviceNowBuildVersion"`
	// The type of authentication used to connect to the ServiceNow instance.
	//
	// If you choose `HTTP_BASIC` , Amazon Kendra is authenticated using the user name and password provided in the AWS Secrets Manager secret in the `SecretArn` field. When you choose `OAUTH2` , Amazon Kendra is authenticated using the OAuth token and secret provided in the Secrets Manager secret, and the user name and password are used to determine which information Amazon Kendra has access to.
	//
	// When you use `OAUTH2` authentication, you must generate a token and a client secret using the ServiceNow console. For more information, see [Using a ServiceNow data source](https://docs.aws.amazon.com/kendra/latest/dg/data-source-servicenow.html) .
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// Configuration information for crawling knowledge articles in the ServiceNow site.
	KnowledgeArticleConfiguration interface{} `field:"optional" json:"knowledgeArticleConfiguration" yaml:"knowledgeArticleConfiguration"`
	// Configuration information for crawling service catalogs in the ServiceNow site.
	ServiceCatalogConfiguration interface{} `field:"optional" json:"serviceCatalogConfiguration" yaml:"serviceCatalogConfiguration"`
}

