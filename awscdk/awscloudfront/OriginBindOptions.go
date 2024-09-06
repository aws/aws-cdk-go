package awscloudfront


// Options passed to Origin.bind().
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   originBindOptions := &OriginBindOptions{
//   	OriginId: jsii.String("originId"),
//
//   	// the properties below are optional
//   	DistributionId: jsii.String("distributionId"),
//   }
//
type OriginBindOptions struct {
	// The identifier of this Origin, as assigned by the Distribution this Origin has been used added to.
	OriginId *string `field:"required" json:"originId" yaml:"originId"`
	// The identifier of the Distribution this Origin is used for.
	//
	// This is used to grant origin access permissions to the distribution for origin access control.
	// Default: - no distribution id.
	//
	DistributionId *string `field:"optional" json:"distributionId" yaml:"distributionId"`
}

