package awscdkamplifyalpha


// Options to add a domain to an application.
//
// Example:
//   var amplifyApp app
//   var main branch
//   var dev branch
//
//
//   domain := amplifyApp.AddDomain(jsii.String("example.com"), &DomainOptions{
//   	EnableAutoSubdomain: jsii.Boolean(true),
//   	 // in case subdomains should be auto registered for branches
//   	AutoSubdomainCreationPatterns: []*string{
//   		jsii.String("*"),
//   		jsii.String("pr*"),
//   	},
//   })
//   domain.MapRoot(main) // map main branch to domain root
//   domain.MapSubDomain(main, jsii.String("www"))
//   domain.MapSubDomain(dev)
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

