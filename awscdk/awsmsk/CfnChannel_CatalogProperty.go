package awsmsk


// Catalog configuration of the destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   catalogProperty := &CatalogProperty{
//   	CatalogArn: jsii.String("catalogArn"),
//   	WarehouseLocation: jsii.String("warehouseLocation"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-catalog.html
//
type CfnChannel_CatalogProperty struct {
	// The ARN of the catalog.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-catalog.html#cfn-msk-channel-catalog-catalogarn
	//
	CatalogArn *string `field:"optional" json:"catalogArn" yaml:"catalogArn"`
	// The warehouse location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-channel-catalog.html#cfn-msk-channel-catalog-warehouselocation
	//
	WarehouseLocation *string `field:"optional" json:"warehouseLocation" yaml:"warehouseLocation"`
}

