package awsroute53

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Construction properties for a DSRecord.
//
// Example:
//   var myZone hostedZone
//
//
//   route53.NewDsRecord(this, jsii.String("DSRecord"), &DsRecordProps{
//   	Zone: myZone,
//   	RecordName: jsii.String("foo"),
//   	Values: []*string{
//   		jsii.String("12345 3 1 123456789abcdef67890123456789abcdef67890"),
//   	},
//   	Ttl: awscdk.Duration_Minutes(jsii.Number(90)),
//   })
//
// Experimental.
type DsRecordProps struct {
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
	// The DS values.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

