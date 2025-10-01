package awsiotsitewise


// Contains information about enforced interface hierarchy and asset model hierarchy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   enforcedAssetModelInterfaceRelationshipProperty := &EnforcedAssetModelInterfaceRelationshipProperty{
//   	InterfaceAssetModelId: jsii.String("interfaceAssetModelId"),
//   	PropertyMappings: []interface{}{
//   		&EnforcedAssetModelInterfacePropertyMappingProperty{
//   			InterfaceAssetModelPropertyExternalId: jsii.String("interfaceAssetModelPropertyExternalId"),
//
//   			// the properties below are optional
//   			AssetModelPropertyExternalId: jsii.String("assetModelPropertyExternalId"),
//   			AssetModelPropertyLogicalId: jsii.String("assetModelPropertyLogicalId"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-enforcedassetmodelinterfacerelationship.html
//
type CfnAssetModel_EnforcedAssetModelInterfaceRelationshipProperty struct {
	// The ID of the interface that is enforced to the asset model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-enforcedassetmodelinterfacerelationship.html#cfn-iotsitewise-assetmodel-enforcedassetmodelinterfacerelationship-interfaceassetmodelid
	//
	InterfaceAssetModelId *string `field:"optional" json:"interfaceAssetModelId" yaml:"interfaceAssetModelId"`
	// Contains information about enforced interface property and asset model property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-enforcedassetmodelinterfacerelationship.html#cfn-iotsitewise-assetmodel-enforcedassetmodelinterfacerelationship-propertymappings
	//
	PropertyMappings interface{} `field:"optional" json:"propertyMappings" yaml:"propertyMappings"`
}

