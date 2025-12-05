package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lockVersionProperty := &LockVersionProperty{
//   	Attribute: jsii.String("attribute"),
//   	DataTable: jsii.String("dataTable"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatableattribute-lockversion.html
//
type CfnDataTableAttribute_LockVersionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatableattribute-lockversion.html#cfn-connect-datatableattribute-lockversion-attribute
	//
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatableattribute-lockversion.html#cfn-connect-datatableattribute-lockversion-datatable
	//
	DataTable *string `field:"optional" json:"dataTable" yaml:"dataTable"`
}

