package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
)

// Domain name configuration for AppSync.
//
// Example:
//   import acm "github.com/aws/aws-cdk-go/awscdk"
//   import route53 "github.com/aws/aws-cdk-go/awscdk"
//
//   // hosted zone and route53 features
//   var hostedZoneId string
//   zoneName := "example.com"
//
//
//   myDomainName := "api.example.com"
//   certificate := acm.NewCertificate(this, jsii.String("cert"), &certificateProps{
//   	domainName: myDomainName,
//   })
//   schema := appsync.NewSchemaFile(&schemaProps{
//   	filePath: jsii.String("mySchemaFile"),
//   })
//   api := appsync.NewGraphqlApi(this, jsii.String("api"), &graphqlApiProps{
//   	name: jsii.String("myApi"),
//   	schema: schema,
//   	domainName: &domainOptions{
//   		certificate: certificate,
//   		domainName: myDomainName,
//   	},
//   })
//
//   // hosted zone for adding appsync domain
//   zone := route53.hostedZone.fromHostedZoneAttributes(this, jsii.String("HostedZone"), &hostedZoneAttributes{
//   	hostedZoneId: jsii.String(hostedZoneId),
//   	zoneName: jsii.String(zoneName),
//   })
//
//   // create a cname to the appsync domain. will map to something like xxxx.cloudfront.net
//   // create a cname to the appsync domain. will map to something like xxxx.cloudfront.net
//   route53.NewCnameRecord(this, jsii.String("CnameApiRecord"), &cnameRecordProps{
//   	recordName: jsii.String("api"),
//   	zone: zone,
//   	domainName: api.appSyncDomainName,
//   })
//
type DomainOptions struct {
	// The certificate to use with the domain name.
	Certificate awscertificatemanager.ICertificate `field:"required" json:"certificate" yaml:"certificate"`
	// The actual domain name.
	//
	// For example, `api.example.com`.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
}

