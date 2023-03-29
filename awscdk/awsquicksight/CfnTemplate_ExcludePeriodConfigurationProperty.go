package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   excludePeriodConfigurationProperty := &ExcludePeriodConfigurationProperty{
//   	Amount: jsii.Number(123),
//   	Granularity: jsii.String("granularity"),
//
//   	// the properties below are optional
//   	Status: jsii.String("status"),
//   }
//
type CfnTemplate_ExcludePeriodConfigurationProperty struct {
	// `CfnTemplate.ExcludePeriodConfigurationProperty.Amount`.
	Amount *float64 `field:"required" json:"amount" yaml:"amount"`
	// `CfnTemplate.ExcludePeriodConfigurationProperty.Granularity`.
	Granularity *string `field:"required" json:"granularity" yaml:"granularity"`
	// `CfnTemplate.ExcludePeriodConfigurationProperty.Status`.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

