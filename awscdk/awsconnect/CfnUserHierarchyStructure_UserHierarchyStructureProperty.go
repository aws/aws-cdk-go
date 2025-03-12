package awsconnect


// Contains information about a hierarchy structure.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userHierarchyStructureProperty := &UserHierarchyStructureProperty{
//   	LevelFive: &LevelFiveProperty{
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		HierarchyLevelArn: jsii.String("hierarchyLevelArn"),
//   		HierarchyLevelId: jsii.String("hierarchyLevelId"),
//   	},
//   	LevelFour: &LevelFourProperty{
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		HierarchyLevelArn: jsii.String("hierarchyLevelArn"),
//   		HierarchyLevelId: jsii.String("hierarchyLevelId"),
//   	},
//   	LevelOne: &LevelOneProperty{
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		HierarchyLevelArn: jsii.String("hierarchyLevelArn"),
//   		HierarchyLevelId: jsii.String("hierarchyLevelId"),
//   	},
//   	LevelThree: &LevelThreeProperty{
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		HierarchyLevelArn: jsii.String("hierarchyLevelArn"),
//   		HierarchyLevelId: jsii.String("hierarchyLevelId"),
//   	},
//   	LevelTwo: &LevelTwoProperty{
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		HierarchyLevelArn: jsii.String("hierarchyLevelArn"),
//   		HierarchyLevelId: jsii.String("hierarchyLevelId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-userhierarchystructure-userhierarchystructure.html
//
type CfnUserHierarchyStructure_UserHierarchyStructureProperty struct {
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

