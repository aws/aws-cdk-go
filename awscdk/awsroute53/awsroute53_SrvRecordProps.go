package awsroute53

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Construction properties for a SrvRecord.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//   var hostedZone hostedZone
//
//   srvRecordProps := &srvRecordProps{
//   	values: []srvRecordValue{
//   		&srvRecordValue{
//   			hostName: jsii.String("hostName"),
//   			port: jsii.Number(123),
//   			priority: jsii.Number(123),
//   			weight: jsii.Number(123),
//   		},
//   	},
//   	zone: hostedZone,
//
//   	// the properties below are optional
//   	comment: jsii.String("comment"),
//   	recordName: jsii.String("recordName"),
//   	ttl: duration,
//   }
//
// Experimental.
type SrvRecordProps struct {
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
	// The values.
	// Experimental.
	Values *[]*SrvRecordValue `field:"required" json:"values" yaml:"values"`
}

