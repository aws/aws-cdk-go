package awsiotsitewise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var computationModelDataBindingValueProperty_ computationModelDataBindingValueProperty
//
//   computationModelDataBindingValueProperty := &computationModelDataBindingValueProperty{
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
type CfnComputationModel_ComputationModelDataBindingValueProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-computationmodel-computationmodeldatabindingvalue.html#cfn-iotsitewise-computationmodel-computationmodeldatabindingvalue-assetmodelproperty
	//
	AssetModelProperty interface{} `field:"optional" json:"assetModelProperty" yaml:"assetModelProperty"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-computationmodel-computationmodeldatabindingvalue.html#cfn-iotsitewise-computationmodel-computationmodeldatabindingvalue-assetproperty
	//
	AssetProperty interface{} `field:"optional" json:"assetProperty" yaml:"assetProperty"`
	// Defines a list of computation model binding values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-computationmodel-computationmodeldatabindingvalue.html#cfn-iotsitewise-computationmodel-computationmodeldatabindingvalue-list
	//
	List interface{} `field:"optional" json:"list" yaml:"list"`
}

