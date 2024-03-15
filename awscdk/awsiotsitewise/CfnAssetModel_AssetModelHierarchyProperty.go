package awsiotsitewise


// Describes an asset hierarchy that contains a hierarchy's name, ID, and child asset model ID that specifies the type of asset that can be in this hierarchy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetModelHierarchyProperty := &AssetModelHierarchyProperty{
//   	ChildAssetModelId: jsii.String("childAssetModelId"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	ExternalId: jsii.String("externalId"),
//   	Id: jsii.String("id"),
//   	LogicalId: jsii.String("logicalId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelhierarchy.html
//
type CfnAssetModel_AssetModelHierarchyProperty struct {
	// The ID of the asset model, in UUID format.
	//
	// All assets in this hierarchy must be instances of the `childAssetModelId` asset model. AWS IoT SiteWise will always return the actual asset model ID for this value. However, when you are specifying this value as part of a call to [UpdateAssetModel](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_UpdateAssetModel.html) , you may provide either the asset model ID or else `externalId:` followed by the asset model's external ID. For more information, see [Using external IDs](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/object-ids.html#external-ids) in the *AWS IoT SiteWise User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelhierarchy.html#cfn-iotsitewise-assetmodel-assetmodelhierarchy-childassetmodelid
	//
	ChildAssetModelId *string `field:"required" json:"childAssetModelId" yaml:"childAssetModelId"`
	// The name of the asset model hierarchy that you specify by using the [CreateAssetModel](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_CreateAssetModel.html) or [UpdateAssetModel](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_UpdateAssetModel.html) API operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelhierarchy.html#cfn-iotsitewise-assetmodel-assetmodelhierarchy-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Customer provided external ID for hierarchy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelhierarchy.html#cfn-iotsitewise-assetmodel-assetmodelhierarchy-externalid
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Customer provided actual ID for hierarchy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelhierarchy.html#cfn-iotsitewise-assetmodel-assetmodelhierarchy-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The `LogicalID` of the asset model hierarchy.
	//
	// This ID is a `hierarchyLogicalId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-assetmodelhierarchy.html#cfn-iotsitewise-assetmodel-assetmodelhierarchy-logicalid
	//
	LogicalId *string `field:"optional" json:"logicalId" yaml:"logicalId"`
}

