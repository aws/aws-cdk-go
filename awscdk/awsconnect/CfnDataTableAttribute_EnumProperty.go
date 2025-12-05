package awsconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   enumProperty := &EnumProperty{
//   	Strict: jsii.Boolean(false),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatableattribute-enum.html
//
type CfnDataTableAttribute_EnumProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatableattribute-enum.html#cfn-connect-datatableattribute-enum-strict
	//
	Strict interface{} `field:"optional" json:"strict" yaml:"strict"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-datatableattribute-enum.html#cfn-connect-datatableattribute-enum-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

