package awslightsail


// A reference to a Distribution resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   distributionReference := &DistributionReference{
//   	DistributionArn: jsii.String("distributionArn"),
//   	DistributionName: jsii.String("distributionName"),
//   }
//
type DistributionReference struct {
	// The ARN of the Distribution resource.
	DistributionArn *string `field:"required" json:"distributionArn" yaml:"distributionArn"`
	// The DistributionName of the Distribution resource.
	DistributionName *string `field:"required" json:"distributionName" yaml:"distributionName"`
}

