package awsquicksight


// The rolling date configuration of a date time filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rollingDateConfigurationProperty := &RollingDateConfigurationProperty{
//   	Expression: jsii.String("expression"),
//
//   	// the properties below are optional
//   	DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-rollingdateconfiguration.html
//
type CfnDashboard_RollingDateConfigurationProperty struct {
	// The expression of the rolling date configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-rollingdateconfiguration.html#cfn-quicksight-dashboard-rollingdateconfiguration-expression
	//
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The data set that is used in the rolling date configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-rollingdateconfiguration.html#cfn-quicksight-dashboard-rollingdateconfiguration-datasetidentifier
	//
	DataSetIdentifier *string `field:"optional" json:"dataSetIdentifier" yaml:"dataSetIdentifier"`
}

