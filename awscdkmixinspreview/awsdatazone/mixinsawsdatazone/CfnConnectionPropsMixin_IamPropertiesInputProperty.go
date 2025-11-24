package mixinsawsdatazone


// The IAM properties of a connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   iamPropertiesInputProperty := &IamPropertiesInputProperty{
//   	GlueLineageSyncEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-iampropertiesinput.html
//
type CfnConnectionPropsMixin_IamPropertiesInputProperty struct {
	// Specifies whether AWS Glue lineage sync is enabled for a connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-iampropertiesinput.html#cfn-datazone-connection-iampropertiesinput-gluelineagesyncenabled
	//
	GlueLineageSyncEnabled interface{} `field:"optional" json:"glueLineageSyncEnabled" yaml:"glueLineageSyncEnabled"`
}

