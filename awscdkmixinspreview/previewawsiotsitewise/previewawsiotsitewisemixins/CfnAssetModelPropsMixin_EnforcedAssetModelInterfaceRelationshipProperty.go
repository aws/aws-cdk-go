package previewawsiotsitewisemixins


// Contains information about applied interface hierarchy and asset model hierarchy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   enforcedAssetModelInterfaceRelationshipProperty := &EnforcedAssetModelInterfaceRelationshipProperty{
//   	InterfaceAssetModelId: jsii.String("interfaceAssetModelId"),
//   	PropertyMappings: []interface{}{
//   		&EnforcedAssetModelInterfacePropertyMappingProperty{
//   			AssetModelPropertyExternalId: jsii.String("assetModelPropertyExternalId"),
//   			AssetModelPropertyLogicalId: jsii.String("assetModelPropertyLogicalId"),
//   			InterfaceAssetModelPropertyExternalId: jsii.String("interfaceAssetModelPropertyExternalId"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-enforcedassetmodelinterfacerelationship.html
//
type CfnAssetModelPropsMixin_EnforcedAssetModelInterfaceRelationshipProperty struct {
	// The ID of the asset model that has the interface applied to it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-enforcedassetmodelinterfacerelationship.html#cfn-iotsitewise-assetmodel-enforcedassetmodelinterfacerelationship-interfaceassetmodelid
	//
	InterfaceAssetModelId *string `field:"optional" json:"interfaceAssetModelId" yaml:"interfaceAssetModelId"`
	// A list of property mappings between the interface asset model and the asset model where the interface is applied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-enforcedassetmodelinterfacerelationship.html#cfn-iotsitewise-assetmodel-enforcedassetmodelinterfacerelationship-propertymappings
	//
	PropertyMappings interface{} `field:"optional" json:"propertyMappings" yaml:"propertyMappings"`
}

