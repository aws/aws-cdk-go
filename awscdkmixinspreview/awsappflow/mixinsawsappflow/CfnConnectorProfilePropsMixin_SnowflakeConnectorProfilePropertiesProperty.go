package mixinsawsappflow


// The connector-specific profile properties required when using Snowflake.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   snowflakeConnectorProfilePropertiesProperty := &SnowflakeConnectorProfilePropertiesProperty{
//   	AccountName: jsii.String("accountName"),
//   	BucketName: jsii.String("bucketName"),
//   	BucketPrefix: jsii.String("bucketPrefix"),
//   	PrivateLinkServiceName: jsii.String("privateLinkServiceName"),
//   	Region: jsii.String("region"),
//   	Stage: jsii.String("stage"),
//   	Warehouse: jsii.String("warehouse"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-snowflakeconnectorprofileproperties.html
//
type CfnConnectorProfilePropsMixin_SnowflakeConnectorProfilePropertiesProperty struct {
	// The name of the account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-snowflakeconnectorprofileproperties.html#cfn-appflow-connectorprofile-snowflakeconnectorprofileproperties-accountname
	//
	AccountName *string `field:"optional" json:"accountName" yaml:"accountName"`
	// The name of the Amazon S3 bucket associated with Snowflake.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-snowflakeconnectorprofileproperties.html#cfn-appflow-connectorprofile-snowflakeconnectorprofileproperties-bucketname
	//
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// The bucket path that refers to the Amazon S3 bucket associated with Snowflake.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-snowflakeconnectorprofileproperties.html#cfn-appflow-connectorprofile-snowflakeconnectorprofileproperties-bucketprefix
	//
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// The Snowflake Private Link service name to be used for private data transfers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-snowflakeconnectorprofileproperties.html#cfn-appflow-connectorprofile-snowflakeconnectorprofileproperties-privatelinkservicename
	//
	PrivateLinkServiceName *string `field:"optional" json:"privateLinkServiceName" yaml:"privateLinkServiceName"`
	// The AWS Region of the Snowflake account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-snowflakeconnectorprofileproperties.html#cfn-appflow-connectorprofile-snowflakeconnectorprofileproperties-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The name of the Amazon S3 stage that was created while setting up an Amazon S3 stage in the Snowflake account.
	//
	// This is written in the following format: < Database>< Schema><Stage Name>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-snowflakeconnectorprofileproperties.html#cfn-appflow-connectorprofile-snowflakeconnectorprofileproperties-stage
	//
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
	// The name of the Snowflake warehouse.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-connectorprofile-snowflakeconnectorprofileproperties.html#cfn-appflow-connectorprofile-snowflakeconnectorprofileproperties-warehouse
	//
	Warehouse *string `field:"optional" json:"warehouse" yaml:"warehouse"`
}

