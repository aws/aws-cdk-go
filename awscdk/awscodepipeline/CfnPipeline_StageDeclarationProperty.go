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
//   stageDeclarationProperty := &StageDeclarationProperty{
//   	Actions: []interface{}{
//   		&ActionDeclarationProperty{
//   			ActionTypeId: &ActionTypeIdProperty{
//   				Category: jsii.String("category"),
//   				Owner: jsii.String("owner"),
//   				Provider: jsii.String("provider"),
//   				Version: jsii.String("version"),
//   			},
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Configuration: configuration,
//   			InputArtifacts: []interface{}{
//   				&InputArtifactProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Namespace: jsii.String("namespace"),
//   			OutputArtifacts: []interface{}{
//   				&OutputArtifactProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Region: jsii.String("region"),
//   			RoleArn: jsii.String("roleArn"),
//   			RunOrder: jsii.Number(123),
//   		},
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Blockers: []interface{}{
//   		&BlockerDeclarationProperty{
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
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

