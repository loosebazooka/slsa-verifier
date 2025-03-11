package common

var trustedBuilderRepository = "https://github.com/slsa-framework/slsa-github-generator"

var (
	// GenericGeneratorBuilderID is the builder ID for the Generic Generator.
	GenericGeneratorBuilderID = trustedBuilderRepository + "/.github/workflows/generator_generic_slsa3.yml"
	// ContainerGeneratorBuilderID is the builder ID for the Container Generator.
	ContainerGeneratorBuilderID = trustedBuilderRepository + "/.github/workflows/generator_container_slsa3.yml"
	// GoBuilderID is the SLSA builder ID for the Go Builder.
	GoBuilderID = trustedBuilderRepository + "/.github/workflows/builder_go_slsa3.yml"
	// ContainerBasedBuilderID is the SLSA builder ID for the Container-Based Builder.
	ContainerBasedBuilderID = trustedBuilderRepository + "/.github/workflows/builder_container-based_slsa3.yml"

	// NpmCLILegacyBuilderID is the legacy builder ID for the npm CLI.
	NpmCLILegacyBuilderID = "https://github.com/actions/runner"
	// NpmCLIHostedBuilderID is the builder ID for the npm CLI on Hosted GitHub Actions.
	NpmCLIHostedBuilderID = NpmCLILegacyBuilderID + "/github-hosted"
	// NpmCLISelfHostedBuilderID is the builder ID for the npm CLI on Self-hosted GitHub Actions.
	NpmCLISelfHostedBuilderID = NpmCLILegacyBuilderID + "/self-hosted"

	// GenericDelegatorBuilderID is the SLSA builder ID for the BYOB Generic Low-Permissions Delegated Builder.
	GenericDelegatorBuilderID = trustedBuilderRepository + "/.github/workflows/delegator_generic_slsa3.yml"
	// GenericLowPermsDelegatorBuilderID is the SLSA builder ID for the BYOB Generic Low-Permissions Delegated Builder.
	GenericLowPermsDelegatorBuilderID = trustedBuilderRepository + "/.github/workflows/delegator_lowperms-generic_slsa3.yml"

	// BCRReusableBuilderID is the bcr resuable workflow that generated provenance for Bazel Central Registry
	BCRBuilderId = "https://github.com/publish-to-bcr-dev/.github/.github/workflows/release_ruleset.yaml"
)
