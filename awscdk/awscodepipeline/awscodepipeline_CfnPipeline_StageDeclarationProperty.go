package awscodepipeline


// Represents information about a stage and its definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configuration interface{}
//
//   stageDeclarationProperty := &stageDeclarationProperty{
//   	actions: []interface{}{
//   		&actionDeclarationProperty{
//   			actionTypeId: &actionTypeIdProperty{
//   				category: jsii.String("category"),
//   				owner: jsii.String("owner"),
//   				provider: jsii.String("provider"),
//   				version: jsii.String("version"),
//   			},
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			configuration: configuration,
//   			inputArtifacts: []interface{}{
//   				&inputArtifactProperty{
//   					name: jsii.String("name"),
//   				},
//   			},
//   			namespace: jsii.String("namespace"),
//   			outputArtifacts: []interface{}{
//   				&outputArtifactProperty{
//   					name: jsii.String("name"),
//   				},
//   			},
//   			region: jsii.String("region"),
//   			roleArn: jsii.String("roleArn"),
//   			runOrder: jsii.Number(123),
//   		},
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	blockers: []interface{}{
//   		&blockerDeclarationProperty{
//   			name: jsii.String("name"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   }
//
type CfnPipeline_StageDeclarationProperty struct {
	// The actions included in a stage.
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// The name of the stage.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Reserved for future use.
	Blockers interface{} `field:"optional" json:"blockers" yaml:"blockers"`
}

