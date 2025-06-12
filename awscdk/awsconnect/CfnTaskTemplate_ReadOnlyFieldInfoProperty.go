package awsconnect


// Indicates a field that is read-only to an agent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   readOnlyFieldInfoProperty := &ReadOnlyFieldInfoProperty{
//   	Id: &FieldIdentifierProperty{
//   		Name: jsii.String("name"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-tasktemplate-readonlyfieldinfo.html
//
type CfnTaskTemplate_ReadOnlyFieldInfoProperty struct {
	// Identifier of the read-only field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-tasktemplate-readonlyfieldinfo.html#cfn-connect-tasktemplate-readonlyfieldinfo-id
	//
	Id interface{} `field:"required" json:"id" yaml:"id"`
}

