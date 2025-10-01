package awscloudfront


// A reference to a Distribution resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   distributionReference := &DistributionReference{
//   	DistributionId: jsii.String("distributionId"),
//   }
//
type DistributionReference struct {
	// The Id of the Distribution resource.
	DistributionId *string `field:"required" json:"distributionId" yaml:"distributionId"`
}

