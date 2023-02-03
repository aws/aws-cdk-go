package awsroute53


// Reference to a hosted zone.
//
// Example:
//   var app app
//
//   stack := awscdk.Newstack(app, jsii.String("Stack"), &stackProps{
//   	crossRegionReferences: jsii.Boolean(true),
//   	env: &environment{
//   		region: jsii.String("us-east-2"),
//   	},
//   })
//
//   patterns.NewHttpsRedirect(this, jsii.String("Redirect"), &httpsRedirectProps{
//   	recordNames: []*string{
//   		jsii.String("foo.example.com"),
//   	},
//   	targetDomain: jsii.String("bar.example.com"),
//   	zone: route53.hostedZone.fromHostedZoneAttributes(this, jsii.String("HostedZone"), &hostedZoneAttributes{
//   		hostedZoneId: jsii.String("ID"),
//   		zoneName: jsii.String("example.com"),
//   	}),
//   })
//
type HostedZoneAttributes struct {
	// Identifier of the hosted zone.
	HostedZoneId *string `field:"required" json:"hostedZoneId" yaml:"hostedZoneId"`
	// Name of the hosted zone.
	ZoneName *string `field:"required" json:"zoneName" yaml:"zoneName"`
}

