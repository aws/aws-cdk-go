package awsec2


// Specifies input parameter values for an SSM document in AWS Systems Manager .
//
// `AssociationParameter` is a property of the [Amazon EC2 Instance SsmAssociation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-ssmassociations.html) property.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   associationParameterProperty := &associationParameterProperty{
//   	key: jsii.String("key"),
//   	value: []*string{
//   		jsii.String("value"),
//   	},
//   }
//
type CfnInstance_AssociationParameterProperty struct {
	// The name of an input parameter that is in the associated SSM document.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value of an input parameter.
	Value *[]*string `field:"required" json:"value" yaml:"value"`
}

