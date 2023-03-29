package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   arcAxisConfigurationProperty := &ArcAxisConfigurationProperty{
//   	Range: &ArcAxisDisplayRangeProperty{
//   		Max: jsii.Number(123),
//   		Min: jsii.Number(123),
//   	},
//   	ReserveRange: jsii.Number(123),
//   }
//
type CfnTemplate_ArcAxisConfigurationProperty struct {
	// `CfnTemplate.ArcAxisConfigurationProperty.Range`.
	Range interface{} `field:"optional" json:"range" yaml:"range"`
	// `CfnTemplate.ArcAxisConfigurationProperty.ReserveRange`.
	ReserveRange *float64 `field:"optional" json:"reserveRange" yaml:"reserveRange"`
}

