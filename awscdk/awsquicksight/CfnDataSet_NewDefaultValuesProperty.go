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
type CfnDataSet_NewDefaultValuesProperty struct {
	// `CfnDataSet.NewDefaultValuesProperty.DateTimeStaticValues`.
	DateTimeStaticValues *[]*string `field:"optional" json:"dateTimeStaticValues" yaml:"dateTimeStaticValues"`
	// `CfnDataSet.NewDefaultValuesProperty.DecimalStaticValues`.
	DecimalStaticValues interface{} `field:"optional" json:"decimalStaticValues" yaml:"decimalStaticValues"`
	// `CfnDataSet.NewDefaultValuesProperty.IntegerStaticValues`.
	IntegerStaticValues interface{} `field:"optional" json:"integerStaticValues" yaml:"integerStaticValues"`
	// `CfnDataSet.NewDefaultValuesProperty.StringStaticValues`.
	StringStaticValues *[]*string `field:"optional" json:"stringStaticValues" yaml:"stringStaticValues"`
}

