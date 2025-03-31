package awsdatazone


// The hyper pod properties of a AWS Glue properties patch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hyperPodPropertiesInputProperty := &HyperPodPropertiesInputProperty{
//   	ClusterName: jsii.String("clusterName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-hyperpodpropertiesinput.html
//
type CfnConnection_HyperPodPropertiesInputProperty struct {
	// The cluster name the hyper pod properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-hyperpodpropertiesinput.html#cfn-datazone-connection-hyperpodpropertiesinput-clustername
	//
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
}

