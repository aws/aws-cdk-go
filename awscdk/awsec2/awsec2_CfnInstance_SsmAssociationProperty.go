package awsec2


// Specifies the SSM document and parameter values in AWS Systems Manager to associate with an instance.
//
// `SsmAssociations` is a property of the [AWS::EC2::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ssmAssociationProperty := &ssmAssociationProperty{
//   	documentName: jsii.String("documentName"),
//
//   	// the properties below are optional
//   	associationParameters: []interface{}{
//   		&associationParameterProperty{
//   			key: jsii.String("key"),
//   			value: []*string{
//   				jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
type CfnInstance_SsmAssociationProperty struct {
	// The name of an SSM document to associate with the instance.
	DocumentName *string `field:"required" json:"documentName" yaml:"documentName"`
	// The input parameter values to use with the associated SSM document.
	AssociationParameters interface{} `field:"optional" json:"associationParameters" yaml:"associationParameters"`
}

