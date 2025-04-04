package awsdatabrew


// Represents a dataset paramater that defines type and conditions for a parameter in the Amazon S3 path of the dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datasetParameterProperty := &DatasetParameterProperty{
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	CreateColumn: jsii.Boolean(false),
//   	DatetimeOptions: &DatetimeOptionsProperty{
//   		Format: jsii.String("format"),
//
//   		// the properties below are optional
//   		LocaleCode: jsii.String("localeCode"),
//   		TimezoneOffset: jsii.String("timezoneOffset"),
//   	},
//   	Filter: &FilterExpressionProperty{
//   		Expression: jsii.String("expression"),
//   		ValuesMap: []interface{}{
//   			&FilterValueProperty{
//   				Value: jsii.String("value"),
//   				ValueReference: jsii.String("valueReference"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-datasetparameter.html
//
type CfnDataset_DatasetParameterProperty struct {
	// The name of the parameter that is used in the dataset's Amazon S3 path.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-datasetparameter.html#cfn-databrew-dataset-datasetparameter-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the dataset parameter, can be one of a 'String', 'Number' or 'Datetime'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-datasetparameter.html#cfn-databrew-dataset-datasetparameter-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// Optional boolean value that defines whether the captured value of this parameter should be loaded as an additional column in the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-datasetparameter.html#cfn-databrew-dataset-datasetparameter-createcolumn
	//
	CreateColumn interface{} `field:"optional" json:"createColumn" yaml:"createColumn"`
	// Additional parameter options such as a format and a timezone.
	//
	// Required for datetime parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-datasetparameter.html#cfn-databrew-dataset-datasetparameter-datetimeoptions
	//
	DatetimeOptions interface{} `field:"optional" json:"datetimeOptions" yaml:"datetimeOptions"`
	// The optional filter expression structure to apply additional matching criteria to the parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-dataset-datasetparameter.html#cfn-databrew-dataset-datasetparameter-filter
	//
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
}

