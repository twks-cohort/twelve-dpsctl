package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
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
		fmt.Printf("Created team: %s ", team.Name)
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
		team, err := GetTeamHandler(apiUrl, teamName)
		if err != nil {
			fmt.Println(err)
			return
		}
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
		fmt.Printf("Deleted team: %s", deletedTeam.Name)
	},
}

func init() {
	getCmd.AddCommand(listTeamCmd)
	getCmd.AddCommand(getTeamCmd)
	createCmd.AddCommand(createTeamCmd)
	deleteCmd.AddCommand(deleteTeamCmd)
	kubeconfigCmd.Flags().StringVarP(&teamName, "team", "t", "", "Get the team name from teams api")
}

type Team struct {
	Name string `json:"name"`
}

type transientError struct {
	err error
}

func (t transientError) Error() string {
	return fmt.Sprintf("%v", t.err)
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

func GetTeamHandler(apiUrl string, teamName string) (*Team, error) {

	team, err := GetTeamHttpHandler(apiUrl, teamName)
	if err != nil {
		terr := transientError{}
		if errors.As(err, &terr) {
			log.Printf("There was a problem getting the team data \n")
			log.Fatal(err.Error())
		} else {
			return nil, err
		}
	}
	return team, nil
}

func GetTeamHttpHandler(url string, teamName string) (*Team, error) {

	team := &Team{Name: teamName}
	request, err := http.NewRequest(
		http.MethodGet,
		url+"/v1/teams/"+teamName,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create new http request object: , %+v", err)
	}

	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, transientError{err: err}
	}

	if response.StatusCode == 404 {
		return nil, fmt.Errorf("Team not found: %s", team.Name)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, transientError{err: err}
	}

	err = json.Unmarshal(responseBytes, &team)
	if err != nil {
		return nil, transientError{err: err}
	}

	return team, nil

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
