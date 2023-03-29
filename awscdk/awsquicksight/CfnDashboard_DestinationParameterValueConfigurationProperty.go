package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationParameterValueConfigurationProperty := &DestinationParameterValueConfigurationProperty{
//   	CustomValuesConfiguration: &CustomValuesConfigurationProperty{
//   		CustomValues: &CustomParameterValuesProperty{
//   			DateTimeValues: []*string{
//   				jsii.String("dateTimeValues"),
//   			},
//   			DecimalValues: []interface{}{
//   				jsii.Number(123),
//   			},
//   			IntegerValues: []interface{}{
//   				jsii.Number(123),
//   			},
//   			StringValues: []*string{
//   				jsii.String("stringValues"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		IncludeNullValue: jsii.Boolean(false),
//   	},
//   	SelectAllValueOptions: jsii.String("selectAllValueOptions"),
//   	SourceField: jsii.String("sourceField"),
//   	SourceParameterName: jsii.String("sourceParameterName"),
//   }
//
type CfnDashboard_DestinationParameterValueConfigurationProperty struct {
	// `CfnDashboard.DestinationParameterValueConfigurationProperty.CustomValuesConfiguration`.
	CustomValuesConfiguration interface{} `field:"optional" json:"customValuesConfiguration" yaml:"customValuesConfiguration"`
	// `CfnDashboard.DestinationParameterValueConfigurationProperty.SelectAllValueOptions`.
	SelectAllValueOptions *string `field:"optional" json:"selectAllValueOptions" yaml:"selectAllValueOptions"`
	// `CfnDashboard.DestinationParameterValueConfigurationProperty.SourceField`.
	SourceField *string `field:"optional" json:"sourceField" yaml:"sourceField"`
	// `CfnDashboard.DestinationParameterValueConfigurationProperty.SourceParameterName`.
	SourceParameterName *string `field:"optional" json:"sourceParameterName" yaml:"sourceParameterName"`
}

