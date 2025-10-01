package awsmediatailor


// A reference to a LiveSource resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   liveSourceReference := &LiveSourceReference{
//   	LiveSourceArn: jsii.String("liveSourceArn"),
//   	LiveSourceName: jsii.String("liveSourceName"),
//   	SourceLocationName: jsii.String("sourceLocationName"),
//   }
//
type LiveSourceReference struct {
	// The ARN of the LiveSource resource.
	LiveSourceArn *string `field:"required" json:"liveSourceArn" yaml:"liveSourceArn"`
	// The LiveSourceName of the LiveSource resource.
	LiveSourceName *string `field:"required" json:"liveSourceName" yaml:"liveSourceName"`
	// The SourceLocationName of the LiveSource resource.
	SourceLocationName *string `field:"required" json:"sourceLocationName" yaml:"sourceLocationName"`
}

