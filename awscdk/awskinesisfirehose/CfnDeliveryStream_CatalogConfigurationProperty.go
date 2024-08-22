package awskinesisfirehose


// Describes the containers where the destination Apache Iceberg Tables are persisted.
//
// Amazon Data Firehose is in preview release and is subject to change.
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
	// Specifies the Glue catalog ARN indentifier of the destination Apache Iceberg Tables.
	//
	// You must specify the ARN in the format `arn:aws:glue:region:account-id:catalog` .
	//
	// Amazon Data Firehose is in preview release and is subject to change.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-catalogconfiguration.html#cfn-kinesisfirehose-deliverystream-catalogconfiguration-catalogarn
	//
	CatalogArn *string `field:"optional" json:"catalogArn" yaml:"catalogArn"`
}

