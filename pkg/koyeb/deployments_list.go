package koyeb

import (
	"strconv"

	"github.com/koyeb/koyeb-api-client-go/api/v1/koyeb"
	"github.com/koyeb/koyeb-cli/pkg/koyeb/idmapper"
	"github.com/koyeb/koyeb-cli/pkg/koyeb/renderer"
	"github.com/spf13/cobra"
)

func (h *DeploymentHandler) List(cmd *cobra.Command, args []string) error {
	list := []koyeb.DeploymentListItem{}

	page := int64(0)
	offset := int64(0)
	limit := int64(100)
	for {
		res, resp, err := h.client.DeploymentsApi.ListDeployments(h.ctx).
			Limit(strconv.FormatInt(limit, 10)).Offset(strconv.FormatInt(offset, 10)).Execute()
		if err != nil {
			fatalApiError(err, resp)
		}
		list = append(list, res.GetDeployments()...)

		page++
		offset = page * limit
		if offset >= res.GetCount() {
			break
		}
	}

	full := GetBoolFlags(cmd, "full")
	output := GetStringFlags(cmd, "output")
	listDeploymentsReply := NewListDeploymentsReply(h.mapper, &koyeb.ListDeploymentsReply{Deployments: list}, full)

	return renderer.NewListRenderer(listDeploymentsReply).Render(output)
}

type ListDeploymentsReply struct {
	mapper *idmapper.Mapper
	value  *koyeb.ListDeploymentsReply
	full   bool
}

func NewListDeploymentsReply(mapper *idmapper.Mapper, value *koyeb.ListDeploymentsReply, full bool) *ListDeploymentsReply {
	return &ListDeploymentsReply{
		mapper: mapper,
		value:  value,
		full:   full,
	}
}

func (ListDeploymentsReply) Title() string {
	return "Deployments"
}

func (r *ListDeploymentsReply) MarshalBinary() ([]byte, error) {
	return r.value.MarshalJSON()
}

func (r *ListDeploymentsReply) Headers() []string {
	return []string{"id", "service", "type", "status", "messages", "regions", "created_at"}
}

func (r *ListDeploymentsReply) Fields() []map[string]string {
	items := r.value.GetDeployments()
	resp := make([]map[string]string, 0, len(items))

	for _, item := range items {
		fields := map[string]string{
			"id":         renderer.FormatDeploymentID(r.mapper, item.GetId(), r.full),
			"service":    renderer.FormatServiceSlug(r.mapper, item.GetServiceId(), r.full),
			"type":       formatDeploymentType(item.Definition.GetType()),
			"status":     formatDeploymentStatus(item.GetStatus()),
			"messages":   formatDeploymentMessages(item.GetMessages(), 80),
			"regions":    renderRegions(item.Definition.Regions),
			"created_at": renderer.FormatTime(item.GetCreatedAt()),
		}
		resp = append(resp, fields)
	}

	return resp
}
