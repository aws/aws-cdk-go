package awsroute53

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Construction properties for a CnameRecord.
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
//   schema := appsync.NewSchemaFile(&SchemaProps{
//   	FilePath: jsii.String("mySchemaFile"),
//   })
//   api := appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
//   	Name: jsii.String("myApi"),
//   	Schema: Schema,
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
//   	DomainName: api.appSyncDomainName,
//   })
//
type CnameRecordProps struct {
	// The hosted zone in which to define the new record.
	Zone IHostedZone `field:"required" json:"zone" yaml:"zone"`
	// A comment to add on the record.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Whether to delete the same record set in the hosted zone if it already exists (dangerous!).
	//
	// This allows to deploy a new record set while minimizing the downtime because the
	// new record set will be created immediately after the existing one is deleted. It
	// also avoids "manual" actions to delete existing record sets.
	//
	// > **N.B.:** this feature is dangerous, use with caution! It can only be used safely when
	// > `deleteExisting` is set to `true` as soon as the resource is added to the stack. Changing
	// > an existing Record Set's `deleteExisting` property from `false -> true` after deployment
	// > will delete the record!
	DeleteExisting *bool `field:"optional" json:"deleteExisting" yaml:"deleteExisting"`
	// The geographical origin for this record to return DNS records based on the user's location.
	GeoLocation GeoLocation `field:"optional" json:"geoLocation" yaml:"geoLocation"`
	// The subdomain name for this record. This should be relative to the zone root name.
	//
	// For example, if you want to create a record for acme.example.com, specify
	// "acme".
	//
	// You can also specify the fully qualified domain name which terminates with a
	// ".". For example, "acme.example.com.".
	RecordName *string `field:"optional" json:"recordName" yaml:"recordName"`
	// The resource record cache time to live (TTL).
	Ttl awscdk.Duration `field:"optional" json:"ttl" yaml:"ttl"`
	// The domain name of the target that this record should point to.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
}

