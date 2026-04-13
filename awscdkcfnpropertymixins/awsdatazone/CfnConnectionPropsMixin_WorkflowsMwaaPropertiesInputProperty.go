package awsdatazone


// Workflows MWAA Properties Input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   workflowsMwaaPropertiesInputProperty := &WorkflowsMwaaPropertiesInputProperty{
//   	MwaaEnvironmentName: jsii.String("mwaaEnvironmentName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-workflowsmwaapropertiesinput.html
//
type CfnConnectionPropsMixin_WorkflowsMwaaPropertiesInputProperty struct {
	// The name of the MWAA environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-workflowsmwaapropertiesinput.html#cfn-datazone-connection-workflowsmwaapropertiesinput-mwaaenvironmentname
	//
	MwaaEnvironmentName *string `field:"optional" json:"mwaaEnvironmentName" yaml:"mwaaEnvironmentName"`
}

