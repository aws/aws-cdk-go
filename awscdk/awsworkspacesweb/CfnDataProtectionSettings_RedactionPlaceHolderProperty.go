package awsworkspacesweb


// The redaction placeholder that will replace the redacted text in session.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redactionPlaceHolderProperty := &RedactionPlaceHolderProperty{
//   	RedactionPlaceHolderType: jsii.String("redactionPlaceHolderType"),
//
//   	// the properties below are optional
//   	RedactionPlaceHolderText: jsii.String("redactionPlaceHolderText"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-redactionplaceholder.html
//
type CfnDataProtectionSettings_RedactionPlaceHolderProperty struct {
	// The redaction placeholder type that will replace the redacted text in session.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-redactionplaceholder.html#cfn-workspacesweb-dataprotectionsettings-redactionplaceholder-redactionplaceholdertype
	//
	RedactionPlaceHolderType *string `field:"required" json:"redactionPlaceHolderType" yaml:"redactionPlaceHolderType"`
	// The redaction placeholder text that will replace the redacted text in session for the custom text redaction placeholder type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspacesweb-dataprotectionsettings-redactionplaceholder.html#cfn-workspacesweb-dataprotectionsettings-redactionplaceholder-redactionplaceholdertext
	//
	RedactionPlaceHolderText *string `field:"optional" json:"redactionPlaceHolderText" yaml:"redactionPlaceHolderText"`
}

