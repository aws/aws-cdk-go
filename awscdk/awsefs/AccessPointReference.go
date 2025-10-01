package awsefs


// A reference to a AccessPoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessPointReference := &AccessPointReference{
//   	AccessPointArn: jsii.String("accessPointArn"),
//   	AccessPointId: jsii.String("accessPointId"),
//   }
//
type AccessPointReference struct {
	// The ARN of the AccessPoint resource.
	AccessPointArn *string `field:"required" json:"accessPointArn" yaml:"accessPointArn"`
	// The AccessPointId of the AccessPoint resource.
	AccessPointId *string `field:"required" json:"accessPointId" yaml:"accessPointId"`
}

