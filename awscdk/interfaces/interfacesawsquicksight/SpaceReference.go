package interfacesawsquicksight


// A reference to a Space resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spaceReference := &SpaceReference{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	SpaceArn: jsii.String("spaceArn"),
//   	SpaceId: jsii.String("spaceId"),
//   }
//
type SpaceReference struct {
	// The AwsAccountId of the Space resource.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// The ARN of the Space resource.
	SpaceArn *string `field:"required" json:"spaceArn" yaml:"spaceArn"`
	// The SpaceId of the Space resource.
	SpaceId *string `field:"required" json:"spaceId" yaml:"spaceId"`
}

