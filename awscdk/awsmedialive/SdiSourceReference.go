package awsmedialive


// A reference to a SdiSource resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sdiSourceReference := &SdiSourceReference{
//   	SdiSourceArn: jsii.String("sdiSourceArn"),
//   	SdiSourceId: jsii.String("sdiSourceId"),
//   }
//
type SdiSourceReference struct {
	// The ARN of the SdiSource resource.
	SdiSourceArn *string `field:"required" json:"sdiSourceArn" yaml:"sdiSourceArn"`
	// The Id of the SdiSource resource.
	SdiSourceId *string `field:"required" json:"sdiSourceId" yaml:"sdiSourceId"`
}

