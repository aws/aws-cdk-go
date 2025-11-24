package mixinsawsglue


// Specifies the version of a schema.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   schemaVersionProperty := &SchemaVersionProperty{
//   	IsLatest: jsii.Boolean(false),
//   	VersionNumber: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-schema-schemaversion.html
//
type CfnSchemaPropsMixin_SchemaVersionProperty struct {
	// Indicates if this version is the latest version of the schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-schema-schemaversion.html#cfn-glue-schema-schemaversion-islatest
	//
	IsLatest interface{} `field:"optional" json:"isLatest" yaml:"isLatest"`
	// The version number of the schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-schema-schemaversion.html#cfn-glue-schema-schemaversion-versionnumber
	//
	VersionNumber *float64 `field:"optional" json:"versionNumber" yaml:"versionNumber"`
}

