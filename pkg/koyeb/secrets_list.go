package koyeb

import (
	"strconv"

	"github.com/koyeb/koyeb-api-client-go/api/v1/koyeb"
	"github.com/koyeb/koyeb-cli/pkg/koyeb/idmapper"
	"github.com/koyeb/koyeb-cli/pkg/koyeb/renderer"
	"github.com/spf13/cobra"
)

func (h *SecretHandler) List(cmd *cobra.Command, args []string) error {
	list := []koyeb.Secret{}

	page := int64(0)
	offset := int64(0)
	limit := int64(100)
	for {
		res, resp, err := h.client.SecretsApi.ListSecrets(h.ctx).
			Limit(strconv.FormatInt(limit, 10)).Offset(strconv.FormatInt(offset, 10)).Execute()
		if err != nil {
			fatalApiError(err, resp)
		}
		list = append(list, res.GetSecrets()...)

		page++
		offset = page * limit
		if offset >= res.GetCount() {
			break
		}
	}

	full := GetBoolFlags(cmd, "full")
	output := GetStringFlags(cmd, "output")
	listSecretsReply := NewListSecretsReply(h.mapper, &koyeb.ListSecretsReply{Secrets: list}, full)

	return renderer.NewListRenderer(listSecretsReply).Render(output)
}

type ListSecretsReply struct {
	mapper *idmapper.Mapper
	value  *koyeb.ListSecretsReply
	full   bool
}

func NewListSecretsReply(mapper *idmapper.Mapper, value *koyeb.ListSecretsReply, full bool) *ListSecretsReply {
	return &ListSecretsReply{
		mapper: mapper,
		value:  value,
		full:   full,
	}
}

func (ListSecretsReply) Title() string {
	return "Secrets"
}

func (r *ListSecretsReply) MarshalBinary() ([]byte, error) {
	return r.value.MarshalJSON()
}

func (r *ListSecretsReply) Headers() []string {
	return []string{"id", "name", "type", "value", "created_at"}
}

func (r *ListSecretsReply) Fields() []map[string]string {
	items := r.value.GetSecrets()
	resp := make([]map[string]string, 0, len(items))

	for _, item := range items {
		fields := map[string]string{
			"id":         renderer.FormatSecretID(r.mapper, item.GetId(), r.full),
			"name":       item.GetName(),
			"type":       formatSecretType(item.GetType()),
			"value":      "*****",
			"created_at": renderer.FormatTime(item.GetCreatedAt()),
		}
		resp = append(resp, fields)
	}

	return resp
}
