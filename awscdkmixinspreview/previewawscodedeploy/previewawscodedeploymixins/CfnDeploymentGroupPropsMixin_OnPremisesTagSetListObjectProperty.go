package previewawscodedeploymixins


// The `OnPremisesTagSetListObject` property type specifies lists of on-premises instance tag groups.
//
// In order for an instance to be included in the deployment group, it must be identified by all the tag groups in the list.
//
// `OnPremisesTagSetListObject` is a property of the [CodeDeploy DeploymentGroup OnPremisesTagSet](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-onpremisestagset.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   onPremisesTagSetListObjectProperty := &OnPremisesTagSetListObjectProperty{
//   	OnPremisesTagGroup: []interface{}{
//   		&TagFilterProperty{
//   			Key: jsii.String("key"),
//   			Type: jsii.String("type"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-onpremisestagsetlistobject.html
//
type CfnDeploymentGroupPropsMixin_OnPremisesTagSetListObjectProperty struct {
	// Information about groups of on-premises instance tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-onpremisestagsetlistobject.html#cfn-codedeploy-deploymentgroup-onpremisestagsetlistobject-onpremisestaggroup
	//
	OnPremisesTagGroup interface{} `field:"optional" json:"onPremisesTagGroup" yaml:"onPremisesTagGroup"`
}

