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
//   domainProps := &domainProps{
//   	app: app,
//
//   	// the properties below are optional
//   	autoSubdomainCreationPatterns: []*string{
//   		jsii.String("autoSubdomainCreationPatterns"),
//   	},
//   	autoSubDomainIamRole: role,
//   	domainName: jsii.String("domainName"),
//   	enableAutoSubdomain: jsii.Boolean(false),
//   	subDomains: []subDomain{
//   		&subDomain{
//   			branch: branch,
//
//   			// the properties below are optional
//   			prefix: jsii.String("prefix"),
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

