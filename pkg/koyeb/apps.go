package koyeb

import (
	"context"

	"github.com/koyeb/koyeb-api-client-go/api/v1/koyeb"
	"github.com/koyeb/koyeb-cli/pkg/koyeb/idmapper"
	"github.com/spf13/cobra"
)

func NewAppCmd() *cobra.Command {
	h := NewAppHandler()

	appCmd := &cobra.Command{
		Use:               "apps ACTION",
		Aliases:           []string{"a", "app"},
		Short:             "Apps",
		PersistentPreRunE: h.InitHandler,
	}

	createAppCmd := &cobra.Command{
		Use:   "create NAME",
		Short: "Create app",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			createApp := koyeb.NewCreateAppWithDefaults()
			SyncFlags(cmd, args, createApp)
			return h.Create(cmd, args, createApp)
		},
	}
	appCmd.AddCommand(createAppCmd)

	initAppCmd := &cobra.Command{
		Use:   "init NAME",
		Short: "Create app and service",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			createApp := koyeb.NewCreateAppWithDefaults()

			createService := koyeb.NewCreateServiceWithDefaults()
			createDefinition := koyeb.NewDeploymentDefinitionWithDefaults()

			err := parseServiceDefinitionFlags(cmd.Flags(), createDefinition, true)
			if err != nil {
				return err
			}
			createDefinition.Name = koyeb.PtrString(args[0])

			createService.SetDefinition(*createDefinition)

			return h.Init(cmd, args, createApp, createService)
		},
	}
	appCmd.AddCommand(initAppCmd)
	addServiceDefinitionFlags(initAppCmd.Flags())

	getAppCmd := &cobra.Command{
		Use:   "get NAME",
		Short: "Get app",
		Args:  cobra.ExactArgs(1),
		RunE:  h.Get,
	}
	appCmd.AddCommand(getAppCmd)

	listAppCmd := &cobra.Command{
		Use:   "list",
		Short: "List apps",
		RunE:  h.List,
	}
	appCmd.AddCommand(listAppCmd)

	describeAppCmd := &cobra.Command{
		Use:   "describe NAME",
		Short: "Describe app",
		Args:  cobra.ExactArgs(1),
		RunE:  h.Describe,
	}
	appCmd.AddCommand(describeAppCmd)

	updateAppCmd := &cobra.Command{
		Use:   "update NAME",
		Short: "Update app",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			updateApp := koyeb.NewUpdateAppWithDefaults()
			SyncFlags(cmd, args, updateApp)
			return h.Update(cmd, args, updateApp)
		},
	}
	appCmd.AddCommand(updateAppCmd)
	updateAppCmd.Flags().StringP("name", "n", "", "Name of the app")

	deleteAppCmd := &cobra.Command{
		Use:   "delete NAME",
		Short: "Delete app",
		Args:  cobra.ExactArgs(1),
		RunE:  h.Delete,
	}
	appCmd.AddCommand(deleteAppCmd)

	pauseServiceCmd := &cobra.Command{
		Use:   "pause NAME",
		Short: "Pause app",
		Args:  cobra.ExactArgs(1),
		RunE:  h.Pause,
	}
	appCmd.AddCommand(pauseServiceCmd)

	resumeServiceCmd := &cobra.Command{
		Use:   "resume NAME",
		Short: "Resume app",
		Args:  cobra.ExactArgs(1),
		RunE:  h.Resume,
	}
	appCmd.AddCommand(resumeServiceCmd)

	return appCmd
}

func NewAppHandler() *AppHandler {
	return &AppHandler{}
}

type AppHandler struct {
	ctx    context.Context
	client *koyeb.APIClient
	mapper *idmapper.Mapper
}

func (h *AppHandler) InitHandler(cmd *cobra.Command, args []string) error {
	h.ctx = getAuth(context.Background())
	h.client = getApiClient()
	h.mapper = idmapper.NewMapper(h.ctx, h.client)
	return nil
}

func (h *AppHandler) ResolveAppArgs(val string) string {
	appMapper := h.mapper.App()
	id, err := appMapper.ResolveID(val)
	if err != nil {
		fatalApiError(err, nil)
	}

	return id
}
