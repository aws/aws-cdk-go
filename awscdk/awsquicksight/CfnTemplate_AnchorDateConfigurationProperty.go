package awsquicksight


// The date configuration of the filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   anchorDateConfigurationProperty := &AnchorDateConfigurationProperty{
//   	AnchorOption: jsii.String("anchorOption"),
//   	ParameterName: jsii.String("parameterName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-anchordateconfiguration.html
//
type CfnTemplate_AnchorDateConfigurationProperty struct {
	// The options for the date configuration. Choose one of the options below:.
	//
	// - `NOW`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-anchordateconfiguration.html#cfn-quicksight-template-anchordateconfiguration-anchoroption
	//
	AnchorOption *string `field:"optional" json:"anchorOption" yaml:"anchorOption"`
	// The name of the parameter that is used for the anchor date configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-anchordateconfiguration.html#cfn-quicksight-template-anchordateconfiguration-parametername
	//
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
}

