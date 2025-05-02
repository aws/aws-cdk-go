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
	// The Border Gateway Protocol (BGP) Autonomous System Number (ASN) for the appliance.
	//
	// Valid values are from 1 to 4294967295. We recommend using a private ASN in the 64512–65534 (16-bit ASN) or 4200000000–4294967294 (32-bit ASN) range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserver.html#cfn-ec2-routeserver-amazonsideasn
	//
	AmazonSideAsn *float64 `field:"required" json:"amazonSideAsn" yaml:"amazonSideAsn"`
	// Indicates whether routes should be persisted after all BGP sessions are terminated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserver.html#cfn-ec2-routeserver-persistroutes
	//
	PersistRoutes *string `field:"optional" json:"persistRoutes" yaml:"persistRoutes"`
	// The number of minutes a route server will wait after BGP is re-established to unpersist the routes in the FIB and RIB.
	//
	// Value must be in the range of 1-5. The default value is 1. Only valid if `persistRoutesState` is 'enabled'.
	//
	// If you set the duration to 1 minute, then when your network appliance re-establishes BGP with route server, it has 1 minute to relearn it's adjacent network and advertise those routes to route server before route server resumes normal functionality. In most cases, 1 minute is probably sufficient. If, however, you have concerns that your BGP network may not be capable of fully re-establishing and re-learning everything in 1 minute, you can increase the duration up to 5 minutes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserver.html#cfn-ec2-routeserver-persistroutesduration
	//
	PersistRoutesDuration *float64 `field:"optional" json:"persistRoutesDuration" yaml:"persistRoutesDuration"`
	// Indicates whether SNS notifications are enabled for the route server.
	//
	// Enabling SNS notifications persists BGP status changes to an SNS topic provisioned by AWS .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserver.html#cfn-ec2-routeserver-snsnotificationsenabled
	//
	SnsNotificationsEnabled interface{} `field:"optional" json:"snsNotificationsEnabled" yaml:"snsNotificationsEnabled"`
	// Any tags assigned to the route server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserver.html#cfn-ec2-routeserver-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

