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
type CfnDashboard_AnchorDateConfigurationProperty struct {
	// The options for the date configuration. Choose one of the options below:.
	//
	// - `NOW`.
	AnchorOption *string `field:"optional" json:"anchorOption" yaml:"anchorOption"`
	// The name of the parameter that is used for the anchor date configuration.
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
}

