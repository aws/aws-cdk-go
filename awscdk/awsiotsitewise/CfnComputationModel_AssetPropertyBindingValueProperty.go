package awsiotsitewise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetPropertyBindingValueProperty := &AssetPropertyBindingValueProperty{
//   	AssetId: jsii.String("assetId"),
//   	PropertyId: jsii.String("propertyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-computationmodel-assetpropertybindingvalue.html
//
type CfnComputationModel_AssetPropertyBindingValueProperty struct {
	// The ID of the asset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-computationmodel-assetpropertybindingvalue.html#cfn-iotsitewise-computationmodel-assetpropertybindingvalue-assetid
	//
	AssetId *string `field:"required" json:"assetId" yaml:"assetId"`
	// The ID of the asset property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-computationmodel-assetpropertybindingvalue.html#cfn-iotsitewise-computationmodel-assetpropertybindingvalue-propertyid
	//
	PropertyId *string `field:"required" json:"propertyId" yaml:"propertyId"`
}

