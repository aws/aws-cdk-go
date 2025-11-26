package previewawskendramixins


// Provides the configuration information to connect to Salesforce as your data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   salesforceConfigurationProperty := &SalesforceConfigurationProperty{
//   	ChatterFeedConfiguration: &SalesforceChatterFeedConfigurationProperty{
//   		DocumentDataFieldName: jsii.String("documentDataFieldName"),
//   		DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   		FieldMappings: []interface{}{
//   			&DataSourceToIndexFieldMappingProperty{
//   				DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				DateFieldFormat: jsii.String("dateFieldFormat"),
//   				IndexFieldName: jsii.String("indexFieldName"),
//   			},
//   		},
//   		IncludeFilterTypes: []*string{
//   			jsii.String("includeFilterTypes"),
//   		},
//   	},
//   	CrawlAttachments: jsii.Boolean(false),
//   	ExcludeAttachmentFilePatterns: []*string{
//   		jsii.String("excludeAttachmentFilePatterns"),
//   	},
//   	IncludeAttachmentFilePatterns: []*string{
//   		jsii.String("includeAttachmentFilePatterns"),
//   	},
//   	KnowledgeArticleConfiguration: &SalesforceKnowledgeArticleConfigurationProperty{
//   		CustomKnowledgeArticleTypeConfigurations: []interface{}{
//   			&SalesforceCustomKnowledgeArticleTypeConfigurationProperty{
//   				DocumentDataFieldName: jsii.String("documentDataFieldName"),
//   				DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   				FieldMappings: []interface{}{
//   					&DataSourceToIndexFieldMappingProperty{
//   						DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						DateFieldFormat: jsii.String("dateFieldFormat"),
//   						IndexFieldName: jsii.String("indexFieldName"),
//   					},
//   				},
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		IncludedStates: []*string{
//   			jsii.String("includedStates"),
//   		},
//   		StandardKnowledgeArticleTypeConfiguration: &SalesforceStandardKnowledgeArticleTypeConfigurationProperty{
//   			DocumentDataFieldName: jsii.String("documentDataFieldName"),
//   			DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   			FieldMappings: []interface{}{
//   				&DataSourceToIndexFieldMappingProperty{
//   					DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					DateFieldFormat: jsii.String("dateFieldFormat"),
//   					IndexFieldName: jsii.String("indexFieldName"),
//   				},
//   			},
//   		},
//   	},
//   	SecretArn: jsii.String("secretArn"),
//   	ServerUrl: jsii.String("serverUrl"),
//   	StandardObjectAttachmentConfiguration: &SalesforceStandardObjectAttachmentConfigurationProperty{
//   		DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   		FieldMappings: []interface{}{
//   			&DataSourceToIndexFieldMappingProperty{
//   				DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				DateFieldFormat: jsii.String("dateFieldFormat"),
//   				IndexFieldName: jsii.String("indexFieldName"),
//   			},
//   		},
//   	},
//   	StandardObjectConfigurations: []interface{}{
//   		&SalesforceStandardObjectConfigurationProperty{
//   			DocumentDataFieldName: jsii.String("documentDataFieldName"),
//   			DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   			FieldMappings: []interface{}{
//   				&DataSourceToIndexFieldMappingProperty{
//   					DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					DateFieldFormat: jsii.String("dateFieldFormat"),
//   					IndexFieldName: jsii.String("indexFieldName"),
//   				},
//   			},
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforceconfiguration.html
//
type CfnDataSourcePropsMixin_SalesforceConfigurationProperty struct {
	// Configuration information for Salesforce chatter feeds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforceconfiguration.html#cfn-kendra-datasource-salesforceconfiguration-chatterfeedconfiguration
	//
	ChatterFeedConfiguration interface{} `field:"optional" json:"chatterFeedConfiguration" yaml:"chatterFeedConfiguration"`
	// Indicates whether Amazon Kendra should index attachments to Salesforce objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforceconfiguration.html#cfn-kendra-datasource-salesforceconfiguration-crawlattachments
	//
	CrawlAttachments interface{} `field:"optional" json:"crawlAttachments" yaml:"crawlAttachments"`
	// A list of regular expression patterns to exclude certain documents in your Salesforce.
	//
	// Documents that match the patterns are excluded from the index. Documents that don't match the patterns are included in the index. If a document matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the document isn't included in the index.
	//
	// The pattern is applied to the name of the attached file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforceconfiguration.html#cfn-kendra-datasource-salesforceconfiguration-excludeattachmentfilepatterns
	//
	ExcludeAttachmentFilePatterns *[]*string `field:"optional" json:"excludeAttachmentFilePatterns" yaml:"excludeAttachmentFilePatterns"`
	// A list of regular expression patterns to include certain documents in your Salesforce.
	//
	// Documents that match the patterns are included in the index. Documents that don't match the patterns are excluded from the index. If a document matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the document isn't included in the index.
	//
	// The pattern is applied to the name of the attached file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforceconfiguration.html#cfn-kendra-datasource-salesforceconfiguration-includeattachmentfilepatterns
	//
	IncludeAttachmentFilePatterns *[]*string `field:"optional" json:"includeAttachmentFilePatterns" yaml:"includeAttachmentFilePatterns"`
	// Configuration information for the knowledge article types that Amazon Kendra indexes.
	//
	// Amazon Kendra indexes standard knowledge articles and the standard fields of knowledge articles, or the custom fields of custom knowledge articles, but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforceconfiguration.html#cfn-kendra-datasource-salesforceconfiguration-knowledgearticleconfiguration
	//
	KnowledgeArticleConfiguration interface{} `field:"optional" json:"knowledgeArticleConfiguration" yaml:"knowledgeArticleConfiguration"`
	// The Amazon Resource Name (ARN) of an AWS Secrets Manager secret that contains the key/value pairs required to connect to your Salesforce instance.
	//
	// The secret must contain a JSON structure with the following keys:
	//
	// - authenticationUrl - The OAUTH endpoint that Amazon Kendra connects to get an OAUTH token.
	// - consumerKey - The application public key generated when you created your Salesforce application.
	// - consumerSecret - The application private key generated when you created your Salesforce application.
	// - password - The password associated with the user logging in to the Salesforce instance.
	// - securityToken - The token associated with the user logging in to the Salesforce instance.
	// - username - The user name of the user logging in to the Salesforce instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforceconfiguration.html#cfn-kendra-datasource-salesforceconfiguration-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
	// The instance URL for the Salesforce site that you want to index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforceconfiguration.html#cfn-kendra-datasource-salesforceconfiguration-serverurl
	//
	ServerUrl *string `field:"optional" json:"serverUrl" yaml:"serverUrl"`
	// Configuration information for processing attachments to Salesforce standard objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforceconfiguration.html#cfn-kendra-datasource-salesforceconfiguration-standardobjectattachmentconfiguration
	//
	StandardObjectAttachmentConfiguration interface{} `field:"optional" json:"standardObjectAttachmentConfiguration" yaml:"standardObjectAttachmentConfiguration"`
	// Configuration of the Salesforce standard objects that Amazon Kendra indexes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforceconfiguration.html#cfn-kendra-datasource-salesforceconfiguration-standardobjectconfigurations
	//
	StandardObjectConfigurations interface{} `field:"optional" json:"standardObjectConfigurations" yaml:"standardObjectConfigurations"`
}

