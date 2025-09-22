package awsfrauddetector


// A reference to a Label resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   labelReference := &LabelReference{
//   	LabelArn: jsii.String("labelArn"),
//   }
//
type LabelReference struct {
	// The Arn of the Label resource.
	LabelArn *string `field:"required" json:"labelArn" yaml:"labelArn"`
}

