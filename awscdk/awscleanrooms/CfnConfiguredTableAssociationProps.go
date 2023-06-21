package awscleanrooms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnConfiguredTableAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfiguredTableAssociationProps := &CfnConfiguredTableAssociationProps{
//   	ConfiguredTableIdentifier: jsii.String("configuredTableIdentifier"),
//   	MembershipIdentifier: jsii.String("membershipIdentifier"),
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnConfiguredTableAssociationProps struct {
	// A unique identifier for the configured table to be associated to.
	//
	// Currently accepts a configured table ID.
	ConfiguredTableIdentifier *string `field:"required" json:"configuredTableIdentifier" yaml:"configuredTableIdentifier"`
	// The unique ID for the membership this configured table association belongs to.
	MembershipIdentifier *string `field:"required" json:"membershipIdentifier" yaml:"membershipIdentifier"`
	// The name of the configured table association, in lowercase.
	//
	// The table is identified by this name when running protected queries against the underlying data.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The service will assume this role to access catalog metadata and query the table.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// A description of the configured table association.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An optional label that you can assign to a resource when you create it.
	//
	// Each tag consists of a key and an optional value, both of which you define. When you use tagging, you can also use tag-based access control in IAM policies to control access to this resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

