package awsappflow


// The connector-specific profile properties required when using Snowflake.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snowflakeConnectorProfilePropertiesProperty := &snowflakeConnectorProfilePropertiesProperty{
//   	bucketName: jsii.String("bucketName"),
//   	stage: jsii.String("stage"),
//   	warehouse: jsii.String("warehouse"),
//
//   	// the properties below are optional
//   	accountName: jsii.String("accountName"),
//   	bucketPrefix: jsii.String("bucketPrefix"),
//   	privateLinkServiceName: jsii.String("privateLinkServiceName"),
//   	region: jsii.String("region"),
//   }
//
type CfnConnectorProfile_SnowflakeConnectorProfilePropertiesProperty struct {
	// The name of the Amazon S3 bucket associated with Snowflake.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The name of the Amazon S3 stage that was created while setting up an Amazon S3 stage in the Snowflake account.
	//
	// This is written in the following format: < Database>< Schema><Stage Name>.
	Stage *string `field:"required" json:"stage" yaml:"stage"`
	// The name of the Snowflake warehouse.
	Warehouse *string `field:"required" json:"warehouse" yaml:"warehouse"`
	// The name of the account.
	AccountName *string `field:"optional" json:"accountName" yaml:"accountName"`
	// The bucket path that refers to the Amazon S3 bucket associated with Snowflake.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// The Snowflake Private Link service name to be used for private data transfers.
	PrivateLinkServiceName *string `field:"optional" json:"privateLinkServiceName" yaml:"privateLinkServiceName"`
	// The AWS Region of the Snowflake account.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

