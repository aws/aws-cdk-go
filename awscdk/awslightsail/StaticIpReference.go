package awslightsail


// A reference to a StaticIp resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   staticIpReference := &StaticIpReference{
//   	StaticIpArn: jsii.String("staticIpArn"),
//   	StaticIpName: jsii.String("staticIpName"),
//   }
//
type StaticIpReference struct {
	// The ARN of the StaticIp resource.
	StaticIpArn *string `field:"required" json:"staticIpArn" yaml:"staticIpArn"`
	// The StaticIpName of the StaticIp resource.
	StaticIpName *string `field:"required" json:"staticIpName" yaml:"staticIpName"`
}

