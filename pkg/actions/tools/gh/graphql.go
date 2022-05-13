package gh

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/rsteube/carapace"
)

func graphQlAction(repo repo, query string, v interface{}, transform func() carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		params := make([]string, 0)
		if strings.Contains(query, "$owner") {
			params = append(params, "$owner: String!")
		}
		if strings.Contains(query, "$repo") {
			params = append(params, "$repo: String!")
		}
		queryParams := strings.Join(params, ",")
		if queryParams != "" {
			queryParams = "(" + queryParams + ")"
		}

		return carapace.ActionExecCommand("gh", "api", "--hostname", repo.host(), "--header", "Accept: application/vnd.github.merge-info-preview+json", "graphql", "-F", "owner="+repo.owner(), "-F", "repo="+repo.name(), "-f", fmt.Sprintf("query=query%v {%v}", queryParams, query))(func(output []byte) carapace.Action {
			if err := json.Unmarshal(output, &v); err != nil {
				return carapace.ActionMessage("failed to unmarshall response: " + err.Error())
			}
			return transform()
		})
	})
}