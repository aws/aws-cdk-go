package awscognito


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routingTypeProperty := &RoutingTypeProperty{
//   	Failover: &FailoverTypeProperty{
//   		PrimaryRoute53HealthCheckId: jsii.String("primaryRoute53HealthCheckId"),
//   		SecondaryRegion: jsii.String("secondaryRegion"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpooldomain-routingtype.html
//
type CfnUserPoolDomain_RoutingTypeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpooldomain-routingtype.html#cfn-cognito-userpooldomain-routingtype-failover
	//
	Failover interface{} `field:"optional" json:"failover" yaml:"failover"`
}

