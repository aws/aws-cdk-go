package interfacesawscloudfront


// A reference to a RealtimeLogConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   realtimeLogConfigReference := &RealtimeLogConfigReference{
//   	RealtimeLogConfigArn: jsii.String("realtimeLogConfigArn"),
//   }
//
type RealtimeLogConfigReference struct {
	// The Arn of the RealtimeLogConfig resource.
	RealtimeLogConfigArn *string `field:"required" json:"realtimeLogConfigArn" yaml:"realtimeLogConfigArn"`
}

