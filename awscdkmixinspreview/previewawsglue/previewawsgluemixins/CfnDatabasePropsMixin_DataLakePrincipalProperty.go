package previewawsgluemixins


// The AWS Lake Formation principal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataLakePrincipalProperty := &DataLakePrincipalProperty{
//   	DataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-database-datalakeprincipal.html
//
type CfnDatabasePropsMixin_DataLakePrincipalProperty struct {
	// An identifier for the AWS Lake Formation principal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-database-datalakeprincipal.html#cfn-glue-database-datalakeprincipal-datalakeprincipalidentifier
	//
	DataLakePrincipalIdentifier *string `field:"optional" json:"dataLakePrincipalIdentifier" yaml:"dataLakePrincipalIdentifier"`
}

