package awsappsync


// Describes a Delta Sync configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deltaSyncConfigProperty := &deltaSyncConfigProperty{
//   	baseTableTtl: jsii.String("baseTableTtl"),
//   	deltaSyncTableName: jsii.String("deltaSyncTableName"),
//   	deltaSyncTableTtl: jsii.String("deltaSyncTableTtl"),
//   }
//
type CfnDataSource_DeltaSyncConfigProperty struct {
	// The number of minutes that an Item is stored in the data source.
	BaseTableTtl *string `field:"required" json:"baseTableTtl" yaml:"baseTableTtl"`
	// The Delta Sync table name.
	DeltaSyncTableName *string `field:"required" json:"deltaSyncTableName" yaml:"deltaSyncTableName"`
	// The number of minutes that a Delta Sync log entry is stored in the Delta Sync table.
	DeltaSyncTableTtl *string `field:"required" json:"deltaSyncTableTtl" yaml:"deltaSyncTableTtl"`
}

