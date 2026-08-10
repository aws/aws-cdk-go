package interfacesawsgreengrassv2


// A reference to a CoreDevice resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   coreDeviceReference := &CoreDeviceReference{
//   	CoreDeviceArn: jsii.String("coreDeviceArn"),
//   }
//
type CoreDeviceReference struct {
	// The Arn of the CoreDevice resource.
	CoreDeviceArn *string `field:"required" json:"coreDeviceArn" yaml:"coreDeviceArn"`
}

