package awsverifiedpermissions


// A reference to a PolicyStore resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyStoreReference := &PolicyStoreReference{
//   	PolicyStoreArn: jsii.String("policyStoreArn"),
//   	PolicyStoreId: jsii.String("policyStoreId"),
//   }
//
type PolicyStoreReference struct {
	// The ARN of the PolicyStore resource.
	PolicyStoreArn *string `field:"required" json:"policyStoreArn" yaml:"policyStoreArn"`
	// The PolicyStoreId of the PolicyStore resource.
	PolicyStoreId *string `field:"required" json:"policyStoreId" yaml:"policyStoreId"`
}

