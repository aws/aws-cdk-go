package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   referenceLineStyleConfigurationProperty := &ReferenceLineStyleConfigurationProperty{
//   	Color: jsii.String("color"),
//   	Pattern: jsii.String("pattern"),
//   }
//
type CfnAnalysis_ReferenceLineStyleConfigurationProperty struct {
	// `CfnAnalysis.ReferenceLineStyleConfigurationProperty.Color`.
	Color *string `field:"optional" json:"color" yaml:"color"`
	// `CfnAnalysis.ReferenceLineStyleConfigurationProperty.Pattern`.
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
}

