package awsappsync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappsync"
)

// The Schema for a GraphQL Api.
//
// If no options are configured, schema will be generated
// code-first.
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
//   	Definition: appsync.Definition_FromSchema(schema),
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
type SchemaFile interface {
	ISchema
	// The definition for this schema.
	Definition() *string
	SetDefinition(val *string)
	// Called when the GraphQL Api is initialized to allow this object to bind to the stack.
	Bind(api interfacesawsappsync.IGraphQLApiRef, options *SchemaBindOptions) ISchemaConfig
}

// The jsii proxy struct for SchemaFile
type jsiiProxy_SchemaFile struct {
	jsiiProxy_ISchema
}

func (j *jsiiProxy_SchemaFile) Definition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}


func NewSchemaFile(options *SchemaProps) SchemaFile {
	_init_.Initialize()

	if err := validateNewSchemaFileParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_SchemaFile{}

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.SchemaFile",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewSchemaFile_Override(s SchemaFile, options *SchemaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appsync.SchemaFile",
		[]interface{}{options},
		s,
	)
}

func (j *jsiiProxy_SchemaFile)SetDefinition(val *string) {
	if err := j.validateSetDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"definition",
		val,
	)
}

// Generate a Schema from file.
//
// Returns: `SchemaAsset` with immutable schema defintion.
func SchemaFile_FromAsset(filePath *string) SchemaFile {
	_init_.Initialize()

	if err := validateSchemaFile_FromAssetParameters(filePath); err != nil {
		panic(err)
	}
	var returns SchemaFile

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appsync.SchemaFile",
		"fromAsset",
		[]interface{}{filePath},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SchemaFile) Bind(api interfacesawsappsync.IGraphQLApiRef, options *SchemaBindOptions) ISchemaConfig {
	if err := s.validateBindParameters(api, options); err != nil {
		panic(err)
	}
	var returns ISchemaConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{api, options},
		&returns,
	)

	return returns
}

