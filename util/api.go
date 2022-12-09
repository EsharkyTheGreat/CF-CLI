package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const CF_BASE_URL string = "https://codeforces.com/api"

type problemset_struct struct {
	Status string `json:"status"`
	Result struct {
		Problems []struct {
			ContestId      *int     `json:"contestId"`
			ProblemsetName string   `json:"problemsetName,omitempty"`
			Index          string   `json:"index"`
			Name           string   `json:"name"`
			Problem_type   string   `json:"type"`
			Points         *float64 `json:"points"`
			Rating         *int     `json:"rating"`
			Tags           []string
		} `json:"problems"`
		ProblemStatistics []struct {
			ContestId   *int   `json:"contestId"`
			Index       string `json:"index"`
			SolvedCount int    `json:"solvedCount"`
		} `json:"problemStatistics"`
	} `json:"result"`
}

func GetProblemset() {
	problemsetURL := fmt.Sprintf("%s/problemset.problems", CF_BASE_URL)
	resp, err := http.Get(problemsetURL)
	if err != nil {
		fmt.Println("Error in making Request")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error Reading Request Body")
	}
	var respJSON problemset_struct = problemset_struct{}
	err = json.Unmarshal(body, &respJSON)
	if err != nil {
		fmt.Println("Error Parsing JSON response")
	}
//	fmt.Printf("Status -> %s\n", respJSON.Status)
//	for index, ele := range respJSON.Result.Problems {
//		fmt.Printf("\tElement -> %d\n", index)
//		fmt.Printf("\tContest ID -> %d\n", *ele.ContestId)
//		fmt.Printf("\tProblemSetName -> %s\n", ele.ProblemsetName)
//		fmt.Printf("\tIndex -> %s\n", ele.Index)
//		fmt.Printf("\tName -> %s\n", ele.Name)
//		fmt.Printf("\tProblem_Type -> %s\n", ele.Problem_type)
//		fmt.Printf("\tPoints-> %v\n", ele.Points)
//		fmt.Printf("\tRating -> %d\n", ele.Rating)
//		fmt.Println(ele.Tags)
//		if index == 20 {
//			break
//		}
//
//	}
}
