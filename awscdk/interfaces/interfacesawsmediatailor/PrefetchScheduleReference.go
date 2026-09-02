package interfacesawsmediatailor


// A reference to a PrefetchSchedule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   prefetchScheduleReference := &PrefetchScheduleReference{
//   	PrefetchScheduleArn: jsii.String("prefetchScheduleArn"),
//   }
//
type PrefetchScheduleReference struct {
	// The Arn of the PrefetchSchedule resource.
	PrefetchScheduleArn *string `field:"required" json:"prefetchScheduleArn" yaml:"prefetchScheduleArn"`
}

