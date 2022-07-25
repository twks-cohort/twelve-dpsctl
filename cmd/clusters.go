package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// TODO - limited to whatever clusters are in tpl/config.go.tpl, does not check remote account(s) for others
// TODO - we should move this to a list command if we start to care about verbs/etc, instead of under `get`
var listClustersCmd = &cobra.Command{
	Use:               "clusters",
	Short:             "List clusters",
	Long:              `Write cluster names to stdout`,
	DisableAutoGenTag: true,
	Args:              cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		clusters := listClusters()
		fmt.Println("Available clusters are: ", strings.Join(clusters, ", "))
	},
}

func init() {
	getCmd.AddCommand(listClustersCmd)
}

func listClusters() []string {
	var clusterNames []string

	for _, c := range clusters {
		clusterNames = append(clusterNames, c.clusterName)
	}

	return clusterNames
}
