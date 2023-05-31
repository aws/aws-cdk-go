package awsamplify

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Properties for a Domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var app app
//   var branch branch
//   var role role
//
//   domainProps := &DomainProps{
//   	App: app,
//
//   	// the properties below are optional
//   	AutoSubdomainCreationPatterns: []*string{
//   		jsii.String("autoSubdomainCreationPatterns"),
//   	},
//   	AutoSubDomainIamRole: role,
//   	DomainName: jsii.String("domainName"),
//   	EnableAutoSubdomain: jsii.Boolean(false),
//   	SubDomains: []subDomain{
//   		&subDomain{
//   			Branch: branch,
//
//   			// the properties below are optional
//   			Prefix: jsii.String("prefix"),
//   		},
//   	},
//   }
//
// Experimental.
type DomainProps struct {
	// Branches which should automatically create subdomains.
	// Experimental.
	AutoSubdomainCreationPatterns *[]*string `field:"optional" json:"autoSubdomainCreationPatterns" yaml:"autoSubdomainCreationPatterns"`
	// The name of the domain.
	// Experimental.
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// Automatically create subdomains for connected branches.
	// Experimental.
	EnableAutoSubdomain *bool `field:"optional" json:"enableAutoSubdomain" yaml:"enableAutoSubdomain"`
	// Subdomains.
	// Experimental.
	SubDomains *[]*SubDomain `field:"optional" json:"subDomains" yaml:"subDomains"`
	// The application to which the domain must be connected.
	// Experimental.
	App IApp `field:"required" json:"app" yaml:"app"`
	// The IAM role with access to Route53 when using enableAutoSubdomain.
	// Experimental.
	AutoSubDomainIamRole awsiam.IRole `field:"optional" json:"autoSubDomainIamRole" yaml:"autoSubDomainIamRole"`
}

