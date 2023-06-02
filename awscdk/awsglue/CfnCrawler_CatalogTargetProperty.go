package awsglue


// Specifies an AWS Glue Data Catalog target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   catalogTargetProperty := &CatalogTargetProperty{
//   	ConnectionName: jsii.String("connectionName"),
//   	DatabaseName: jsii.String("databaseName"),
//   	DlqEventQueueArn: jsii.String("dlqEventQueueArn"),
//   	EventQueueArn: jsii.String("eventQueueArn"),
//   	Tables: []*string{
//   		jsii.String("tables"),
//   	},
//   }
//
type CfnCrawler_CatalogTargetProperty struct {
	// `CfnCrawler.CatalogTargetProperty.ConnectionName`.
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// The name of the database to be synchronized.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// `CfnCrawler.CatalogTargetProperty.DlqEventQueueArn`.
	DlqEventQueueArn *string `field:"optional" json:"dlqEventQueueArn" yaml:"dlqEventQueueArn"`
	// `CfnCrawler.CatalogTargetProperty.EventQueueArn`.
	EventQueueArn *string `field:"optional" json:"eventQueueArn" yaml:"eventQueueArn"`
	// A list of the tables to be synchronized.
	Tables *[]*string `field:"optional" json:"tables" yaml:"tables"`
}

