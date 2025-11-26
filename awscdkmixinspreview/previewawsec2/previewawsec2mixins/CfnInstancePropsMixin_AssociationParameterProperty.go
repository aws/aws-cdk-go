package previewawsec2mixins


// Specifies input parameter values for an SSM document in AWS Systems Manager .
//
// `AssociationParameter` is a property of the [SsmAssociation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociation.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   associationParameterProperty := &AssociationParameterProperty{
//   	Key: jsii.String("key"),
//   	Value: []*string{
//   		jsii.String("value"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-associationparameter.html
//
type CfnInstancePropsMixin_AssociationParameterProperty struct {
	// The name of an input parameter that is in the associated SSM document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-associationparameter.html#cfn-ec2-instance-associationparameter-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of an input parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-associationparameter.html#cfn-ec2-instance-associationparameter-value
	//
	Value *[]*string `field:"optional" json:"value" yaml:"value"`
}

