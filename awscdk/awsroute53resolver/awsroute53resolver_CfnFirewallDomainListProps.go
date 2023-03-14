package awsroute53resolver

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnFirewallDomainList`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFirewallDomainListProps := &cfnFirewallDomainListProps{
//   	domainFileUrl: jsii.String("domainFileUrl"),
//   	domains: []*string{
//   		jsii.String("domains"),
//   	},
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFirewallDomainListProps struct {
	// The fully qualified URL or URI of the file stored in Amazon Simple Storage Service (Amazon S3) that contains the list of domains to import.
	//
	// The file must be in an S3 bucket that's in the same Region as your DNS Firewall. The file must be a text file and must contain a single domain per line.
	DomainFileUrl *string `field:"optional" json:"domainFileUrl" yaml:"domainFileUrl"`
	// A list of the domain lists that you have defined.
	Domains *[]*string `field:"optional" json:"domains" yaml:"domains"`
	// The name of the domain list.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of the tag keys and values that you want to associate with the domain list.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

