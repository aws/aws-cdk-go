package awscontrolcatalog


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   implementationDetailsProperty := &ImplementationDetailsProperty{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Identifier: jsii.String("identifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-controlcatalog-control-implementationdetails.html
//
type CfnControl_ImplementationDetailsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-controlcatalog-control-implementationdetails.html#cfn-controlcatalog-control-implementationdetails-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-controlcatalog-control-implementationdetails.html#cfn-controlcatalog-control-implementationdetails-identifier
	//
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
}

