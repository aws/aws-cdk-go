package previewawsquicksightmixins


// The custom actions for a layer.
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallayerjoindefinition.html
//
type CfnDashboardPropsMixin_GeospatialLayerJoinDefinitionProperty struct {
	// The geospatial color field for the join definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallayerjoindefinition.html#cfn-quicksight-dashboard-geospatiallayerjoindefinition-colorfield
	//
	ColorField interface{} `field:"optional" json:"colorField" yaml:"colorField"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallayerjoindefinition.html#cfn-quicksight-dashboard-geospatiallayerjoindefinition-datasetkeyfield
	//
	DatasetKeyField interface{} `field:"optional" json:"datasetKeyField" yaml:"datasetKeyField"`
	// The name of the field or property in the geospatial data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatiallayerjoindefinition.html#cfn-quicksight-dashboard-geospatiallayerjoindefinition-shapekeyfield
	//
	ShapeKeyField *string `field:"optional" json:"shapeKeyField" yaml:"shapeKeyField"`
}

