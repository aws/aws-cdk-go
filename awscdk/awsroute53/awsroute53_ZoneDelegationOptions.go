package awsroute53

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Options available when creating a delegation relationship from one PublicHostedZone to another.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//
//   zoneDelegationOptions := &zoneDelegationOptions{
//   	comment: jsii.String("comment"),
//   	ttl: duration,
//   }
//
// Experimental.
type ZoneDelegationOptions struct {
	// A comment to add on the DNS record created to incorporate the delegation.
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The TTL (Time To Live) of the DNS delegation record in DNS caches.
	// Experimental.
	Ttl awscdk.Duration `field:"optional" json:"ttl" yaml:"ttl"`
}

