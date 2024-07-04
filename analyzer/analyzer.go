package requests

import "github.com/nhatnam1198/interface-compliance-linter/types"

type Validator interface {
	Validate() error
}

func (c CreateEtcdCluster) Validate() error {
	return nil
}

type CreateEtcdCluster struct {
	ResourceID string
	Name       string
	Val        int64
}

func (c CreateEtcdCluster) GetResourceMappings() []types.ResourceMapping {
	return []types.ResourceMapping{
		{
			ResourceId:   types.Id(1),
			ResourceType: "random string",
		},
	}
}

func (c CreateEtcdCluster) AnotherFunc() int {
	return 3
}

type TruncateTableRequest struct {
}

func (t TruncateTableRequest) Validate() error {
	return nil
}

type MigrateDatabaseRequest struct {
}

type ResetPasswordRequest struct {
}

func (t MigrateDatabaseRequest) Validate() error {
	return nil
}

func (t ResetPasswordRequest) Validate() error {
	return nil
}