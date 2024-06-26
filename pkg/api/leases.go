package api

// LeasesForTest aggregates all the lease configurations in a test.
// It is assumed that they have been validated and contain only valid and
// unique values.
func LeasesForTest(s *MultiStageTestConfigurationLiteral) (ret []StepLease) {
	if p := s.ClusterProfile; p != "" {
		ret = append(ret, StepLease{
			ResourceType: p.LeaseType(),
			Env:          DefaultLeaseEnv,
			Count:        1,
		})
	}
	for _, step := range append(s.Pre, append(s.Test, s.Post...)...) {
		ret = append(ret, step.Leases...)
	}
	ret = append(ret, s.Leases...)
	return
}

func IPPoolLeaseForTest(s *MultiStageTestConfigurationLiteral) (ret StepLease) {
	if p := s.ClusterProfile; p == "aws" { //TODO(sgoeddel): Hardcoded to only work on aws, eventually this will be available as a configuration
		ret = StepLease{
			ResourceType: p.IPPoolLeaseType(),
			Env:          DefaultIPPoolLeaseEnv,
			Count:        13,
		}
	}
	return
}
