package awsresiliencehubv2


// Properties for defining a `CfnServiceFunction`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServiceFunctionProps := &CfnServiceFunctionProps{
//   	Criticality: jsii.String("criticality"),
//   	Name: jsii.String("name"),
//   	ServiceArn: jsii.String("serviceArn"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-servicefunction.html
//
type CfnServiceFunctionProps struct {
	// The criticality of the service function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-servicefunction.html#cfn-resiliencehubv2-servicefunction-criticality
	//
	Criticality *string `field:"required" json:"criticality" yaml:"criticality"`
	// The name of the service function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-servicefunction.html#cfn-resiliencehubv2-servicefunction-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of the parent service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-servicefunction.html#cfn-resiliencehubv2-servicefunction-servicearn
	//
	ServiceArn *string `field:"required" json:"serviceArn" yaml:"serviceArn"`
	// The description of the service function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-servicefunction.html#cfn-resiliencehubv2-servicefunction-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

