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
//   bucketWebsite := s3.NewBucket(this, jsii.String("BucketWebsite"), &bucketProps{
//   	bucketName: []*string{
//   		recordName,
//   		domainName,
//   	}.join(jsii.String(".")),
//   	 // www.example.com
//   	publicReadAccess: jsii.Boolean(true),
//   	websiteIndexDocument: jsii.String("index.html"),
//   })
//
//   zone := route53.hostedZone.fromLookup(this, jsii.String("Zone"), &hostedZoneProviderProps{
//   	domainName: jsii.String(domainName),
//   }) // example.com
//
//    // example.com
//   route53.NewARecord(this, jsii.String("AliasRecord"), &aRecordProps{
//   	zone: zone,
//   	recordName: jsii.String(recordName),
//   	 // www
//   	target: route53.recordTarget.fromAlias(targets.NewBucketWebsiteTarget(bucketWebsite)),
//   })
//
type HostedZoneProviderProps struct {
	// The zone domain e.g. example.com.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Whether the zone that is being looked up is a private hosted zone.
	PrivateZone *bool `field:"optional" json:"privateZone" yaml:"privateZone"`
	// Specifies the ID of the VPC associated with a private hosted zone.
	//
	// If a VPC ID is provided and privateZone is false, no results will be returned
	// and an error will be raised.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

