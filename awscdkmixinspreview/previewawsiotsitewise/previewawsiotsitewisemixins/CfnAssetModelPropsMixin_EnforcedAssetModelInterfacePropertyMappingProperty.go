package previewawsiotsitewisemixins


// Contains information about applied interface property and asset model property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   enforcedAssetModelInterfacePropertyMappingProperty := &EnforcedAssetModelInterfacePropertyMappingProperty{
//   	AssetModelPropertyExternalId: jsii.String("assetModelPropertyExternalId"),
//   	AssetModelPropertyLogicalId: jsii.String("assetModelPropertyLogicalId"),
//   	InterfaceAssetModelPropertyExternalId: jsii.String("interfaceAssetModelPropertyExternalId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-enforcedassetmodelinterfacepropertymapping.html
//
type CfnAssetModelPropsMixin_EnforcedAssetModelInterfacePropertyMappingProperty struct {
	// The external ID of the linked asset model property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-enforcedassetmodelinterfacepropertymapping.html#cfn-iotsitewise-assetmodel-enforcedassetmodelinterfacepropertymapping-assetmodelpropertyexternalid
	//
	AssetModelPropertyExternalId *string `field:"optional" json:"assetModelPropertyExternalId" yaml:"assetModelPropertyExternalId"`
	// The logical ID of the linked asset model property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-enforcedassetmodelinterfacepropertymapping.html#cfn-iotsitewise-assetmodel-enforcedassetmodelinterfacepropertymapping-assetmodelpropertylogicalid
	//
	AssetModelPropertyLogicalId *string `field:"optional" json:"assetModelPropertyLogicalId" yaml:"assetModelPropertyLogicalId"`
	// The external ID of the applied interface property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-enforcedassetmodelinterfacepropertymapping.html#cfn-iotsitewise-assetmodel-enforcedassetmodelinterfacepropertymapping-interfaceassetmodelpropertyexternalid
	//
	InterfaceAssetModelPropertyExternalId *string `field:"optional" json:"interfaceAssetModelPropertyExternalId" yaml:"interfaceAssetModelPropertyExternalId"`
}

