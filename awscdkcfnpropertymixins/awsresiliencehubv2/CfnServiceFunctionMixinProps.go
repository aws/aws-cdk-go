package awsresiliencehubv2


// Properties for CfnServiceFunctionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnServiceFunctionMixinProps := &CfnServiceFunctionMixinProps{
//   	Criticality: jsii.String("criticality"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	ServiceArn: jsii.String("serviceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-servicefunction.html
//
type CfnServiceFunctionMixinProps struct {
	// The criticality of the service function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-servicefunction.html#cfn-resiliencehubv2-servicefunction-criticality
	//
	Criticality *string `field:"optional" json:"criticality" yaml:"criticality"`
	// The description of the service function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-servicefunction.html#cfn-resiliencehubv2-servicefunction-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the service function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-servicefunction.html#cfn-resiliencehubv2-servicefunction-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ARN of the parent service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resiliencehubv2-servicefunction.html#cfn-resiliencehubv2-servicefunction-servicearn
	//
	ServiceArn *string `field:"optional" json:"serviceArn" yaml:"serviceArn"`
}

