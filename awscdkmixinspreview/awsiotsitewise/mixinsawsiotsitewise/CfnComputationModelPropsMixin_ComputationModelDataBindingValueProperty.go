package mixinsawsiotsitewise


// Contains computation model data binding value information, which can be one of `assetModelProperty` , `list` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var computationModelDataBindingValueProperty_ ComputationModelDataBindingValueProperty
//
//   computationModelDataBindingValueProperty := &ComputationModelDataBindingValueProperty{
//   	AssetModelProperty: &AssetModelPropertyBindingValueProperty{
//   		AssetModelId: jsii.String("assetModelId"),
//   		PropertyId: jsii.String("propertyId"),
//   	},
//   	AssetProperty: &AssetPropertyBindingValueProperty{
//   		AssetId: jsii.String("assetId"),
//   		PropertyId: jsii.String("propertyId"),
//   	},
//   	List: []interface{}{
//   		computationModelDataBindingValueProperty_,
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-computationmodel-computationmodeldatabindingvalue.html
//
type CfnComputationModelPropsMixin_ComputationModelDataBindingValueProperty struct {
	// Specifies an asset model property data binding value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-computationmodel-computationmodeldatabindingvalue.html#cfn-iotsitewise-computationmodel-computationmodeldatabindingvalue-assetmodelproperty
	//
	AssetModelProperty interface{} `field:"optional" json:"assetModelProperty" yaml:"assetModelProperty"`
	// The asset property value used for computation model data binding.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-computationmodel-computationmodeldatabindingvalue.html#cfn-iotsitewise-computationmodel-computationmodeldatabindingvalue-assetproperty
	//
	AssetProperty interface{} `field:"optional" json:"assetProperty" yaml:"assetProperty"`
	// Specifies a list of data binding value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-computationmodel-computationmodeldatabindingvalue.html#cfn-iotsitewise-computationmodel-computationmodeldatabindingvalue-list
	//
	List interface{} `field:"optional" json:"list" yaml:"list"`
}

