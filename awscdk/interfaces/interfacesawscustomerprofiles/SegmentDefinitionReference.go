package interfacesawscustomerprofiles


// A reference to a SegmentDefinition resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   segmentDefinitionReference := &SegmentDefinitionReference{
//   	DomainName: jsii.String("domainName"),
//   	SegmentDefinitionArn: jsii.String("segmentDefinitionArn"),
//   	SegmentDefinitionName: jsii.String("segmentDefinitionName"),
//   }
//
type SegmentDefinitionReference struct {
	// The DomainName of the SegmentDefinition resource.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The ARN of the SegmentDefinition resource.
	SegmentDefinitionArn *string `field:"required" json:"segmentDefinitionArn" yaml:"segmentDefinitionArn"`
	// The SegmentDefinitionName of the SegmentDefinition resource.
	SegmentDefinitionName *string `field:"required" json:"segmentDefinitionName" yaml:"segmentDefinitionName"`
}

