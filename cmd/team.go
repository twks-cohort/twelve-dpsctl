package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// TODO - handle api versioning (such as /v1 /v2) instead of static strings

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
			teams, err := ListTeams(apiUrl)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("Teams: %v", strings.Join(teams, " "))
			return
		}

		team := &Team{Name: args[0]}
		team, err := GetTeam(apiUrl, team)
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
		teams, err := ListTeams(teams_url)
		if err != nil {
			fmt.Println(err)
			return
		}
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
		teams_url := viper.GetString("TeamsApi")
		team := &Team{Name: args[0]}
		deletedTeam, err := DeleteTeam(teams_url, team)
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

type StatusCodeErr struct {
	code       int
	errMessage error
}

type StatusCodes struct {
	statusErrors []StatusCodeErr
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

func CreateTeam(apiUrl string, teamName string) (*Team, error) {
	team, err := createTeamHandler(apiUrl, teamName)
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

func createTeamHandler(apiUrl string, teamName string) (*Team, error) {
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

	statusCodeErrs := &StatusCodes{[]StatusCodeErr{
		{code: 409, errMessage: fmt.Errorf("could not create team %s, it already exists", team.Name)},
	}}

	responseBytes, err := HttpHandler(request, statusCodeErrs)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &team)
	if err != nil {
		return nil, transientError{err: err}
	}

	return team, nil
}

func DeleteTeam(apiUrl string, team *Team) (*Team, error) {
	deletedTeam, err := deleteTeamHandler(apiUrl, team)
	if err != nil {
		terr := transientError{}
		if errors.As(err, &terr) {
			log.Printf("There was a problem deleteing the team \n")
			log.Fatal(err.Error())
		} else {
			return nil, err
		}
	}

	return deletedTeam, nil
}

func deleteTeamHandler(apiUrl string, team *Team) (*Team, error) {
	deleteUrl := apiUrl + "/v1/teams/" + team.Name
	request, err := http.NewRequest(
		http.MethodDelete,
		deleteUrl,
		nil,
	)
	if err != nil {
		return nil, transientError{err: err}
	}

	statusCodeErrs := &StatusCodes{[]StatusCodeErr{
		{code: 404, errMessage: fmt.Errorf("could not delete team %s, does not exist", team.Name)},
	}}

	_, err = HttpHandler(request, statusCodeErrs)
	if err != nil {
		return nil, err
	}

	return team, nil
}

func ListTeams(apiUrl string) ([]string, error) {
	teams, err := ListTeamsHandler(apiUrl)
	if err != nil {
		terr := transientError{}
		if errors.As(err, &terr) {
			log.Printf("There was a problem getting the teams data \n")
			log.Fatal(err.Error())
		} else {
			return nil, err
		}
	}

	return teams, nil
}

func ListTeamsHandler(apiUrl string) ([]string, error) {
	teamsUrl := apiUrl + "/v1/teams"
	request, err := http.NewRequest(
		http.MethodGet,
		teamsUrl,
		nil,
	)
	if err != nil {
		return nil, transientError{err: err}
	}

	responseBytes, err := HttpHandler(request, nil)
	if err != nil {
		return nil, err
	}

	var teams []Team
	err = json.Unmarshal(responseBytes, &teams)
	if err != nil {
		return nil, transientError{err: err}
	}

	if len(teams) == 0 {
		return nil, fmt.Errorf("no teams found at %s", teamsUrl)
	}

	var teamNames []string
	for _, c := range teams {
		teamNames = append(teamNames, c.Name)
	}

	return teamNames, nil
}

func GetTeam(apiUrl string, team *Team) (*Team, error) {
	team, err := getTeamHandler(apiUrl, team)
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

func getTeamHandler(url string, team *Team) (*Team, error) {
	teamUrl := url + "/v1/teams/" + team.Name
	request, err := http.NewRequest(
		http.MethodGet,
		teamUrl,
		nil,
	)
	if err != nil {
		return nil, transientError{err: err}
	}

	statusCodeErrs := &StatusCodes{[]StatusCodeErr{
		{code: 404, errMessage: fmt.Errorf("could not find team %s at url %s", team.Name, teamUrl)},
	}}

	responseBytes, err := HttpHandler(request, statusCodeErrs)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(responseBytes, &team)
	if err != nil {
		return nil, transientError{err: err}
	}

	return team, nil

}
