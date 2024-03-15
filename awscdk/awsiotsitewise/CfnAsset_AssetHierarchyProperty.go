package awsiotsitewise


// Describes an asset hierarchy that contains a hierarchy's name and ID.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetHierarchyProperty := &AssetHierarchyProperty{
//   	ChildAssetId: jsii.String("childAssetId"),
//
//   	// the properties below are optional
//   	ExternalId: jsii.String("externalId"),
//   	Id: jsii.String("id"),
//   	LogicalId: jsii.String("logicalId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-asset-assethierarchy.html
//
type CfnAsset_AssetHierarchyProperty struct {
	// The Id of the child asset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-asset-assethierarchy.html#cfn-iotsitewise-asset-assethierarchy-childassetid
	//
	ChildAssetId *string `field:"required" json:"childAssetId" yaml:"childAssetId"`
	// String-friendly customer provided external ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-asset-assethierarchy.html#cfn-iotsitewise-asset-assethierarchy-externalid
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Customer provided actual UUID for property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-asset-assethierarchy.html#cfn-iotsitewise-asset-assethierarchy-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ID of the hierarchy.
	//
	// This ID is a `hierarchyId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-asset-assethierarchy.html#cfn-iotsitewise-asset-assethierarchy-logicalid
	//
	LogicalId *string `field:"optional" json:"logicalId" yaml:"logicalId"`
}

