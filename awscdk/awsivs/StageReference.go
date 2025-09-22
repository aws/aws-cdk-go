package awsivs


// A reference to a Stage resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stageReference := &StageReference{
//   	StageArn: jsii.String("stageArn"),
//   }
//
type StageReference struct {
	// The Arn of the Stage resource.
	StageArn *string `field:"required" json:"stageArn" yaml:"stageArn"`
}

