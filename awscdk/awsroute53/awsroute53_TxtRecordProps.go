package awsroute53

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Construction properties for a TxtRecord.
//
// Example:
//   var myZone hostedZone
//
//
//   route53.NewTxtRecord(this, jsii.String("TXTRecord"), &txtRecordProps{
//   	zone: myZone,
//   	recordName: jsii.String("_foo"),
//   	 // If the name ends with a ".", it will be used as-is;
//   	// if it ends with a "." followed by the zone name, a trailing "." will be added automatically;
//   	// otherwise, a ".", the zone name, and a trailing "." will be added automatically.
//   	// Defaults to zone root if not specified.
//   	values: []*string{
//   		jsii.String("Bar!"),
//   		jsii.String("Baz?"),
//   	},
//   	ttl: awscdk.Duration.minutes(jsii.Number(90)),
//   })
//
// Experimental.
type TxtRecordProps struct {
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
	// The text values.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

