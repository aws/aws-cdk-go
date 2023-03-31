package awsglue


// Specifies an AWS Glue Data Catalog target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   catalogTargetProperty := &catalogTargetProperty{
//   	databaseName: jsii.String("databaseName"),
//   	tables: []*string{
//   		jsii.String("tables"),
//   	},
//   }
//
type CfnCrawler_CatalogTargetProperty struct {
	// The name of the database to be synchronized.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// A list of the tables to be synchronized.
	Tables *[]*string `field:"optional" json:"tables" yaml:"tables"`
}

