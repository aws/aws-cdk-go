package awskendra


// Provides the configuration information to connect to Confluence as your data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   confluenceConfigurationProperty := &ConfluenceConfigurationProperty{
//   	SecretArn: jsii.String("secretArn"),
//   	ServerUrl: jsii.String("serverUrl"),
//   	Version: jsii.String("version"),
//
//   	// the properties below are optional
//   	AttachmentConfiguration: &ConfluenceAttachmentConfigurationProperty{
//   		AttachmentFieldMappings: []interface{}{
//   			&ConfluenceAttachmentToIndexFieldMappingProperty{
//   				DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				IndexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				DateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   		CrawlAttachments: jsii.Boolean(false),
//   	},
//   	BlogConfiguration: &ConfluenceBlogConfigurationProperty{
//   		BlogFieldMappings: []interface{}{
//   			&ConfluenceBlogToIndexFieldMappingProperty{
//   				DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				IndexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				DateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   	},
//   	ExclusionPatterns: []*string{
//   		jsii.String("exclusionPatterns"),
//   	},
//   	InclusionPatterns: []*string{
//   		jsii.String("inclusionPatterns"),
//   	},
//   	PageConfiguration: &ConfluencePageConfigurationProperty{
//   		PageFieldMappings: []interface{}{
//   			&ConfluencePageToIndexFieldMappingProperty{
//   				DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				IndexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				DateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   	},
//   	SpaceConfiguration: &ConfluenceSpaceConfigurationProperty{
//   		CrawlArchivedSpaces: jsii.Boolean(false),
//   		CrawlPersonalSpaces: jsii.Boolean(false),
//   		ExcludeSpaces: []*string{
//   			jsii.String("excludeSpaces"),
//   		},
//   		IncludeSpaces: []*string{
//   			jsii.String("includeSpaces"),
//   		},
//   		SpaceFieldMappings: []interface{}{
//   			&ConfluenceSpaceToIndexFieldMappingProperty{
//   				DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				IndexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				DateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   	},
//   	VpcConfiguration: &DataSourceVpcConfigurationProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
type CfnDataSource_ConfluenceConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of an AWS Secrets Manager secret that contains the user name and password required to connect to the Confluence instance.
	//
	// If you use Confluence Cloud, you use a generated API token as the password.
	//
	// You can also provide authentication credentials in the form of a personal access token. For more information, see [Using a Confluence data source](https://docs.aws.amazon.com/kendra/latest/dg/data-source-confluence.html) .
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// The URL of your Confluence instance.
	//
	// Use the full URL of the server. For example, *https://server.example.com:port/* . You can also use an IP address, for example, *https://192.168.1.113/* .
	ServerUrl *string `field:"required" json:"serverUrl" yaml:"serverUrl"`
	// The version or the type of Confluence installation to connect to.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Configuration information for indexing attachments to Confluence blogs and pages.
	AttachmentConfiguration interface{} `field:"optional" json:"attachmentConfiguration" yaml:"attachmentConfiguration"`
	// Configuration information for indexing Confluence blogs.
	BlogConfiguration interface{} `field:"optional" json:"blogConfiguration" yaml:"blogConfiguration"`
	// A list of regular expression patterns to exclude certain blog posts, pages, spaces, or attachments in your Confluence.
	//
	// Content that matches the patterns are excluded from the index. Content that doesn't match the patterns is included in the index. If content matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the content isn't included in the index.
	ExclusionPatterns *[]*string `field:"optional" json:"exclusionPatterns" yaml:"exclusionPatterns"`
	// A list of regular expression patterns to include certain blog posts, pages, spaces, or attachments in your Confluence.
	//
	// Content that matches the patterns are included in the index. Content that doesn't match the patterns is excluded from the index. If content matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the content isn't included in the index.
	InclusionPatterns *[]*string `field:"optional" json:"inclusionPatterns" yaml:"inclusionPatterns"`
	// Configuration information for indexing Confluence pages.
	PageConfiguration interface{} `field:"optional" json:"pageConfiguration" yaml:"pageConfiguration"`
	// Configuration information for indexing Confluence spaces.
	SpaceConfiguration interface{} `field:"optional" json:"spaceConfiguration" yaml:"spaceConfiguration"`
	// Configuration information for an Amazon Virtual Private Cloud to connect to your Confluence.
	//
	// For more information, see [Configuring a VPC](https://docs.aws.amazon.com/kendra/latest/dg/vpc-configuration.html) .
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

