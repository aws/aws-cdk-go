package awsvpclattice

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnAccessLogSubscription`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccessLogSubscriptionProps := &cfnAccessLogSubscriptionProps{
//   	destinationArn: jsii.String("destinationArn"),
//
//   	// the properties below are optional
//   	resourceIdentifier: jsii.String("resourceIdentifier"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAccessLogSubscriptionProps struct {
	// `AWS::VpcLattice::AccessLogSubscription.DestinationArn`.
	DestinationArn *string `field:"required" json:"destinationArn" yaml:"destinationArn"`
	// `AWS::VpcLattice::AccessLogSubscription.ResourceIdentifier`.
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// `AWS::VpcLattice::AccessLogSubscription.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

