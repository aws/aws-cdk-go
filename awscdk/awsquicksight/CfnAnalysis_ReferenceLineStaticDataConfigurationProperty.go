package awsquicksight


// The static data configuration of the reference line data configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   referenceLineStaticDataConfigurationProperty := &ReferenceLineStaticDataConfigurationProperty{
//   	Value: jsii.Number(123),
//   }
//
type CfnAnalysis_ReferenceLineStaticDataConfigurationProperty struct {
	// The double input of the static data.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

