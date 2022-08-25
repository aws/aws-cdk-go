package awskendra


// Configuration information for indexing Confluence spaces.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   confluenceSpaceConfigurationProperty := &confluenceSpaceConfigurationProperty{
//   	crawlArchivedSpaces: jsii.Boolean(false),
//   	crawlPersonalSpaces: jsii.Boolean(false),
//   	excludeSpaces: []*string{
//   		jsii.String("excludeSpaces"),
//   	},
//   	includeSpaces: []*string{
//   		jsii.String("includeSpaces"),
//   	},
//   	spaceFieldMappings: []interface{}{
//   		&confluenceSpaceToIndexFieldMappingProperty{
//   			dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			indexFieldName: jsii.String("indexFieldName"),
//
//   			// the properties below are optional
//   			dateFieldFormat: jsii.String("dateFieldFormat"),
//   		},
//   	},
//   }
//
type CfnDataSource_ConfluenceSpaceConfigurationProperty struct {
	// Specifies whether Amazon Kendra should index archived spaces.
	CrawlArchivedSpaces interface{} `field:"optional" json:"crawlArchivedSpaces" yaml:"crawlArchivedSpaces"`
	// Specifies whether Amazon Kendra should index personal spaces.
	//
	// Users can add restrictions to items in personal spaces. If personal spaces are indexed, queries without user context information may return restricted items from a personal space in their results. For more information, see [Filtering on user context](https://docs.aws.amazon.com/kendra/latest/dg/user-context-filter.html) .
	CrawlPersonalSpaces interface{} `field:"optional" json:"crawlPersonalSpaces" yaml:"crawlPersonalSpaces"`
	// A list of space keys of Confluence spaces.
	//
	// If you include a key, the blogs, documents, and attachments in the space are not indexed. If a space is in both the `ExcludeSpaces` and the `IncludeSpaces` list, the space is excluded.
	ExcludeSpaces *[]*string `field:"optional" json:"excludeSpaces" yaml:"excludeSpaces"`
	// A list of space keys for Confluence spaces.
	//
	// If you include a key, the blogs, documents, and attachments in the space are indexed. Spaces that aren't in the list aren't indexed. A space in the list must exist. Otherwise, Amazon Kendra logs an error when the data source is synchronized. If a space is in both the `IncludeSpaces` and the `ExcludeSpaces` list, the space is excluded.
	IncludeSpaces *[]*string `field:"optional" json:"includeSpaces" yaml:"includeSpaces"`
	// Maps attributes or field names of Confluence spaces to Amazon Kendra index field names.
	//
	// To create custom fields, use the `UpdateIndex` API before you map to Confluence fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The Confluence data source field names must exist in your Confluence custom metadata.
	//
	// If you specify the `SpaceFieldMappings` parameter, you must specify at least one field mapping.
	SpaceFieldMappings interface{} `field:"optional" json:"spaceFieldMappings" yaml:"spaceFieldMappings"`
}

