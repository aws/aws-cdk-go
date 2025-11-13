package interfacesawscloudtrail


// A reference to a Trail resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   trailReference := &TrailReference{
//   	TrailArn: jsii.String("trailArn"),
//   	TrailName: jsii.String("trailName"),
//   }
//
type TrailReference struct {
	// The ARN of the Trail resource.
	TrailArn *string `field:"required" json:"trailArn" yaml:"trailArn"`
	// The TrailName of the Trail resource.
	TrailName *string `field:"required" json:"trailName" yaml:"trailName"`
}

