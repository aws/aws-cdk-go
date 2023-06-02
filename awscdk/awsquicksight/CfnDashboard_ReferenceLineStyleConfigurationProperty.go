package awsquicksight


// The style configuration of the reference line.
//
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
type CfnDashboard_ReferenceLineStyleConfigurationProperty struct {
	// The hex color of the reference line.
	Color *string `field:"optional" json:"color" yaml:"color"`
	// The pattern type of the line style. Choose one of the following options:.
	//
	// - `SOLID`
	// - `DASHED`
	// - `DOTTED`.
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
}

