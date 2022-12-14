package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timeoutConfigProperty := &timeoutConfigProperty{
//   	inProgressTimeoutInMinutes: jsii.Number(123),
//   }
//
type CfnJobTemplate_TimeoutConfigProperty struct {
	// `CfnJobTemplate.TimeoutConfigProperty.InProgressTimeoutInMinutes`.
	InProgressTimeoutInMinutes *float64 `field:"required" json:"inProgressTimeoutInMinutes" yaml:"inProgressTimeoutInMinutes"`
}

