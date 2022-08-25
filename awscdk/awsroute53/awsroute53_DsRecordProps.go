package awsroute53

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Construction properties for a DSRecord.
//
// Example:
//   var myZone hostedZone
//
//
//   route53.NewDsRecord(this, jsii.String("DSRecord"), &dsRecordProps{
//   	zone: myZone,
//   	recordName: jsii.String("foo"),
//   	values: []*string{
//   		jsii.String("12345 3 1 123456789abcdef67890123456789abcdef67890"),
//   	},
//   	ttl: awscdk.Duration.minutes(jsii.Number(90)),
//   })
//
type DsRecordProps struct {
	// The hosted zone in which to define the new record.
	Zone IHostedZone `field:"required" json:"zone" yaml:"zone"`
	// A comment to add on the record.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Whether to delete the same record set in the hosted zone if it already exists.
	//
	// This allows to deploy a new record set while minimizing the downtime because the
	// new record set will be created immediately after the existing one is deleted. It
	// also avoids "manual" actions to delete existing record sets.
	DeleteExisting *bool `field:"optional" json:"deleteExisting" yaml:"deleteExisting"`
	// The domain name for this record.
	RecordName *string `field:"optional" json:"recordName" yaml:"recordName"`
	// The resource record cache time to live (TTL).
	Ttl awscdk.Duration `field:"optional" json:"ttl" yaml:"ttl"`
	// The DS values.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

