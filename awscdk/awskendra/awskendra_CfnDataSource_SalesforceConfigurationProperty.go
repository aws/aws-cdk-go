package awskendra


// Provides the configuration information to connect to Salesforce as your data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   salesforceConfigurationProperty := &salesforceConfigurationProperty{
//   	secretArn: jsii.String("secretArn"),
//   	serverUrl: jsii.String("serverUrl"),
//
//   	// the properties below are optional
//   	chatterFeedConfiguration: &salesforceChatterFeedConfigurationProperty{
//   		documentDataFieldName: jsii.String("documentDataFieldName"),
//
//   		// the properties below are optional
//   		documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   		fieldMappings: []interface{}{
//   			&dataSourceToIndexFieldMappingProperty{
//   				dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				indexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				dateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   		includeFilterTypes: []*string{
//   			jsii.String("includeFilterTypes"),
//   		},
//   	},
//   	crawlAttachments: jsii.Boolean(false),
//   	excludeAttachmentFilePatterns: []*string{
//   		jsii.String("excludeAttachmentFilePatterns"),
//   	},
//   	includeAttachmentFilePatterns: []*string{
//   		jsii.String("includeAttachmentFilePatterns"),
//   	},
//   	knowledgeArticleConfiguration: &salesforceKnowledgeArticleConfigurationProperty{
//   		includedStates: []*string{
//   			jsii.String("includedStates"),
//   		},
//
//   		// the properties below are optional
//   		customKnowledgeArticleTypeConfigurations: []interface{}{
//   			&salesforceCustomKnowledgeArticleTypeConfigurationProperty{
//   				documentDataFieldName: jsii.String("documentDataFieldName"),
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   				fieldMappings: []interface{}{
//   					&dataSourceToIndexFieldMappingProperty{
//   						dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   						indexFieldName: jsii.String("indexFieldName"),
//
//   						// the properties below are optional
//   						dateFieldFormat: jsii.String("dateFieldFormat"),
//   					},
//   				},
//   			},
//   		},
//   		standardKnowledgeArticleTypeConfiguration: &salesforceStandardKnowledgeArticleTypeConfigurationProperty{
//   			documentDataFieldName: jsii.String("documentDataFieldName"),
//
//   			// the properties below are optional
//   			documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   			fieldMappings: []interface{}{
//   				&dataSourceToIndexFieldMappingProperty{
//   					dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					indexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					dateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   		},
//   	},
//   	standardObjectAttachmentConfiguration: &salesforceStandardObjectAttachmentConfigurationProperty{
//   		documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   		fieldMappings: []interface{}{
//   			&dataSourceToIndexFieldMappingProperty{
//   				dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				indexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				dateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   	},
//   	standardObjectConfigurations: []interface{}{
//   		&salesforceStandardObjectConfigurationProperty{
//   			documentDataFieldName: jsii.String("documentDataFieldName"),
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   			fieldMappings: []interface{}{
//   				&dataSourceToIndexFieldMappingProperty{
//   					dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					indexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					dateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnDataSource_SalesforceConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of an AWS Secrets Manager secret that contains the key/value pairs required to connect to your Salesforce instance.
	//
	// The secret must contain a JSON structure with the following keys:
	//
	// - authenticationUrl - The OAUTH endpoint that Amazon Kendra connects to get an OAUTH token.
	// - consumerKey - The application public key generated when you created your Salesforce application.
	// - consumerSecret - The application private key generated when you created your Salesforce application.
	// - password - The password associated with the user logging in to the Salesforce instance.
	// - securityToken - The token associated with the user account logging in to the Salesforce instance.
	// - username - The user name of the user logging in to the Salesforce instance.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// The instance URL for the Salesforce site that you want to index.
	ServerUrl *string `field:"required" json:"serverUrl" yaml:"serverUrl"`
	// Configuration information for Salesforce chatter feeds.
	ChatterFeedConfiguration interface{} `field:"optional" json:"chatterFeedConfiguration" yaml:"chatterFeedConfiguration"`
	// Indicates whether Amazon Kendra should index attachments to Salesforce objects.
	CrawlAttachments interface{} `field:"optional" json:"crawlAttachments" yaml:"crawlAttachments"`
	// A list of regular expression patterns to exclude certain documents in your Salesforce.
	//
	// Documents that match the patterns are excluded from the index. Documents that don't match the patterns are included in the index. If a document matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the document isn't included in the index.
	//
	// The pattern is applied to the name of the attached file.
	ExcludeAttachmentFilePatterns *[]*string `field:"optional" json:"excludeAttachmentFilePatterns" yaml:"excludeAttachmentFilePatterns"`
	// A list of regular expression patterns to include certain documents in your Salesforce.
	//
	// Documents that match the patterns are included in the index. Documents that don't match the patterns are excluded from the index. If a document matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the document isn't included in the index.
	//
	// The pattern is applied to the name of the attached file.
	IncludeAttachmentFilePatterns *[]*string `field:"optional" json:"includeAttachmentFilePatterns" yaml:"includeAttachmentFilePatterns"`
	// Configuration information for the knowledge article types that Amazon Kendra indexes.
	//
	// Amazon Kendra indexes standard knowledge articles and the standard fields of knowledge articles, or the custom fields of custom knowledge articles, but not both.
	KnowledgeArticleConfiguration interface{} `field:"optional" json:"knowledgeArticleConfiguration" yaml:"knowledgeArticleConfiguration"`
	// Configuration information for processing attachments to Salesforce standard objects.
	StandardObjectAttachmentConfiguration interface{} `field:"optional" json:"standardObjectAttachmentConfiguration" yaml:"standardObjectAttachmentConfiguration"`
	// Configuration of the Salesforce standard objects that Amazon Kendra indexes.
	StandardObjectConfigurations interface{} `field:"optional" json:"standardObjectConfigurations" yaml:"standardObjectConfigurations"`
}

