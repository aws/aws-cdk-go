package awscognito


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   failoverTypeProperty := &FailoverTypeProperty{
//   	PrimaryRoute53HealthCheckId: jsii.String("primaryRoute53HealthCheckId"),
//   	SecondaryRegion: jsii.String("secondaryRegion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpooldomain-failovertype.html
//
type CfnUserPoolDomainPropsMixin_FailoverTypeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpooldomain-failovertype.html#cfn-cognito-userpooldomain-failovertype-primaryroute53healthcheckid
	//
	PrimaryRoute53HealthCheckId *string `field:"optional" json:"primaryRoute53HealthCheckId" yaml:"primaryRoute53HealthCheckId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpooldomain-failovertype.html#cfn-cognito-userpooldomain-failovertype-secondaryregion
	//
	SecondaryRegion *string `field:"optional" json:"secondaryRegion" yaml:"secondaryRegion"`
}

