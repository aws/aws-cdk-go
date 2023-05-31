package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/awscertificatemanager"
)

// Domain name configuration for AppSync.
//
// Example:
//   import acm "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   // hosted zone and route53 features
//   var hostedZoneId string
//   zoneName := "example.com"
//
//
//   myDomainName := "api.example.com"
//   certificate := acm.NewCertificate(this, jsii.String("cert"), &CertificateProps{
//   	DomainName: myDomainName,
//   })
//   api := appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
//   	Name: jsii.String("myApi"),
//   	DomainName: &DomainOptions{
//   		Certificate: *Certificate,
//   		DomainName: myDomainName,
//   	},
//   })
//
//   // hosted zone for adding appsync domain
//   zone := route53.HostedZone_FromHostedZoneAttributes(this, jsii.String("HostedZone"), &HostedZoneAttributes{
//   	HostedZoneId: jsii.String(HostedZoneId),
//   	ZoneName: jsii.String(ZoneName),
//   })
//
//   // create a cname to the appsync domain. will map to something like xxxx.cloudfront.net
//   // create a cname to the appsync domain. will map to something like xxxx.cloudfront.net
//   route53.NewCnameRecord(this, jsii.String("CnameApiRecord"), &CnameRecordProps{
//   	RecordName: jsii.String("api"),
//   	Zone: Zone,
//   	DomainName: myDomainName,
//   })
//
// Experimental.
type DomainOptions struct {
	// The certificate to use with the domain name.
	// Experimental.
	Certificate awscertificatemanager.ICertificate `field:"required" json:"certificate" yaml:"certificate"`
	// The actual domain name.
	//
	// For example, `api.example.com`.
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
}

