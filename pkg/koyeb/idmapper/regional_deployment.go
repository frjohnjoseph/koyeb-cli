package idmapper

import (
	"context"
	"fmt"
	"strconv"

	"github.com/koyeb/koyeb-api-client-go/api/v1/koyeb"
	"github.com/pkg/errors"
)

type RegionalDeploymentMapper struct {
	ctx     context.Context
	client  *koyeb.APIClient
	fetched bool
	sidMap  *IDMap
}

func NewRegionalDeploymentMapper(ctx context.Context, client *koyeb.APIClient) *RegionalDeploymentMapper {
	return &RegionalDeploymentMapper{
		ctx:     ctx,
		client:  client,
		fetched: false,
		sidMap:  NewIDMap(),
	}
}

func (mapper *RegionalDeploymentMapper) ResolveID(val string) (string, error) {
	if IsUUIDv4(val) {
		return val, nil
	}

	if !mapper.fetched {
		err := mapper.fetch()
		if err != nil {
			return "", err
		}
	}

	id, ok := mapper.sidMap.GetID(val)
	if ok {
		return id, nil
	}

	return "", fmt.Errorf("id not found %q", val)
}

func (mapper *RegionalDeploymentMapper) GetShortID(id string) (string, error) {
	if !mapper.fetched {
		err := mapper.fetch()
		if err != nil {
			return "", err
		}
	}

	sid, ok := mapper.sidMap.GetValue(id)
	if !ok {
		return "", fmt.Errorf("deployment short id not found for %q", id)
	}

	return sid, nil
}

func (mapper *RegionalDeploymentMapper) fetch() error {
	radix := NewRadixTree()

	page := int64(0)
	offset := int64(0)
	limit := int64(100)
	for {

		resp, _, err := mapper.client.RegionalDeploymentsApi.ListRegionalDeployments(mapper.ctx).
			Limit(strconv.FormatInt(limit, 10)).
			Offset(strconv.FormatInt(offset, 10)).
			Execute()
		if err != nil {
			return errors.Wrap(err, "cannot list regional deployments from API")
		}

		deployments := resp.GetRegionalDeployments()
		for i := range deployments {
			deployment := &deployments[i]
			radix.Insert(getKey(deployment.GetId()), deployment)
		}

		page++
		offset = page * limit
		if offset >= resp.GetCount() {
			break
		}
	}

	minLength := radix.MinimalLength(8)
	err := radix.ForEach(func(key Key, value Value) error {
		deployment := value.(*koyeb.RegionalDeploymentListItem)
		id := deployment.GetId()
		sid := getShortID(id, minLength)

		mapper.sidMap.Set(id, sid)

		return nil
	})
	if err != nil {
		return err
	}

	mapper.fetched = true

	return nil
}
