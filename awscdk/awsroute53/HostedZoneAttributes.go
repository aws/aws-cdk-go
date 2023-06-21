package awsroute53


// Reference to a hosted zone.
//
// Example:
//   var app app
//
//   stack := awscdk.Newstack(app, jsii.String("Stack"), &StackProps{
//   	CrossRegionReferences: jsii.Boolean(true),
//   	Env: &Environment{
//   		Region: jsii.String("us-east-2"),
//   	},
//   })
//
//   patterns.NewHttpsRedirect(this, jsii.String("Redirect"), &HttpsRedirectProps{
//   	RecordNames: []*string{
//   		jsii.String("foo.example.com"),
//   	},
//   	TargetDomain: jsii.String("bar.example.com"),
//   	Zone: route53.HostedZone_FromHostedZoneAttributes(this, jsii.String("HostedZone"), &HostedZoneAttributes{
//   		HostedZoneId: jsii.String("ID"),
//   		ZoneName: jsii.String("example.com"),
//   	}),
//   })
//
type HostedZoneAttributes struct {
	// Identifier of the hosted zone.
	HostedZoneId *string `field:"required" json:"hostedZoneId" yaml:"hostedZoneId"`
	// Name of the hosted zone.
	ZoneName *string `field:"required" json:"zoneName" yaml:"zoneName"`
}

