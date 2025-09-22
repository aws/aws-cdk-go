package awsquicksight


// The configuration of custom values for the destination parameter in `DestinationParameterValueConfiguration` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customValuesConfigurationProperty := &CustomValuesConfigurationProperty{
//   	CustomValues: &CustomParameterValuesProperty{
//   		DateTimeValues: []*string{
//   			jsii.String("dateTimeValues"),
//   		},
//   		DecimalValues: []interface{}{
//   			jsii.Number(123),
//   		},
//   		IntegerValues: []interface{}{
//   			jsii.Number(123),
//   		},
//   		StringValues: []*string{
//   			jsii.String("stringValues"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	IncludeNullValue: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customvaluesconfiguration.html
//
type CfnDashboard_CustomValuesConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customvaluesconfiguration.html#cfn-quicksight-dashboard-customvaluesconfiguration-customvalues
	//
	CustomValues interface{} `field:"required" json:"customValues" yaml:"customValues"`
	// Includes the null value in custom action parameter values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-customvaluesconfiguration.html#cfn-quicksight-dashboard-customvaluesconfiguration-includenullvalue
	//
	IncludeNullValue interface{} `field:"optional" json:"includeNullValue" yaml:"includeNullValue"`
}

