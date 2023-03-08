package awskendra


// Provides the configuration information to connect to Amazon WorkDocs as your data source.
//
// Amazon WorkDocs connector is available in Oregon, North Virginia, Sydney, Singapore and Ireland regions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workDocsConfigurationProperty := &workDocsConfigurationProperty{
//   	organizationId: jsii.String("organizationId"),
//
//   	// the properties below are optional
//   	crawlComments: jsii.Boolean(false),
//   	exclusionPatterns: []*string{
//   		jsii.String("exclusionPatterns"),
//   	},
//   	fieldMappings: []interface{}{
//   		&dataSourceToIndexFieldMappingProperty{
//   			dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   			indexFieldName: jsii.String("indexFieldName"),
//
//   			// the properties below are optional
//   			dateFieldFormat: jsii.String("dateFieldFormat"),
//   		},
//   	},
//   	inclusionPatterns: []*string{
//   		jsii.String("inclusionPatterns"),
//   	},
//   	useChangeLog: jsii.Boolean(false),
//   }
//
type CfnDataSource_WorkDocsConfigurationProperty struct {
	// The identifier of the directory corresponding to your Amazon WorkDocs site repository.
	//
	// You can find the organization ID in the [AWS Directory Service](https://docs.aws.amazon.com/directoryservicev2/) by going to *Active Directory* , then *Directories* . Your Amazon WorkDocs site directory has an ID, which is the organization ID. You can also set up a new Amazon WorkDocs directory in the AWS Directory Service console and enable a Amazon WorkDocs site for the directory in the Amazon WorkDocs console.
	OrganizationId *string `field:"required" json:"organizationId" yaml:"organizationId"`
	// `TRUE` to include comments on documents in your index.
	//
	// Including comments in your index means each comment is a document that can be searched on.
	//
	// The default is set to `FALSE` .
	CrawlComments interface{} `field:"optional" json:"crawlComments" yaml:"crawlComments"`
	// A list of regular expression patterns to exclude certain files in your Amazon WorkDocs site repository.
	//
	// Files that match the patterns are excluded from the index. Files that don’t match the patterns are included in the index. If a file matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the file isn't included in the index.
	ExclusionPatterns *[]*string `field:"optional" json:"exclusionPatterns" yaml:"exclusionPatterns"`
	// A list of `DataSourceToIndexFieldMapping` objects that map Amazon WorkDocs data source attributes or field names to Amazon Kendra index field names.
	//
	// To create custom fields, use the `UpdateIndex` API before you map to Amazon WorkDocs fields. For more information, see [Mapping data source fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) . The Amazon WorkDocs data source field names must exist in your Amazon WorkDocs custom metadata.
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
	// A list of regular expression patterns to include certain files in your Amazon WorkDocs site repository.
	//
	// Files that match the patterns are included in the index. Files that don't match the patterns are excluded from the index. If a file matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the file isn't included in the index.
	InclusionPatterns *[]*string `field:"optional" json:"inclusionPatterns" yaml:"inclusionPatterns"`
	// `TRUE` to use the Amazon WorkDocs change log to determine which documents require updating in the index.
	//
	// Depending on the change log's size, it may take longer for Amazon Kendra to use the change log than to scan all of your documents in Amazon WorkDocs.
	UseChangeLog interface{} `field:"optional" json:"useChangeLog" yaml:"useChangeLog"`
}

