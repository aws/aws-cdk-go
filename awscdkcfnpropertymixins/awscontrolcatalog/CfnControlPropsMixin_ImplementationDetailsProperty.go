package awscontrolcatalog


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   implementationDetailsProperty := &ImplementationDetailsProperty{
//   	Identifier: jsii.String("identifier"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-controlcatalog-control-implementationdetails.html
//
type CfnControlPropsMixin_ImplementationDetailsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-controlcatalog-control-implementationdetails.html#cfn-controlcatalog-control-implementationdetails-identifier
	//
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-controlcatalog-control-implementationdetails.html#cfn-controlcatalog-control-implementationdetails-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

