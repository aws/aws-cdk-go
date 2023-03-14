package awsglue


// The AWS Lake Formation principal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataLakePrincipalProperty := &dataLakePrincipalProperty{
//   	dataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   }
//
type CfnDatabase_DataLakePrincipalProperty struct {
	// An identifier for the AWS Lake Formation principal.
	DataLakePrincipalIdentifier *string `field:"optional" json:"dataLakePrincipalIdentifier" yaml:"dataLakePrincipalIdentifier"`
}

