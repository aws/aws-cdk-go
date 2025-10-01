package awsappsync


// Describes a Delta Sync configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deltaSyncConfigProperty := &DeltaSyncConfigProperty{
//   	BaseTableTtl: jsii.String("baseTableTtl"),
//   	DeltaSyncTableName: jsii.String("deltaSyncTableName"),
//   	DeltaSyncTableTtl: jsii.String("deltaSyncTableTtl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-deltasyncconfig.html
//
type CfnDataSource_DeltaSyncConfigProperty struct {
	// The number of minutes that an Item is stored in the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-deltasyncconfig.html#cfn-appsync-datasource-deltasyncconfig-basetablettl
	//
	BaseTableTtl *string `field:"required" json:"baseTableTtl" yaml:"baseTableTtl"`
	// The Delta Sync table name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-deltasyncconfig.html#cfn-appsync-datasource-deltasyncconfig-deltasynctablename
	//
	DeltaSyncTableName *string `field:"required" json:"deltaSyncTableName" yaml:"deltaSyncTableName"`
	// The number of minutes that a Delta Sync log entry is stored in the Delta Sync table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-deltasyncconfig.html#cfn-appsync-datasource-deltasyncconfig-deltasynctablettl
	//
	DeltaSyncTableTtl *string `field:"required" json:"deltaSyncTableTtl" yaml:"deltaSyncTableTtl"`
}

