package previewawsdevopsgurumixins


// Properties for CfnResourceCollectionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResourceCollectionMixinProps := &CfnResourceCollectionMixinProps{
//   	ResourceCollectionFilter: &ResourceCollectionFilterProperty{
//   		CloudFormation: &CloudFormationCollectionFilterProperty{
//   			StackNames: []*string{
//   				jsii.String("stackNames"),
//   			},
//   		},
//   		Tags: []TagCollectionProperty{
//   			&TagCollectionProperty{
//   				AppBoundaryKey: jsii.String("appBoundaryKey"),
//   				TagValues: []*string{
//   					jsii.String("tagValues"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsguru-resourcecollection.html
//
type CfnResourceCollectionMixinProps struct {
	// Information about a filter used to specify which AWS resources are analyzed for anomalous behavior by DevOps Guru.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsguru-resourcecollection.html#cfn-devopsguru-resourcecollection-resourcecollectionfilter
	//
	ResourceCollectionFilter interface{} `field:"optional" json:"resourceCollectionFilter" yaml:"resourceCollectionFilter"`
}

