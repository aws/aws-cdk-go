package awskendra


// Provides the configuration information to connect to Microsoft SharePoint as your data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sharePointConfigurationProperty := &sharePointConfigurationProperty{
//   	secretArn: jsii.String("secretArn"),
//   	sharePointVersion: jsii.String("sharePointVersion"),
//   	urls: []*string{
//   		jsii.String("urls"),
//   	},
//
//   	// the properties below are optional
//   	crawlAttachments: jsii.Boolean(false),
//   	disableLocalGroups: jsii.Boolean(false),
//   	documentTitleFieldName: jsii.String("documentTitleFieldName"),
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
//   	sslCertificateS3Path: &s3PathProperty{
//   		bucket: jsii.String("bucket"),
//   		key: jsii.String("key"),
//   	},
//   	useChangeLog: jsii.Boolean(false),
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
type CfnDataSource_SharePointConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of credentials stored in AWS Secrets Manager .
	//
	// The credentials should be a user/password pair. If you use SharePoint Server, you also need to provide the sever domain name as part of the credentials. For more information, see [Using a Microsoft SharePoint Data Source](https://docs.aws.amazon.com/kendra/latest/dg/data-source-sharepoint.html) . For more information about AWS Secrets Manager see [What Is AWS Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html) in the *AWS Secrets Manager* user guide.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// The version of Microsoft SharePoint that you are using as a data source.
	SharePointVersion *string `field:"required" json:"sharePointVersion" yaml:"sharePointVersion"`
	// The URLs of the Microsoft SharePoint site that contains the documents that should be indexed.
	Urls *[]*string `field:"required" json:"urls" yaml:"urls"`
	// `TRUE` to include attachments to documents stored in your Microsoft SharePoint site in the index;
	//
	// otherwise, `FALSE` .
	CrawlAttachments interface{} `field:"optional" json:"crawlAttachments" yaml:"crawlAttachments"`
	// A Boolean value that specifies whether local groups are disabled ( `True` ) or enabled ( `False` ).
	DisableLocalGroups interface{} `field:"optional" json:"disableLocalGroups" yaml:"disableLocalGroups"`
	// The Microsoft SharePoint attribute field that contains the title of the document.
	DocumentTitleFieldName *string `field:"optional" json:"documentTitleFieldName" yaml:"documentTitleFieldName"`
	// A list of regular expression patterns.
	//
	// Documents that match the patterns are excluded from the index. Documents that don't match the patterns are included in the index. If a document matches both an exclusion pattern and an inclusion pattern, the document is not included in the index.
	//
	// The regex is applied to the display URL of the SharePoint document.
	ExclusionPatterns *[]*string `field:"optional" json:"exclusionPatterns" yaml:"exclusionPatterns"`
	// A list of `DataSourceToIndexFieldMapping` objects that map Microsoft SharePoint attributes to custom fields in the Amazon Kendra index.
	//
	// You must first create the index fields using the [UpdateIndex](https://docs.aws.amazon.com/kendra/latest/dg/API_UpdateIndex.html) operation before you map SharePoint attributes. For more information, see [Mapping Data Source Fields](https://docs.aws.amazon.com/kendra/latest/dg/field-mapping.html) .
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
	// A list of regular expression patterns to include certain documents in your SharePoint.
	//
	// Documents that match the patterns are included in the index. Documents that don't match the patterns are excluded from the index. If a document matches both an inclusion and exclusion pattern, the exclusion pattern takes precedence and the document isn't included in the index.
	//
	// The regex is applied to the display URL of the SharePoint document.
	InclusionPatterns *[]*string `field:"optional" json:"inclusionPatterns" yaml:"inclusionPatterns"`
	// Information required to find a specific file in an Amazon S3 bucket.
	SslCertificateS3Path interface{} `field:"optional" json:"sslCertificateS3Path" yaml:"sslCertificateS3Path"`
	// `TRUE` to use the SharePoint change log to determine which documents require updating in the index.
	//
	// Depending on the change log's size, it may take longer for Amazon Kendra to use the change log than to scan all of your documents in SharePoint.
	UseChangeLog interface{} `field:"optional" json:"useChangeLog" yaml:"useChangeLog"`
	// Provides information for connecting to an Amazon VPC.
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

