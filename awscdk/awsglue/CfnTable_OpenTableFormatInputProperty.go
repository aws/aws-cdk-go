package awsglue


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   openTableFormatInputProperty := &OpenTableFormatInputProperty{
//   	IcebergInput: &IcebergInputProperty{
//   		MetadataOperation: jsii.String("metadataOperation"),
//   		Version: jsii.String("version"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-opentableformatinput.html
//
type CfnTable_OpenTableFormatInputProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-table-opentableformatinput.html#cfn-glue-table-opentableformatinput-iceberginput
	//
	IcebergInput interface{} `field:"optional" json:"icebergInput" yaml:"icebergInput"`
}

