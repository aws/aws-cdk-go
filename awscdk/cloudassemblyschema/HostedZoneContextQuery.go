package cloudassemblyschema


// Query to hosted zone context provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hostedZoneContextQuery := &HostedZoneContextQuery{
//   	Account: jsii.String("account"),
//   	DomainName: jsii.String("domainName"),
//   	Region: jsii.String("region"),
//
//   	// the properties below are optional
//   	LookupRoleArn: jsii.String("lookupRoleArn"),
//   	PrivateZone: jsii.Boolean(false),
//   	VpcId: jsii.String("vpcId"),
//   }
//
type HostedZoneContextQuery struct {
	// Query account.
	Account *string `field:"required" json:"account" yaml:"account"`
	// The domain name e.g. example.com to lookup.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Query region.
	Region *string `field:"required" json:"region" yaml:"region"`
	// The ARN of the role that should be used to look up the missing values.
	// Default: - None.
	//
	LookupRoleArn *string `field:"optional" json:"lookupRoleArn" yaml:"lookupRoleArn"`
	// True if the zone you want to find is a private hosted zone.
	// Default: false.
	//
	PrivateZone *bool `field:"optional" json:"privateZone" yaml:"privateZone"`
	// The VPC ID to that the private zone must be associated with.
	//
	// If you provide VPC ID and privateZone is false, this will return no results
	// and raise an error.
	// Default: - Required if privateZone=true.
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

