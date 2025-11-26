package previewawsquicksightmixins


// A date time parameter that is created in the dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dateTimeDatasetParameterProperty := &DateTimeDatasetParameterProperty{
//   	DefaultValues: &DateTimeDatasetParameterDefaultValuesProperty{
//   		StaticValues: []*string{
//   			jsii.String("staticValues"),
//   		},
//   	},
//   	Id: jsii.String("id"),
//   	Name: jsii.String("name"),
//   	TimeGranularity: jsii.String("timeGranularity"),
//   	ValueType: jsii.String("valueType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datetimedatasetparameter.html
//
type CfnDataSetPropsMixin_DateTimeDatasetParameterProperty struct {
	// A list of default values for a given date time parameter.
	//
	// This structure only accepts static values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datetimedatasetparameter.html#cfn-quicksight-dataset-datetimedatasetparameter-defaultvalues
	//
	DefaultValues interface{} `field:"optional" json:"defaultValues" yaml:"defaultValues"`
	// An identifier for the parameter that is created in the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datetimedatasetparameter.html#cfn-quicksight-dataset-datetimedatasetparameter-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the date time parameter that is created in the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datetimedatasetparameter.html#cfn-quicksight-dataset-datetimedatasetparameter-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The time granularity of the date time parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datetimedatasetparameter.html#cfn-quicksight-dataset-datetimedatasetparameter-timegranularity
	//
	TimeGranularity *string `field:"optional" json:"timeGranularity" yaml:"timeGranularity"`
	// The value type of the dataset parameter.
	//
	// Valid values are `single value` or `multi value` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datetimedatasetparameter.html#cfn-quicksight-dataset-datetimedatasetparameter-valuetype
	//
	ValueType *string `field:"optional" json:"valueType" yaml:"valueType"`
}

