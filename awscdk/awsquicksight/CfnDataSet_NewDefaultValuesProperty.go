package awsquicksight


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-newdefaultvalues.html#cfn-quicksight-dataset-newdefaultvalues-datetimestaticvalues
	//
	DateTimeStaticValues *[]*string `field:"optional" json:"dateTimeStaticValues" yaml:"dateTimeStaticValues"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-newdefaultvalues.html#cfn-quicksight-dataset-newdefaultvalues-decimalstaticvalues
	//
	DecimalStaticValues interface{} `field:"optional" json:"decimalStaticValues" yaml:"decimalStaticValues"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-newdefaultvalues.html#cfn-quicksight-dataset-newdefaultvalues-integerstaticvalues
	//
	IntegerStaticValues interface{} `field:"optional" json:"integerStaticValues" yaml:"integerStaticValues"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-newdefaultvalues.html#cfn-quicksight-dataset-newdefaultvalues-stringstaticvalues
	//
	StringStaticValues *[]*string `field:"optional" json:"stringStaticValues" yaml:"stringStaticValues"`
}

