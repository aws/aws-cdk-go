package awslightsail


// Properties for defining a `CfnStaticIp`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStaticIpProps := &cfnStaticIpProps{
//   	staticIpName: jsii.String("staticIpName"),
//
//   	// the properties below are optional
//   	attachedTo: jsii.String("attachedTo"),
//   }
//
type CfnStaticIpProps struct {
	// The name of the static IP.
	StaticIpName *string `field:"required" json:"staticIpName" yaml:"staticIpName"`
	// The instance that the static IP is attached to.
	AttachedTo *string `field:"optional" json:"attachedTo" yaml:"attachedTo"`
}

