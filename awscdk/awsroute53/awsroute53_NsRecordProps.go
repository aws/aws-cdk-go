package awsroute53

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Construction properties for a NSRecord.
//
// Example:
//   var myZone hostedZone
//
//
//   route53.NewNsRecord(this, jsii.String("NSRecord"), &nsRecordProps{
//   	zone: myZone,
//   	recordName: jsii.String("foo"),
//   	values: []*string{
//   		jsii.String("ns-1.awsdns.co.uk."),
//   		jsii.String("ns-2.awsdns.com."),
//   	},
//   	ttl: awscdk.Duration.minutes(jsii.Number(90)),
//   })
//
// Experimental.
type NsRecordProps struct {
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
	// The NS values.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

