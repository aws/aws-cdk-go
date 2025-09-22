package awscleanrooms


// A reference to a ConfiguredTableAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   configuredTableAssociationReference := &ConfiguredTableAssociationReference{
//   	ConfiguredTableAssociationArn: jsii.String("configuredTableAssociationArn"),
//   	ConfiguredTableAssociationIdentifier: jsii.String("configuredTableAssociationIdentifier"),
//   	MembershipIdentifier: jsii.String("membershipIdentifier"),
//   }
//
type ConfiguredTableAssociationReference struct {
	// The ARN of the ConfiguredTableAssociation resource.
	ConfiguredTableAssociationArn *string `field:"required" json:"configuredTableAssociationArn" yaml:"configuredTableAssociationArn"`
	// The ConfiguredTableAssociationIdentifier of the ConfiguredTableAssociation resource.
	ConfiguredTableAssociationIdentifier *string `field:"required" json:"configuredTableAssociationIdentifier" yaml:"configuredTableAssociationIdentifier"`
	// The MembershipIdentifier of the ConfiguredTableAssociation resource.
	MembershipIdentifier *string `field:"required" json:"membershipIdentifier" yaml:"membershipIdentifier"`
}

