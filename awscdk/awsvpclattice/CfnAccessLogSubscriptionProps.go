package awsvpclattice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAccessLogSubscription`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccessLogSubscriptionProps := &CfnAccessLogSubscriptionProps{
//   	DestinationArn: jsii.String("destinationArn"),
//
//   	// the properties below are optional
//   	ResourceIdentifier: jsii.String("resourceIdentifier"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
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

