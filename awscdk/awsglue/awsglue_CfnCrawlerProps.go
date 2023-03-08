package awsglue


// Properties for defining a `CfnCrawler`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnCrawlerProps := &cfnCrawlerProps{
//   	role: jsii.String("role"),
//   	targets: &targetsProperty{
//   		catalogTargets: []interface{}{
//   			&catalogTargetProperty{
//   				databaseName: jsii.String("databaseName"),
//   				tables: []*string{
//   					jsii.String("tables"),
//   				},
//   			},
//   		},
//   		dynamoDbTargets: []interface{}{
//   			&dynamoDBTargetProperty{
//   				path: jsii.String("path"),
//   			},
//   		},
//   		jdbcTargets: []interface{}{
//   			&jdbcTargetProperty{
//   				connectionName: jsii.String("connectionName"),
//   				exclusions: []*string{
//   					jsii.String("exclusions"),
//   				},
//   				path: jsii.String("path"),
//   			},
//   		},
//   		mongoDbTargets: []interface{}{
//   			&mongoDBTargetProperty{
//   				connectionName: jsii.String("connectionName"),
//   				path: jsii.String("path"),
//   			},
//   		},
//   		s3Targets: []interface{}{
//   			&s3TargetProperty{
//   				connectionName: jsii.String("connectionName"),
//   				dlqEventQueueArn: jsii.String("dlqEventQueueArn"),
//   				eventQueueArn: jsii.String("eventQueueArn"),
//   				exclusions: []*string{
//   					jsii.String("exclusions"),
//   				},
//   				path: jsii.String("path"),
//   				sampleSize: jsii.Number(123),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	classifiers: []*string{
//   		jsii.String("classifiers"),
//   	},
//   	configuration: jsii.String("configuration"),
//   	crawlerSecurityConfiguration: jsii.String("crawlerSecurityConfiguration"),
//   	databaseName: jsii.String("databaseName"),
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	recrawlPolicy: &recrawlPolicyProperty{
//   		recrawlBehavior: jsii.String("recrawlBehavior"),
//   	},
//   	schedule: &scheduleProperty{
//   		scheduleExpression: jsii.String("scheduleExpression"),
//   	},
//   	schemaChangePolicy: &schemaChangePolicyProperty{
//   		deleteBehavior: jsii.String("deleteBehavior"),
//   		updateBehavior: jsii.String("updateBehavior"),
//   	},
//   	tablePrefix: jsii.String("tablePrefix"),
//   	tags: tags,
//   }
//
type CfnCrawlerProps struct {
	// The Amazon Resource Name (ARN) of an IAM role that's used to access customer resources, such as Amazon Simple Storage Service (Amazon S3) data.
	Role *string `field:"required" json:"role" yaml:"role"`
	// A collection of targets to crawl.
	Targets interface{} `field:"required" json:"targets" yaml:"targets"`
	// A list of UTF-8 strings that specify the names of custom classifiers that are associated with the crawler.
	Classifiers *[]*string `field:"optional" json:"classifiers" yaml:"classifiers"`
	// Crawler configuration information.
	//
	// This versioned JSON string allows users to specify aspects of a crawler's behavior. For more information, see [Configuring a Crawler](https://docs.aws.amazon.com/glue/latest/dg/crawler-configuration.html) .
	Configuration *string `field:"optional" json:"configuration" yaml:"configuration"`
	// The name of the `SecurityConfiguration` structure to be used by this crawler.
	CrawlerSecurityConfiguration *string `field:"optional" json:"crawlerSecurityConfiguration" yaml:"crawlerSecurityConfiguration"`
	// The name of the database in which the crawler's output is stored.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// A description of the crawler.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the crawler.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A policy that specifies whether to crawl the entire dataset again, or to crawl only folders that were added since the last crawler run.
	RecrawlPolicy interface{} `field:"optional" json:"recrawlPolicy" yaml:"recrawlPolicy"`
	// For scheduled crawlers, the schedule when the crawler runs.
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
	// The policy that specifies update and delete behaviors for the crawler.
	//
	// The policy tells the crawler what to do in the event that it detects a change in a table that already exists in the customer's database at the time of the crawl. The `SchemaChangePolicy` does not affect whether or how new tables and partitions are added. New tables and partitions are always created regardless of the `SchemaChangePolicy` on a crawler.
	//
	// The SchemaChangePolicy consists of two components, `UpdateBehavior` and `DeleteBehavior` .
	SchemaChangePolicy interface{} `field:"optional" json:"schemaChangePolicy" yaml:"schemaChangePolicy"`
	// The prefix added to the names of tables that are created.
	TablePrefix *string `field:"optional" json:"tablePrefix" yaml:"tablePrefix"`
	// The tags to use with this crawler.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

