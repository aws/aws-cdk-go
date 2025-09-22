package awsservicecatalog


// A reference to a CloudFormationProduct resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudFormationProductReference := &CloudFormationProductReference{
//   	CloudFormationProductId: jsii.String("cloudFormationProductId"),
//   }
//
type CloudFormationProductReference struct {
	// The Id of the CloudFormationProduct resource.
	CloudFormationProductId *string `field:"required" json:"cloudFormationProductId" yaml:"cloudFormationProductId"`
}

