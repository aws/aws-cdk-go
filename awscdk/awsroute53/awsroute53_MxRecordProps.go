package awsroute53

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Construction properties for a MxRecord.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var hostedZone hostedZone
//
//   mxRecordProps := &mxRecordProps{
//   	values: []mxRecordValue{
//   		&mxRecordValue{
//   			hostName: jsii.String("hostName"),
//   			priority: jsii.Number(123),
//   		},
//   	},
//   	zone: hostedZone,
//
//   	// the properties below are optional
//   	comment: jsii.String("comment"),
//   	deleteExisting: jsii.Boolean(false),
//   	recordName: jsii.String("recordName"),
//   	ttl: cdk.duration.minutes(jsii.Number(30)),
//   }
//
type MxRecordProps struct {
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
	// The values.
	Values *[]*MxRecordValue `field:"required" json:"values" yaml:"values"`
}

