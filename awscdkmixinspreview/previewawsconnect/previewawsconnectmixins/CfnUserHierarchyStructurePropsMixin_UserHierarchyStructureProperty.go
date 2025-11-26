package previewawsconnectmixins


// Contains information about a hierarchy structure.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   userHierarchyStructureProperty := &UserHierarchyStructureProperty{
//   	LevelFive: &LevelFiveProperty{
//   		HierarchyLevelArn: jsii.String("hierarchyLevelArn"),
//   		HierarchyLevelId: jsii.String("hierarchyLevelId"),
//   		Name: jsii.String("name"),
//   	},
//   	LevelFour: &LevelFourProperty{
//   		HierarchyLevelArn: jsii.String("hierarchyLevelArn"),
//   		HierarchyLevelId: jsii.String("hierarchyLevelId"),
//   		Name: jsii.String("name"),
//   	},
//   	LevelOne: &LevelOneProperty{
//   		HierarchyLevelArn: jsii.String("hierarchyLevelArn"),
//   		HierarchyLevelId: jsii.String("hierarchyLevelId"),
//   		Name: jsii.String("name"),
//   	},
//   	LevelThree: &LevelThreeProperty{
//   		HierarchyLevelArn: jsii.String("hierarchyLevelArn"),
//   		HierarchyLevelId: jsii.String("hierarchyLevelId"),
//   		Name: jsii.String("name"),
//   	},
//   	LevelTwo: &LevelTwoProperty{
//   		HierarchyLevelArn: jsii.String("hierarchyLevelArn"),
//   		HierarchyLevelId: jsii.String("hierarchyLevelId"),
//   		Name: jsii.String("name"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-userhierarchystructure-userhierarchystructure.html
//
type CfnUserHierarchyStructurePropsMixin_UserHierarchyStructureProperty struct {
	// Information about level five.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-userhierarchystructure-userhierarchystructure.html#cfn-connect-userhierarchystructure-userhierarchystructure-levelfive
	//
	LevelFive interface{} `field:"optional" json:"levelFive" yaml:"levelFive"`
	// The update for level four.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-userhierarchystructure-userhierarchystructure.html#cfn-connect-userhierarchystructure-userhierarchystructure-levelfour
	//
	LevelFour interface{} `field:"optional" json:"levelFour" yaml:"levelFour"`
	// The update for level one.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-userhierarchystructure-userhierarchystructure.html#cfn-connect-userhierarchystructure-userhierarchystructure-levelone
	//
	LevelOne interface{} `field:"optional" json:"levelOne" yaml:"levelOne"`
	// The update for level three.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-userhierarchystructure-userhierarchystructure.html#cfn-connect-userhierarchystructure-userhierarchystructure-levelthree
	//
	LevelThree interface{} `field:"optional" json:"levelThree" yaml:"levelThree"`
	// The update for level two.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-userhierarchystructure-userhierarchystructure.html#cfn-connect-userhierarchystructure-userhierarchystructure-leveltwo
	//
	LevelTwo interface{} `field:"optional" json:"levelTwo" yaml:"levelTwo"`
}

