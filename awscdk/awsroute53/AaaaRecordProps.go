package awsroute53

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Construction properties for a AaaaRecord.
//
// Example:
//   import cloudfront "github.com/aws/aws-cdk-go/awscdk"
//
//   var myZone hostedZone
//   var distribution cloudFrontWebDistribution
//
//   route53.NewAaaaRecord(this, jsii.String("Alias"), &AaaaRecordProps{
//   	Zone: myZone,
//   	Target: route53.RecordTarget_FromAlias(targets.NewCloudFrontTarget(distribution)),
//   })
//
type AaaaRecordProps struct {
	// The hosted zone in which to define the new record.
	Zone IHostedZone `field:"required" json:"zone" yaml:"zone"`
	// A comment to add on the record.
	// Default: no comment.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Whether to delete the same record set in the hosted zone if it already exists (dangerous!).
	//
	// This allows to deploy a new record set while minimizing the downtime because the
	// new record set will be created immediately after the existing one is deleted. It
	// also avoids "manual" actions to delete existing record sets.
	//
	// > **N.B.:** this feature is dangerous, use with caution! It can only be used safely when
	// > `deleteExisting` is set to `true` as soon as the resource is added to the stack. Changing
	// > an existing Record Set's `deleteExisting` property from `false -> true` after deployment
	// > will delete the record!
	// Default: false.
	//
	DeleteExisting *bool `field:"optional" json:"deleteExisting" yaml:"deleteExisting"`
	// The geographical origin for this record to return DNS records based on the user's location.
	GeoLocation GeoLocation `field:"optional" json:"geoLocation" yaml:"geoLocation"`
	// Whether to return multiple values, such as IP addresses for your web servers, in response to DNS queries.
	// Default: false.
	//
	MultiValueAnswer *bool `field:"optional" json:"multiValueAnswer" yaml:"multiValueAnswer"`
	// The subdomain name for this record. This should be relative to the zone root name.
	//
	// For example, if you want to create a record for acme.example.com, specify
	// "acme".
	//
	// You can also specify the fully qualified domain name which terminates with a
	// ".". For example, "acme.example.com.".
	// Default: zone root.
	//
	RecordName *string `field:"optional" json:"recordName" yaml:"recordName"`
	// The Amazon EC2 Region where you created the resource that this resource record set refers to.
	//
	// The resource typically is an AWS resource, such as an EC2 instance or an ELB load balancer,
	// and is referred to by an IP address or a DNS domain name, depending on the record type.
	//
	// When Amazon Route 53 receives a DNS query for a domain name and type for which you have created latency resource record sets,
	// Route 53 selects the latency resource record set that has the lowest latency between the end user and the associated Amazon EC2 Region.
	// Route 53 then returns the value that is associated with the selected resource record set.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordset.html#cfn-route53-recordset-region
	//
	// Default: - Do not set latency based routing.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// A string used to distinguish between different records with the same combination of DNS name and type.
	//
	// It can only be set when either weight or geoLocation is defined.
	//
	// This parameter must be between 1 and 128 characters in length.
	// Default: - Auto generated string.
	//
	SetIdentifier *string `field:"optional" json:"setIdentifier" yaml:"setIdentifier"`
	// The resource record cache time to live (TTL).
	// Default: Duration.minutes(30)
	//
	Ttl awscdk.Duration `field:"optional" json:"ttl" yaml:"ttl"`
	// Among resource record sets that have the same combination of DNS name and type, a value that determines the proportion of DNS queries that Amazon Route 53 responds to using the current resource record set.
	//
	// Route 53 calculates the sum of the weights for the resource record sets that have the same combination of DNS name and type.
	// Route 53 then responds to queries based on the ratio of a resource's weight to the total.
	//
	// This value can be a number between 0 and 255.
	// See: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy-weighted.html
	//
	// Default: - Do not set weighted routing.
	//
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
	// The target.
	Target RecordTarget `field:"required" json:"target" yaml:"target"`
}

