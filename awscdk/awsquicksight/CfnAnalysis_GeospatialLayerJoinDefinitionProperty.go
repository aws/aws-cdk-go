package awsquicksight


// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallayerjoindefinition.html
//
type CfnAnalysis_GeospatialLayerJoinDefinitionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallayerjoindefinition.html#cfn-quicksight-analysis-geospatiallayerjoindefinition-colorfield
	//
	ColorField interface{} `field:"optional" json:"colorField" yaml:"colorField"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallayerjoindefinition.html#cfn-quicksight-analysis-geospatiallayerjoindefinition-datasetkeyfield
	//
	DatasetKeyField interface{} `field:"optional" json:"datasetKeyField" yaml:"datasetKeyField"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallayerjoindefinition.html#cfn-quicksight-analysis-geospatiallayerjoindefinition-shapekeyfield
	//
	ShapeKeyField *string `field:"optional" json:"shapeKeyField" yaml:"shapeKeyField"`
}

