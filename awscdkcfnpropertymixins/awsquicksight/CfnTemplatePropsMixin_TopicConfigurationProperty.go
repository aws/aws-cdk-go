package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   topicConfigurationProperty := &TopicConfigurationProperty{
//   	ColumnGroupSchemaList: []interface{}{
//   		&ColumnGroupSchemaProperty{
//   			ColumnGroupColumnSchemaList: []interface{}{
//   				&ColumnGroupColumnSchemaProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	DataSetSchema: &DataSetSchemaProperty{
//   		ColumnSchemaList: []interface{}{
//   			&ColumnSchemaProperty{
//   				DataType: jsii.String("dataType"),
//   				GeographicRole: jsii.String("geographicRole"),
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	Placeholder: jsii.String("placeholder"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topicconfiguration.html
//
type CfnTemplatePropsMixin_TopicConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topicconfiguration.html#cfn-quicksight-template-topicconfiguration-columngroupschemalist
	//
	ColumnGroupSchemaList interface{} `field:"optional" json:"columnGroupSchemaList" yaml:"columnGroupSchemaList"`
	// <p>Dataset schema.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topicconfiguration.html#cfn-quicksight-template-topicconfiguration-datasetschema
	//
	DataSetSchema interface{} `field:"optional" json:"dataSetSchema" yaml:"dataSetSchema"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topicconfiguration.html#cfn-quicksight-template-topicconfiguration-placeholder
	//
	Placeholder *string `field:"optional" json:"placeholder" yaml:"placeholder"`
}

