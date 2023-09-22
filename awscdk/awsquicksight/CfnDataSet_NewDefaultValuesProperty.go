package awsquicksight


// The new default values for the parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   newDefaultValuesProperty := &NewDefaultValuesProperty{
//   	DateTimeStaticValues: []*string{
//   		jsii.String("dateTimeStaticValues"),
//   	},
//   	DecimalStaticValues: []interface{}{
//   		jsii.Number(123),
//   	},
//   	IntegerStaticValues: []interface{}{
//   		jsii.Number(123),
//   	},
//   	StringStaticValues: []*string{
//   		jsii.String("stringStaticValues"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-newdefaultvalues.html
//
type CfnDataSet_NewDefaultValuesProperty struct {
	// A list of static default values for a given date time parameter.
	//
	// The valid format for this property is `yyyy-MM-dd’T’HH:mm:ss’Z’` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-newdefaultvalues.html#cfn-quicksight-dataset-newdefaultvalues-datetimestaticvalues
	//
	DateTimeStaticValues *[]*string `field:"optional" json:"dateTimeStaticValues" yaml:"dateTimeStaticValues"`
	// A list of static default values for a given decimal parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-newdefaultvalues.html#cfn-quicksight-dataset-newdefaultvalues-decimalstaticvalues
	//
	DecimalStaticValues interface{} `field:"optional" json:"decimalStaticValues" yaml:"decimalStaticValues"`
	// A list of static default values for a given integer parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-newdefaultvalues.html#cfn-quicksight-dataset-newdefaultvalues-integerstaticvalues
	//
	IntegerStaticValues interface{} `field:"optional" json:"integerStaticValues" yaml:"integerStaticValues"`
	// A list of static default values for a given string parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-newdefaultvalues.html#cfn-quicksight-dataset-newdefaultvalues-stringstaticvalues
	//
	StringStaticValues *[]*string `field:"optional" json:"stringStaticValues" yaml:"stringStaticValues"`
}

