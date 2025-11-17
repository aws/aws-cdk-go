package awsroute53


// Options for the delegation permissions granted.
//
// Example:
//   var betaCrossAccountRole Role
//
//   var prodCrossAccountRole Role
//   parentZone := route53.NewPublicHostedZone(this, jsii.String("HostedZone"), &PublicHostedZoneProps{
//   	ZoneName: jsii.String("someexample.com"),
//   })
//   parentZone.GrantDelegation(betaCrossAccountRole, &GrantDelegationOptions{
//   	DelegatedZoneNames: []*string{
//   		jsii.String("beta.someexample.com"),
//   	},
//   })
//   parentZone.GrantDelegation(prodCrossAccountRole, &GrantDelegationOptions{
//   	DelegatedZoneNames: []*string{
//   		jsii.String("prod.someexample.com"),
//   	},
//   })
//
type GrantDelegationOptions struct {
	// List of hosted zone names to allow delegation to in the grant permissions.
	//
	// If the delegated zone name contains an unresolved token,
	// it must resolve to a zone name that satisfies the requirements according to the documentation:
	// https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/specifying-conditions-route53.html#route53_rrset_conditionkeys_normalization
	//
	// > All letters must be lowercase.
	// > The DNS name must be without the trailing dot.
	// > Characters other than a–z, 0–9, - (hyphen), _ (underscore), and . (period, as a delimiter between labels) must use escape codes in the format \three-digit octal code. For example, \052 is the octal code for character *.
	// Default: the grant allows delegation to any hosted zone.
	//
	DelegatedZoneNames *[]*string `field:"optional" json:"delegatedZoneNames" yaml:"delegatedZoneNames"`
}

