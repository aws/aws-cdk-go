package awsroute53


// Reference to a private hosted zone.
//
// Example:
//   privateZoneFromAttributes := route53.PrivateHostedZone_FromPrivateHostedZoneAttributes(this, jsii.String("MyPrivateZone"), &PrivateHostedZoneAttributes{
//   	ZoneName: jsii.String("example.local"),
//   	HostedZoneId: jsii.String("ZOJJZC49E0EPZ"),
//   })
//
//   // Does not know zoneName
//   privateZoneFromId := route53.PrivateHostedZone_FromPrivateHostedZoneId(this, jsii.String("MyPrivateZone"), jsii.String("ZOJJZC49E0EPZ"))
//
type PrivateHostedZoneAttributes struct {
	// Identifier of the hosted zone.
	HostedZoneId *string `field:"required" json:"hostedZoneId" yaml:"hostedZoneId"`
	// Name of the hosted zone.
	ZoneName *string `field:"required" json:"zoneName" yaml:"zoneName"`
}

