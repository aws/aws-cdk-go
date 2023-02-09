package awsconnect


// Properties for defining a `CfnSecurityKey`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSecurityKeyProps := &cfnSecurityKeyProps{
//   	instanceId: jsii.String("instanceId"),
//   	key: jsii.String("key"),
//   }
//
type CfnSecurityKeyProps struct {
	// The Amazon Resource Name (ARN) of the instance.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `100`.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// A valid security key in PEM format.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `1024`.
	Key *string `field:"required" json:"key" yaml:"key"`
}

