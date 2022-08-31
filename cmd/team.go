package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var teamName string

var createTeamCmd = &cobra.Command{
	Use:               "team",
	Short:             "create team",
	Long:              `Create new team with teams api`,
	DisableAutoGenTag: true,
	Args:              cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		teamName := args[0]
		apiUrl := viper.GetString("TeamsApi")
		team, err := CreateTeam(apiUrl, teamName)
		if err != nil {
			err = fmt.Errorf("unable to create team - %+v", err)
			fmt.Println(err)
			return
		}
		fmt.Println(team)
	},
}

var getTeamCmd = &cobra.Command{
	Use:               "team",
	Short:             "Get team",
	Long:              `Get team from teams api`,
	DisableAutoGenTag: true,
	Args:              cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiUrl := viper.GetString("TeamsApi")
		if len(args) < 1 {
			teams := listTeams(apiUrl)
			fmt.Printf("Teams: %v", strings.Join(teams, " "))
			return
		}

		teamName := args[0]
		team := GetTeam(apiUrl, teamName)
		// We'll want to return the entire object if there are other fields in the future
		fmt.Printf(team.Name)
	},
}

var listTeamCmd = &cobra.Command{
	Use:               "teams",
	Short:             "get teams",
	Long:              `List teams from teams api`,
	DisableAutoGenTag: true,
	Args:              cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		teams_url := viper.GetString("TeamsApi")
		teams := listTeams(teams_url)
		fmt.Printf("Teams: %v ", strings.Join(teams, " "))
	},
}

var deleteTeamCmd = &cobra.Command{
	Use:               "team",
	Short:             "delete team",
	Long:              `Delete team from teams api`,
	DisableAutoGenTag: true,
	Args:              cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		teamName := args[0]
		teams_url := viper.GetString("TeamsApi")
		deletedTeam, err := DeleteTeam(teams_url, teamName)
		if err != nil {
			err = fmt.Errorf("unable to delete team - %+v", err)
			fmt.Println(err)
			return
		}
		fmt.Printf("Deleted team: %v ", *deletedTeam)
	},
}

func init() {
	getCmd.AddCommand(listTeamCmd)
	getCmd.AddCommand(getTeamCmd)
	createCmd.AddCommand(createTeamCmd)
	deleteCmd.AddCommand(deleteTeamCmd)
	kubeconfigCmd.Flags().StringVarP(&teamName, "team", "t", "", "Get the team name from teams api")
}

type Teams []struct {
	Name string `json:"name"`
}

type Team struct {
	Name string `json:"name"`
}

func listTeams(apiUrl string) []string {
	responseBytes := GetTeamsData(apiUrl)
	var teams []Team

	err := json.Unmarshal(responseBytes, &teams)
	if err != nil {
		fmt.Printf("could not unmarshal %v", err)
	}

	var teamNames []string
	for _, c := range teams {
		teamNames = append(teamNames, c.Name)
	}

	return teamNames

}

func CreateTeam(apiUrl string, teamName string) (*Team, error) {

	team := &Team{Name: teamName}
	jsonValue, _ := json.Marshal(*team)

	request, err := http.NewRequest(
		http.MethodPost,
		apiUrl+"/v1/teams",
		bytes.NewBuffer(jsonValue),
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%q", err)
		return nil, err
	}

	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == 409 {
		return nil, fmt.Errorf("the team %+v already exists", *team)
	}

	if response.StatusCode != 201 {
		return nil, fmt.Errorf("%+v", response.Status)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &team)
	if err != nil {
		fmt.Printf("could not unmarshal %v", err)
		return nil, err
	}

	return team, nil
}

func GetTeam(apiUrl string, teamName string) Team {
	responseBytes, err := TeamData(apiUrl, teamName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%q", err)
		os.Exit(1)
	}

	var team Team
	err = json.Unmarshal(responseBytes, &team)
	if err != nil {
		fmt.Printf("could not unmarshal %v", err)
	}

	return team
}

func DeleteTeam(apiUrl string, teamName string) (*Team, error) {
	team := &Team{Name: teamName}
	request, err := http.NewRequest(
		http.MethodDelete,
		apiUrl+"/v1/teams/"+team.Name,
		nil,
	)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Accept", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == 404 {
		return nil, fmt.Errorf("error from server (%s): team '%+v' not found", response.Status, team.Name)
	}

	if response.StatusCode != 204 {
		return nil, fmt.Errorf("failed to delete %s, api response: %+v", team.Name, response.Status)
	}

	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return team, nil
}

func TeamData(url string, teamName string) ([]byte, error) {
	request, err := http.NewRequest(
		http.MethodGet,
		url+"/v1/teams/"+teamName,
		nil,
	)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Accept", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBytes, nil
}

func GetTeamsData(apiUrl string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		apiUrl+"/v1/teams",
		nil,
	)

	if err != nil {
		fmt.Printf("Failed to get team: %v", err)
	}

	request.Header.Add("Accept", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("Failed to get HTTP response: %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Failed to parse http response body: %v", err)
	}

	return responseBytes
}
