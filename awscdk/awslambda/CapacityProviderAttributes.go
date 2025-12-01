package awslambda


// Attributes for importing an existing Lambda capacity provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityProviderAttributes := &CapacityProviderAttributes{
//   	CapacityProviderArn: jsii.String("capacityProviderArn"),
//   }
//
type CapacityProviderAttributes struct {
	// The Amazon Resource Name (ARN) of the capacity provider.
	//
	// Format: arn:<partition>:lambda:<region>:<account-id>:capacity-provider:<capacity-provider-name>.
	CapacityProviderArn *string `field:"required" json:"capacityProviderArn" yaml:"capacityProviderArn"`
}

