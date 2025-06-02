package awsconnect


// The update for level two.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   levelTwoProperty := &LevelTwoProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	HierarchyLevelArn: jsii.String("hierarchyLevelArn"),
//   	HierarchyLevelId: jsii.String("hierarchyLevelId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-userhierarchystructure-leveltwo.html
//
type CfnUserHierarchyStructure_LevelTwoProperty struct {
	// The name of the hierarchy level.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-userhierarchystructure-leveltwo.html#cfn-connect-userhierarchystructure-leveltwo-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the hierarchy level.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-userhierarchystructure-leveltwo.html#cfn-connect-userhierarchystructure-leveltwo-hierarchylevelarn
	//
	HierarchyLevelArn *string `field:"optional" json:"hierarchyLevelArn" yaml:"hierarchyLevelArn"`
	// The identifier of the hierarchy level.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-userhierarchystructure-leveltwo.html#cfn-connect-userhierarchystructure-leveltwo-hierarchylevelid
	//
	HierarchyLevelId *string `field:"optional" json:"hierarchyLevelId" yaml:"hierarchyLevelId"`
}

