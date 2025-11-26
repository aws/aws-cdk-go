package previewawsiotsitewisemixins


// Contains information about an `assetModelProperty` binding value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   assetModelPropertyBindingValueProperty := &AssetModelPropertyBindingValueProperty{
//   	AssetModelId: jsii.String("assetModelId"),
//   	PropertyId: jsii.String("propertyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-computationmodel-assetmodelpropertybindingvalue.html
//
type CfnComputationModelPropsMixin_AssetModelPropertyBindingValueProperty struct {
	// The ID of the asset model, in UUID format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-computationmodel-assetmodelpropertybindingvalue.html#cfn-iotsitewise-computationmodel-assetmodelpropertybindingvalue-assetmodelid
	//
	AssetModelId *string `field:"optional" json:"assetModelId" yaml:"assetModelId"`
	// The ID of the asset model property used in data binding value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-computationmodel-assetmodelpropertybindingvalue.html#cfn-iotsitewise-computationmodel-assetmodelpropertybindingvalue-propertyid
	//
	PropertyId *string `field:"optional" json:"propertyId" yaml:"propertyId"`
}

