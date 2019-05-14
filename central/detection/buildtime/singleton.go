package buildtime

import (
	"context"

	"github.com/stackrox/rox/central/detection"
	policyDataStore "github.com/stackrox/rox/central/policy/datastore"
	"github.com/stackrox/rox/central/searchbasedpolicies/matcher"
	policyUtils "github.com/stackrox/rox/pkg/policies"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/sync"
)

var (
	once sync.Once

	policySet detection.PolicySet
	detector  Detector
)

// SingletonDetector returns the singleton instance of a Detector.
func SingletonDetector() Detector {
	once.Do(initialize)
	return detector
}

// SingletonPolicySet returns the singleton instance of a PolicySet.
func SingletonPolicySet() detection.PolicySet {
	once.Do(initialize)
	return policySet
}

func initialize() {
	ctx := sac.WithAllAccess(context.TODO())
	policySet = detection.NewPolicySet(policyDataStore.Singleton(), detection.NewPolicyCompiler(matcher.ImageBuilderSingleton()))
	policies, err := policyDataStore.Singleton().GetPolicies(ctx)
	if err != nil {
		panic(err)
	}
	for _, policy := range policies {
		if policyUtils.AppliesAtBuildTime(policy) {
			if err := policySet.UpsertPolicy(policy); err != nil {
				panic(err)
			}
		}
	}

	detector = NewDetector(policySet)
}
