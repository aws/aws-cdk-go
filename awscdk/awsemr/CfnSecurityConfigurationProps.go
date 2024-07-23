package awsemr


// Properties for defining a `CfnSecurityConfiguration`.
//
// Example:
//   import emr "github.com/aws/aws-cdk-go/awscdk"
//
//
//   cfnSecurityConfiguration := emr.NewCfnSecurityConfiguration(this, jsii.String("EmrSecurityConfiguration"), &CfnSecurityConfigurationProps{
//   	Name: jsii.String("AddStepRuntimeRoleSecConfig"),
//   	SecurityConfiguration: jSON.parse(jsii.String(`
//   	    {
//   	      "AuthorizationConfiguration": {
//   	          "IAMConfiguration": {
//   	              "EnableApplicationScopedIAMRole": true,
//   	              "ApplicationScopedIAMRoleConfiguration":
//   	                  {
//   	                      "PropagateSourceIdentity": true
//   	                  }
//   	          },
//   	          "LakeFormationConfiguration": {
//   	              "AuthorizedSessionTagValue": "Amazon EMR"
//   	          }
//   	      }
//   	    }`)),
//   })
//
//   task := tasks.NewEmrCreateCluster(this, jsii.String("Create Cluster"), &EmrCreateClusterProps{
//   	Instances: &InstancesConfigProperty{
//   	},
//   	Name: sfn.TaskInput_FromJsonPathAt(jsii.String("$.ClusterName")).value,
//   	SecurityConfiguration: cfnSecurityConfiguration.Name,
//   })
//
//   executionRole := iam.NewRole(this, jsii.String("Role"), &RoleProps{
//   	AssumedBy: iam.NewArnPrincipal(task.clusterRole.RoleArn),
//   })
//
//   executionRole.AssumeRolePolicy.AddStatements(
//   iam.NewPolicyStatement(&PolicyStatementProps{
//   	Effect: iam.Effect_ALLOW,
//   	Principals: []iPrincipal{
//   		task.clusterRole,
//   	},
//   	Actions: []*string{
//   		jsii.String("sts:SetSourceIdentity"),
//   	},
//   }),
//   iam.NewPolicyStatement(&PolicyStatementProps{
//   	Effect: iam.Effect_ALLOW,
//   	Principals: []*iPrincipal{
//   		task.clusterRole,
//   	},
//   	Actions: []*string{
//   		jsii.String("sts:TagSession"),
//   	},
//   	Conditions: map[string]interface{}{
//   		"StringEquals": map[string]*string{
//   			"aws:RequestTag/LakeFormationAuthorizedCaller": jsii.String("Amazon EMR"),
//   		},
//   	},
//   }))
//
//   tasks.NewEmrAddStep(this, jsii.String("Task"), &EmrAddStepProps{
//   	ClusterId: jsii.String("ClusterId"),
//   	ExecutionRoleArn: executionRole.RoleArn,
//   	Name: jsii.String("StepName"),
//   	Jar: jsii.String("Jar"),
//   	ActionOnFailure: tasks.ActionOnFailure_CONTINUE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-securityconfiguration.html
//
type CfnSecurityConfigurationProps struct {
	// The security configuration details in JSON format.
	//
	// For JSON parameters and examples, see [Use Security Configurations to Set Up Cluster Security](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-security-configurations.html) in the *Amazon EMR Management Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-securityconfiguration.html#cfn-emr-securityconfiguration-securityconfiguration
	//
	SecurityConfiguration interface{} `field:"required" json:"securityConfiguration" yaml:"securityConfiguration"`
	// The name of the security configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-securityconfiguration.html#cfn-emr-securityconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

