package awsemr


// A reference to a Studio resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   studioReference := &StudioReference{
//   	StudioArn: jsii.String("studioArn"),
//   	StudioId: jsii.String("studioId"),
//   }
//
type StudioReference struct {
	// The ARN of the Studio resource.
	StudioArn *string `field:"required" json:"studioArn" yaml:"studioArn"`
	// The StudioId of the Studio resource.
	StudioId *string `field:"required" json:"studioId" yaml:"studioId"`
}

