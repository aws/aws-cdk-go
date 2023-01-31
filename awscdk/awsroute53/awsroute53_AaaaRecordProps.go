package awsroute53

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Construction properties for a AaaaRecord.
//
// Example:
//   import cloudfront "github.com/aws/aws-cdk-go/awscdk"
//
//   var myZone hostedZone
//   var distribution cloudFrontWebDistribution
//
//   route53.NewAaaaRecord(this, jsii.String("Alias"), &aaaaRecordProps{
//   	zone: myZone,
//   	target: route53.recordTarget.fromAlias(targets.NewCloudFrontTarget(distribution)),
//   })
//
// Experimental.
type AaaaRecordProps struct {
	// The hosted zone in which to define the new record.
	// Experimental.
	Zone IHostedZone `field:"required" json:"zone" yaml:"zone"`
	// A comment to add on the record.
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The domain name for this record.
	// Experimental.
	RecordName *string `field:"optional" json:"recordName" yaml:"recordName"`
	// The resource record cache time to live (TTL).
	// Experimental.
	Ttl awscdk.Duration `field:"optional" json:"ttl" yaml:"ttl"`
	// The target.
	// Experimental.
	Target RecordTarget `field:"required" json:"target" yaml:"target"`
}

