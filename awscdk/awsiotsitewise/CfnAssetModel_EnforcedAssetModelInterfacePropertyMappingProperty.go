package awsiotsitewise


// Contains information about applied interface property and asset model property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   enforcedAssetModelInterfacePropertyMappingProperty := &EnforcedAssetModelInterfacePropertyMappingProperty{
//   	InterfaceAssetModelPropertyExternalId: jsii.String("interfaceAssetModelPropertyExternalId"),
//
//   	// the properties below are optional
//   	AssetModelPropertyExternalId: jsii.String("assetModelPropertyExternalId"),
//   	AssetModelPropertyLogicalId: jsii.String("assetModelPropertyLogicalId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-enforcedassetmodelinterfacepropertymapping.html
//
type CfnAssetModel_EnforcedAssetModelInterfacePropertyMappingProperty struct {
	// The external ID of the applied interface property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-enforcedassetmodelinterfacepropertymapping.html#cfn-iotsitewise-assetmodel-enforcedassetmodelinterfacepropertymapping-interfaceassetmodelpropertyexternalid
	//
	InterfaceAssetModelPropertyExternalId *string `field:"required" json:"interfaceAssetModelPropertyExternalId" yaml:"interfaceAssetModelPropertyExternalId"`
	// The external ID of the linked asset model property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-enforcedassetmodelinterfacepropertymapping.html#cfn-iotsitewise-assetmodel-enforcedassetmodelinterfacepropertymapping-assetmodelpropertyexternalid
	//
	AssetModelPropertyExternalId *string `field:"optional" json:"assetModelPropertyExternalId" yaml:"assetModelPropertyExternalId"`
	// The logical ID of the linked asset model property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-enforcedassetmodelinterfacepropertymapping.html#cfn-iotsitewise-assetmodel-enforcedassetmodelinterfacepropertymapping-assetmodelpropertylogicalid
	//
	AssetModelPropertyLogicalId *string `field:"optional" json:"assetModelPropertyLogicalId" yaml:"assetModelPropertyLogicalId"`
}

