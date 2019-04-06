package transformer

import (
	"github.com/sirupsen/logrus"
	"hq-stash.corp.proofpoint.com/del/manifest-transformer/pkg/parser"
)

type Alerting struct {
	//Is Any default alert enabled
	IsDefaultAlerts bool `yaml:"isDefaultAlerts"`
	//Is general alerts enabled
	IsGeneralAlerts bool `yaml:"isGeneralAlerts"`
	//Is Ingress Alerts
	IsIngressAlerts bool `yaml:"isIngressAlerts"`
	//Does custom alerts exist
	IsCustomAlerts bool `yaml:"isCustomAlerts"`
	//Default environment specific alert thresholds
	DefaultRulesValues parser.DefaultAlerts `yaml:"defaultRulesValues"`
	//Custom environment specific alert thresholds
	CustomRulesValues map[string]string `yaml:"custom"`
}

func generateAlertingObject(data *parser.ManifestData, environment parser.EnvironmentValues, namespace string) Alerting {
	var alerting Alerting

	//Get alert profile
	profileName := environment.Alerting
	if profileName == "" {
		return alerting
	}
	profile, ok := data.AlertingProfiles.Profiles[profileName]
	if !ok {
		logrus.Errorf("The alerting profile:%s is not defined", profileName)
		return alerting
	}
	logrus.Infof("The alerting profile is:: %v", profile)

	if &profile.DefaultAlerts != nil {
		//If defaults alerts section is not empty, enable all default alerts, and then check for single alert threshold
		//If a particular alert threshold info is not specified, that alert is turned off automatically
		//I initially split default alerts into general alerts and Ingress alerts.
		//But as that distinction is no longer necessary, I enable all so that I dont have to touch the template files for now
		//Do the clean up later
		alerting.IsDefaultAlerts = true
		alerting.IsGeneralAlerts = true
		alerting.IsIngressAlerts = true
		alerting.DefaultRulesValues = profile.DefaultAlerts
	}

	if data.AlertingProfiles.CustomRules.Prometheus != nil {
		alerting.IsCustomAlerts = true
		alerting.CustomRulesValues = make(map[string]string)
	}

	if len(profile.CustomAlerts.Prometheus.CustomVariables) != 0 {
		alerting.CustomRulesValues = profile.CustomAlerts.Prometheus.CustomVariables
	}

	if alerting.IsCustomAlerts {
		alerting.CustomRulesValues["namespace"] = namespace
	}

	return alerting
}
