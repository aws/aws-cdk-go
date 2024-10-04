package awskinesisfirehose


// Describes the containers where the destination Apache Iceberg Tables are persisted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   catalogConfigurationProperty := &CatalogConfigurationProperty{
//   	CatalogArn: jsii.String("catalogArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-catalogconfiguration.html
//
type CfnDeliveryStream_CatalogConfigurationProperty struct {
	// Specifies the Glue catalog ARN identifier of the destination Apache Iceberg Tables.
	//
	// You must specify the ARN in the format `arn:aws:glue:region:account-id:catalog` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-catalogconfiguration.html#cfn-kinesisfirehose-deliverystream-catalogconfiguration-catalogarn
	//
	CatalogArn *string `field:"optional" json:"catalogArn" yaml:"catalogArn"`
}

