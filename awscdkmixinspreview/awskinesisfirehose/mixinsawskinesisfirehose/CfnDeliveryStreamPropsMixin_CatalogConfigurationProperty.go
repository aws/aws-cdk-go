package mixinsawskinesisfirehose


// Describes the containers where the destination Apache Iceberg Tables are persisted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   catalogConfigurationProperty := &CatalogConfigurationProperty{
//   	CatalogArn: jsii.String("catalogArn"),
//   	WarehouseLocation: jsii.String("warehouseLocation"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-catalogconfiguration.html
//
type CfnDeliveryStreamPropsMixin_CatalogConfigurationProperty struct {
	// Specifies the Glue catalog ARN identifier of the destination Apache Iceberg Tables.
	//
	// You must specify the ARN in the format `arn:aws:glue:region:account-id:catalog` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-catalogconfiguration.html#cfn-kinesisfirehose-deliverystream-catalogconfiguration-catalogarn
	//
	CatalogArn *string `field:"optional" json:"catalogArn" yaml:"catalogArn"`
	// The warehouse location for Apache Iceberg tables. You must configure this when schema evolution and table creation is enabled.
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-catalogconfiguration.html#cfn-kinesisfirehose-deliverystream-catalogconfiguration-warehouselocation
	//
	WarehouseLocation *string `field:"optional" json:"warehouseLocation" yaml:"warehouseLocation"`
}

