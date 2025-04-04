package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRouteServer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRouteServerProps := &CfnRouteServerProps{
//   	AmazonSideAsn: jsii.Number(123),
//
//   	// the properties below are optional
//   	PersistRoutes: jsii.String("persistRoutes"),
//   	PersistRoutesDuration: jsii.Number(123),
//   	SnsNotificationsEnabled: jsii.Boolean(false),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserver.html
//
type CfnRouteServerProps struct {
	// The Amazon-side ASN of the Route Server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserver.html#cfn-ec2-routeserver-amazonsideasn
	//
	AmazonSideAsn *float64 `field:"required" json:"amazonSideAsn" yaml:"amazonSideAsn"`
	// Whether to enable persistent routes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserver.html#cfn-ec2-routeserver-persistroutes
	//
	PersistRoutes *string `field:"optional" json:"persistRoutes" yaml:"persistRoutes"`
	// The duration of persistent routes in minutes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserver.html#cfn-ec2-routeserver-persistroutesduration
	//
	PersistRoutesDuration *float64 `field:"optional" json:"persistRoutesDuration" yaml:"persistRoutesDuration"`
	// Whether to enable SNS notifications.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserver.html#cfn-ec2-routeserver-snsnotificationsenabled
	//
	SnsNotificationsEnabled interface{} `field:"optional" json:"snsNotificationsEnabled" yaml:"snsNotificationsEnabled"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserver.html#cfn-ec2-routeserver-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

