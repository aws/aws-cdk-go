package previewawssagemakermixins


// Defines the instance capacity requirements for an instance group, including configurations for both Spot and On-Demand capacity types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var onDemand interface{}
//   var spot interface{}
//
//   clusterCapacityRequirementsProperty := &ClusterCapacityRequirementsProperty{
//   	OnDemand: onDemand,
//   	Spot: spot,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clustercapacityrequirements.html
//
type CfnClusterPropsMixin_ClusterCapacityRequirementsProperty struct {
	// Configuration options specific to On-Demand instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clustercapacityrequirements.html#cfn-sagemaker-cluster-clustercapacityrequirements-ondemand
	//
	OnDemand interface{} `field:"optional" json:"onDemand" yaml:"onDemand"`
	// Configuration options specific to Spot instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-cluster-clustercapacityrequirements.html#cfn-sagemaker-cluster-clustercapacityrequirements-spot
	//
	Spot interface{} `field:"optional" json:"spot" yaml:"spot"`
}

