package awsimagebuilder


// A reference to a DistributionConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   distributionConfigurationReference := &DistributionConfigurationReference{
//   	DistributionConfigurationArn: jsii.String("distributionConfigurationArn"),
//   }
//
type DistributionConfigurationReference struct {
	// The Arn of the DistributionConfiguration resource.
	DistributionConfigurationArn *string `field:"required" json:"distributionConfigurationArn" yaml:"distributionConfigurationArn"`
}

