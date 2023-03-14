package awscodedeploy


// The `OnPremisesTagSetListObject` property type specifies lists of on-premises instance tag groups.
//
// In order for an instance to be included in the deployment group, it must be identified by all the tag groups in the list.
//
// `OnPremisesTagSetListObject` is a property of the [CodeDeploy DeploymentGroup OnPremisesTagSet](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-onpremisestagset.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onPremisesTagSetListObjectProperty := &onPremisesTagSetListObjectProperty{
//   	onPremisesTagGroup: []interface{}{
//   		&tagFilterProperty{
//   			key: jsii.String("key"),
//   			type: jsii.String("type"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDeploymentGroup_OnPremisesTagSetListObjectProperty struct {
	// Information about groups of on-premises instance tags.
	OnPremisesTagGroup interface{} `field:"optional" json:"onPremisesTagGroup" yaml:"onPremisesTagGroup"`
}

