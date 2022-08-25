package awskendra


// Provides the configuration information to connect to Confluence as your data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   confluenceConfigurationProperty := &confluenceConfigurationProperty{
//   	secretArn: jsii.String("secretArn"),
//   	serverUrl: jsii.String("serverUrl"),
//   	version: jsii.String("version"),
//
//   	// the properties below are optional
//   	attachmentConfiguration: &confluenceAttachmentConfigurationProperty{
//   		attachmentFieldMappings: []interface{}{
//   			&confluenceAttachmentToIndexFieldMappingProperty{
//   				dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				indexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				dateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   		crawlAttachments: jsii.Boolean(false),
//   	},
//   	blogConfiguration: &confluenceBlogConfigurationProperty{
//   		blogFieldMappings: []interface{}{
//   			&confluenceBlogToIndexFieldMappingProperty{
//   				dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				indexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				dateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   	},
//   	exclusionPatterns: []*string{
//   		jsii.String("exclusionPatterns"),
//   	},
//   	inclusionPatterns: []*string{
//   		jsii.String("inclusionPatterns"),
//   	},
//   	pageConfiguration: &confluencePageConfigurationProperty{
//   		pageFieldMappings: []interface{}{
//   			&confluencePageToIndexFieldMappingProperty{
//   				dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				indexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				dateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   	},
//   	spaceConfiguration: &confluenceSpaceConfigurationProperty{
//   		crawlArchivedSpaces: jsii.Boolean(false),
//   		crawlPersonalSpaces: jsii.Boolean(false),
//   		excludeSpaces: []*string{
//   			jsii.String("excludeSpaces"),
//   		},
//   		includeSpaces: []*string{
//   			jsii.String("includeSpaces"),
//   		},
//   		spaceFieldMappings: []interface{}{
//   			&confluenceSpaceToIndexFieldMappingProperty{
//   				dataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				indexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				dateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   	},
//   	vpcConfiguration: &dataSourceVpcConfigurationProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
type CfnDataSource_ConfluenceConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of an AWS Secrets Manager secret that contains the key-value pairs required to connect to your Confluence server.
	//
	// The secret must contain a JSON structure with the following keys:
	//
	// - username—The user name or email address of a user with administrative privileges for the Confluence server.
	// - password—The password associated with the user logging in to the Confluence server.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// The URL of your Confluence instance.
	//
	// Use the full URL of the server. For example, *https://server.example.com:port/* . You can also use an IP address, for example, *https://192.168.1.113/* .
	ServerUrl *string `field:"required" json:"serverUrl" yaml:"serverUrl"`
	// Specifies the version of the Confluence installation that you are connecting to.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Configuration information for indexing attachments to Confluence blogs and pages.
	AttachmentConfiguration interface{} `field:"optional" json:"attachmentConfiguration" yaml:"attachmentConfiguration"`
	// Configuration information for indexing Confluence blogs.
	BlogConfiguration interface{} `field:"optional" json:"blogConfiguration" yaml:"blogConfiguration"`
	// >A list of regular expression patterns to exclude certain blog posts, pages, spaces, or attachments in your Confluence.
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

