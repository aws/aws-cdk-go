package mixinsawsiotsitewise


// Represents a data binding value referencing a specific asset property.
//
// It's used to bind computation model variables to actual asset property values for processing.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   assetPropertyBindingValueProperty := &AssetPropertyBindingValueProperty{
//   	AssetId: jsii.String("assetId"),
//   	PropertyId: jsii.String("propertyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-computationmodel-assetpropertybindingvalue.html
//
type CfnComputationModelPropsMixin_AssetPropertyBindingValueProperty struct {
	// The ID of the asset containing the property.
	//
	// This identifies the specific asset instance's property value used in the computation model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-computationmodel-assetpropertybindingvalue.html#cfn-iotsitewise-computationmodel-assetpropertybindingvalue-assetid
	//
	AssetId *string `field:"optional" json:"assetId" yaml:"assetId"`
	// The ID of the property within the asset.
	//
	// This identifies the specific property's value used in the computation model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-computationmodel-assetpropertybindingvalue.html#cfn-iotsitewise-computationmodel-assetpropertybindingvalue-propertyid
	//
	PropertyId *string `field:"optional" json:"propertyId" yaml:"propertyId"`
}

