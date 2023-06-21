package awsroute53


// Reference to a public hosted zone.
//
// Example:
//   zoneFromAttributes := route53.PublicHostedZone_FromPublicHostedZoneAttributes(this, jsii.String("MyZone"), &PublicHostedZoneAttributes{
//   	ZoneName: jsii.String("example.com"),
//   	HostedZoneId: jsii.String("ZOJJZC49E0EPZ"),
//   })
//
//   // Does not know zoneName
//   zoneFromId := route53.PublicHostedZone_FromPublicHostedZoneId(this, jsii.String("MyZone"), jsii.String("ZOJJZC49E0EPZ"))
//
type PublicHostedZoneAttributes struct {
	// Identifier of the hosted zone.
	HostedZoneId *string `field:"required" json:"hostedZoneId" yaml:"hostedZoneId"`
	// Name of the hosted zone.
	ZoneName *string `field:"required" json:"zoneName" yaml:"zoneName"`
}

