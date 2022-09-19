// The CDK Construct Library for AWS::Amplify
package awscdkamplifyalpha


// Options to add a domain to an application.
//
// Example:
//   var amplifyApp app
//   var main branch
//   var dev branch
//
//
//   domain := amplifyApp.addDomain(jsii.String("example.com"), &domainOptions{
//   	enableAutoSubdomain: jsii.Boolean(true),
//   	 // in case subdomains should be auto registered for branches
//   	autoSubdomainCreationPatterns: []*string{
//   		jsii.String("*"),
//   		jsii.String("pr*"),
//   	},
//   })
//   domain.mapRoot(main) // map main branch to domain root
//   domain.mapSubDomain(main, jsii.String("www"))
//   domain.mapSubDomain(dev)
//
// Experimental.
type DomainOptions struct {
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
}

