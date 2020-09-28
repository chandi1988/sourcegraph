package licensing

// The list of plans.
const (
	// oldEnterpriseStarter is the old "Enterprise Starter" plan.
	oldEnterpriseStarter Plan = "old-starter-0"
	// oldEnterprise is the old "Enterprise" plan.
	oldEnterprise Plan = "old-enterprise-0"

	// team is the "Team" plan.
	team Plan = "team-0"
	// enterprise is the "Enterprise" plan.
	enterprise Plan = "enterprise-0"
)

var allPlans = []Plan{
	oldEnterpriseStarter,
	oldEnterprise,
	team,
	enterprise,
}

// The list of features. For each feature, add a new const here and the checking logic in
// isFeatureEnabled.
const (
	// FeatureACLs is whether ACLs may be used, such as GitHub, GitLab or Bitbucket Server repository
	// permissions and integration with GitHub, GitLab or Bitbucket Server for user authentication.
	FeatureACLs Feature = "acls"

	// FeatureExtensionRegistry is whether publishing extensions to this Sourcegraph instance is
	// allowed. If not, then extensions must be published to Sourcegraph.com. All instances may use
	// extensions published to Sourcegraph.com.
	FeatureExtensionRegistry Feature = "private-extension-registry"

	// FeatureRemoteExtensionsAllowDisallow is whether the site admin may explicitly specify a list
	// of allowed remote extensions and prevent any other remote extensions from being used. It does
	// not apply to locally published extensions.
	FeatureRemoteExtensionsAllowDisallow Feature = "remote-extensions-allow-disallow"

	// FeatureBranding is whether custom branding of this Sourcegraph instance is allowed.
	FeatureBranding Feature = "branding"

	// FeatureCampaigns is whether campaigns may be used on this Sourcegraph instance.
	FeatureCampaigns Feature = "campaigns"

	// FeatureMonitoring is whether monitoring may be used on this Sourcegraph instance.
	FeatureMonitoring Feature = "monitoring"

	// FeatureBackupAndRestore is whether builtin backup and restore may be used this Sourcegraph instance
	FeatureBackupAndRestore Feature = "backup-and-restore"
)

// planFeatures defines the features that are enabled for each plan.
var planFeatures = map[Plan][]Feature{
	oldEnterpriseStarter: {},
	oldEnterprise: {
		FeatureACLs,
		FeatureExtensionRegistry,
		FeatureRemoteExtensionsAllowDisallow,
		FeatureBranding,
		FeatureCampaigns,
		FeatureMonitoring,
		FeatureBackupAndRestore,
	},
	team:       {},
	enterprise: {},
}
