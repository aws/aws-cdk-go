package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secretProperty := &secretProperty{
//   	name: jsii.String("name"),
//   	valueFrom: jsii.String("valueFrom"),
//   }
//
type CfnService_SecretProperty struct {
	// `CfnService.SecretProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnService.SecretProperty.ValueFrom`.
	ValueFrom *string `field:"required" json:"valueFrom" yaml:"valueFrom"`
}

