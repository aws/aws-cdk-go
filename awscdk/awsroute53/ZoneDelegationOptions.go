package awsroute53

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options available when creating a delegation relationship from one PublicHostedZone to another.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   zoneDelegationOptions := &ZoneDelegationOptions{
//   	Comment: jsii.String("comment"),
//   	Ttl: cdk.Duration_Minutes(jsii.Number(30)),
//   }
//
type ZoneDelegationOptions struct {
	// A comment to add on the DNS record created to incorporate the delegation.
	// Default: none.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The TTL (Time To Live) of the DNS delegation record in DNS caches.
	// Default: 172800.
	//
	Ttl awscdk.Duration `field:"optional" json:"ttl" yaml:"ttl"`
}

