package mixinsawsamazonmq


// A list of information about the configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   configurationIdProperty := &ConfigurationIdProperty{
//   	Id: jsii.String("id"),
//   	Revision: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-configurationassociation-configurationid.html
//
type CfnConfigurationAssociationPropsMixin_ConfigurationIdProperty struct {
	// Required.
	//
	// The unique ID that Amazon MQ generates for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-configurationassociation-configurationid.html#cfn-amazonmq-configurationassociation-configurationid-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The revision number of the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-configurationassociation-configurationid.html#cfn-amazonmq-configurationassociation-configurationid-revision
	//
	Revision *float64 `field:"optional" json:"revision" yaml:"revision"`
}

