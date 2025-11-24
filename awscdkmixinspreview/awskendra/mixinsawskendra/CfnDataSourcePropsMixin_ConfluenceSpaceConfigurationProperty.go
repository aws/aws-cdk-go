package mixinsawskendra


// Configuration information for indexing Confluence spaces.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   confluenceSpaceConfigurationProperty := &ConfluenceSpaceConfigurationProperty{
//   	CrawlArchivedSpaces: jsii.Boolean(false),
//   	CrawlPersonalSpaces: jsii.Boolean(false),
//   	ExcludeSpaces: []*string{
//   		jsii.String("excludeSpaces"),
//   	},
//   	IncludeSpaces: []*string{
//   		jsii.String("includeSpaces"),
//   	},
//   	SpaceFieldMappings: []interface{}{
//   		&ConfluenceSpaceToIndexFieldMappingProperty{
//   			DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			DateFieldFormat: jsii.String("dateFieldFormat"),
//   			IndexFieldName: jsii.String("indexFieldName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencespaceconfiguration.html
//
type CfnDataSourcePropsMixin_ConfluenceSpaceConfigurationProperty struct {
	// `TRUE` to index archived spaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencespaceconfiguration.html#cfn-kendra-datasource-confluencespaceconfiguration-crawlarchivedspaces
	//
	CrawlArchivedSpaces interface{} `field:"optional" json:"crawlArchivedSpaces" yaml:"crawlArchivedSpaces"`
	// `TRUE` to index personal spaces.
	//
	// You can add restrictions to items in personal spaces. If personal spaces are indexed, queries without user context information may return restricted items from a personal space in their results. For more information, see [Filtering on user context](https://docs.aws.amazon.com/kendra/latest/dg/user-context-filter.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencespaceconfiguration.html#cfn-kendra-datasource-confluencespaceconfiguration-crawlpersonalspaces
	//
	CrawlPersonalSpaces interface{} `field:"optional" json:"crawlPersonalSpaces" yaml:"crawlPersonalSpaces"`
	// A list of space keys of Confluence spaces.
	//
	// If you include a key, the blogs, documents, and attachments in the space are not indexed. If a space is in both the `ExcludeSpaces` and the `IncludeSpaces` list, the space is excluded.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencespaceconfiguration.html#cfn-kendra-datasource-confluencespaceconfiguration-excludespaces
	//
	ExcludeSpaces *[]*string `field:"optional" json:"excludeSpaces" yaml:"excludeSpaces"`
	// A list of space keys for Confluence spaces.
	//
	// If you include a key, the blogs, documents, and attachments in the space are indexed. Spaces that aren't in the list aren't indexed. A space in the list must exist. Otherwise, Amazon Kendra logs an error when the data source is synchronized. If a space is in both the `IncludeSpaces` and the `ExcludeSpaces` list, the space is excluded.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencespaceconfiguration.html#cfn-kendra-datasource-confluencespaceconfiguration-includespaces
	//
	IncludeSpaces *[]*string `field:"optional" json:"includeSpaces" yaml:"includeSpaces"`
	// Maps attributes or field names of Confluence spaces to Amazon Kendra index field names.
	//
	// To create custom fields, use the `UpdateIndex` API before you map to Confluence fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The Confluence data source field names must exist in your Confluence custom metadata.
	//
	// If you specify the `SpaceFieldMappings` parameter, you must specify at least one field mapping.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencespaceconfiguration.html#cfn-kendra-datasource-confluencespaceconfiguration-spacefieldmappings
	//
	SpaceFieldMappings interface{} `field:"optional" json:"spaceFieldMappings" yaml:"spaceFieldMappings"`
}

