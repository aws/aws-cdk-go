package awsimagebuilder


// Image status and the reason for that status.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   stateProperty := &StateProperty{
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-allimagebuildversions-state.html
//
type CfnAllImageBuildVersionsPropsMixin_StateProperty struct {
	// The status of the image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-allimagebuildversions-state.html#cfn-imagebuilder-allimagebuildversions-state-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

