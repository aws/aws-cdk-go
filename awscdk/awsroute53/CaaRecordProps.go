package awsroute53

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Construction properties for a CaaRecord.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var hostedZone hostedZone
//
//   caaRecordProps := &CaaRecordProps{
//   	Values: []caaRecordValue{
//   		&caaRecordValue{
//   			Flag: jsii.Number(123),
//   			Tag: awscdk.Aws_route53.CaaTag_ISSUE,
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Zone: hostedZone,
//
//   	// the properties below are optional
//   	Comment: jsii.String("comment"),
//   	DeleteExisting: jsii.Boolean(false),
//   	RecordName: jsii.String("recordName"),
//   	Ttl: cdk.Duration_Minutes(jsii.Number(30)),
//   }
//
type CaaRecordProps struct {
	// The hosted zone in which to define the new record.
	Zone IHostedZone `field:"required" json:"zone" yaml:"zone"`
	// A comment to add on the record.
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
	DeleteExisting *bool `field:"optional" json:"deleteExisting" yaml:"deleteExisting"`
	// The domain name for this record.
	RecordName *string `field:"optional" json:"recordName" yaml:"recordName"`
	// The resource record cache time to live (TTL).
	Ttl awscdk.Duration `field:"optional" json:"ttl" yaml:"ttl"`
	// The values.
	Values *[]*CaaRecordValue `field:"required" json:"values" yaml:"values"`
}

