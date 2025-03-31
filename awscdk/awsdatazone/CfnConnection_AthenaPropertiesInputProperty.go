package awsdatazone


// The Amazon Athena properties of a connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   athenaPropertiesInputProperty := &AthenaPropertiesInputProperty{
//   	WorkgroupName: jsii.String("workgroupName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-athenapropertiesinput.html
//
type CfnConnection_AthenaPropertiesInputProperty struct {
	// The Amazon Athena workgroup name of a connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-athenapropertiesinput.html#cfn-datazone-connection-athenapropertiesinput-workgroupname
	//
	WorkgroupName *string `field:"required" json:"workgroupName" yaml:"workgroupName"`
}

