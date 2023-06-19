package awsroute53


// Reference to a hosted zone.
//
// Example:
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
// Experimental.
type HostedZoneAttributes struct {
	// Identifier of the hosted zone.
	// Experimental.
	HostedZoneId *string `field:"required" json:"hostedZoneId" yaml:"hostedZoneId"`
	// Name of the hosted zone.
	// Experimental.
	ZoneName *string `field:"required" json:"zoneName" yaml:"zoneName"`
}

