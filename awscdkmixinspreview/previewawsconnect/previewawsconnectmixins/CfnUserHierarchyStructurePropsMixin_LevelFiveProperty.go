package previewawsconnectmixins


// The update for level five.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   levelFiveProperty := &LevelFiveProperty{
//   	HierarchyLevelArn: jsii.String("hierarchyLevelArn"),
//   	HierarchyLevelId: jsii.String("hierarchyLevelId"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-userhierarchystructure-levelfive.html
//
type CfnUserHierarchyStructurePropsMixin_LevelFiveProperty struct {
	// The Amazon Resource Name (ARN) of the hierarchy level.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-userhierarchystructure-levelfive.html#cfn-connect-userhierarchystructure-levelfive-hierarchylevelarn
	//
	HierarchyLevelArn *string `field:"optional" json:"hierarchyLevelArn" yaml:"hierarchyLevelArn"`
	// The identifier of the hierarchy level.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-userhierarchystructure-levelfive.html#cfn-connect-userhierarchystructure-levelfive-hierarchylevelid
	//
	HierarchyLevelId *string `field:"optional" json:"hierarchyLevelId" yaml:"hierarchyLevelId"`
	// The name of the hierarchy level.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-userhierarchystructure-levelfive.html#cfn-connect-userhierarchystructure-levelfive-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

