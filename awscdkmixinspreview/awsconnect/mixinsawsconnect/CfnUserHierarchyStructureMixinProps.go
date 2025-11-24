package mixinsawsconnect


// Properties for CfnUserHierarchyStructurePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnUserHierarchyStructureMixinProps := &CfnUserHierarchyStructureMixinProps{
//   	InstanceArn: jsii.String("instanceArn"),
//   	UserHierarchyStructure: &UserHierarchyStructureProperty{
//   		LevelFive: &LevelFiveProperty{
//   			HierarchyLevelArn: jsii.String("hierarchyLevelArn"),
//   			HierarchyLevelId: jsii.String("hierarchyLevelId"),
//   			Name: jsii.String("name"),
//   		},
//   		LevelFour: &LevelFourProperty{
//   			HierarchyLevelArn: jsii.String("hierarchyLevelArn"),
//   			HierarchyLevelId: jsii.String("hierarchyLevelId"),
//   			Name: jsii.String("name"),
//   		},
//   		LevelOne: &LevelOneProperty{
//   			HierarchyLevelArn: jsii.String("hierarchyLevelArn"),
//   			HierarchyLevelId: jsii.String("hierarchyLevelId"),
//   			Name: jsii.String("name"),
//   		},
//   		LevelThree: &LevelThreeProperty{
//   			HierarchyLevelArn: jsii.String("hierarchyLevelArn"),
//   			HierarchyLevelId: jsii.String("hierarchyLevelId"),
//   			Name: jsii.String("name"),
//   		},
//   		LevelTwo: &LevelTwoProperty{
//   			HierarchyLevelArn: jsii.String("hierarchyLevelArn"),
//   			HierarchyLevelId: jsii.String("hierarchyLevelId"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-userhierarchystructure.html
//
type CfnUserHierarchyStructureMixinProps struct {
	// The Amazon Resource Name (ARN) of the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-userhierarchystructure.html#cfn-connect-userhierarchystructure-instancearn
	//
	InstanceArn *string `field:"optional" json:"instanceArn" yaml:"instanceArn"`
	// Contains information about a hierarchy structure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-userhierarchystructure.html#cfn-connect-userhierarchystructure-userhierarchystructure
	//
	UserHierarchyStructure interface{} `field:"optional" json:"userHierarchyStructure" yaml:"userHierarchyStructure"`
}

