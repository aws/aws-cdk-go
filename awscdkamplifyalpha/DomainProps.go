package awscdkamplifyalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for a Domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import amplify_alpha "github.com/aws/aws-cdk-go/awscdkamplifyalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var app App
//   var branch Branch
//   var certificate Certificate
//   var role Role
//
//   domainProps := &DomainProps{
//   	App: app,
//
//   	// the properties below are optional
//   	AutoSubdomainCreationPatterns: []*string{
//   		jsii.String("autoSubdomainCreationPatterns"),
//   	},
//   	AutoSubDomainIamRole: role,
//   	CustomCertificate: certificate,
//   	DomainName: jsii.String("domainName"),
//   	EnableAutoSubdomain: jsii.Boolean(false),
//   	SubDomains: []SubDomain{
//   		&SubDomain{
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
	// Default: - all repository branches ['*', 'pr*'].
	//
	// Experimental.
	AutoSubdomainCreationPatterns *[]*string `field:"optional" json:"autoSubdomainCreationPatterns" yaml:"autoSubdomainCreationPatterns"`
	// The type of SSL/TLS certificate to use for your custom domain.
	// Default: - Amplify uses the default certificate that it provisions and manages for you.
	//
	// Experimental.
	CustomCertificate awscertificatemanager.ICertificate `field:"optional" json:"customCertificate" yaml:"customCertificate"`
	// The name of the domain.
	// Default: - the construct's id.
	//
	// Experimental.
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// Automatically create subdomains for connected branches.
	// Default: false.
	//
	// Experimental.
	EnableAutoSubdomain *bool `field:"optional" json:"enableAutoSubdomain" yaml:"enableAutoSubdomain"`
	// Subdomains.
	// Default: - use `addSubDomain()` to add subdomains.
	//
	// Experimental.
	SubDomains *[]*SubDomain `field:"optional" json:"subDomains" yaml:"subDomains"`
	// The application to which the domain must be connected.
	// Experimental.
	App IApp `field:"required" json:"app" yaml:"app"`
	// The IAM role with access to Route53 when using enableAutoSubdomain.
	// Default: - the IAM role from App.grantPrincipal
	//
	// Experimental.
	AutoSubDomainIamRole awsiam.IRole `field:"optional" json:"autoSubDomainIamRole" yaml:"autoSubDomainIamRole"`
}

