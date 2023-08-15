package awsroute53


// Zone properties for looking up the Hosted Zone.
//
// Example:
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   recordName := "www"
//   domainName := "example.com"
//
//   bucketWebsite := s3.NewBucket(this, jsii.String("BucketWebsite"), &BucketProps{
//   	BucketName: []*string{
//   		recordName,
//   		domainName,
//   	}.join(jsii.String(".")),
//   	 // www.example.com
//   	PublicReadAccess: jsii.Boolean(true),
//   	WebsiteIndexDocument: jsii.String("index.html"),
//   })
//
//   zone := route53.HostedZone_FromLookup(this, jsii.String("Zone"), &HostedZoneProviderProps{
//   	DomainName: jsii.String(DomainName),
//   }) // example.com
//
//    // example.com
//   route53.NewARecord(this, jsii.String("AliasRecord"), &ARecordProps{
//   	Zone: Zone,
//   	RecordName: jsii.String(RecordName),
//   	 // www
//   	Target: route53.RecordTarget_FromAlias(targets.NewBucketWebsiteTarget(bucketWebsite)),
//   })
//
type HostedZoneProviderProps struct {
	// The zone domain e.g. example.com.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Whether the zone that is being looked up is a private hosted zone.
	// Default: false.
	//
	PrivateZone *bool `field:"optional" json:"privateZone" yaml:"privateZone"`
	// Specifies the ID of the VPC associated with a private hosted zone.
	//
	// If a VPC ID is provided and privateZone is false, no results will be returned
	// and an error will be raised.
	// Default: - No VPC ID.
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

