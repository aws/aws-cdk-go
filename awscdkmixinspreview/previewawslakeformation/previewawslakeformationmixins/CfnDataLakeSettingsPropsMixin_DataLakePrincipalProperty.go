package previewawslakeformationmixins


// The Lake Formation principal.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-datalakesettings-datalakeprincipal.html
//
type CfnDataLakeSettingsPropsMixin_DataLakePrincipalProperty struct {
	// An identifier for the Lake Formation principal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-datalakesettings-datalakeprincipal.html#cfn-lakeformation-datalakesettings-datalakeprincipal-datalakeprincipalidentifier
	//
	DataLakePrincipalIdentifier *string `field:"optional" json:"dataLakePrincipalIdentifier" yaml:"dataLakePrincipalIdentifier"`
}

