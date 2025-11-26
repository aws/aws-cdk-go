package previewawskendramixins


// The configuration information for syncing a Salesforce chatter feed.
//
// The contents of the object comes from the Salesforce FeedItem table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   salesforceChatterFeedConfigurationProperty := &SalesforceChatterFeedConfigurationProperty{
//   	DocumentDataFieldName: jsii.String("documentDataFieldName"),
//   	DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   	FieldMappings: []interface{}{
//   		&DataSourceToIndexFieldMappingProperty{
//   			DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			DateFieldFormat: jsii.String("dateFieldFormat"),
//   			IndexFieldName: jsii.String("indexFieldName"),
//   		},
//   	},
//   	IncludeFilterTypes: []*string{
//   		jsii.String("includeFilterTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcechatterfeedconfiguration.html
//
type CfnDataSourcePropsMixin_SalesforceChatterFeedConfigurationProperty struct {
	// The name of the column in the Salesforce FeedItem table that contains the content to index.
	//
	// Typically this is the `Body` column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcechatterfeedconfiguration.html#cfn-kendra-datasource-salesforcechatterfeedconfiguration-documentdatafieldname
	//
	DocumentDataFieldName *string `field:"optional" json:"documentDataFieldName" yaml:"documentDataFieldName"`
	// The name of the column in the Salesforce FeedItem table that contains the title of the document.
	//
	// This is typically the `Title` column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcechatterfeedconfiguration.html#cfn-kendra-datasource-salesforcechatterfeedconfiguration-documenttitlefieldname
	//
	DocumentTitleFieldName *string `field:"optional" json:"documentTitleFieldName" yaml:"documentTitleFieldName"`
	// Maps fields from a Salesforce chatter feed into Amazon Kendra index fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcechatterfeedconfiguration.html#cfn-kendra-datasource-salesforcechatterfeedconfiguration-fieldmappings
	//
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
	// Filters the documents in the feed based on status of the user.
	//
	// When you specify `ACTIVE_USERS` only documents from users who have an active account are indexed. When you specify `STANDARD_USER` only documents for Salesforce standard users are documented. You can specify both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-salesforcechatterfeedconfiguration.html#cfn-kendra-datasource-salesforcechatterfeedconfiguration-includefiltertypes
	//
	IncludeFilterTypes *[]*string `field:"optional" json:"includeFilterTypes" yaml:"includeFilterTypes"`
}

