package awsconnect


// Properties for defining a `CfnUserHierarchyStructure`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserHierarchyStructureProps := &CfnUserHierarchyStructureProps{
//   	InstanceArn: jsii.String("instanceArn"),
//
//   	// the properties below are optional
//   	UserHierarchyStructure: &UserHierarchyStructureProperty{
//   		LevelFive: &LevelFiveProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			HierarchyLevelArn: jsii.String("hierarchyLevelArn"),
//   			HierarchyLevelId: jsii.String("hierarchyLevelId"),
//   		},
//   		LevelFour: &LevelFourProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			HierarchyLevelArn: jsii.String("hierarchyLevelArn"),
//   			HierarchyLevelId: jsii.String("hierarchyLevelId"),
//   		},
//   		LevelOne: &LevelOneProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			HierarchyLevelArn: jsii.String("hierarchyLevelArn"),
//   			HierarchyLevelId: jsii.String("hierarchyLevelId"),
//   		},
//   		LevelThree: &LevelThreeProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			HierarchyLevelArn: jsii.String("hierarchyLevelArn"),
//   			HierarchyLevelId: jsii.String("hierarchyLevelId"),
//   		},
//   		LevelTwo: &LevelTwoProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			HierarchyLevelArn: jsii.String("hierarchyLevelArn"),
//   			HierarchyLevelId: jsii.String("hierarchyLevelId"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-userhierarchystructure.html
//
type CfnUserHierarchyStructureProps struct {
	// The Amazon Resource Name (ARN) of the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-userhierarchystructure.html#cfn-connect-userhierarchystructure-instancearn
	//
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// Contains information about a hierarchy structure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-userhierarchystructure.html#cfn-connect-userhierarchystructure-userhierarchystructure
	//
	UserHierarchyStructure interface{} `field:"optional" json:"userHierarchyStructure" yaml:"userHierarchyStructure"`
}

