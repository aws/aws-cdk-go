package mixinsawsquicksight


// The rolling date configuration of a date time filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   rollingDateConfigurationProperty := &RollingDateConfigurationProperty{
//   	DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	Expression: jsii.String("expression"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-rollingdateconfiguration.html
//
type CfnTemplatePropsMixin_RollingDateConfigurationProperty struct {
	// The data set that is used in the rolling date configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-rollingdateconfiguration.html#cfn-quicksight-template-rollingdateconfiguration-datasetidentifier
	//
	DataSetIdentifier *string `field:"optional" json:"dataSetIdentifier" yaml:"dataSetIdentifier"`
	// The expression of the rolling date configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-rollingdateconfiguration.html#cfn-quicksight-template-rollingdateconfiguration-expression
	//
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

