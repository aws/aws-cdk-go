package awsquicksight


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
type CfnAnalysis_AnchorDateConfigurationProperty struct {
	// `CfnAnalysis.AnchorDateConfigurationProperty.AnchorOption`.
	AnchorOption *string `field:"optional" json:"anchorOption" yaml:"anchorOption"`
	// `CfnAnalysis.AnchorDateConfigurationProperty.ParameterName`.
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
}

