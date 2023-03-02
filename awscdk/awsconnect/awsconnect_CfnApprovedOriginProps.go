package awsconnect


// Properties for defining a `CfnApprovedOrigin`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApprovedOriginProps := &cfnApprovedOriginProps{
//   	instanceId: jsii.String("instanceId"),
//   	origin: jsii.String("origin"),
//   }
//
type CfnApprovedOriginProps struct {
	// The Amazon Resource Name (ARN) of the instance.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `100`.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// Domain name to be added to the allow-list of the instance.
	//
	// *Maximum* : `267`.
	Origin *string `field:"required" json:"origin" yaml:"origin"`
}

