package awsgreengrassv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ioTJobTimeoutConfigProperty := &ioTJobTimeoutConfigProperty{
//   	inProgressTimeoutInMinutes: jsii.Number(123),
//   }
//
type CfnDeployment_IoTJobTimeoutConfigProperty struct {
	// `CfnDeployment.IoTJobTimeoutConfigProperty.InProgressTimeoutInMinutes`.
	InProgressTimeoutInMinutes *float64 `field:"optional" json:"inProgressTimeoutInMinutes" yaml:"inProgressTimeoutInMinutes"`
}

