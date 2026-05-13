package awsdatazone


// Lakehouse Properties Input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lakehousePropertiesInputProperty := &LakehousePropertiesInputProperty{
//   	GlueLineageSyncEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-lakehousepropertiesinput.html
//
type CfnConnection_LakehousePropertiesInputProperty struct {
	// Specifies whether Glue lineage sync is enabled for the lakehouse connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-lakehousepropertiesinput.html#cfn-datazone-connection-lakehousepropertiesinput-gluelineagesyncenabled
	//
	GlueLineageSyncEnabled interface{} `field:"optional" json:"glueLineageSyncEnabled" yaml:"glueLineageSyncEnabled"`
}

