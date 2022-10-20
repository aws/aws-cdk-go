package awscertificatemanager


// Properties for your certificate.
//
// Example:
//   import cloudwatch "github.com/aws/aws-cdk-go/awscdk"
//
//   var myHostedZone hostedZone
//
//   certificate := acm.NewCertificate(this, jsii.String("Certificate"), &certificateProps{
//   	domainName: jsii.String("hello.example.com"),
//   	validation: acm.certificateValidation.fromDns(myHostedZone),
//   })
//   certificate.metricDaysToExpiry().createAlarm(this, jsii.String("Alarm"), &createAlarmOptions{
//   	comparisonOperator: cloudwatch.comparisonOperator_LESS_THAN_THRESHOLD,
//   	evaluationPeriods: jsii.Number(1),
//   	threshold: jsii.Number(45),
//   })
//
// Experimental.
type CertificateProps struct {
	// Fully-qualified domain name to request a certificate for.
	//
	// May contain wildcards, such as ``*.domain.com``.
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Alternative domain names on your certificate.
	//
	// Use this to register alternative domain names that represent the same site.
	// Experimental.
	SubjectAlternativeNames *[]*string `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
	// How to validate this certificate.
	// Experimental.
	Validation CertificateValidation `field:"optional" json:"validation" yaml:"validation"`
	// What validation domain to use for every requested domain.
	//
	// Has to be a superdomain of the requested domain.
	// Deprecated: use `validation` instead.
	ValidationDomains *map[string]*string `field:"optional" json:"validationDomains" yaml:"validationDomains"`
	// Validation method used to assert domain ownership.
	// Deprecated: use `validation` instead.
	ValidationMethod ValidationMethod `field:"optional" json:"validationMethod" yaml:"validationMethod"`
}

