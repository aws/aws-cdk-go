package awsbedrockagentcorealpha


// Configuration for episodic memory reflection.
//
// Example:
//   // Create memory with custom strategies
//   memory := agentcore.NewMemory(this, jsii.String("MyMemory"), &MemoryProps{
//   	MemoryName: jsii.String("my_memory"),
//   	Description: jsii.String("Memory with custom strategies"),
//   	ExpirationDuration: cdk.Duration_Days(jsii.Number(90)),
//   	MemoryStrategies: []IMemoryStrategy{
//   		agentcore.MemoryStrategy_UsingUserPreference(&ManagedStrategyProps{
//   			Name: jsii.String("CustomerPreferences"),
//   			Namespaces: []*string{
//   				jsii.String("support/customer/{actorId}/preferences"),
//   			},
//   		}),
//   		agentcore.MemoryStrategy_UsingSemantic(&ManagedStrategyProps{
//   			Name: jsii.String("CustomerSupportSemantic"),
//   			Namespaces: []*string{
//   				jsii.String("support/customer/{actorId}/semantic"),
//   			},
//   		}),
//   		agentcore.MemoryStrategy_UsingEpisodic(&ManagedStrategyProps{
//   			Name: jsii.String("customerJourneyEpisodic"),
//   			Namespaces: []*string{
//   				jsii.String("/journey/customer/{actorId}/episodes"),
//   			},
//   			ReflectionConfiguration: &EpisodicReflectionConfiguration{
//   				Namespaces: []*string{
//   					jsii.String("/journey/customer/{actorId}/reflections"),
//   				},
//   			},
//   		}),
//   	},
//   })
//
// Experimental.
type EpisodicReflectionConfiguration struct {
	// Namespaces for episodic reflection Minimum 1 namespace required.
	// Experimental.
	Namespaces *[]*string `field:"required" json:"namespaces" yaml:"namespaces"`
}

