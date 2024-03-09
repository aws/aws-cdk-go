package awsiotsitewise


// Identifies a property value used in an expression.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   variableValueProperty := &VariableValueProperty{
//   	HierarchyExternalId: jsii.String("hierarchyExternalId"),
//   	HierarchyId: jsii.String("hierarchyId"),
//   	HierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   	PropertyExternalId: jsii.String("propertyExternalId"),
//   	PropertyId: jsii.String("propertyId"),
//   	PropertyLogicalId: jsii.String("propertyLogicalId"),
//   	PropertyPath: []interface{}{
//   		&PropertyPathDefinitionProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-variablevalue.html
//
type CfnAssetModel_VariableValueProperty struct {
	// The External ID of the hierarchy that is trying to be referenced.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-variablevalue.html#cfn-iotsitewise-assetmodel-variablevalue-hierarchyexternalid
	//
	HierarchyExternalId *string `field:"optional" json:"hierarchyExternalId" yaml:"hierarchyExternalId"`
	// The ID of the hierarchy that is trying to be referenced.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-variablevalue.html#cfn-iotsitewise-assetmodel-variablevalue-hierarchyid
	//
	HierarchyId *string `field:"optional" json:"hierarchyId" yaml:"hierarchyId"`
	// The `LogicalID` of the hierarchy to query for the `PropertyLogicalID` .
	//
	// You use a `hierarchyLogicalID` instead of a model ID because you can have several hierarchies using the same model and therefore the same property. For example, you might have separately grouped assets that come from the same asset model. For more information, see [Defining relationships between assets](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-hierarchies.html) in the *AWS IoT SiteWise User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-variablevalue.html#cfn-iotsitewise-assetmodel-variablevalue-hierarchylogicalid
	//
	HierarchyLogicalId *string `field:"optional" json:"hierarchyLogicalId" yaml:"hierarchyLogicalId"`
	// The External ID of the property that is trying to be referenced.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-variablevalue.html#cfn-iotsitewise-assetmodel-variablevalue-propertyexternalid
	//
	PropertyExternalId *string `field:"optional" json:"propertyExternalId" yaml:"propertyExternalId"`
	// The ID of the property that is trying to be referenced.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-variablevalue.html#cfn-iotsitewise-assetmodel-variablevalue-propertyid
	//
	PropertyId *string `field:"optional" json:"propertyId" yaml:"propertyId"`
	// The `LogicalID` of the property to use as the variable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-variablevalue.html#cfn-iotsitewise-assetmodel-variablevalue-propertylogicalid
	//
	PropertyLogicalId *string `field:"optional" json:"propertyLogicalId" yaml:"propertyLogicalId"`
	// The path of the property that is trying to be referenced.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-variablevalue.html#cfn-iotsitewise-assetmodel-variablevalue-propertypath
	//
	PropertyPath interface{} `field:"optional" json:"propertyPath" yaml:"propertyPath"`
}

