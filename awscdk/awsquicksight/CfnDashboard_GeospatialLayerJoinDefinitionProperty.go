package awsquicksight


// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallayerjoindefinition.html
//
type CfnDashboard_GeospatialLayerJoinDefinitionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallayerjoindefinition.html#cfn-quicksight-dashboard-geospatiallayerjoindefinition-colorfield
	//
	ColorField interface{} `field:"optional" json:"colorField" yaml:"colorField"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallayerjoindefinition.html#cfn-quicksight-dashboard-geospatiallayerjoindefinition-datasetkeyfield
	//
	DatasetKeyField interface{} `field:"optional" json:"datasetKeyField" yaml:"datasetKeyField"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallayerjoindefinition.html#cfn-quicksight-dashboard-geospatiallayerjoindefinition-shapekeyfield
	//
	ShapeKeyField *string `field:"optional" json:"shapeKeyField" yaml:"shapeKeyField"`
}

