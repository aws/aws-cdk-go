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
	// `AWS::Connect::SecurityKey.InstanceId`.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The key of the security key.
	Key *string `field:"required" json:"key" yaml:"key"`
}

