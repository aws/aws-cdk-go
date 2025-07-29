package awsquicksight


// The option that determines the decimal places configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   decimalPlacesConfigurationProperty := &DecimalPlacesConfigurationProperty{
//   	DecimalPlaces: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-decimalplacesconfiguration.html
//
type CfnAnalysis_DecimalPlacesConfigurationProperty struct {
	// The values of the decimal places.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-decimalplacesconfiguration.html#cfn-quicksight-analysis-decimalplacesconfiguration-decimalplaces
	//
	DecimalPlaces *float64 `field:"required" json:"decimalPlaces" yaml:"decimalPlaces"`
}

