package awssagemaker


// A reference to a Space resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spaceReference := &SpaceReference{
//   	DomainId: jsii.String("domainId"),
//   	SpaceArn: jsii.String("spaceArn"),
//   	SpaceName: jsii.String("spaceName"),
//   }
//
type SpaceReference struct {
	// The DomainId of the Space resource.
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// The ARN of the Space resource.
	SpaceArn *string `field:"required" json:"spaceArn" yaml:"spaceArn"`
	// The SpaceName of the Space resource.
	SpaceName *string `field:"required" json:"spaceName" yaml:"spaceName"`
}

