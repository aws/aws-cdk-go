package previewawskendramixins


// Provides the configuration information to connect to Confluence as your data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   confluenceConfigurationProperty := &ConfluenceConfigurationProperty{
//   	AttachmentConfiguration: &ConfluenceAttachmentConfigurationProperty{
//   		AttachmentFieldMappings: []interface{}{
//   			&ConfluenceAttachmentToIndexFieldMappingProperty{
//   				DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				DateFieldFormat: jsii.String("dateFieldFormat"),
//   				IndexFieldName: jsii.String("indexFieldName"),
//   			},
//   		},
//   		CrawlAttachments: jsii.Boolean(false),
//   	},
//   	BlogConfiguration: &ConfluenceBlogConfigurationProperty{
//   		BlogFieldMappings: []interface{}{
//   			&ConfluenceBlogToIndexFieldMappingProperty{
//   				DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				DateFieldFormat: jsii.String("dateFieldFormat"),
//   				IndexFieldName: jsii.String("indexFieldName"),
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
//   				DateFieldFormat: jsii.String("dateFieldFormat"),
//   				IndexFieldName: jsii.String("indexFieldName"),
//   			},
//   		},
//   	},
//   	SecretArn: jsii.String("secretArn"),
//   	ServerUrl: jsii.String("serverUrl"),
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
//   				DateFieldFormat: jsii.String("dateFieldFormat"),
//   				IndexFieldName: jsii.String("indexFieldName"),
//   			},
//   		},
//   	},
//   	Version: jsii.String("version"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceconfiguration.html
//
type CfnDataSourcePropsMixin_ConfluenceConfigurationProperty struct {
	// Configuration information for indexing attachments to Confluence blogs and pages.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceconfiguration.html#cfn-kendra-datasource-confluenceconfiguration-attachmentconfiguration
	//
	AttachmentConfiguration interface{} `field:"optional" json:"attachmentConfiguration" yaml:"attachmentConfiguration"`
	// Configuration information for indexing Confluence blogs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceconfiguration.html#cfn-kendra-datasource-confluenceconfiguration-blogconfiguration
	//
	BlogConfiguration interface{} `field:"optional" json:"blogConfiguration" yaml:"blogConfiguration"`
	// A list of regular expression patterns to exclude certain blog posts, pages, spaces, or attachments in your Confluence.
	//
	// Content that matches the patterns are excluded from the index. Content that doesn't match the patterns is included in the index. If content matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the content isn't included in the index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceconfiguration.html#cfn-kendra-datasource-confluenceconfiguration-exclusionpatterns
	//
	ExclusionPatterns *[]*string `field:"optional" json:"exclusionPatterns" yaml:"exclusionPatterns"`
	// A list of regular expression patterns to include certain blog posts, pages, spaces, or attachments in your Confluence.
	//
	// Content that matches the patterns are included in the index. Content that doesn't match the patterns is excluded from the index. If content matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the content isn't included in the index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceconfiguration.html#cfn-kendra-datasource-confluenceconfiguration-inclusionpatterns
	//
	InclusionPatterns *[]*string `field:"optional" json:"inclusionPatterns" yaml:"inclusionPatterns"`
	// Configuration information for indexing Confluence pages.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceconfiguration.html#cfn-kendra-datasource-confluenceconfiguration-pageconfiguration
	//
	PageConfiguration interface{} `field:"optional" json:"pageConfiguration" yaml:"pageConfiguration"`
	// The Amazon Resource Name (ARN) of an AWS Secrets Manager secret that contains the user name and password required to connect to the Confluence instance.
	//
	// If you use Confluence Cloud, you use a generated API token as the password.
	//
	// You can also provide authentication credentials in the form of a personal access token. For more information, see [Using a Confluence data source](https://docs.aws.amazon.com/kendra/latest/dg/data-source-confluence.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceconfiguration.html#cfn-kendra-datasource-confluenceconfiguration-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
	// The URL of your Confluence instance.
	//
	// Use the full URL of the server. For example, *https://server.example.com:port/* . You can also use an IP address, for example, *https://192.168.1.113/* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceconfiguration.html#cfn-kendra-datasource-confluenceconfiguration-serverurl
	//
	ServerUrl *string `field:"optional" json:"serverUrl" yaml:"serverUrl"`
	// Configuration information for indexing Confluence spaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceconfiguration.html#cfn-kendra-datasource-confluenceconfiguration-spaceconfiguration
	//
	SpaceConfiguration interface{} `field:"optional" json:"spaceConfiguration" yaml:"spaceConfiguration"`
	// The version or the type of Confluence installation to connect to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceconfiguration.html#cfn-kendra-datasource-confluenceconfiguration-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
	// Configuration information for an Amazon Virtual Private Cloud to connect to your Confluence.
	//
	// For more information, see [Configuring a VPC](https://docs.aws.amazon.com/kendra/latest/dg/vpc-configuration.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceconfiguration.html#cfn-kendra-datasource-confluenceconfiguration-vpcconfiguration
	//
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

