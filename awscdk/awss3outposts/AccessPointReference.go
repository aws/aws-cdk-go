package awss3outposts


// A reference to a AccessPoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessPointReference := &AccessPointReference{
//   	AccessPointArn: jsii.String("accessPointArn"),
//   }
//
type AccessPointReference struct {
	// The Arn of the AccessPoint resource.
	AccessPointArn *string `field:"required" json:"accessPointArn" yaml:"accessPointArn"`
}

