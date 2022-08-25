package awskendra


// The configuration information for syncing a Salesforce chatter feed.
//
// The contents of the object comes from the Salesforce FeedItem table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   salesforceChatterFeedConfigurationProperty := &salesforceChatterFeedConfigurationProperty{
//   	documentDataFieldName: jsii.String("documentDataFieldName"),
//
//   	// the properties below are optional
//   	documentTitleFieldName: jsii.String("documentTitleFieldName"),
//   	fieldMappings: []interface{}{
//   		&dataSourceToIndexFieldMappingProperty{
//   			dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			indexFieldName: jsii.String("indexFieldName"),
//
//   			// the properties below are optional
//   			dateFieldFormat: jsii.String("dateFieldFormat"),
//   		},
//   	},
//   	includeFilterTypes: []*string{
//   		jsii.String("includeFilterTypes"),
//   	},
//   }
//
type CfnDataSource_SalesforceChatterFeedConfigurationProperty struct {
	// The name of the column in the Salesforce FeedItem table that contains the content to index.
	//
	// Typically this is the `Body` column.
	DocumentDataFieldName *string `field:"required" json:"documentDataFieldName" yaml:"documentDataFieldName"`
	// The name of the column in the Salesforce FeedItem table that contains the title of the document.
	//
	// This is typically the `Title` column.
	DocumentTitleFieldName *string `field:"optional" json:"documentTitleFieldName" yaml:"documentTitleFieldName"`
	// Maps fields from a Salesforce chatter feed into Amazon Kendra index fields.
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
	// Filters the documents in the feed based on status of the user.
	//
	// When you specify `ACTIVE_USERS` only documents from users who have an active account are indexed. When you specify `STANDARD_USER` only documents for Salesforce standard users are documented. You can specify both.
	IncludeFilterTypes *[]*string `field:"optional" json:"includeFilterTypes" yaml:"includeFilterTypes"`
}

